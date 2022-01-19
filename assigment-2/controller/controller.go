package controller

import (
	"assignment-7/entity"
	"assignment-7/service"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// coba ubah
func OrdersCreate(c *gin.Context) {
	var data map[string]interface{}

	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "invalid json body")
		return
	}

	var createdAt string
	_ = createdAt
	if val, ok := data["orderedAt"].(string); ok {
		createdAt = val
	}

	t, err := time.Parse("2006-01-02", createdAt)
	if err != nil {
		fmt.Println(err)
	}

	items := []entity.Items{}
	for i := 0; i < len(data["items"].([]interface{})); i++ {
		item := data["items"].([]interface{})[i].(map[string]interface{})

		itemReq := entity.Items{

			ItemCode:    item["itemCode"].(string),
			Description: item["description"].(string),
			Quantity:    int(item["quantity"].(float64)),
		}
		items = append(items, itemReq)
	}

	req := entity.Orders{
		CustomerName: data["customerName"].(string),
		OrderedAt:    t,
		Items:        items,
	}

	ordersData, err := service.Serv.Create(&req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	webResponse := entity.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   ordersData,
	}

	c.JSON(http.StatusCreated, webResponse)
}

func GetById(c *gin.Context) {
	var orders entity.Orders

	ordersId, _ := orders.GetOrdersIdParams(c)

	res, err := service.Serv.GetById(ordersId)
	if err != nil {
		webResponse := entity.WebResponse{
			Code:   500,
			Status: "bad request",
			Data:   nil,
		}
		c.JSON(http.StatusBadRequest, webResponse)
		return
	}

	webResponse := entity.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   res,
	}

	c.JSON(http.StatusOK, webResponse)
}

func GetAll(c *gin.Context) {
	data, err := service.Serv.GetAll()

	if err != nil {
		c.JSON(500, err)
	}

	webResponse := entity.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   data,
	}

	c.JSON(http.StatusOK, webResponse)
}

func Update(c *gin.Context) {
	var orders entity.Orders
	var item entity.Items

	OrderId, err := orders.GetOrdersIdParams(c)

	if err != nil {
		webResponse := entity.WebResponse{
			Code:   500,
			Status: "bad request",
			Data:   nil,
		}
		c.JSON(http.StatusBadRequest, webResponse)
		return
	}

	itemId, err := item.GetItemIdParams(c)

	if err != nil {
		webResponse := entity.WebResponse{
			Code:   500,
			Status: "bad request",
			Data:   nil,
		}
		c.JSON(http.StatusBadRequest, webResponse)
		return
	}

	if err := c.ShouldBindJSON(&orders); err != nil {
		webResponse := entity.WebResponse{
			Code:   500,
			Status: "eror when should bind json",
			Data:   nil,
		}
		c.JSON(http.StatusBadRequest, webResponse)
		return
	}

	orders.OrderId = OrderId

	item.ItemId = itemId

	orders.Items[0].ItemId = item.ItemId

	res, err := service.Serv.Update(&orders)

	if err != nil {
		webResponse := entity.WebResponse{
			Code:   500,
			Status: "eror when access method service update",
			Data:   nil,
		}
		c.JSON(http.StatusBadRequest, webResponse)
		return
	}

	webResponse := entity.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   res,
	}

	c.JSON(http.StatusOK, webResponse)
}

func Delete(c *gin.Context) {
	var orders entity.Orders
	ordersId, _ := orders.GetOrdersIdParams(c)

	fmt.Println("Coba hapus")
	service.Serv.Delete(ordersId)

}
