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

		st := graphUtils.RailmlToGraph(t, db, epsg, ln)
		_ = st

		counter++
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
