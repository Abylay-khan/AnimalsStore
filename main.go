package main

import (
	// "fmt"
	// "time"
	// "errors"
	"net/http"
	// "os"
	// "time"
	"goAnimals/handlers"
)

// func hello(w http.ResponseWriter, req *http.Request){

// 	fmt.Fprintf(w, "hello\n")
// }

// func headers(w http.ResponseWriter, req *http.Request){

// 	for name, headers := range req.Header{
// 		for _, h := range headers {
// 			fmt.Fprintf(w, "%v: %v\n", name, h)
// 		}
// 	}
// }

func main(){

	http.HandleFunc("/animals", handlers.GetRequest)

	http.ListenAndServe(":8080", nil)
} 

// type Animal struct {
// 	Name    string `json:"name"`
// 	Country string `json:"country"`
// 	Breed   string `json:"breed"`
// 	Age     string `json:"age"`
// }

// func main(){


// 	Arr := []Animal{
// 		{Name: "Garfield", Country: "UK", Breed: "Orange Tabby Cat", Age: "6 y.o."},
// 		{Name: "Scooby Doo", Country: "USA", Breed: "Great Dane", Age: "7 y.o."},
// 	}



// 	fmt.Println(Arr)

// }
