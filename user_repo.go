package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var users = []user{
	{ID: 0, Email: "user0@gmail.com", Username: "User0"},
	{ID: 1, Email: "user1@gmail.com", Username: "User1"},
	{ID: 2, Email: "user2@gmail.com", Username: "User2"},
}

func GetUsers(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, users)
}

func GetUser(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	for i := 0; i < len(users); i++ {
		if users[i].ID == id {
			ctx.IndentedJSON(http.StatusOK, users[i])

			return
		}
	}

	ctx.Status(http.StatusNotFound)
}

func SaveUser(ctx *gin.Context) {
	var newUser user

	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err := ctx.BindJSON(&newUser); err != nil {
		return
	}

	for index := 0; index < len(users); index++ {
		if users[index].ID == id {
			ctx.Status(http.StatusConflict)

			return
		}
	}

	users = append(users, newUser)

	ctx.IndentedJSON(http.StatusCreated, newUser)
}

func UpdateUser(ctx *gin.Context) {
	
}

func DeleteUser(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	for index := 0; index < len(users); index++ {
		if users[index].ID == id {
			users = append(users[:index], users[index+1:]...)

			ctx.Status(http.StatusNoContent)

			return
		}
	}

	ctx.Status(http.StatusNotFound)
}
