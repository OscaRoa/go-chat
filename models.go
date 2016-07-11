package main

type User struct {
	Id	int			`json:"id"`
	Name	string			`json:"name"`
}

type Users []User

type Message struct {
	Id		int		`json:"id"`
	Message		string		`json:"message"`
	User		User		`json:"user"`
}

type Messages []Message

type Chat struct {
	Id		int		`json:"id"`
	Users		Users		`json:"users"`
	Messages	Messages	`json:"messages"`
	Seen		bool		`json:"seen"`
}

type Chats []Chat

type Response struct {
	Status		int		`json:"status_code"`
	Data		interface{}	`json:"data"`
}
