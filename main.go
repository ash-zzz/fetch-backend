package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type item struct {
	Description string `json:"shortDescription"`
	Price       string `json:"price"`
}

type receipt struct {
	ID           uuid.UUID `json:"id"`
	Points       int       `json:"points"`
	Retailer     string    `json:"retailer"`
	PurchaseDate string    `json:"purchaseDate"`
	PurchaseTime string    `json:"purchaseTime"`
	Items        []item    `json:"items"`
	Total        string    `json:"total"`
}

type proccess_response struct {
	ID uuid.UUID `json:"id"`
}

type points_response struct {
	Points int `json:"points"`
}

var receipts = []receipt{}

func getReceipts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, receipts)
}

func processReceipt(c *gin.Context) {
	var data receipt

	if err := c.BindJSON(&data); err != nil {
		return
	}

	var id = uuid.New()
	data.ID = id
	data.Points = calculateAllPoints(data)

	receipts = append(receipts, data)
	c.IndentedJSON(http.StatusCreated, proccess_response{ID: id})
}

func countPoints(c *gin.Context) {
	id := c.Param("id")

	for _, r := range receipts {
		if r.ID.String() == id {
			points := r.Points
			c.IndentedJSON(http.StatusOK, points_response{Points: points})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, map[string]string{"error": "provided id does not exist"})

}

func main() {
	router := gin.Default()
	router.GET("/receipts", getReceipts)

	router.POST("/receipts/process", processReceipt)

	router.GET("/receipts/:id/points", countPoints)

	router.Run()
}
