package main

import (
	"encoding/json"
	"net/http"
	"gofr.dev/pkg/gofr"
)

type Data struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

type DataStore struct {
	data []Data
}

func NewDataStore() *DataStore {
	return &DataStore{
		data: []Data{
			{ID: 1, Name: "John"},
			{ID: 2, Name: "Jane"},
		},
	}
}

func (ds *DataStore) GetAllData() []Data {
	return ds.data
}

func (ds *DataStore) GetDataByID(id int) *Data {
	for _, data := range ds.data {
		if data.ID == id {
			return &data
		}
	}
	return nil
}

func (ds *DataStore) AddData(data Data) {
	ds.data = append(ds.data, data)
}

func WriteJson( _ w http.ResponseWriter) (interface{}, error) {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(data)
}

func ReadJson(r *http.Request) (interface{}, error){
	return json.NewDecoder(r.Body).Decode(data)
}


func getAllHandler(w http.ResponseWriter, r *http.Request) error{
	data:= datastore.GetAllData()
	return WriteJson(w, data)
}

// func getHandler(w http.ResponseWriter, r *http.Request) {
// 	id:= ("id")
// 	data:= datastore.GetDataByID(id)
// 	if data == nil {
// 		http.Error(w, fmt.Sprintf("Data with id %d not found", id), http.StatusNotFound)
// 		return
// 	}
// 		WriteJson(w, data)
// }	

func addDataHandler(w http.ResponseWriter, r *http.Request) error {
	var data Data
	err:= ReadJson(r, &data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	datastore.AddData(data)
	return WriteJson(w, data)
}

var datastore = NewDataStore()





func main() {
	app:= gofr.New()

	app.GET("/data", getAllHandler)
	// app.GET("/data/{id}", getHandler)
	app.POST("/data", addDataHandler)

	app.Start()
}



