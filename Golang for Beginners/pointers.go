package main

import "log"

func main() {
	var old int
	var oldString string
	old = 11
	oldString = "This is old string"
	log.Println("OLD: ", old, oldString)
	changeValueUsingReference(&old, &oldString)
	log.Println("NEW: ", old, oldString)

}
func changeValueUsingReference(i *int, s *string) {
	var new int
	var newString string
	new = 22
	newString = "This is a new string"
	*i = new
	*s = newString
}
