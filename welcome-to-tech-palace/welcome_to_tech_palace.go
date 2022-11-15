package techpalace
import "strings"

// WelcomeMessage returns a welcome message for the customer.
const GenericWelcome = "Welcome to the Tech Palace, "
func WelcomeMessage(customer string) string {
	var customerName = strings.ToUpper(customer)

	return GenericWelcome + customerName
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	var border = strings.Repeat("*", numStarsPerLine)
	return border + "\n" + welcomeMsg + "\n" + border
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	var noStars =	strings.ReplaceAll(oldMsg, "*", "")
	var newMsg = strings.TrimSpace(noStars)
	return newMsg
}
