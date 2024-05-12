package main

import (
	"fmt"
	"strings"
)

var user_map map[string]user
var vaccination_centre_map map[string]vaccination_centre
var capacity_map map[string]int
var bookings_made map[string]string

func take_input() {
	var cmd string
	fmt.Scan(&cmd)
	res := strings.Split(cmd, " ")
	if res[0] == "ADD_USER" {
		add_user()
	} else if res[0] == "ADD_VACCINATION_CENTER" {
		add_vaccination_centre()
	} else if res[0] == "ADD_CAPACITY" {
		add_centre_capacity()
	} else if res[0] == "LIST_VACCINATION_CENTERS" {
		list_vaccination_centres()
	} else if res[0] == "BOOK_VACCINATION" {
		book_vaccination()
	} else if res[0] == "LIST_BOOKINGS" {
		list_bookings()
	} else if res[0] == "CANCEL_BOOKING" {
		cancel_booking()
	} else if res[0] == "LIST_ALL_BOOKINGS" {
		// search_vaccination_centre()
	}
}
func main() {
	user_map = make(map[string]user)
	vaccination_centre_map = make(map[string]vaccination_centre)
	capacity_map = make(map[string]int)
	bookings_made = make(map[string]string)
	var queries int
	fmt.Print("Enter no of queries\n")
	fmt.Scan(&queries)
	for i := 1; i <= queries; i++ {
		take_input()
	}

	// if len(user_map) > 0 {
	// 	fmt.Print("USERS\n")
	// 	for _, value := range user_map {
	// 		fmt.Println(value)
	// 	}
	// }

	// fmt.Print("Vaccination_centre\n")
	// for _, value := range vaccination_centre_map {
	// 	fmt.Println(value)
	// }

	// fmt.Print("Capacity_MAP\n")
	// for key, val := range capacity_map {
	// 	fmt.Println(key, val)
	// }

}
