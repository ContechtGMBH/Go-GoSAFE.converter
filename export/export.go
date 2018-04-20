package export

import (
	"encoding/xml"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"Go-GoSAFE.converter/config"

	"github.com/jmcvetta/neoism"
)

type BufferStop struct {
	XMLName     xml.Name `xml:"bufferStop"`
	Id          string   `xml:"id,attr"`
	Code        string   `xml:"code,attr,omitempty"`
	Name        string   `xml:"name,attr,omitempty"`
	Description string   `xml:"description,attr,omitempty"`
}

type Connection struct {
	XMLName xml.Name `xml:"connection"`
	Id      string   `xml:"id,attr"`
	Ref     string   `xml:"ref,attr,omitempty"`
}

type OpenEnd struct {
	XMLName     xml.Name `xml:"openEnd"`
	Id          string   `xml:"id,attr"`
	Code        string   `xml:"code,attr,omitempty"`
	Name        string   `xml:"name,attr,omitempty"`
	Description string   `xml:"description,attr,omitempty"`
}

type MacroscopicNode struct {
	XMLName       xml.Name `xml:"macroscopicNode"`
	OcpRef        string   `xml:"ocpRef,attr"`
	FlowDirection string   `xml:"flowDirection,attr,omitempty"`
}

type TrackEdge struct {
	XMLName         xml.Name
	Id              string `xml:"id,attr"`
	Pos             string `xml:"pos,attr,omitempty"`
	AbsPos          string `xml:"absPos,attr,omitempty"`
	AbsDir          string `xml:"absDir,attr,omitempty"`
	BufferStop      []BufferStop
	Connection      []Connection
	OpenEnd         []OpenEnd
	MacroscopicNode []MacroscopicNode
}

type TrackTopology struct {
	XMLName    xml.Name `xml:"trackTopology"`
	TrackBegin TrackEdge
	TrackEnd   TrackEdge
}

type SpeedChange struct {
	XMLName       xml.Name `xml:"speedChange"`
	Id            string   `xml:"id,attr"`
	Code          string   `xml:"code,attr,omitempty"`
	Name          string   `xml:"name,attr,omitempty"`
	Description   string   `xml:"description,attr,omitempty"`
	Pos           string   `xml:"pos,attr,omitempty"`
	AbsPos        string   `xml:"absPos,attr,omitempty"`
	Dir           string   `xml:"dir,attr,omitempty"`
	ProfileRef    string   `xml:"profileRef,attr,omitempty"`
	Status        string   `xml:"status,attr,omitempty"`
	VMax          string   `xml:"vMax,attr,omitempty"`
	TrainRelation string   `xml:"trainRelation,attr,omitempty"`
	MandatoryStop string   `xml:"mandatoryStop,attr,omitempty"`
	Signalised    string   `xml:"signalised,attr,omitempty"`
}

type SpeedChanges struct {
	XMLName     xml.Name `xml:"speedChanges"`
	SpeedChange []SpeedChange
}

type TrackElements struct {
	XMLName      xml.Name      `xml:"trackElements"`
	SpeedChanges *SpeedChanges // pointer to the struct means that if there is no object by default (empty parent), must by created by passing a reference
}

type Track struct {
	XMLName       xml.Name `xml:"track"`
	Id            string   `xml:"id,attr"`
	TrackTopology TrackTopology
	TrackElements TrackElements
}

type Tracks struct {
	XMLName xml.Name `xml:"tracks"`
	Tracks  []Track
}

type InfraAttributes struct {
	XMLName xml.Name `xml:"infraAttributes"`
	Id      string   `xml:"id,attr"`
}

type InfraAttrGroups struct {
	XMLName         xml.Name `xml:"infraAttrGroups"`
	InfraAttributes []InfraAttributes
}

type Infrastructure struct {
	XMLName         xml.Name `xml:"infrastructure"`
	Id              string   `xml:"id,attr"`
	Name            string   `xml:"name,attr,omitempty"`
	InfraAttrGroups InfraAttrGroups
	Tracks          Tracks
}

type Metadata struct {
	XMLName xml.Name `xml:"metadata"`
	Source  string   `xml:"dc:source"`
	Creator string   `xml:"dc:creator"`
	Date    string   `xml:"dc:date"`
}

type Railml struct {
	XMLName        xml.Name `xml:"railml"`
	Version        string   `xml:"version,attr"`
	Xmlns          string   `xml:"xmlns,attr"`
	Xsi            string   `xml:"xmlns:xsi,attr"`
	SchemaLocation string   `xml:"xsi:schemaLocation,attr"`
	Metadata       Metadata
	Infrastructure Infrastructure
}

func ExportLine(lineId string) interface{} {
	db := config.GetDBConnection()
	query := "MATCH (n:Line {id:{lineId}})-[:HAS_TRACK]-(t) RETURN t.id"
	tid := []struct {
		ID string `json:"t.id"`
	}{}
	cq := neoism.CypherQuery{
		Statement:  query,
		Parameters: neoism.Props{"lineId": lineId},
		Result:     &tid,
	}

	err := db.Cypher(&cq)
	if err != nil {
		panic(err)
	}

	ts := Tracks{}
	for _, t := range tid {
		v := ExportTrack(t.ID)
		ts.Tracks = append(ts.Tracks, v)

	}
	in := Infrastructure{
		Id:              lineId + "-" + time.Now().Format("20060102150405"), // UNIX timestamp format
		Name:            lineId,
		InfraAttrGroups: InfraAttrGroups{},
		Tracks:          ts,
	}
	meta := Metadata{
		Source:  "GoSAFE Converter v0.1",
		Creator: "Damian Harasymczuk harasymczuk_at_contecht.eu",
		Date:    time.Now().Format("2006-01-02 15:04:05"),
	}
	rm := Railml{
		Version:        "2.2",
		Xmlns:          "http://www.railml.org/schemas/2013",
		Xsi:            "http://www.w3.org/2001/XMLSchema-instance",
		SchemaLocation: "http://www.railml.org/schemas/2013 http://schemas.railml.org/2013/railML-2.2/railML.xsd",
		Metadata:       meta,
		Infrastructure: in,
	}

	return rm
}

