package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	message := fmt.Sprintf("Welcome to my party, %s!", name)
	return message
	// panic("Please implement the Welcome function")
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	message := fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
	return message
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	tableAssignedMessage := fmt.Sprintf("You have been assigned to table %03d. Your table is %s, exactly %.1f meters from here.", table, direction, distance)
	neighborMessage := fmt.Sprintf("You will be sitting next to %s.", neighbor)
	message := fmt.Sprintf("%s\n%s\n%s", Welcome(name), tableAssignedMessage, neighborMessage)
	return message
}
