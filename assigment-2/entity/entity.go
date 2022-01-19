package entity

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Items struct {
	ItemId      int    `json:"itemId`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderId     int
}

type Orders struct {
	OrderId      int
	OrderedAt    time.Time
	CustomerName string  `json:"customerName"`
	Items        []Items `json:"items"`
}

func (o *Orders) GetOrdersIdParams(c *gin.Context) (int, error) {
	paramId := c.Param("ordersId")

	// fmt.Println("ini Id orders =>", paramId)

	ordersId, err := strconv.Atoi(paramId)

	if err != nil {
		return 0, err
	}

	return ordersId, nil
}

func (i *Items) GetItemIdParams(c *gin.Context) (int, error) {
	paramId := c.Param("itemId")

	// fmt.Println("ini Id orders =>", paramId)

	itemId, err := strconv.Atoi(paramId)

	if err != nil {
		return 0, err
	}

	return itemId, nil
}