type UnmarshalTrack struct {
	Track        neoism.Node         `json:"t"`
	Relationship neoism.Relationship `json:"r"`
	Node         neoism.Node         `json:"n"`
	Label        []string            `json:"labels(n)"`
}

func ExportTrack(id string) Track {
	db := config.GetDBConnection()
	query := "MATCH (t:Track {id:{trackId}})-[r:BEGINS|ENDS|HAS_TRACK_ELEMENT]-(n) RETURN t,r,n,labels(n)"
	track := []UnmarshalTrack{}
	cq := neoism.CypherQuery{
		Statement:  query,
		Parameters: neoism.Props{"trackId": id},
		Result:     &track,
	}

	e := db.Cypher(&cq)
	_ = e

	xtb := TrackEdge{XMLName: xml.Name{Local: "trackBegin"}}
	xte := TrackEdge{XMLName: xml.Name{Local: "trackEnd"}}
	xtel := TrackElements{}

	for _, t := range track {
		if len(t.Label) < 1 {
			continue
		}
		lb := t.Label[0]
		ty := t.Relationship.Type
		switch ty {
		case "BEGINS":
			createTrackEdge(lb, &xtb, t)
		case "ENDS":
			createTrackEdge(lb, &xte, t)
		case "HAS_TRACK_ELEMENT":
			createTrackElement(lb, &xtel, t)
		}
	}
	xtt := TrackTopology{TrackBegin: xtb, TrackEnd: xte}
	xt := Track{Id: id, TrackTopology: xtt, TrackElements: xtel}
	/*
		output, err := xml.MarshalIndent(xt, "  ", "    ")
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}

		os.Stdout.Write(output)
	*/
	return xt
}

// TRACK TOPOLOGIES
func createTrackEdge(lb string, xteg *TrackEdge, t UnmarshalTrack) {
	switch lb {
	case "BufferStop":
		xteg.BufferStop = append(
			xteg.BufferStop,
			createBufferStop(&t.Node),
		)
	case "Connection":
		xteg.Connection = append(
			xteg.Connection,
			createConnection(&t.Node),
		)
	case "OpenEnd":
		xteg.OpenEnd = append(
			xteg.OpenEnd,
			createOpenEnd(&t.Node),
		)
	case "MacroscopicNode":
		xteg.MacroscopicNode = append(
			xteg.MacroscopicNode,
			createMacroscopicNode(&t.Node),
		)
	}
}

func createBufferStop(n *neoism.Node) BufferStop {
	nbs := BufferStop{}
	for k, v := range n.Data {
		switch k {
		case "id":
			nbs.Id = v.(string)
		case "code":
			nbs.Code = v.(string)
		case "name":
			nbs.Name = v.(string)
		case "description":
			nbs.Description = v.(string)
		}
	}
	return nbs
}

func createConnection(n *neoism.Node) Connection {
	nc := Connection{}
	for k, v := range n.Data {
		switch k {
		case "id":
			nc.Id = v.(string)
		case "ref":
			nc.Ref = v.(string)
		}
	}
	return nc
}

func createOpenEnd(n *neoism.Node) OpenEnd {
	noe := OpenEnd{}
	for k, v := range n.Data {
		switch k {
		case "id":
			noe.Id = v.(string)
		case "code":
			noe.Code = v.(string)
		case "name":
			noe.Name = v.(string)
		case "description":
			noe.Description = v.(string)
		}
	}
	return noe
}

func createMacroscopicNode(n *neoism.Node) MacroscopicNode {
	nmn := MacroscopicNode{}
	for k, v := range n.Data {
		switch k {
		case "ocpRef":
			nmn.OcpRef = v.(string)
		case "flowDirection":
			nmn.FlowDirection = v.(string)
		}
	}
	return nmn
}

// TRACK ELEMENTS
func createTrackElement(lb string, xtel *TrackElements, t UnmarshalTrack) {
	switch lb {
	case "SpeedChange":
		if xtel.SpeedChanges == nil {
			xtel.SpeedChanges = &SpeedChanges{}
		}
		nsc := &SpeedChange{}
		xtel.SpeedChanges.SpeedChange = append(
			xtel.SpeedChanges.SpeedChange,
			*createElementFromNode(&t.Node, nsc).(*SpeedChange),
		)
	}
}

func createElementFromNode(n *neoism.Node, nif interface{}) interface{} {
	for k, v := range n.Data {
		capitalized := strings.Title(k)
		err := setField(nif, capitalized, v.(string))
		if err != nil {
			continue
		}
	}
	return nif
}

// setField is a simple helper, sets a value in the given struct by its name (string)
func setField(v interface{}, name string, value string) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.Elem().Kind() != reflect.Struct {
		return errors.New("v must be pointer to struct")
	}
	rv = rv.Elem()

	fv := rv.FieldByName(name)
	if !fv.IsValid() {
		return fmt.Errorf("not a field name: %s", name)
	}

	if !fv.CanSet() {
		return fmt.Errorf("cannot set field %s", name)
	}

	if fv.Kind() != reflect.String {
		return fmt.Errorf("%s is not a string field", name)
	}

	fv.SetString(value)
	return nil
}
