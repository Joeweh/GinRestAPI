package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var users = []user{
	{ID: "1", Email: "user1@gmail.com", Username: "User1"},
	{ID: "2", Email: "user2@gmail.com", Username: "User2"},
	{ID: "3", Email: "user3@gmail.com", Username: "User3"},
}

func GetUsers(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, users)
}

func SaveUser(ctx *gin.Context) {
	var newUser user

	if err := ctx.BindJSON(&newUser); err != nil {
		return
	}

	

	users = append(users, newUser)

	ctx.IndentedJSON(http.StatusCreated, newUser)
}

func UpdateUser(ctx *gin.Context) {

}

func DeleteUser(ctx *gin.Context) {

}