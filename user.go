package main

import "fmt"

type user struct {
	name     string
	gender   string
	age      int
	district string
	state    string
}

func add_user() {
	var unique_identification string
	var age int
	var gender, name, state, district string
	fmt.Scan(&unique_identification)
	fmt.Scan(&name)
	fmt.Scan(&gender)
	fmt.Scan(&age)
	fmt.Scan(&state)
	fmt.Scan(&district)
	var temp_user user
	// fmt.Println(name, gender, age, district)
	temp_user.name = name
	temp_user.gender = gender
	temp_user.age = age
	temp_user.state = state
	temp_user.district = district
	user_map[unique_identification] = temp_user

}
