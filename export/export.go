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

type AxleWeightChange struct {
	XMLName     xml.Name `xml:"axleWeightChange"`
	Id          string   `xml:"id,attr"`
	Code        string   `xml:"code,attr,omitempty"`
	Name        string   `xml:"name,attr,omitempty"`
	Description string   `xml:"description,attr,omitempty"`
	Pos         string   `xml:"pos,attr,omitempty"`
	AbsPos      string   `xml:"absPos,attr,omitempty"`
	Dir         string   `xml:"dir,attr,omitempty"`
	Value       string   `xml:"value,attr,omitempty"`
	Meterload   string   `xml:"meterload,attr,omitempty"`
}

type AxleWeightChanges struct {
	XMLName          xml.Name `xml:"axleWeightChanges"`
	AxleWeightChange []AxleWeightChange
}

type Brigde struct {
	XMLName     xml.Name `xml:"brigde"` // yes, this is bri-G-D-e, what is going on with those people????
	Id          string   `xml:"id,attr"`
	Code        string   `xml:"code,attr,omitempty"`
	Name        string   `xml:"name,attr,omitempty"`
	Description string   `xml:"description,attr,omitempty"`
	Pos         string   `xml:"pos,attr,omitempty"`
	AbsPos      string   `xml:"absPos,attr,omitempty"`
	Dir         string   `xml:"dir,attr,omitempty"`
	Length      string   `xml:"length,attr,omitempty"`
	Meterload   string   `xml:"meterload,attr,omitempty"`
	Kind        string   `xml:"kind,attr,omitempty"`
}

type Bridges struct {
	XMLName xml.Name `xml:"bridges"` // and yes, this is a normal bri-D-G-es container
	Brigde  []Brigde
}

type ClearanceGaugeChange struct {
	XMLName     xml.Name `xml:"clearanceGaugeChange"`
	Id          string   `xml:"id,attr"`
	Code        string   `xml:"code,attr,omitempty"`
	Name        string   `xml:"name,attr,omitempty"`
	Description string   `xml:"description,attr,omitempty"`
	Pos         string   `xml:"pos,attr,omitempty"`
	AbsPos      string   `xml:"absPos,attr,omitempty"`
	Dir         string   `xml:"dir,attr,omitempty"`
}

type ClearanceGaugeChanges struct {
	XMLName              xml.Name `xml:"clearanceGaugeChanges"`
	ClearanceGaugeChange []ClearanceGaugeChange
}

type ElectrificationChange struct {
	XMLName         xml.Name `xml:"electrificationChange"`
	Id              string   `xml:"id,attr"`
	Code            string   `xml:"code,attr,omitempty"`
	Name            string   `xml:"name,attr,omitempty"`
	Description     string   `xml:"description,attr,omitempty"`
	Pos             string   `xml:"pos,attr,omitempty"`
	AbsPos          string   `xml:"absPos,attr,omitempty"`
	Dir             string   `xml:"dir,attr,omitempty"`
	Type            string   `xml:"type,attr,omitempty"`
	Voltage         string   `xml:"voltage,attr,omitempty"`
	Frequency       string   `xml:"frequency,attr,omitempty"`
	VMax            string   `xml:"vMax,attr,omitempty"`
	IsolatedSection string   `xml:"isolatedSection,attr,omitempty"`
}

type ElectrificationChanges struct {
	XMLName               xml.Name `xml:"electrificationChanges"`
	ElectrificationChange []ElectrificationChange
}

type GaugeChange struct {
	XMLName     xml.Name `xml:"gaugeChange"`
	Id          string   `xml:"id,attr"`
	Code        string   `xml:"code,attr,omitempty"`
	Name        string   `xml:"name,attr,omitempty"`
	Description string   `xml:"description,attr,omitempty"`
	Pos         string   `xml:"pos,attr,omitempty"`
	AbsPos      string   `xml:"absPos,attr,omitempty"`
	Dir         string   `xml:"dir,attr,omitempty"`
	Value       string   `xml:"value,attr,omitempty"`
}

