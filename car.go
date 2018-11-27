package car

import "sort"

// Car of struct type that holds make and model as fields
type Car struct {
	Make  string
	Model string
}

// CarShow of struct type that holds name and Car as fields
type CarShow struct {
	Name string
	Car
}



// function that prepares the carshow types and return as a single array/slice
func LoadData() []CarShow {
	// creating multiple cars with the shows attended
	c1 := CarShow{Car: Car{Make: "George Motors", Model: "George 15"}, Name: "New York Show"}
	c2 := CarShow{Car: Car{Make: "Hondaka", Model: "Elisa"}, Name: "New York Show"}

	// Creating a slice to hold all the different objects and their data
	carShow := []CarShow{}
	carShow = append(carShow, c1, c2)
	return carShow
}

// a sort function that will sort the data by Make and return the sorted array of CarShow type result
func SortData(carShow []CarShow) []CarShow {
	sort.Slice(carShow, func(i, j int) bool {
		return carShow[i].Make < carShow[j].Make
	})
	return carShow
}