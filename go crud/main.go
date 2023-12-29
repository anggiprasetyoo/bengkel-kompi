package main

import (
	"net/http"

	"github.com/jeypc/go-crud/controllers/customercontroller"
)

func main() {

	http.HandleFunc("/", customercontroller.Index)
	http.HandleFunc("/customer", customercontroller.Index)
	http.HandleFunc("/customer/index", customercontroller.Index)
	http.HandleFunc("/customer/add", customercontroller.Add)
	http.HandleFunc("/customer/edit", customercontroller.Edit)
	http.HandleFunc("/customer/delete", customercontroller.Delete)

	http.ListenAndServe(":3000", nil)
}
