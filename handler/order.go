package handler

import (
	"fmt"
	"net/http"
	"time"
)

type Order struct {
}

func (order *Order) Create(w http.ResponseWriter, r *http.Request) {
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
