package routes

import (
	"github.com/gorilla/mux"
	"github.com/philipokiokio/book-ms-api/pkg/controllers"
)

var RegisterBookStoreRouter = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}/", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}/update/", controllers.UpdateBookById).Methods("PUT")
	router.HandleFunc("/book/{bookId}/", controllers.DeleteBook).Methods("DELETE")

}
