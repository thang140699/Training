package controller

import (
	"github.com/gin-gonic/gin"
	"mongo-with-golang/models"

	"mongo-with-golang/services"
	"net/http"
)

type Controll struct {
	UserService services.UserService
}

func New(userservice services.UserService) Controll {
	return Controll{
		UserService: userservice,
	}
}

// create domain, time ,..
func (uc *Controll) CreateUser(ctx *gin.Context) {
	var user models.SetTime
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		return
	}
	err := uc.UserService.CreateDomain(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"Message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"messaga": "success"})
}

func (uc *Controll) GetDomain(ctx *gin.Context) {
	Domains := ctx.Param("domain")
	user, err := uc.UserService.GetDomain(&Domains)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"Message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (uc *Controll) GetAll(ctx *gin.Context) {
	users, err := uc.UserService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"Message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func RegisterUserRoutes(rg *gin.RouterGroup) {
	var uc *Controll
	userroute := rg.Group("/user")
	userroute.POST("/create", uc.CreateUser)
	userroute.GET("/get/:name", uc.GetDomain)
	userroute.GET("/getall ", uc.GetAll)
}
