package main

import (
	"food-delivery/common"
	"food-delivery/module/note/notemodel"
	"food-delivery/module/note/notetransport/ginnote"
	"github.com/gin-gonic/gin"
	"net/http"
)

func setupRouter(router *gin.Engine, ctx common.AppContext) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

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

	notes := router.Group("/notes")
	{
		notes.PUT("/:note-id", ginnote.UpdateNote(ctx))
		//notes.GET("/:note-id", func(c *gin.Context) {
		//	id := c.Param("note-id")
		//
		//	var note notemodel.Note
		//
		//	if err := db.Table(note.TableName()).
		//		Where("id = ?", id).First(&note).Error; err != nil {
		//		log.Println("can't get a note", err)
		//		c.JSON(http.StatusOK, gin.H{"error": "err"})
		//	} else {
		//		c.JSON(http.StatusOK, gin.H{"data": note})
		//	}
		//})
		notes.GET("", ginnote.ListNote(ctx))
		notes.DELETE("/:note-id", ginnote.DeleteNote(ctx))
	}
}
