package main

import (
	"food-delivery/common"
	"food-delivery/middleware"
	"food-delivery/module/note/notemodel"
	"food-delivery/module/note/notetransport/ginnote"
	"food-delivery/module/user/usertransport/ginuser"
	"github.com/gin-gonic/gin"
	"net/http"
)

func setupRouter(router *gin.Engine, ctx common.AppContext) {
	//router.Use(middleware.Recover(ctx))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/register", ginuser.Register(ctx))
	router.POST("/login", ginuser.Login(ctx))
	router.GET("/profile", middleware.RequiredAuth(ctx), ginuser.GetProfile(ctx))

	router.GET("/demo", func(c *gin.Context) {
		c.JSON(http.StatusOK, notemodel.Note{})
	})

	router.POST("/demo", func(c *gin.Context) {
		var data LoginData

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, data)
	})

	notes := router.Group("/notes", middleware.Recover(ctx))
	{
		notes.PUT("/:note-id", ginnote.UpdateNote(ctx))
		notes.POST("", ginnote.CreateNote(ctx))
		notes.GET("", ginnote.ListNote(ctx))
		notes.DELETE("/:note-id", ginnote.DeleteNote(ctx))
	}
}