type GaugeChanges struct {
	XMLName     xml.Name `xml:"gaugeChanges"`
	GaugeChange []GaugeChange
}

type GradientChange struct {
	XMLName          xml.Name `xml:"gradientChange"`
	Id               string   `xml:"id,attr"`
	Code             string   `xml:"code,attr,omitempty"`
	Name             string   `xml:"name,attr,omitempty"`
	Description      string   `xml:"description,attr,omitempty"`
	Pos              string   `xml:"pos,attr,omitempty"`
	AbsPos           string   `xml:"absPos,attr,omitempty"`
	Dir              string   `xml:"dir,attr,omitempty"`
	Slope            string   `xml:"slope,attr,omitempty"`
	TransitionLenght string   `xml:"transitionLenght,attr,omitempty"`
	TransitionRadius string   `xml:"transitionRadius,attr,omitempty"`
}

type GradientChanges struct {
	XMLName        xml.Name `xml:"gradientChanges"`
	GradientChange []GradientChange
}

type LevelCrossing struct {
	XMLName       xml.Name `xml:"levelCrossing"`
	Id            string   `xml:"id,attr"`
	Code          string   `xml:"code,attr,omitempty"`
	Name          string   `xml:"name,attr,omitempty"`
	Description   string   `xml:"description,attr,omitempty"`
	Pos           string   `xml:"pos,attr,omitempty"`
	AbsPos        string   `xml:"absPos,attr,omitempty"`
	Dir           string   `xml:"dir,attr,omitempty"`
	OcpStationRef string   `xml:"ocpStationRef,attr,omitempty"`
	ControllerRef string   `xml:"controllerRef,attr,omitempty"`
	Length        string   `xml:"length,attr,omitempty"`
	Angle         string   `xml:"angle,attr,omitempty"`
	Protection    string   `xml:"protection,attr,omitempty"`
}

type LevelCrossings struct {
	XMLName       xml.Name `xml:"levelCrossings"`
	LevelCrossing []LevelCrossing
}

type OperationModeChange struct {
	XMLName           xml.Name `xml:"operationModeChange"`
	Id                string   `xml:"id,attr"`
	Code              string   `xml:"code,attr,omitempty"`
	Name              string   `xml:"name,attr,omitempty"`
	Description       string   `xml:"description,attr,omitempty"`
	Pos               string   `xml:"pos,attr,omitempty"`
	AbsPos            string   `xml:"absPos,attr,omitempty"`
	Dir               string   `xml:"dir,attr,omitempty"`
	ModeLegislative   string   `xml:"modeLegislative,attr,omitempty"`
	ModeExecutive     string   `xml:"modeEvecutive,attr,omitempty"`
	ClearanceManaging string   `xml:"clereanceMamaging,attr,omitempty"`
}

type OperationModeChanges struct {
	XMLName             xml.Name `xml:"operationModeChanges"`
	OperationModeChange []OperationModeChange
}

type OwnerChange struct {
	XMLName                  xml.Name `xml:"ownerChange"`
	Id                       string   `xml:"id,attr"`
	Code                     string   `xml:"code,attr,omitempty"`
	Name                     string   `xml:"name,attr,omitempty"`
	Description              string   `xml:"description,attr,omitempty"`
	Pos                      string   `xml:"pos,attr,omitempty"`
	AbsPos                   string   `xml:"absPos,attr,omitempty"`
	Dir                      string   `xml:"dir,attr,omitempty"`
	OwnerName                string   `xml:"ownerName,attr,omitempty"`
	InfrastructureManagerRef string   `xml:"infrastructureManagerRef,attr,omitempty"`
}

type OwnerChanges struct {
	XMLName     xml.Name `xml:"ownerChanges"`
	OwnerChange []OwnerChange
}

