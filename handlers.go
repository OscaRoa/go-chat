package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	response := Response{
		Status: http.StatusOK,
		Data: "Welcome!",
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

func ChatList(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	chats := Chats{
		Chat{
			Id: 1,
			Users: Users{
				User{
					Id: 1,
					Name: "Oscar Roa",
				},
				User{
					Id: 2,
					Name: "Daniela Millán",
				},
			},
			Messages: Messages{
				Message{
					Id: 1,
					Message: "Hello world!",
					User: User{
						Id: 1,
						Name: "Oscar Roa",
					},
				},
				Message{
					Id: 2,
					Message: "Hello world... again!",
					User: User{
						Id: 2,
						Name: "Daniela Millán",
					},
				},
			},
			Seen: true,
		},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	response := Response{
		Status: http.StatusOK,
		Data: chats,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

func ChatDetail(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	data := fmt.Sprintf("Chat number %s", ps.ByName("chatId"))

	response := Response{
		Status: http.StatusOK,
		Data: data,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}
