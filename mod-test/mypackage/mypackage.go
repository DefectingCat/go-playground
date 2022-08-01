package mypackage

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Hello(name string) (string, error) {
	// If no name given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}
	return fmt.Sprintf(randomFormat(), name), nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, v := range names {
		message, err := Hello(v)
		if err != nil {
			return nil, err
		}
		messages[v] = message
	}
	return messages, nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