type PlatformEdge struct {
	XMLName               xml.Name `xml:"platformEdge"`
	Id                    string   `xml:"id,attr"`
	Code                  string   `xml:"code,attr,omitempty"`
	Name                  string   `xml:"name,attr,omitempty"`
	Description           string   `xml:"description,attr,omitempty"`
	Pos                   string   `xml:"pos,attr,omitempty"`
	AbsPos                string   `xml:"absPos,attr,omitempty"`
	Dir                   string   `xml:"dir,attr,omitempty"`
	OcpRef                string   `xml:"ocpRef,attr,omitempty"`
	Length                string   `xml:"length,attr,omitempty"`
	Height                string   `xml:"height,attr,omitempty"`
	Side                  string   `xml:"side,attr,omitempty"`
	ParentPlatformEdgeRef string   `xml:"parentPlatformEdgeRef,attr,omitempty"`
}

type PlatformEdges struct {
	XMLName      xml.Name `xml:"platformEdges"`
	PlatformEdge []PlatformEdge
}

type PowerTransmissionChange struct {
	XMLName     xml.Name `xml:"powerTransmissionChange"`
	Id          string   `xml:"id,attr"`
	Code        string   `xml:"code,attr,omitempty"`
	Name        string   `xml:"name,attr,omitempty"`
	Description string   `xml:"description,attr,omitempty"`
	Pos         string   `xml:"pos,attr,omitempty"`
	AbsPos      string   `xml:"absPos,attr,omitempty"`
	Dir         string   `xml:"dir,attr,omitempty"`
	Type        string   `xml:"type,attr,omitempty"`
	Style       string   `xml:"style,attr,omitempty"`
}

type PowerTransmissionChanges struct {
	XMLName                 xml.Name `xml:"powerTransmissionChanges"`
	PowerTransmissionChange []PowerTransmissionChange
}

type RadiusChange struct {
	XMLName                    xml.Name `xml:"radiusChange"`
	Id                         string   `xml:"id,attr"`
	Code                       string   `xml:"code,attr,omitempty"`
	Name                       string   `xml:"name,attr,omitempty"`
	Description                string   `xml:"description,attr,omitempty"`
	Pos                        string   `xml:"pos,attr,omitempty"`
	AbsPos                     string   `xml:"absPos,attr,omitempty"`
	Dir                        string   `xml:"dir,attr,omitempty"`
	Radius                     string   `xml:"radius,attr,omitempty"`
	Superelevation             string   `xml:"superelevation,attr,omitempty"`
	GeometryElementDescription string   `xml:"geometryElementDescription,attr,omitempty"`
}

type RadiusChanges struct {
	XMLName      xml.Name `xml:"radiusChanges"`
	RadiusChange []RadiusChange
}

type ServiceSection struct {
	XMLName                 xml.Name `xml:"serviceSection"`
	Id                      string   `xml:"id,attr"`
	Code                    string   `xml:"code,attr,omitempty"`
	Name                    string   `xml:"name,attr,omitempty"`
	Description             string   `xml:"description,attr,omitempty"`
	Pos                     string   `xml:"pos,attr,omitempty"`
	AbsPos                  string   `xml:"absPos,attr,omitempty"`
	Dir                     string   `xml:"dir,attr,omitempty"`
	OcpRef                  string   `xml:"ocpRef,attr,omitempty"`
	Length                  string   `xml:"length,attr,omitempty"`
	Height                  string   `xml:"height,attr,omitempty"`
	Side                    string   `xml:"side,attr,omitempty"`
	ParentServiceSectionRef string   `xml:"parentServiceSectionRef,attr,omitempty"`
	Ramp                    string   `xml:"ramp,attr,omitempty"`
	Maintenance             string   `xml:"maintenance,attr,omitempty"`
	LoadingFacility         string   `xml:"loadingFacility,attr,omitempty"`
	Cleaning                string   `xml:"cleaning,attr,omitempty"`
	Fueling                 string   `xml:"fueling,attr,omitempty"`
	Parking                 string   `xml:"parking,attr,omitempty"`
	Preheating              string   `xml:"preheating,attr,omitempty"`
}

