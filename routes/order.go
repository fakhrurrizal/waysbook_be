package routes

import (
	"waysbuck/handlers"
	"waysbuck/pkg/middleware"
	"waysbuck/pkg/mysql"
	"waysbuck/repositories"

	"github.com/gorilla/mux"
)

func OrderRoutes(r *mux.Router) {
	orderRepository := repositories.RepositoryOrder(mysql.DB)
	h := handlers.HandlerOrder(orderRepository)
	r.HandleFunc("/order", middleware.Auth(h.AddOrder)).Methods("POST")
	r.HandleFunc("/orders", middleware.Auth(h.FindOrders)).Methods("GET")
	r.HandleFunc("/order/{id}", middleware.Auth(h.DeleteOrder)).Methods("DELETE")
	r.HandleFunc("/order/{id}", middleware.Auth(h.GetOrder)).Methods("GET")
	r.HandleFunc("/order/{id}", middleware.Auth(h.UpdateOrder)).Methods("PATCH")
	r.HandleFunc("/orders-id", middleware.Auth(h.GetOrdersByID)).Methods("GET")
}