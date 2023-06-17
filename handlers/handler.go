package handlers

import (
	"encoding/json"
	"goAnimals/data"
	"net/http"
)




func GetRequest(w http.ResponseWriter, r *http.Request){
	
	// fmt.Fprintf(w, "hello\n")
	
	json.NewEncoder(w).Encode(data.Cat)
	json.NewEncoder(w).Encode(data.Dog)

}

// func postRequest(){

// }