package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

//Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name, returns error
	if name == "" {
		return "", errors.New("Empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos return messages
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	//loop
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}

	return messages, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	// A slice of message formats.

	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"How far guy, %v well met!",
	}

	return formats[rand.Intn(len(formats))]
}
