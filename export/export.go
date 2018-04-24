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

type SwitchOrCrossing struct { // exactly the same attrs
	XMLName             xml.Name
	Id                  string `xml:"id,attr"`
	Code                string `xml:"code,attr,omitempty"`
	Name                string `xml:"name,attr,omitempty"`
	Description         string `xml:"description,attr,omitempty"`
	Pos                 string `xml:"pos,attr,omitempty"`
	AbsPos              string `xml:"absPos,attr,omitempty"`
	OcpStationRef       string `xml:"ocpStationRef,attr,omitempty"`
	ControllerRef       string `xml:"controllerRef,attr,omitempty"`
	TrackContinueCourse string `xml:"trackContinueCourse,attr,omitempty"`
	TrackContinueRadius string `xml:"trackContinueRadius,attr,omitempty"`
	NormalPosition      string `xml:"normalPosition,attr,omitempty"`
	Model               string `xml:"model,attr,omitempty"`
	Length              string `xml:"length,attr,omitempty"`
	Type                string `xml:"type,attr,omitempty"`
	Connection          Connection
}

type Connections struct {
	XMLName     xml.Name `xml:"connections"`
	Connections []SwitchOrCrossing
}

type TrackTopology struct {
	XMLName     xml.Name `xml:"trackTopology"`
	TrackBegin  TrackEdge
	TrackEnd    TrackEdge
	Connections Connections
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

type Signal struct {
	XMLName       xml.Name `xml:"signal"`
	Id            string   `xml:"id,attr"`
	Code          string   `xml:"code,attr,omitempty"`
	Name          string   `xml:"name,attr,omitempty"`
	Description   string   `xml:"description,attr,omitempty"`
	Pos           string   `xml:"pos,attr,omitempty"`
	AbsPos        string   `xml:"absPos,attr,omitempty"`
	Dir           string   `xml:"dir,attr,omitempty"`
	OcpStationRef string   `xml:"ocpStationRef,attr,omitempty"`
	ControllerRef string   `xml:"controllerRef,attr,omitempty"`
	Type          string   `xml:"type,attr,omitempty"`
	Function      string   `xml:"function,attr,omitempty"`
	Virtual       string   `xml:"virtual,attr,omitempty"`
	RuleCode      string   `xml:"ruleCode,attr,omitempty"`
	TrackDist     string   `xml:"trackDist,attr,omitempty"`
	Height        string   `xml:"height,attr,omitempty"`
}

type Signals struct {
	XMLName xml.Name `xml:"signals"`
	Signal  []Signal
}

// TrainDetectionElements skipped for now

type Balise struct {
	XMLName                xml.Name `xml:"balise"`
	Id                     string   `xml:"id,attr"`
	Code                   string   `xml:"code,attr,omitempty"`
	Name                   string   `xml:"name,attr,omitempty"`
	Description            string   `xml:"description,attr,omitempty"`
	Pos                    string   `xml:"pos,attr,omitempty"`
	AbsPos                 string   `xml:"absPos,attr,omitempty"`
	Dir                    string   `xml:"dir,attr,omitempty"`
	CountryID              string   `xml:"countryID,attr,omitempty"`
	GroupID                string   `xml:"groupID,attr,omitempty"`
	LinkingAccuracy        string   `xml:"linkingAccuracy,attr,omitempty"`
	LinkReactionAscending  string   `xml:"linkReactionAscending,attr,omitempty"`
	LinkReactionDescending string   `xml:"linkReactionDescending,attr,omitempty"`
	StaticTelegram         string   `xml:"staticTelegram,attr,omitempty"`
	Ndx                    string   `xml:"ndx,attr,omitempty"`
}

type Balises struct {
	XMLName xml.Name `xml:"balises"`
	Balise  []Balise
}

type TrainProtectionElement struct {
	XMLName               xml.Name `xml:"trainProtectionElement"`
	Id                    string   `xml:"id,attr"`
	Code                  string   `xml:"code,attr,omitempty"`
	Name                  string   `xml:"name,attr,omitempty"`
	Description           string   `xml:"description,attr,omitempty"`
	Pos                   string   `xml:"pos,attr,omitempty"`
	AbsPos                string   `xml:"absPos,attr,omitempty"`
	Dir                   string   `xml:"dir,attr,omitempty"`
	Medium                string   `xml:"medium,attr,omitempty"`
	System                string   `xml:"system,attr,omitempty"`
	TrainProtectionSystem string   `xml:"trainProtectionSystem,attr,omitempty"`
	Model                 string   `xml:"model,attr,omitempty"`
}

type TrainProtectionElements struct {
	XMLName                xml.Name `xml:"trainProtectionElements"`
	TrainProtectionElement []TrainProtectionElement
}

type StopPost struct {
	XMLName           xml.Name `xml:"stopPost"`
	Id                string   `xml:"id,attr"`
	Code              string   `xml:"code,attr,omitempty"`
	Name              string   `xml:"name,attr,omitempty"`
	Description       string   `xml:"description,attr,omitempty"`
	Pos               string   `xml:"pos,attr,omitempty"`
	AbsPos            string   `xml:"absPos,attr,omitempty"`
	Dir               string   `xml:"dir,attr,omitempty"`
	RuleCode          string   `xml:"ruleCode,attr,omitempty"`
	TrainRelation     string   `xml:"trainRelation,attr,omitempty"`
	PlatformEdgeRef   string   `xml:"platformEdgeRef,attr,omitempty"`
	TrainLength       string   `xml:"trainLength,attr,omitempty"`
	AxleCount         string   `xml:"axleCoount,attr,omitempty"`
	WagonCount        string   `xml:"wagonCount,attr,omitempty"`
	VerbalConstraints string   `xml:"verbalConstraints,attr,omitempty"`
	Virtual           string   `xml:"virtual,attr,omitempty"`
	OcpRef            string   `xml:"ocpRef,attr,omitempty"`
}

type StopPosts struct {
	XMLName  xml.Name `xml:"stopPosts"`
	StopPost []StopPost
}

type Derailer struct {
	XMLName     xml.Name `xml:"derailer"`
	Id          string   `xml:"id,attr"`
	Code        string   `xml:"code,attr,omitempty"`
	Name        string   `xml:"name,attr,omitempty"`
	Description string   `xml:"description,attr,omitempty"`
	Pos         string   `xml:"pos,attr,omitempty"`
	AbsPos      string   `xml:"absPos,attr,omitempty"`
	Dir         string   `xml:"dir,attr,omitempty"`
	RuleCode    string   `xml:"ruleCode,attr,omitempty"`
	DerailSide  string   `xml:"derailSide,attr,omitempty"`
	Kind        string   `xml:"kind,attr,omitempty"`
	Model       string   `xml:"model,attr,omitempty"`
}

type Derailers struct {
	XMLName  xml.Name `xml:"derailers"`
	Derailer []Derailer
}

type TrainRadioChange struct {
	XMLName              xml.Name `xml:"trainRadioChange"`
	Id                   string   `xml:"id,attr"`
	Code                 string   `xml:"code,attr,omitempty"`
	Name                 string   `xml:"name,attr,omitempty"`
	Description          string   `xml:"description,attr,omitempty"`
	Pos                  string   `xml:"pos,attr,omitempty"`
	AbsPos               string   `xml:"absPos,attr,omitempty"`
	Dir                  string   `xml:"dir,attr,omitempty"`
	RadioSystem          string   `xml:"radioSystem,attr,omitempty"`
	NetworkSelection     string   `xml:"networkSelection,attr,omitempty"`
	PublicEmergency      string   `xml:"publicEmergency,attr,omitempty"`
	BroadcastCalls       string   `xml:"broadcastCalls,attr,omitempty"`
	TextMessageService   string   `xml:"textMessageService,attr,omitempty"`
	DirectMode           string   `xml:"directMode,attr,omitempty"`
	PublicNetworkRoaming string   `xml:"publicNetworkRoaming,attr,omitempty"`
}

type TrainRadioChanges struct {
	XMLName          xml.Name `xml:"trainRadioChanges"`
	TrainRadioChange []TrainRadioChange
}

type OcsElements struct {
	XMLName                 xml.Name `xml:"ocsElements"`
	Signals                 *Signals
	Balises                 *Balises
	TrainProtectionElements *TrainProtectionElements
	StopPosts               *StopPosts
	Derailers               *Derailers
	TrainRadioChanges       *TrainRadioChanges
}

type Track struct {
	XMLName       xml.Name `xml:"track"`
	Id            string   `xml:"id,attr"`
	TrackTopology TrackTopology
	TrackElements TrackElements
	OcsElements   OcsElements
}

type Tracks struct {
	XMLName xml.Name `xml:"tracks"`
	Tracks  []Track
}

type AxleWeight struct {
	XMLName   xml.Name `xml:"axleWeight"`
	Value     string   `xml:"value,attr"`
	Meterload string   `xml:"meterload,attr"`
}

type Electrification struct {
	XMLName   xml.Name `xml:"electrification"`
	Type      string   `xml:"type,attr"`
	Voltage   string   `xml:"voltage,attr"`
	Frequency string   `xml:"frequency,attr"`
}

type EpsgCode struct {
	XMLName     xml.Name `xml:"epsgCode"`
	Default     string   `xml:"default,attr"`
	ExtraHeight string   `xml:"extraHeight,attr"`
}

type Gauge struct {
	XMLName xml.Name `xml:"gauge"`
	Value   string   `xml:"value,attr"`
}

type ClearanceGauge struct {
	XMLName xml.Name `xml:"clearanceGauge"`
	Code    string   `xml:"code,attr"`
}

type OperationMode struct {
	XMLName           xml.Name `xml:"operationMode"`
	ModeLegislative   string   `xml:"modeLegislative,attr"`
	ModeExecutive     string   `xml:"modeExecutive,attr"`
	ClearanceManaging string   `xml:"clereanceManaging,attr"`
}

type Owner struct {
	XMLName                  xml.Name `xml:"owner"`
	OwnerName                string   `xml:"ownerName,attr"`
	InfrastructureManagerRef string   `xml:"infrastructureManagerRef,attr"`
}

type PowerTransmission struct {
	XMLName xml.Name `xml:"powerTransmission"`
	Type    string   `xml:"type,attr"`
	Style   string   `xml:"style,attr"`
}

type Speed struct {
	XMLName           xml.Name `xml:"speed"`
	TrainCategory     string   `xml:"trainCategory,attr"`
	EtcsTrainCategory string   `xml:"etcsTrainCategory,attr"`
	ProfileRef        string   `xml:"profileRef,attr"`
	Status            string   `xml:"status,attr"`
	VMax              string   `xml:"vMax,attr"`
}

type Speeds struct {
	XMLName xml.Name `xml:"speeds"`
	Speed   []Speed
}

type TrainRadio struct {
	XMLName              xml.Name `xml:"trainRadio"`
	RadioSystem          string   `xml:"radioSystem,attr"`
	NetworkSelection     string   `xml:"networkSelection,attr"`
	PublicEmergency      string   `xml:"publicEmergency,attr"`
	BroadcastCalls       string   `xml:"broadcastCalls,attr"`
	TextMessageService   string   `xml:"textMessageService,attr"`
	DirectMode           string   `xml:"directMode,attr"`
	PublicNetworkRoaming string   `xml:"publicNetworkRoaming,attr"`
}

type TrainProtection struct {
	XMLName    xml.Name `xml:"trainProtection"`
	Medium     string   `xml:"medium,attr"`
	Monitoring string   `xml:"monitoring,attr"`
}

type InfraAttributes struct {
	XMLName           xml.Name `xml:"infraAttributes"`
	Id                string   `xml:"id,attr"`
	Speeds            *Speeds
	AxleWeight        *AxleWeight
	Electrification   *Electrification
	EpsgCode          *EpsgCode
	Gauge             *Gauge
	ClearanceGauge    *ClearanceGauge
	OperationMode     *OperationMode
	Owner             *Owner
	PowerTransmission *PowerTransmission
	TrainRadio        *TrainRadio
	TrainProtection   *TrainProtection
}

type InfraAttrGroups struct {
	XMLName         xml.Name `xml:"infraAttrGroups"`
	InfraAttributes []InfraAttributes
}

type Infrastructure struct {
	XMLName         xml.Name `xml:"infrastructure"`
	Id              string   `xml:"id,attr"`
	Name            string   `xml:"name,attr,omitempty"`
	InfraAttrGroups []InfraAttrGroups
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
		v := ExportTrack(t.ID, lineId)
		ts.Tracks = append(ts.Tracks, v)

	}
	iag := ExportInfraAttrs(lineId)
	in := Infrastructure{
		Id:              lineId + "-" + time.Now().Format("20060102150405"), // UNIX timestamp format
		Name:            lineId,
		InfraAttrGroups: iag,
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
	Connection   neoism.Node         `json:"c"`
}

func ExportTrack(id string, line string) Track {
	db := config.GetDBConnection()
	query := "MATCH (l:Line {id:{lineId}})-[]-(t:Track {id:{trackId}})-[r:BEGINS|ENDS|HAS_TRACK_ELEMENT|HAS_OCS_ELEMENT|HAS_SWITCH|HAS_CROSSING]-(n)-[:HAS_CONNECTION*0..1]-(c) RETURN t,r,n,labels(n),c"
	track := []UnmarshalledTrack{}
	cq := neoism.CypherQuery{
		Statement:  query,
		Parameters: neoism.Props{"trackId": id, "lineId": line},
		Result:     &track,
	}

	e := db.Cypher(&cq)
	_ = e

	xtb := TrackEdge{XMLName: xml.Name{Local: "trackBegin"}}
	xte := TrackEdge{XMLName: xml.Name{Local: "trackEnd"}}
	xc := Connections{}
	xtel := TrackElements{}
	xoel := OcsElements{}

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
		case "HAS_OCS_ELEMENT":
			createOcsElement(lb, &xoel, t)
		case "HAS_SWITCH", "HAS_CROSSING":
			if t.Node.Data["id"] != t.Connection.Data["id"] { // weird behaviour query returns track-switch-connection and track-switch-switch - to investigate
				createConnection(lb, &xc, t)
			}
		}
	}
	xtt := TrackTopology{TrackBegin: xtb, TrackEnd: xte, Connections: xc}
	xt := Track{Id: id, TrackTopology: xtt, TrackElements: xtel, OcsElements: xoel}

	return xt
}

