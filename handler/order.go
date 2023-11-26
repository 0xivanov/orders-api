package handler

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/0xivanov/orders-api/model"
	"github.com/0xivanov/orders-api/repository"
	"github.com/google/uuid"
)

type Order struct {
	Repo *repository.RedisRepo
}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	var body struct {
		CustomerId uuid.UUID
		Items      []model.Item
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	order := model.Order{
		OrderId:    rand.Uint64(),
		CustomerId: body.CustomerId,
		Items:      body.Items,
		Status:     model.Processing,
	}

	err := o.Repo.Insert(r.Context(), &order)
	if err != nil {
		fmt.Println("failed to insert: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(order)
	if err != nil {
		fmt.Println("failed to marshal: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(res)
	w.WriteHeader(http.StatusCreated)
	fmt.Println("create was called")
}
func (order *Order) List(w http.ResponseWriter, r *http.Request) {
	time.Sleep(3 * time.Second)
	fmt.Println("list was called")
}
func (order *Order) GetById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get was called")
}
func (order *Order) UpdateById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update was called")
}
func (order *Order) DeleteById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete was called")
}
