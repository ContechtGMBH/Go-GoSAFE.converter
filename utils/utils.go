package utils

import (
	"bytes"
	"reflect"
	"strconv"
	"strings"

	"github.com/beevik/etree"
	"github.com/jmcvetta/neoism"
	"github.com/pebbe/go-proj-4/proj"
)

// Unknown is a simple helper
const Unknown = "unknown"

// Edge represents <trackEnd /> or <trackBegin /> of the particular track.
type Edge struct {
	Label      string
	Properties neoism.Props
}

// Switch structure extracted from the <connections /> container
type Switch struct {
	Properties neoism.Props
	Connection neoism.Props
}

// Crossing structure extracted from the <connections /> container
type Crossing struct {
	Properties neoism.Props
	Connection neoism.Props
}

// TrackTopologies structure represents complete <trackTopology /> section.
type TrackTopologies struct {
	Begin          Edge
	End            Edge
	Switch         []Switch
	Crossing       []Crossing
	CrossSection   []neoism.Props // TODO
	MileageChanges []neoism.Props // TODO
}

// TrackElements structure represents complete <trackElements /> section.
// Only <geoMappings /> from this section are not here because they belong directly to the track node.
type TrackElements struct {
	AxleWeightChanges        []neoism.Props
	Bridges                  []neoism.Props
	ClearanceGaugeChanges    []neoism.Props
	ElectrificationChanges   []neoism.Props
	GaugeChanges             []neoism.Props
	GradientChanges          []neoism.Props
	LevelCrossings           []neoism.Props
	OperationModeChanges     []neoism.Props
	OwnerChanges             []neoism.Props
	PlatformEdges            []neoism.Props
	PowerTransmissionChanges []neoism.Props
	RadiusChanges            []neoism.Props
	ServiceSections          []neoism.Props
	SpeedChanges             []neoism.Props
	TrackConditions          []neoism.Props
	TrainProtectionChanges   []neoism.Props
	Tunnels                  []neoism.Props
}

// OCSElements structure represents complete <ocsElements /> section.
type OCSElements struct {
	Signals                 []neoism.Props
	TrainDetectionElements  []neoism.Props // TODO
	Balises                 []neoism.Props
	TrainProtectionElements []neoism.Props
	StopPosts               []neoism.Props
	Derailers               []neoism.Props
	TrainRadioChanges       []neoism.Props
}

var wgs84, _ = proj.NewProj("+init=epsg:4326")

// Converts <geoMapings /> section to the WKTLinestring.
// If there are no coordinates, returns "unknown" string.
// If epsg is different than 4326 (wgs84), coords are transformed
func toWKTLinestring(t *etree.Element, epsg string) string {

	projection, err := proj.NewProj("+init=epsg:" + epsg)
	defer projection.Close()
	if err != nil {
		panic(err)
	}

	var coordsArray bytes.Buffer
	for _, g := range t.FindElements("trackElements/geoMappings/geoMapping/geoCoord") {
		c := g.SelectAttrValue("coord", Unknown)
		if c != Unknown {
			if epsg != "4326" {
				s := strings.Split(c, " ")
				x, _ := strconv.ParseFloat(s[0], 64)
				y, _ := strconv.ParseFloat(s[1], 64)
				xt, yt, _ := proj.Transform2(projection, wgs84, x, y)
				xts := strconv.FormatFloat(proj.RadToDeg(xt), 'f', 6, 64)
				yts := strconv.FormatFloat(proj.RadToDeg(yt), 'f', 6, 64)
				ct := xts + " " + yts
				coordsArray.WriteString(ct + ",")
			} else {
				coordsArray.WriteString(c + ",")
			}
		}
	}

	cas := coordsArray.String()

	if len(cas) > 0 {
		wkt := "LINESTRING(" + cas[:len(cas)-1] + ")"

		return wkt
	}

	return "unknown"
}

// Converts <geoCoord /> section to the WKTPoint.
// If there are no coordinates, returns "unknown" string.
// If epsg is different than 4326 (wgs84), coords are transformed
func toWKTPoint(e *etree.Element, epsg string) string {

	projection, err := proj.NewProj("+init=epsg:" + epsg)
	defer projection.Close()
	if err != nil {
		panic(err)
	}

	c := e.SelectAttrValue("coord", Unknown)
	if c != Unknown {
		if epsg != "4326" {
			s := strings.Split(c, " ")
			x, _ := strconv.ParseFloat(s[0], 64)
			y, _ := strconv.ParseFloat(s[1], 64)
			xt, yt, _ := proj.Transform2(projection, wgs84, x, y)
			xts := strconv.FormatFloat(proj.RadToDeg(xt), 'f', 6, 64)
			yts := strconv.FormatFloat(proj.RadToDeg(yt), 'f', 6, 64)
			wkt := "POINT(" + xts + " " + yts + ")"

			return wkt

		} else {
			wkt := "POINT(" + c + ")"

			return wkt
		}
	}

	return c
}

// ElementsUtils - XML extraction utilities.
// Extract tracks and all related elements.
type ElementsUtils struct{}

// Utility function, removes all unknown key:value pairs from the given map object
func removeUnknown(m neoism.Props) {
	for k, v := range m {
		if v == Unknown {
			delete(m, k)
		}
	}
}