// <infraAttrGroups />
type UnmarshalledGroup struct {
	ID int `json:"ID(g)"`
}

// <infraAttributes />
type UnmarshalledInfraAttrs struct {
	AttrGroup neoism.Node `json:"a"`
}

// <speeds />, <gauge /> etc.
type UnmarshalledAttrsGroup struct {
	Node  neoism.Node `json:"n"`
	Label []string    `json:"labels(n)"`
	ID    int         `json:"ID(n)"`
}

// <speed />
type UnmarshalledSpeeds struct {
	Node neoism.Node `json:"s"`
}

// INFRA ATTR GROUPS
func ExportInfraAttrs(id string) []InfraAttrGroups {
	db := config.GetDBConnection()
	query := "MATCH (l:Line {id:{lineId}})-[r:HAS_ATTR_GROUP]-(g:InfraAttrGroup) RETURN ID(g)"
	ug := []UnmarshalledGroup{}
	cq := neoism.CypherQuery{
		Statement:  query,
		Parameters: neoism.Props{"lineId": id},
		Result:     &ug,
	}

	e := db.Cypher(&cq)
	_ = e

	xiags := []InfraAttrGroups{} // <infraAttrGroups /> - no idea how much times this section can occur, so it is an array
	for _, i := range ug {
		xiag := InfraAttrGroups{}                                                                                   // <infraAttrGroups />
		query2 := "MATCH (g:InfraAttrGroup)-[r:HAS_INFRA_ATTRS]-(a:InfraAttributes) WHERE ID(g)={groupId} RETURN a" // get all <infraAttributes/> from the given <infraAttrGroups/>
		uia := []UnmarshalledInfraAttrs{}
		cq2 := neoism.CypherQuery{
			Statement:  query2,
			Parameters: neoism.Props{"groupId": i.ID},
			Result:     &uia,
		}

		e := db.Cypher(&cq2)
		_ = e

		for _, ia := range uia {
			xia := &InfraAttributes{} // <infraAttributes/>

			iaid := ia.AttrGroup.Data["id"]
			query3 := "MATCH (g:InfraAttrGroup)-[]-(a:InfraAttributes {id:{iaid}})-[r:INFRA_ATTR]-(n) WHERE ID(g)={groupId}  RETURN n, ID(n), labels(n)" // get all groups from the given <infraAttributes/>
			uag := []UnmarshalledAttrsGroup{}
			cq3 := neoism.CypherQuery{
				Statement:  query3,
				Parameters: neoism.Props{"iaid": iaid, "groupId": i.ID},
				Result:     &uag,
			}
			e := db.Cypher(&cq3)
			_ = e

			for _, j := range uag { // iterate over array of groups
				switch j.Label[0] { // create group ex. <axleWeight/>, <gauge/> and only the <speeds/> has childs
				case "AxleWeight":
					aw := &AxleWeight{}
					xia.AxleWeight = createElementFromNode(&j.Node, aw).(*AxleWeight)
				case "Electrification":
					el := &Electrification{}
					xia.Electrification = createElementFromNode(&j.Node, el).(*Electrification)
				case "EpsgCode":
					ec := &EpsgCode{}
					xia.EpsgCode = createElementFromNode(&j.Node, ec).(*EpsgCode)
				case "Gauge":
					ga := &Gauge{}
					xia.Gauge = createElementFromNode(&j.Node, ga).(*Gauge)
				case "ClearanceGauge":
					cg := &ClearanceGauge{}
					xia.ClearanceGauge = createElementFromNode(&j.Node, cg).(*ClearanceGauge)
				case "OperationMode":
					om := &OperationMode{}
					xia.OperationMode = createElementFromNode(&j.Node, om).(*OperationMode)
				case "Owner":
					ow := &Owner{}
					xia.Owner = createElementFromNode(&j.Node, ow).(*Owner)
				case "PowerTransmission":
					pt := &PowerTransmission{}
					xia.PowerTransmission = createElementFromNode(&j.Node, pt).(*PowerTransmission)
				case "Speeds":
					ns := &Speeds{}                                                            // <speeds/>
					query4 := "MATCH (n:Speeds)-[r:HAS_SPEED]-(s) WHERE ID(n)={sid}  RETURN s" // grabs all <speed/> from the given <speeds/>
					us := []UnmarshalledSpeeds{}
					cq4 := neoism.CypherQuery{
						Statement:  query4,
						Parameters: neoism.Props{"sid": j.ID},
						Result:     &us,
					}
					e := db.Cypher(&cq4)
					_ = e
					for _, k := range us {
						ss := &Speed{} // <speed/>
						ns.Speed = append(
							ns.Speed,
							*createElementFromNode(&k.Node, ss).(*Speed),
						)
					}
					xia.Speeds = createElementFromNode(&j.Node, ns).(*Speeds)
				case "TrainRadio":
					tr := &TrainRadio{}
					xia.TrainRadio = createElementFromNode(&j.Node, tr).(*TrainRadio)
				case "TrainProtection":
					tp := &TrainProtection{}
					xia.TrainProtection = createElementFromNode(&j.Node, tp).(*TrainProtection)
				}
			}

			xiag.InfraAttributes = append(
				xiag.InfraAttributes,
				*createElementFromNode(&ia.AttrGroup, xia).(*InfraAttributes),
			)

		}
		xiags = append(xiags, xiag)
	}
	return xiags
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

// createConnection creates an XML connection node inside switch or crossing parent
func createConnection(lb string, xc *Connections, t UnmarshalledTrack) {
	nsc := &SwitchOrCrossing{}
	switch lb {
	case "Switch":
		nsc.XMLName = xml.Name{Local: "switch"}

	case "Crossing":
		nsc.XMLName = xml.Name{Local: "crossing"}
	}
	nco := &Connection{}
	nsc.Connection = *createElementFromNode(&t.Connection, nco).(*Connection)
	xc.Connections = append(
		xc.Connections,
		*createElementFromNode(&t.Node, nsc).(*SwitchOrCrossing),
	)
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

// OCS Elements
func createOcsElement(lb string, xoel *OcsElements, t UnmarshalledTrack) {
	switch lb {
	case "Signal":
		if xoel.Signals == nil {
			xoel.Signals = &Signals{}
		}
		ns := &Signal{}
		xoel.Signals.Signal = append(
			xoel.Signals.Signal,
			*createElementFromNode(&t.Node, ns).(*Signal),
		)
	case "Balise":
		if xoel.Balises == nil {
			xoel.Balises = &Balises{}
		}
		nb := &Balise{}
		xoel.Balises.Balise = append(
			xoel.Balises.Balise,
			*createElementFromNode(&t.Node, nb).(*Balise),
		)
	case "TrainProtectionElement":
		if xoel.TrainProtectionElements == nil {
			xoel.TrainProtectionElements = &TrainProtectionElements{}
		}
		ntpe := &TrainProtectionElement{}
		xoel.TrainProtectionElements.TrainProtectionElement = append(
			xoel.TrainProtectionElements.TrainProtectionElement,
			*createElementFromNode(&t.Node, ntpe).(*TrainProtectionElement),
		)
	case "StopPost":
		if xoel.StopPosts == nil {
			xoel.StopPosts = &StopPosts{}
		}
		nsp := &StopPost{}
		xoel.StopPosts.StopPost = append(
			xoel.StopPosts.StopPost,
			*createElementFromNode(&t.Node, nsp).(*StopPost),
		)
	case "Derailers":
		if xoel.Derailers == nil {
			xoel.Derailers = &Derailers{}
		}
		nd := &Derailer{}
		xoel.Derailers.Derailer = append(
			xoel.Derailers.Derailer,
			*createElementFromNode(&t.Node, nd).(*Derailer),
		)
	case "TrainRadioChange":
		if xoel.TrainRadioChanges == nil {
			xoel.TrainRadioChanges = &TrainRadioChanges{}
		}
		ntrc := &TrainRadioChange{}
		xoel.TrainRadioChanges.TrainRadioChange = append(
			xoel.TrainRadioChanges.TrainRadioChange,
			*createElementFromNode(&t.Node, ntrc).(*TrainRadioChange),
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
