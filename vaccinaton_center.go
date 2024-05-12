package main

import (
	"fmt"
)

type vaccination_centre struct {
	centerid string
	state    string
	district string
}

func add_vaccination_centre() {
	var state, district, centerid string
	fmt.Scan(&state)
	fmt.Scan(&district)
	fmt.Scan(&centerid)
	fmt.Print(state, district, centerid)
	var temp_vaccination_centre vaccination_centre
	temp_vaccination_centre.centerid = centerid
	temp_vaccination_centre.state = state
	temp_vaccination_centre.district = district
	vaccination_centre_map[centerid] = temp_vaccination_centre

}

func add_centre_capacity() {
	var centerid string
	var day, capacity int
	fmt.Scan(&centerid)
	fmt.Scan(&day)
	fmt.Scan(&capacity)
	unique_string := centerid + " " + fmt.Sprint(day)
	capacity_map[unique_string] = capacity

}

func list_vaccination_centres() {

	var district string
	fmt.Scan(&district)

	for _, ele := range vaccination_centre_map {
		if ele.district == district {
			fmt.Println(ele)
		}
	}
}
