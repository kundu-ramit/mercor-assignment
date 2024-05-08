package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kundu-ramit/mercor_assignment/service"
)

type QueryController interface {
	ProcessNLPQuery(c *gin.Context)
	ProcessOrderedQuery(c *gin.Context)
}

type queryController struct {
	service service.QueryService
}

type NLPQueryBody struct {
	Query string `json:"query"`
	// Add other fields as needed
}

type OrderedQueryBody struct {
	Skills     []string `json:"skills"`
	Budget     int      `json:"budget"`
	Experience string   `json:"experience"`
	// Add other fields as needed
}

func NewQueryController() QueryController {
	return queryController{
		service: service.NewQueryService(),
	}
}

func (q queryController) ProcessNLPQuery(c *gin.Context) {
	var requestBody NLPQueryBody

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := q.service.Fetch(c, requestBody.Query)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}

func (q queryController) ProcessOrderedQuery(c *gin.Context) {
	var requestBody OrderedQueryBody

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := q.service.FetchOrdered(c, requestBody.Skills, requestBody.Budget, requestBody.Experience)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}