type ServiceSections struct {
	XMLName        xml.Name `xml:"serviceSections"`
	ServiceSection []ServiceSection
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

type TrackCondition struct {
	XMLName     xml.Name `xml:"trackCondition"`
	Id          string   `xml:"id,attr"`
	Code        string   `xml:"code,attr,omitempty"`
	Name        string   `xml:"name,attr,omitempty"`
	Description string   `xml:"description,attr,omitempty"`
	Pos         string   `xml:"pos,attr,omitempty"`
	AbsPos      string   `xml:"absPos,attr,omitempty"`
	Dir         string   `xml:"dir,attr,omitempty"`
	Length      string   `xml:"length,attr,omitempty"`
	Type        string   `xml:"type,attr,omitempty"`
}

type TrackConditions struct {
	XMLName        xml.Name `xml:"trackConditions"`
	TrackCondition []TrackCondition
}

type TrainProtectionChange struct {
	XMLName     xml.Name `xml:"trainProtectionChange"`
	Id          string   `xml:"id,attr"`
	Code        string   `xml:"code,attr,omitempty"`
	Name        string   `xml:"name,attr,omitempty"`
	Description string   `xml:"description,attr,omitempty"`
	Pos         string   `xml:"pos,attr,omitempty"`
	AbsPos      string   `xml:"absPos,attr,omitempty"`
	Dir         string   `xml:"dir,attr,omitempty"`
	Medium      string   `xml:"medium,attr,omitempty"`
	Monitoring  string   `xml:"monitoring,attr,omitempty"`
}

type TrainProtectionChanges struct {
	XMLName               xml.Name `xml:"trainProtectionChanges"`
	TrainProtectionChange []TrainProtectionChange
}

type Tunnel struct {
	XMLName      xml.Name `xml:"tunnel"`
	Id           string   `xml:"id,attr"`
	Code         string   `xml:"code,attr,omitempty"`
	Name         string   `xml:"name,attr,omitempty"`
	Description  string   `xml:"description,attr,omitempty"`
	Pos          string   `xml:"pos,attr,omitempty"`
	AbsPos       string   `xml:"absPos,attr,omitempty"`
	Dir          string   `xml:"dir,attr,omitempty"`
	Length       string   `xml:"length,attr,omitempty"`
	CrossSection string   `xml:"crossSection,attr,omitempty"`
	Kind         string   `xml:"kind,attr,omitempty"`
}

type Tunnels struct {
	XMLName xml.Name `xml:"tunnels"`
	Tunnel  []Tunnel
}

type TrackElements struct {
	XMLName                  xml.Name           `xml:"trackElements"`
	AxleWeightChanges        *AxleWeightChanges // pointer to the struct means that if there is no object by default (empty parent), must by created by passing a reference
	Bridges                  *Bridges
	ClearanceGaugeChanges    *ClearanceGaugeChanges
	ElectrificationChanges   *ElectrificationChanges
	GaugeChanges             *GaugeChanges
	GradientChanges          *GradientChanges
	LevelCrossings           *LevelCrossings
	OperationModeChanges     *OperationModeChanges
	OwnerChanges             *OwnerChanges
	PlatformEdges            *PlatformEdges
	PowerTransmissionChanges *PowerTransmissionChanges
	RadiusChanges            *RadiusChanges
	ServiceSections          *ServiceSections
	SpeedChanges             *SpeedChanges
	TrackConditions          *TrackConditions
	TrainProtectionChanges   *TrainProtectionChanges
	Tunnels                  *Tunnels
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

type UnmarshalledTrack struct {
	Track        neoism.Node         `json:"t"`
	Relationship neoism.Relationship `json:"r"`
	Node         neoism.Node         `json:"n"`
	Label        []string            `json:"labels(n)"`
}

func ExportTrack(id string) Track {
	db := config.GetDBConnection()
	query := "MATCH (t:Track {id:{trackId}})-[r:BEGINS|ENDS|HAS_TRACK_ELEMENT]-(n) RETURN t,r,n,labels(n)"
	track := []UnmarshalledTrack{}
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
func createTrackEdge(lb string, xteg *TrackEdge, t UnmarshalledTrack) {
	switch lb {
	case "BufferStop":
		nbs := &BufferStop{}
		xteg.BufferStop = append(
			xteg.BufferStop,
			*createElementFromNode(&t.Node, nbs).(*BufferStop),
		)
	case "Connection":
		nco := &Connection{}
		xteg.Connection = append(
			xteg.Connection,
			*createElementFromNode(&t.Node, nco).(*Connection),
		)
	case "OpenEnd":
		noe := &OpenEnd{}
		xteg.OpenEnd = append(
			xteg.OpenEnd,
			*createElementFromNode(&t.Node, noe).(*OpenEnd),
		)
	case "MacroscopicNode":
		nmn := &MacroscopicNode{}
		xteg.MacroscopicNode = append(
			xteg.MacroscopicNode,
			*createElementFromNode(&t.Node, nmn).(*MacroscopicNode),
		)
	}
}

// TRACK ELEMENTS
func createTrackElement(lb string, xtel *TrackElements, t UnmarshalledTrack) {
	switch lb {
	case "AxleWeightChange":
		if xtel.AxleWeightChanges == nil {
			xtel.AxleWeightChanges = &AxleWeightChanges{}
		}
		nawc := &AxleWeightChange{}
		xtel.AxleWeightChanges.AxleWeightChange = append(
			xtel.AxleWeightChanges.AxleWeightChange,
			*createElementFromNode(&t.Node, nawc).(*AxleWeightChange),
		)
	case "Brigde":
		if xtel.Bridges == nil {
			xtel.Bridges = &Bridges{}
		}
		nb := &Brigde{}
		xtel.Bridges.Brigde = append(
			xtel.Bridges.Brigde,
			*createElementFromNode(&t.Node, nb).(*Brigde),
		)
	case "ClearanceGaugeChange":
		if xtel.ClearanceGaugeChanges == nil {
			xtel.ClearanceGaugeChanges = &ClearanceGaugeChanges{}
		}
		ncgc := &ClearanceGaugeChange{}
		xtel.ClearanceGaugeChanges.ClearanceGaugeChange = append(
			xtel.ClearanceGaugeChanges.ClearanceGaugeChange,
			*createElementFromNode(&t.Node, ncgc).(*ClearanceGaugeChange),
		)
	case "ElectrificationChange":
		if xtel.ElectrificationChanges == nil {
			xtel.ElectrificationChanges = &ElectrificationChanges{}
		}
		ec := &ElectrificationChange{}
		xtel.ElectrificationChanges.ElectrificationChange = append(
			xtel.ElectrificationChanges.ElectrificationChange,
			*createElementFromNode(&t.Node, ec).(*ElectrificationChange),
		)
	case "GaugeChange":
		if xtel.GaugeChanges == nil {
			xtel.GaugeChanges = &GaugeChanges{}
		}
		ngc := &GaugeChange{}
		xtel.GaugeChanges.GaugeChange = append(
			xtel.GaugeChanges.GaugeChange,
			*createElementFromNode(&t.Node, ngc).(*GaugeChange),
		)
	case "GradientChange":
		if xtel.GradientChanges == nil {
			xtel.GradientChanges = &GradientChanges{}
		}
		ngrc := &GradientChange{}
		xtel.GradientChanges.GradientChange = append(
			xtel.GradientChanges.GradientChange,
			*createElementFromNode(&t.Node, ngrc).(*GradientChange),
		)
	case "LevelCrossing":
		if xtel.LevelCrossings == nil {
			xtel.LevelCrossings = &LevelCrossings{}
		}
		nlc := &LevelCrossing{}
		xtel.LevelCrossings.LevelCrossing = append(
			xtel.LevelCrossings.LevelCrossing,
			*createElementFromNode(&t.Node, nlc).(*LevelCrossing),
		)
	case "OperationModeChange":
		if xtel.OperationModeChanges == nil {
			xtel.OperationModeChanges = &OperationModeChanges{}
		}
		nomc := &OperationModeChange{}
		xtel.OperationModeChanges.OperationModeChange = append(
			xtel.OperationModeChanges.OperationModeChange,
			*createElementFromNode(&t.Node, nomc).(*OperationModeChange),
		)
	case "OwnerChange":
		if xtel.OwnerChanges == nil {
			xtel.OwnerChanges = &OwnerChanges{}
		}
		noc := &OwnerChange{}
		xtel.OwnerChanges.OwnerChange = append(
			xtel.OwnerChanges.OwnerChange,
			*createElementFromNode(&t.Node, noc).(*OwnerChange),
		)
	case "PlatformEdge":
		if xtel.PlatformEdges == nil {
			xtel.PlatformEdges = &PlatformEdges{}
		}
		npe := &PlatformEdge{}
		xtel.PlatformEdges.PlatformEdge = append(
			xtel.PlatformEdges.PlatformEdge,
			*createElementFromNode(&t.Node, npe).(*PlatformEdge),
		)
	case "PowerTransmissionChange":
		if xtel.PowerTransmissionChanges == nil {
			xtel.PowerTransmissionChanges = &PowerTransmissionChanges{}
		}
		nptc := &PowerTransmissionChange{}
		xtel.PowerTransmissionChanges.PowerTransmissionChange = append(
			xtel.PowerTransmissionChanges.PowerTransmissionChange,
			*createElementFromNode(&t.Node, nptc).(*PowerTransmissionChange),
		)
	case "RadiusChange":
		if xtel.RadiusChanges == nil {
			xtel.RadiusChanges = &RadiusChanges{}
		}
		nrc := &RadiusChange{}
		xtel.RadiusChanges.RadiusChange = append(
			xtel.RadiusChanges.RadiusChange,
			*createElementFromNode(&t.Node, nrc).(*RadiusChange),
		)
	case "ServiceSection":
		if xtel.ServiceSections == nil {
			xtel.ServiceSections = &ServiceSections{}
		}
		nss := &ServiceSection{}
		xtel.ServiceSections.ServiceSection = append(
			xtel.ServiceSections.ServiceSection,
			*createElementFromNode(&t.Node, nss).(*ServiceSection),
		)
	case "SpeedChange":
		if xtel.SpeedChanges == nil {
			xtel.SpeedChanges = &SpeedChanges{}
		}
		nsc := &SpeedChange{}
		xtel.SpeedChanges.SpeedChange = append(
			xtel.SpeedChanges.SpeedChange,
			*createElementFromNode(&t.Node, nsc).(*SpeedChange),
		)
	case "TrackCondition":
		if xtel.TrackConditions == nil {
			xtel.TrackConditions = &TrackConditions{}
		}
		ntc := &TrackCondition{}
		xtel.TrackConditions.TrackCondition = append(
			xtel.TrackConditions.TrackCondition,
			*createElementFromNode(&t.Node, ntc).(*TrackCondition),
		)
	case "TrainProtectionChange":
		if xtel.TrainProtectionChanges == nil {
			xtel.TrainProtectionChanges = &TrainProtectionChanges{}
		}
		ntpc := &TrainProtectionChange{}
		xtel.TrainProtectionChanges.TrainProtectionChange = append(
			xtel.TrainProtectionChanges.TrainProtectionChange,
			*createElementFromNode(&t.Node, ntpc).(*TrainProtectionChange),
		)
	case "Tunnel":
		if xtel.Tunnels == nil {
			xtel.Tunnels = &Tunnels{}
		}
		nt := &Tunnel{}
		xtel.Tunnels.Tunnel = append(
			xtel.Tunnels.Tunnel,
			*createElementFromNode(&t.Node, nt).(*Tunnel),
		)
	}
}

// createElementFromNode converts a neoism.Node to the interface that can be passed as a struct - *nif.(*StructType)
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
