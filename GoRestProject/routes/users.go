package routes

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suraj8108/learngo/Gorestproject/models"
	"github.com/suraj8108/learngo/Gorestproject/utils"
)

type UserHandler struct {
	DB *sql.DB
}

func (u UserHandler) signUp(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the data"})
		return
	}

	err = user.Save(u.DB)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse request data"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User created", "event": user})
}

func (u UserHandler) login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the data"})
		return
	}

	err = user.ValidateCredentials(u.DB)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not authenticate user", "errorMessage": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not authenticate user", "errorMessage": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Login Successfully", "token": token})

}
