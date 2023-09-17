package techpalace

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	customer = strings.ToUpper(customer)
	message := fmt.Sprintf("Welcome to the Tech Palace, %s", customer)
	return message
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	stars := strings.Repeat("*", numStarsPerLine)
	fancyMessage := fmt.Sprintf("%s\n%s\n%s", stars, welcomeMsg, stars)
	return fancyMessage
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	cleanMessage := strings.ReplaceAll(oldMsg, "*", "")
	cleanMessage = strings.TrimSpace(cleanMessage)
	return cleanMessage
}
