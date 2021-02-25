package main

import (
	"food-delivery/common"
	"food-delivery/module/note/notestorage"
	"log"
	"net/http"
	"os"
	"strconv"

	"food-delivery/module/note/notebusiness"
	"food-delivery/module/note/notemodel"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// LoginData is login body
type LoginData struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func main() {
	dsn := os.Getenv("DB_CONN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Connected to DB.", db)

	// if err := db.Create(&note).Error; err != nil {
	// 	log.Println("Can not create a note", db)
	// }

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/demo", func(c *gin.Context) {
		// c.JSON(http.StatusOK, LoginData{
		// 	Username: "chungbui",
		// 	Password: "123456",
		// })
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
		notes.GET("", func(c *gin.Context) {
			var notes []notemodel.Note

			var limit int
			var page int

			limitStr := c.DefaultQuery("limit", "10")
			rawPage := c.DefaultQuery("page", "0")

			if rLimit, err := strconv.Atoi(limitStr); err == nil {
				limit = rLimit
			}

			if rPage, err := strconv.Atoi(rawPage); err == nil {
				page = rPage
			}

			if err := db.Table(notemodel.NoteTableName).
				Limit(limit).
				Offset(page).
				Find(&notes).Error; err != nil {
				log.Println("can't get any notes", err)
				c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "err"})
			} else {
				c.JSON(http.StatusOK, gin.H{"data": notes})
			}
		})
		notes.POST("", func(c *gin.Context) {
			var note notemodel.Note

			if err := c.ShouldBind(&note); err != nil {
				c.AbortWithStatusJSON(400, gin.H{"message": err})
				return
			}

			if err := db.Create(&note).Error; err != nil {
				log.Println("cant create a note", err)
				c.AbortWithStatusJSON(400, gin.H{"message": err})
				return
			}

			c.JSON(http.StatusOK, gin.H{"id": note.ID})
		})
		notes.GET("/:note-id", func(c *gin.Context) {
			id := c.Param("note-id")

			var note notemodel.Note

			if err := db.Table(note.TableName()).
				Where("id = ?", id).First(&note).Error; err != nil {
				log.Println("can't get a note", err)
				c.JSON(http.StatusOK, gin.H{"error": "err"})
			} else {
				c.JSON(http.StatusOK, gin.H{"data": note})
			}
		})
		notes.DELETE("/:note-id", func(c *gin.Context) {
			id, _ := strconv.Atoi(c.Param("note-id"))

			store := notestorage.NewSQLStore(db)
			biz := notebusiness.NewDeleteNoteBiz(store)

			if err := biz.DeleteNote(id); err != nil {
				c.JSON(401, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{"data": 1})
		})
	}
	_ = router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

type fakeDeleteNoteStore struct{}

func (fakeDeleteNoteStore) FindDataWithCondition(condition map[string]interface{}) (*notemodel.Note, error) {
	return &notemodel.Note{
		SQLModel: common.SQLModel{ID: 1, Status: 1},
		Title:    "",
		Content:  "",
	}, nil
}

func (fakeDeleteNoteStore) Delete(id int) error {
	return nil
}
