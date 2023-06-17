package data

import "goAnimals/structs"

// type Animal struct {
// 	Name    string `json:"name"`
// 	Country string `json:"country"`
// 	Breed   string `json:"breed"`
// 	Age     string `json:"age"`
// }

var Animals = []structs.Animal{
	{Name: "Garfield", Country: "UK", Breed: "Orange Tabby Cat", Age: "6 y.o."},
	{Name: "Scooby Doo", Country: "USA", Breed: "Great Dane", Age: "7 y.o."},
}

var Cat = structs.Animal{Name: "Garfield", 
Country: "UK", 
Breed: "Orange Tabby Cat", 
Age: "6 y.o."}

var Dog = structs.Animal{Name: "Scooby Doo", 
Country: "USA", 
Breed: "Great Dane", 
Age: "7 y.o."}

// var cats = new(Animal)

// cats.Name = "Garfield"
// cats.Country = "UK"
// cats.Breed = "Siamese cat"
// cas.Age = "2 Month"




