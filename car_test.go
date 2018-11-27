package test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"reflect"
	"fmt"
	"encoding/json"
	"../car"
)

func TestLoadData(t *testing.T)  {
	result := car.LoadData()
	assert.Equal(t, 2, len(result))
}

func TestValidData(t *testing.T)  {
	c1 := car.CarShow{Car: car.Car{Make: "2020", Model: "George 15"}, Name: "New York Show"}
	c2 := car.CarShow{Car: car.Car{Make: "1010", Model: "Elisa"}, Name: "Las Vegas Show"}
	c3 := car.CarShow{Car: car.Car{Make: "3010", Model: "Elisa"}, Name: "San Francisco Show"}
	c4 := car.CarShow{Car: car.Car{Make: "1910", Model: "Elisa"}, Name: "Los Angeles Show"}
	if reflect.TypeOf(c1) != reflect.TypeOf(car.CarShow{}) {
		t.Error("unmatched type")
		t.Fail()
	}
	assert.Equal(t, reflect.TypeOf(car.CarShow{}), reflect.TypeOf(c1))
	if reflect.TypeOf(c2) != reflect.TypeOf(car.CarShow{}) {
		t.Error("unmatched type")
		t.Fail()
	}
	assert.Equal(t, reflect.TypeOf(car.CarShow{}), reflect.TypeOf(c2))
	if reflect.TypeOf(c3) != reflect.TypeOf(car.CarShow{}) {
		t.Error("unmatched type")
		t.Fail()
	}
	assert.Equal(t, reflect.TypeOf(car.CarShow{}), reflect.TypeOf(c3))
	if reflect.TypeOf(c4) != reflect.TypeOf(car.CarShow{}) {
		t.Error("unmatched type")
		t.Fail()
	}
	assert.Equal(t, reflect.TypeOf(car.CarShow{}), reflect.TypeOf(c4))
	}

// Test case to sort slice by make year/numerically and return make, model and show attended
func TestSortByYear(t *testing.T)  {
	c := car.LoadData()
	sortCarsByMakeShowSlice := car.SortData(c)
	bytes, err := json.Marshal(sortCarsByMakeShowSlice)
	if err != nil {
		fmt.Println(err)
	}
	carShows := []car.CarShow{}
	err = json.Unmarshal(bytes, &carShows)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(carShows)
	assert.Equal(t, 2, len(carShows))
}

// Test case to sort slice by make alphabetically and return make, model and show attended
func TestSortByMakeAlphabetically(t *testing.T)  {
	c1 := car.CarShow{Car: car.Car{Make: "abc2020", Model: "George 15"}, Name: "New York Show"}
	c2 := car.CarShow{Car: car.Car{Make: "nop1010", Model: "Elisa"}, Name: "Las Vegas Show"}
	c3 := car.CarShow{Car: car.Car{Make: "yes3010", Model: "Elisa"}, Name: "San Francisco Show"}
	c4 := car.CarShow{Car: car.Car{Make: "mmm1910", Model: "Elisa"}, Name: "Los Angeles Show"}

	// Creating a slice to hold all the different objects and their data
	carShow := []car.CarShow{}
	sortCarsByMakeShowSlice := append(append(append(append(carShow, c1), c2), c3), c4)
	bytes, err := json.Marshal(sortCarsByMakeShowSlice)
	if err != nil {
		fmt.Println(err)
	}
	carShows := []car.CarShow{}
	err = json.Unmarshal(bytes, &carShows)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(carShows)
	assert.Equal(t, 4, len(carShows))
}

// Test case to return sort for an empty slice
func TestSortByMakeEmptySlice(t *testing.T)  {

	// Creating a slice to hold all the different objects and their data
	carShow := []car.CarShow{}
	sortCarsByMakeShowSlice := car.SortData(carShow)
	bytes, err := json.Marshal(sortCarsByMakeShowSlice)
	if err != nil {
		fmt.Println(err)
	}
	carShows := []car.CarShow{}
	err = json.Unmarshal(bytes, &carShows)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(carShows)
	assert.Equal(t, 0, len(sortCarsByMakeShowSlice))
}