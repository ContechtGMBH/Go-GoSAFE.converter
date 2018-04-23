package controllers

import (
	"io/ioutil"
	"strconv"

	"Go-GoSAFE.converter/config"
	"Go-GoSAFE.converter/export"
	"Go-GoSAFE.converter/graph"

	"github.com/beevik/etree"
	"github.com/gin-gonic/gin"
	"github.com/jmcvetta/neoism"
)

/**
* @api {POST} /api/v1/converter/railml
* @apiDescription Converts a RailML file to neo4j graph
* @apiGroup Railml
* @apiName ConvertRailml
* @apiParam {string} line A line name
* @apiParam {string} epsg CRS EPSG number, should match geooCoords CRS
* @apiParam {file} file A valid RailML file that contains the Infrastructure subschema
* @apiSuccess (200) {json} object Response message with the number of extracted tracks
 */
func ConvertRailml(c *gin.Context) {

	epsg := c.PostForm("epsg")
	lineName := c.PostForm("line")
	file, _ := c.FormFile("file")
	xmlFile, _ := file.Open()
	byteValue, _ := ioutil.ReadAll(xmlFile)
	doc := etree.NewDocument()
	if err := doc.ReadFromBytes(byteValue); err != nil {
		panic(err)
	}

	graphUtils := graph.GraphUtils{}
	db := config.GetDBConnection()

	ln, _ := db.CreateNode(neoism.Props{"id": lineName})
	ln.AddLabel("Line")

	counter := 0

	for _, t := range doc.FindElements("//track") {

		st := graphUtils.TrackToGraph(t, db, epsg, ln)
		_ = st

		counter++
	}

	for _, a := range doc.FindElements("//infraAttrGroups") {
		ag, _ := db.CreateNode(neoism.Props{})
		ag.AddLabel("InfraAttrGroup")
		ln.Relate("HAS_ATTR_GROUP", ag.Id(), neoism.Props{})
		st := graphUtils.InfraAttributesToGraph(a, db, ag)
		_ = st
	}
	// Create relationships between connection nodes

	connect := `MATCH ()-[:BEGINS|ENDS]-(s:Connection),(e:Connection) WHERE s.id=e.ref AND not ((s)--(e)) MERGE (s)-[r:CONNECTS]->(e)`
	cq := neoism.CypherQuery{Statement: connect}

	e := db.Cypher(&cq)
	_ = e

	x := map[string]string{"status": "ok", "number of tracks": strconv.Itoa(counter)}

	c.JSON(200, gin.H{
		"response": x,
	})
}

/**
* @api {POST} /api/v1/export/railml
* @apiDescription Export a subgraph to the railml file
* @apiGroup Railml
* @apiName ExportRailml
* @apiParam {string} line A line name
* @apiSuccess (200) {XML} RailML A valid RailML document that describes the given line
 */
func ExportRailml(c *gin.Context) {

	lineId := c.PostForm("line")

	rm := export.ExportLine(lineId)
	/*
		output, err := xml.MarshalIndent(rm, "  ", "    ")
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}

		os.Stdout.Write(output)
	*/

	c.XML(200, rm)
}
