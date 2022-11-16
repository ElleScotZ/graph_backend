package main

import (
	"encoding/json"
	"strconv"

	"github.com/ellescotz/graph_backend/pkg/core"
	"github.com/gin-gonic/gin"
)

const (
	malformedNodesErrorMessage string = "Nodes key in request header is malformed"
	malformedGraphErrorMessage string = "malformed graph"
)

// getPathsWithMaxSteps calls GeneratePathsWithMaxSteps.
// Header requirements:
// Initial and end nodes: "Nodes": "<node1><node2>" / for example: "Nodes": "AB"
// (Maximum) number of steps: "MaxEdges": "<a positive integer>"
// Exact or up to certain number of edges: "Exact": "<true/false>"
// In case of malformed graph or header file the function exits,
// and it gives an error response.
func getPathsWithMaxSteps(c *gin.Context) {
	var (
		graph                core.Graph
		endNodesString       string
		initialNode, endNode core.Node
		maxEdgesString       string
		maxEdges             int
		exactNumberString    string
		exactNumber          bool
		relevantPaths        []core.Path
	)

	// Decoding request body that contains the graph
	err := json.NewDecoder(c.Request.Body).Decode(&graph)
	if err != nil {
		c.JSON(500, gin.H{
			"error": malformedGraphErrorMessage,
		})
		return
	}

	// Identifying initial and end node from request header
	// Request header has to contain information about the initial and end nodes in the following way:
	// "Nodes": "AB"
	endNodesString = c.Query("Nodes")

	if len(endNodesString) != 2 {
		c.JSON(500, gin.H{
			"error": malformedNodesErrorMessage,
		})
		return
	}

	initialNode.Name = endNodesString[:1]
	endNode.Name = endNodesString[1:]

	// Identifying maximum number of steps (edges) from request header
	// Request header has to contain information in the following way:
	// "MaxEdges": "<a positive integer>"
	maxEdgesString = c.Query("MaxEdges")
	if len(maxEdgesString) == 0 {
		c.JSON(500, gin.H{
			"error": "wrong MaxEdges",
		})
		return
	}

	maxEdges, err = strconv.Atoi(maxEdgesString[:len(maxEdgesString)-1])
	if err != nil {
		c.JSON(500, gin.H{
			"error": "wrong MaxEdges",
		})
		return
	}

	// Identifying whether exact or maximum number of steps is required
	// Request header has to contain information in the following way:
	// "Exact": "<true/false>"
	exactNumberString = c.Query("Exact")

	exactNumber, err = strconv.ParseBool(exactNumberString)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "wrong Exact",
		})
		return
	}

	// Calculating relevant paths
	relevantPaths = graph.GeneratePathsWithMaxSteps(initialNode, endNode, maxEdges, exactNumber)

	// Binding relevantPaths with request
	c.JSON(200, relevantPaths)
}

// getPathsWithMaxWeight calls GeneratePathsWithMaxWeight.
// Header requirements:
// Initial and end nodes: "Nodes": "<node1><node2>" / for example: "Nodes": "AB"
// (Maximum) sum weight of a path: "MaxWeight": "<a positive floating point number>"
// Exact or up to certain sum weight: "Exact": "<true/false>"
// In case of malformed graph or header file the function exits,
// and it gives an error response.
func getPathsWithMaxWeight(c *gin.Context) {
	var (
		graph                core.Graph
		endNodesString       string
		initialNode, endNode core.Node
		maxWeightString      string
		maxWeight            float64
		exactWeightString    string
		exactWeight          bool
		relevantPaths        []core.Path
	)

	// Decoding request body that contains the graph
	err := json.NewDecoder(c.Request.Body).Decode(&graph)
	if err != nil {
		c.JSON(500, gin.H{
			"error": malformedGraphErrorMessage,
		})
		return
	}

	// Identifying initial and end node from request header
	// Request header has to contain information about the initial and end nodes in the following way:
	// "Nodes": "AB"
	endNodesString = c.Query("Nodes")

	if len(endNodesString) != 2 {
		c.JSON(500, gin.H{
			"error": malformedNodesErrorMessage,
		})
		return
	}

	initialNode.Name = endNodesString[:1]
	endNode.Name = endNodesString[1:]

	// Identifying maximum sum weight of a path from request header
	// Request header has to contain information in the following way:
	// "MaxWeight": "<a positive floating point number>T"
	maxWeightString = c.Query("MaxWeight")
	if len(maxWeightString) == 0 {
		c.JSON(500, gin.H{
			"error": "wrong MaxWeight length",
		})
		return
	}

	maxWeight, err = strconv.ParseFloat(maxWeightString[:len(maxWeightString)-1], 64)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "wrong MaxWeight float",
		})
		return
	}

	// Identifying whether exact or maximum sum weight of path is required
	// Request header has to contain information in the following way:
	// "Exact": "<true/false>"
	exactWeightString = c.Query("Exact")

	exactWeight, err = strconv.ParseBool(exactWeightString)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "wrong Exact",
		})
		return
	}

	// Calculating relevant paths
	relevantPaths = graph.GeneratePathsWithMaxWeight(initialNode, endNode, maxWeight, exactWeight)

	// Binding relevantPaths with request
	c.JSON(200, relevantPaths)
}

