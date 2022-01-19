package repository

import (
	"assignment-7/db"
	"assignment-7/entity"
	"fmt"
	"log"
)

type RepositoryInterface interface {
	Create(*entity.Orders) (*entity.Orders, error)
	GetById(int) (*entity.Orders, error)
	GetAll() (*[]entity.Orders, error)
	Update(*entity.Orders) (*entity.Orders, error)
	Delete(int)
}

type repository struct{}

var Repo RepositoryInterface = &repository{}

// ini belum
func (repo *repository) Create(orderReq *entity.Orders) (*entity.Orders, error) {

	db := db.NewDB()

	sqlOrders := "INSERT INTO orders (customer_name, ordered_at) VALUES ($1, $2) RETURNING order_id"
	rowOrders := db.QueryRow(sqlOrders, orderReq.CustomerName, orderReq.OrderedAt)

	var orderId = 0
	err := rowOrders.Scan(&orderId)

	if err != nil {
		log.Fatal(err.Error())
	}

	items := []entity.Items{}
	sqlItems := "INSERT INTO items (item_code, description, quantity, order_id) VALUES ($1, $2, $3, $4) RETURNING item_id, item_code, description, quantity, order_id"
	for i := 0; i < len(orderReq.Items); i++ {
		itemReq := orderReq.Items[i]
		rowItems := db.QueryRow(sqlItems, itemReq.ItemCode, itemReq.Description, itemReq.Quantity, orderId)
		item := entity.Items{}
		err = rowItems.Scan(&item.ItemId, &item.ItemCode, &item.Description, &item.Quantity, &item.OrderId)
		items = append(items, item)
	}

	orderReq.Items = items
	orderReq.OrderId = orderId

	return orderReq, nil
}

func (repo *repository) Update(orderReq *entity.Orders) (*entity.Orders, error) {
	db := db.NewDB()

	sqlOrders := "UPDATE orders SET customer_name = $2 WHERE order_id = $1 RETURNING order_id, customer_name, ordered_at"
	rowOrders := db.QueryRow(sqlOrders, orderReq.OrderId, orderReq.CustomerName)

	var orders entity.Orders
	err := rowOrders.Scan(&orders.OrderId, &orders.CustomerName, &orders.OrderedAt)

	if err != nil {
		log.Fatal(err.Error())
	}

	sqlItems := "UPDATE items SET item_code=$1, description=$2, quantity=$3 WHERE item_id = $4 RETURNING item_id, item_code, description, quantity, order_id"
	itemReq := orderReq.Items[0]
	rowItems := db.QueryRow(sqlItems, itemReq.ItemCode, itemReq.Description, itemReq.Quantity, itemReq.ItemId)

	item := entity.Items{}
	err = rowItems.Scan(&item.ItemId, &item.ItemCode, &item.Description, &item.Quantity, &item.OrderId)

	if err != nil {
		log.Fatal(err)
	}

	// var item entity.Items
	// err = rowItems.Scan(&item.ItemId, &item.ItemCode, &item.Description, &item.Quantity, &item.OrderId)

	orderReq.Items = []entity.Items{item}
	fmt.Println(itemReq)
	// orders.OrderId = orders.OrderId
	return orderReq, nil
}

func (repo *repository) GetById(orderId int) (*entity.Orders, error) {
	db := db.NewDB()

	var orders entity.Orders

	sqlOrders := "SELECT order_id, customer_name, ordered_at from orders WHERE order_id = $1"
	rowOrders := db.QueryRow(sqlOrders, orderId)

	err := rowOrders.Scan(&orders.OrderId, &orders.CustomerName, &orders.OrderedAt)

	if err != nil {
		log.Fatal(err.Error())
	}

	sqlItems := "SELECT item_id, item_code, description, quantity, order_id FROM items WHERE order_id = $1"
	rowItems, _ := db.Query(sqlItems, orderId)

	items := []entity.Items{}

	for rowItems.Next() {
		item := entity.Items{}
		err = rowItems.Scan(&item.ItemId, &item.ItemCode, &item.Description, &item.Quantity, &item.OrderId)
		if err != nil {
			log.Fatal(err)
		}

		items = append(items, item)
	}

	orders.Items = items
	orders.OrderId = orderId

	return &orders, nil
}

func (repo *repository) GetAll() (*[]entity.Orders, error) {
	db := db.NewDB()

	var orders []entity.Orders
	sqlOrders := "SELECT order_id, customer_name, ordered_at from orders"
	rowOrders, err := db.Query(sqlOrders)

	if err != nil {
		log.Fatal(err)
	}

	for rowOrders.Next() {
		order := entity.Orders{}
		err = rowOrders.Scan(&order.OrderId, &order.CustomerName, &order.OrderedAt)
		if err != nil {
			log.Fatal(err)
		}

		sqlItems := "SELECT item_id, item_code, description, quantity, order_id FROM items WHERE order_id = $1"
		rowItems, _ := db.Query(sqlItems, order.OrderId)

		items := []entity.Items{}
		for rowItems.Next() {
			item := entity.Items{}
			err = rowItems.Scan(&item.ItemId, &item.ItemCode, &item.Description, &item.Quantity, &item.OrderId)
			if err != nil {
				log.Fatal(err)
			}

			items = append(items, item)
		}

		order.Items = items

		orders = append(orders, order)
	}

	// orders.OrderId = orderId

	return &orders, nil
}

func (repo *repository) Delete(ordersId int) {
	db := db.NewDB()

	sql := "DELETE FROM items WHERE order_id = $1"
	// row := db.QueryRow(sql, ordersId)

	db.Exec(sql, ordersId)
}
