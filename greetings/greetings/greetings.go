package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("Empty name provided")
	}

	var message2 string
	message2 = name

	message := fmt.Sprintf(randomFormat(), message2)

	return message, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	messages := []string{
		"Hi, %v, Welcome",
		"Hello %v",
	}

	return messages[rand.Intn(len(messages))]
}
