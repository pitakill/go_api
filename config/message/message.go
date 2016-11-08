package message

import (
	"encoding/json"
	"fmt"
	"os"
)

type Message struct {
	User User
}

type User struct {
	Login Login
}

type Login struct {
	Error Error
}

type Error struct {
	Code    int
	Message string
}

func GetMessagesUser() Message {
	file, _ := os.Open("messages.json")
	decoder := json.NewDecoder(file)
	var message Message

	err := decoder.Decode(&message)
	if err != nil {
		fmt.Println("error:", err)
	}

	return message
}