// getLowestHighestWeightPath calls GenerateLowestHighestWeightPath.
// Header requirements:
// Initial and end nodes: "Nodes": "<node1><node2>" / for example: "Nodes": "AB"
// Lowest or highest weighted path: "Lowest": "<true/false>"
// In case of malformed graph or header file the function exits,
// and it gives an error response.
func getLowestHighestWeightPath(c *gin.Context) {
	var (
		graph                core.Graph
		endNodesString       string
		initialNode, endNode core.Node
		lowestString         string
		lowest               bool
		relevantPaths        []core.Path
	)

	// Decoding request body that contains the graph
	err := json.NewDecoder(c.Request.Body).Decode(&graph)
	if err != nil {
		c.JSON(500, gin.H{
			"error": malformedGraphErrorMessage,
		})
		return
	}

	// Identifying initial and end node from request header
	// Request header has to contain information about the initial and end nodes in the following way:
	// "Nodes": "AB"
	endNodesString = c.Query("Nodes")

	if len(endNodesString) != 2 {
		c.JSON(500, gin.H{
			"error": malformedNodesErrorMessage,
		})
		return
	}

	initialNode.Name = endNodesString[:1]
	endNode.Name = endNodesString[1:]

	// Identifying maximum or minimum sum weight of a path from request header
	// Request header has to contain information in the following way:
	// "Lowest": "true"
	// true, if calculation is for lowest weighted path.
	// false if calculation is for highest weighted path.
	lowestString = c.Query("Lowest")
	if len(lowestString) == 0 {
		c.JSON(500, gin.H{
			"error": "wrong Lowest",
		})
		return
	}

	lowest, err = strconv.ParseBool(lowestString)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "wrong Lowest",
		})
		return
	}

	// Calculating relevant paths
	relevantPaths = graph.GenerateLowestHighestWeightPath(initialNode, endNode, lowest)

	// Binding relevantPaths with request
	c.JSON(200, relevantPaths)
}

// getShortestLongestPath calls GenerateShortestLongestPath.
// Header requirements:
// Initial and end nodes: "Nodes": "<node1><node2>" / for example: "Nodes": "AB"
// Shortest or longest path: "Shortest": "<true/false>"
// In case of malformed graph or header file the function exits,
// and it gives an error response.
func getShortestLongestPath(c *gin.Context) {
	var (
		graph                core.Graph
		endNodesString       string
		initialNode, endNode core.Node
		shortestString       string
		shortest             bool
		relevantPaths        []core.Path
	)

	// Decoding request body that contains the graph
	err := json.NewDecoder(c.Request.Body).Decode(&graph)
	if err != nil {
		c.JSON(500, gin.H{
			"error": malformedGraphErrorMessage,
		})
		return
	}

	// Identifying initial and end node from request header
	// Request header has to contain information about the initial and end nodes in the following way:
	// "Nodes": "AB"
	endNodesString = c.Query("Nodes")

	if len(endNodesString) != 2 {
		c.JSON(500, gin.H{
			"error": malformedNodesErrorMessage,
		})
		return
	}

	initialNode.Name = endNodesString[:1]
	endNode.Name = endNodesString[1:]

	// Identifying shortest or longest path from request header
	// Request header has to contain information in the following way:
	// "Shortest": "true"
	// true, if calculation is for shortest path.
	// false if calculation is for longest path.
	shortestString = c.Query("Shortest")
	if len(shortestString) == 0 {
		c.JSON(500, gin.H{
			"error": "wrong Shortest",
		})
		return
	}

	shortest, err = strconv.ParseBool(shortestString)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "wrong Shortest",
		})
		return
	}

	// Calculating relevant paths
	relevantPaths = graph.GenerateShortestLongestPath(initialNode, endNode, shortest)

	// Binding relevantPaths with request
	c.JSON(200, relevantPaths)
}

// cors enables middleware handling.
// It connects frontend to backend with certain settings.
// Backend application is allowed to be reached from the origin defined here.
func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://34.76.180.95:3000")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	router := gin.Default()

	router.SetTrustedProxies([]string{"http://34.76.180.95"})

	router.Use(cors())

	// Generating all the paths with <= #steps
	router.POST("/maxSteps", getPathsWithMaxSteps)

	// Generating all the paths with <= weight
	router.POST("/maxWeight", getPathsWithMaxWeight)

	// Generating lowest/highest weighted paths
	router.POST("/highLowWeight", getLowestHighestWeightPath)

	// Generating shoertest/longest paths
	router.POST("/shortLong", getShortestLongestPath)

	router.Run("http://34.78.47.255:8080")
}
