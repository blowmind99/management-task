package main

import (
	"net/http"

	"github.com/blowmind99/go-crud/controllers/taskcontroller"
)

func main() {
	http.HandleFunc("/", taskcontroller.Index)
	http.HandleFunc("/task/get_form", taskcontroller.GetForm)
	http.HandleFunc("/task/store", taskcontroller.Store)

	http.ListenAndServe(":8000", nil)

}