// Utility function to extract attributes
func extractAttributes(e *etree.Element) neoism.Props {
	props := neoism.Props{}

	for _, p := range e.Attr {
		props[p.Key] = p.Value
	}

	return props
}

// GetTrackProperties creates track properties with valid railml properties and the geometry.
// A <track> represents one of possibly multiple tracks (= "pair of rails") that make up a line.
func (eu *ElementsUtils) GetTrackProperties(t *etree.Element, epsg string) neoism.Props {
	track := extractAttributes(t)
	geom := toWKTLinestring(t, epsg)
	if geom != Unknown {
		track["geometry"] = geom
	}
	return track
}

// GetTrackTopologies extracts topology elements like begin, end, connections.
// The element <trackTopology> works as a "container element" for several topology-related elements that "can't be touched in real life".
func (eu *ElementsUtils) GetTrackTopologies(t *etree.Element, epsg string) TrackTopologies {
	tt := TrackTopologies{}
	for _, topologies := range t.SelectElement("trackTopology").ChildElements() { // <trackBegin />, <trackEnd />, <connections />
		if (topologies.Tag == "trackBegin") || (topologies.Tag == "trackEnd") { // TRACK BEGINS AND ENDS
			te := Edge{}
			props := neoism.Props{}
			for _, child := range topologies.ChildElements() { // <bufferStop />, <openEnd />, <macroscopicNode />, <connection />, <geoCoord />

				if child.Tag == "geoCoord" {
					geom := toWKTPoint(child, epsg)
					if geom != Unknown {
						props["geometry"] = geom
					}
				} else {
					attrs := extractAttributes(child)
					for k, v := range attrs { // maybe it should be a separate function?
						props[k] = v
					}
					te.Label = strings.Title(child.Tag)
				}
			}

			te.Properties = props

			if topologies.Tag == "trackBegin" {
				tt.Begin = te
			} else if topologies.Tag == "trackEnd" {
				tt.End = te
			}

		} else if topologies.Tag == "connections" { // SWITCHES AND CROSSINGS
			var switches []Switch
			var crossings []Crossing
			for _, child := range topologies.ChildElements() { // <switch />, <crossing />
				props := extractAttributes(child)
				con := neoism.Props{}
				for _, nested := range child.ChildElements() { // <geoCoord />, <connection />
					if nested.Tag == "geoCoord" {
						geom := toWKTPoint(nested, epsg)
						if geom != Unknown {
							props["geometry"] = geom
						}
					} else if nested.Tag == "connection" {
						con = extractAttributes(nested)
					}
				}

				if child.Tag == "switch" {
					sw := Switch{}
					sw.Properties = props
					sw.Connection = con
					switches = append(switches, sw)
				} else if child.Tag == "crossing" {
					cr := Crossing{}
					cr.Properties = props
					cr.Connection = con
					crossings = append(crossings, cr)
				}
			}

			tt.Switch = switches
			tt.Crossing = crossings
		}

		// TODO <mileageChanges /> and <CrossSections />
		// These are not the most common tags and we don't even have objects like that in our railml files.

	}

	return tt
}

// GetTrackElements extracts all track elements related to the given track.
// The element <trackElements> works as a "container element" for elements which can be (more or less) "touched in real life".
func (eu *ElementsUtils) GetTrackElements(t *etree.Element, epsg string) TrackElements {

	te := TrackElements{}

	elements := t.SelectElement("trackElements")

	if elements == nil {
		return te
	}

	for _, element := range elements.ChildElements() {
		capitalized := strings.Title(element.Tag)
		if capitalized != "GeoMappings" {
			a := reflect.ValueOf(&te).Elem().FieldByName(capitalized) // grabs an array from the existing struct
			ae := []neoism.Props{}
			for _, child := range element.ChildElements() {
				attr := extractAttributes(child)
				g := child.SelectElement("geoCoord")
				if g != nil {
					attr["geometry"] = toWKTPoint(g, epsg)
				}
				ae = append(ae, attr)
			}
			v := reflect.ValueOf(ae) // a new array must be a value
			a.Set(v)                 // the value can be set as a new struct property
		}
	}
	return te
}

// GetOCSElements extracts all OCS elements related to the given track.
// The element <ocsElements> works as a "container element" for operation and control system elements.
// It doesn't group elements.
func (eu *ElementsUtils) GetOCSElements(t *etree.Element, epsg string) OCSElements {

	oe := OCSElements{}

	elements := t.SelectElement("ocsElements")

	if elements == nil {
		return oe
	}

	for _, element := range elements.ChildElements() {
		capitalized := strings.Title(element.Tag)
		if capitalized == "TrainDetectionElements" {
			// has two possible childs <trackCircuitBorder />  and <trainDetector />,
			// so for now it is skipped
			continue
		}
		a := reflect.ValueOf(&oe).Elem().FieldByName(capitalized) // grabs an array from the existing struct
		ae := []neoism.Props{}
		ct := strings.TrimSuffix(element.Tag, "s")
		for _, child := range element.ChildElements() {
			if child.Tag == ct { // because there can be another child <nameGroup />
				attr := extractAttributes(child)
				g := child.SelectElement("geoCoord")
				if g != nil {
					attr["geometry"] = toWKTPoint(g, epsg)
				}
				ae = append(ae, attr)
			}
		}
		v := reflect.ValueOf(ae) // new array must be a value
		a.Set(v)

	}

	return oe
}
