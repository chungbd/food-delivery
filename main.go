package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// LoginData is login body
type LoginData struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

const noteTableName = "notes"

// Note is a entity for note
type Note struct {
	ID      int    `json:"id" gorm:"column:id"`
	Title   string `json:"title" gorm:"column:title"`
	Content string `json:"content" gorm:"column:content"`
	// Status   uint      `json:"status" gorm:"column:status"`
	// CreateAt time.Time `json:"created_at" gorm:"column:created_at"`
	// UpdateAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

// TableName is the table name of note from DB
func (Note) TableName() string {
	return noteTableName
}

func main() {

	dsn := "root:chungbd@tcp(127.0.0.1:3306)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local"
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
		c.JSON(http.StatusOK, Note{})
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
			var notes []Note

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

			if err := db.Table(noteTableName).
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
			var note Note

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

			var note Note

			if err := db.Table(note.TableName()).
				Where("id = ?", id).First(&note).Error; err != nil {
				log.Println("can't get a note", err)
				c.JSON(http.StatusOK, gin.H{"error": "err"})
			} else {
				c.JSON(http.StatusOK, gin.H{"data": note})
			}
		})
	}
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
