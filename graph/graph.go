package graph

import (
	"reflect"
	"strings"

	"Go-GoSAFE.converter/utils"

	"github.com/beevik/etree"
	"github.com/jmcvetta/neoism"
)

type GraphUtils struct{}

func (g *GraphUtils) TrackToGraph(t *etree.Element, db *neoism.Database, epsg string, ln *neoism.Node) string {

	elementsUtils := utils.ElementsUtils{}
	// TRACK
	tr := elementsUtils.GetTrackProperties(t, epsg)
	tn, _ := db.CreateNode(tr)
	tn.AddLabel("Track")

	ln.Relate("HAS_TRACK", tn.Id(), neoism.Props{})
	// TRACK TOPOLOGIES
	tt := elementsUtils.GetTrackTopologies(t, epsg)

	tb, _ := db.CreateNode(tt.Begin.Properties)
	tb.AddLabel(tt.Begin.Label)
	tn.Relate("BEGINS", tb.Id(), neoism.Props{})

	te, _ := db.CreateNode(tt.End.Properties)
	te.AddLabel(tt.End.Label)
	tn.Relate("ENDS", te.Id(), neoism.Props{})

	for _, sw := range tt.Switch {
		swn, _ := db.CreateNode(sw.Properties)
		con, _ := db.CreateNode(sw.Connection)
		swn.AddLabel("Switch")
		con.AddLabel("Connection")
		tn.Relate("HAS_SWITCH", swn.Id(), neoism.Props{})
		swn.Relate("HAS_CONNECTION", con.Id(), neoism.Props{})
	}

	for _, cr := range tt.Crossing {
		crn, _ := db.CreateNode(cr.Properties)
		con, _ := db.CreateNode(cr.Connection)
		crn.AddLabel("Crossing")
		con.AddLabel("Connection")
		tn.Relate("HAS_CROSSING", crn.Id(), neoism.Props{})
		crn.Relate("HAS_CONNECTION", con.Id(), neoism.Props{})
	}
	// TRACK ELEMENTS
	tre := elementsUtils.GetTrackElements(t, epsg)

	st := reflect.ValueOf(&tre).Elem()
	for i := 0; i < st.NumField(); i++ {
		a := st.Field(i).Interface().([]neoism.Props)
		lb := strings.TrimSuffix(st.Type().Field(i).Name, "s")

		for _, e := range a {
			en, _ := db.CreateNode(e)
			en.AddLabel(lb)
			tn.Relate("HAS_TRACK_ELEMENT", en.Id(), neoism.Props{})
		}
	}

	// OCS ELEMENTS
	oe := elementsUtils.GetOCSElements(t, epsg)

	so := reflect.ValueOf(&oe).Elem()
	for i := 0; i < so.NumField(); i++ {
		a := so.Field(i).Interface().([]neoism.Props)
		lb := strings.TrimSuffix(so.Type().Field(i).Name, "s")

		for _, e := range a {
			on, _ := db.CreateNode(e)
			on.AddLabel(lb)
			tn.Relate("HAS_OCS_ELEMENT", on.Id(), neoism.Props{})
		}
	}

	return "ok"
}

func (g *GraphUtils) InfraAttributesToGraph(a *etree.Element, db *neoism.Database, ag *neoism.Node) string {
	elementsUtils := utils.ElementsUtils{}

	elements := a.SelectElements("infraAttributes")
	if elements == nil {
		return "no attributes"
	}
	for _, element := range elements {
		ia := elementsUtils.GetInfraAttributes(element)
		at, _ := db.CreateNode(neoism.Props{"id": ia.ID})
		at.AddLabel("InfraAttributes")
		ag.Relate("HAS_INFRA_ATTRS", at.Id(), neoism.Props{})

		sa := reflect.ValueOf(&ia).Elem()
		for i := 0; i < sa.NumField(); i++ {
			name := sa.Type().Field(i).Name
			if name == "ID" {
				continue
			}
			if name == "Speeds" {
				a := sa.Field(i).Interface().([]neoism.Props)
				if len(a) == 0 {
					continue
				}
				nss, _ := db.CreateNode(neoism.Props{})
				nss.AddLabel(name)
				at.Relate("INFRA_ATTR", nss.Id(), neoism.Props{})
				for _, e := range a {
					ns, _ := db.CreateNode(e)
					ns.AddLabel("Speed")
					nss.Relate("HAS_SPEED", ns.Id(), neoism.Props{})
				}
				continue
			}

			p := sa.Field(i).Interface().(neoism.Props)
			if len(p) == 0 {
				continue
			}
			a, _ := db.CreateNode(p)
			a.AddLabel(name)
			at.Relate("INFRA_ATTR", a.Id(), neoism.Props{})
		}

	}
	return "ok"
}
