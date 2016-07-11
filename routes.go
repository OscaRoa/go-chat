package main

import (
	"github.com/julienschmidt/httprouter"
)

type Route struct {
	Name		string
	Method		string
	Pattern		string
	HandlerFunc	httprouter.Handle
}

type Routes []Route

var routes = Routes{
    Route{
	"Index",
	"GET",
	"/",
	Index,
    },
    Route{
	"ChatList",
	"GET",
	"/chats/",
	ChatList,
    },
    Route{
	"ChatDetail",
	"GET",
	"/chats/:chatId/",
	ChatDetail,
    },
}
