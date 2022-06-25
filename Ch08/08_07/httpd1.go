package main

import(
	"net/http"
	"encoding/json"
	"log"
	"fmt"
)
type Entry struct{
	Key string `json:"key"`
	Value any `json:"value"`
}
var database=map[string]interface{}{}
func postRequestHandler(w http.ResponseWriter,r *http.Request){
	entry:=&Entry{}
	defer r.Body.Close()
	dec:=json.NewDecoder(r.Body)
	if err:=dec.Decode(entry);err != nil {
		http.Error(w,err.Error(),http.StatusBadRequest)	
		return 
	}

	database[entry.Key]=entry.Value;
	w.Header().Set("Content-type","application/json")
	enc:=json.NewEncoder(w)
	if err:=enc.Encode(entry);err != nil {
		log.Fatalf("%s",err)	
	}
}

func getRequestHandler(w http.ResponseWriter,r *http.Request){

	key:= r.URL.Path[4:]
	
	defer r.Body.Close()
	value,ok:=database[key]
	if !ok {
		http.Error(w,fmt.Sprintf("Key %s not found",key),http.StatusNotFound) 
		return
	}
	entry:=&Entry{
		Key: key,
		Value:value,
	}
	w.Header().Set("Content-type","application/json")
	enc:=json.NewEncoder(w)
	if err:=enc.Encode(entry);err != nil {
		log.Fatalf("%s",err)	
	}
}


func main() {
	http.HandleFunc("/db",postRequestHandler)
	http.HandleFunc("/db/",getRequestHandler)
	if err := http.ListenAndServe(":8080",nil); err!=nil {
		log.Fatalf("%s",err) 
	}
}