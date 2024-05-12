package main

import (
	"fmt"
	"strings"
)

func list_bookings() {

	var centerid string
	var day int
	fmt.Scan(&centerid)
	fmt.Scan(&day)

	for unique_string, user := range bookings_made {
		tmp_string := strings.Split(unique_string, " ")
		if tmp_string[1] == centerid && tmp_string[2] == fmt.Sprint(day) {
			fmt.Println(user_map[user])
		}

	}
}

func cancel_booking() {
	var centerid, bookingid, userid string
	fmt.Scan(&centerid)
	fmt.Scan(&bookingid)
	fmt.Scan(&userid)
	for string, user := range bookings_made {
		if user == "-1" {
			continue
		}
		tmp_string := strings.Split(string, " ")
		if tmp_string[1] == centerid {
			day := (tmp_string[2])
			unique_string := centerid + " " + (day)
			capacity_map[unique_string]++
			bookings_made[(string)] = "-1"
			return
		}
	}
}

func book_vaccination() {
	var centerid, user_id string
	var day int
	fmt.Scan(&centerid)
	fmt.Scan(&day)
	fmt.Scan(&user_id)
	unique_string := user_id + " " + centerid + " " + fmt.Sprint(day)
	if bookings_made[unique_string] == (user_id) {
		fmt.Println("Can't Book once more already booked")
		return
	}
	if capacity_map[unique_string] > 0 {
		fmt.Println("Booked Succesfullly")
		capacity_map[unique_string]--
		bookings_made[unique_string] = (user_id)
	} else {
		fmt.Println("Slots Full!")
	}
}
