package main

import (
    "net/http"
    "log"
    "../src/car"
	"fmt"
	"encoding/json"
)

func main() {
	http.HandleFunc("/api/v1/cars", handler)
		log.Fatal(http.ListenAndServe(":8080", nil))
}
func handler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(r.URL.Path)
	if r.URL.Path == "/" {
		http.Error(w, "API to retreive cars and the show they attended, please call /api/v1/cars endpoint",
			http.StatusNotFound)
		return
	}
	fmt.Println(r.Method)
	switch r.Method {
	case "GET":
		carShowSlice := car.LoadData()
		sortCarsByMakeShowSlice := car.SortData(carShowSlice)
		bytes, _ := MarshalCarSlice(sortCarsByMakeShowSlice)
		carShows := []car.CarShow{}
		carcarShows := UnMarshalCarSlice( bytes, carShows)
		fmt.Fprintf(w, "%+v\n", carcarShows)
	}
}

//unmarshaling bytes to the carshows slice type of CarShow, printing out error if any
func UnMarshalCarSlice(bytes []byte, carShows []car.CarShow) []car.CarShow {
	err := json.Unmarshal(bytes, &carShows)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", carShows)
	return carShows
}

// marshal data to get bytes required for unMarshaling, printing out error if any
func MarshalCarSlice(sortCarsByMakeShowSlice []car.CarShow) ([]byte, error) {
	bytes, err := json.Marshal(sortCarsByMakeShowSlice)
	if err != nil {
		fmt.Println(err)
	}
	return bytes, err
}
