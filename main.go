package main

import (
	"food-delivery/common"
	"food-delivery/component/appcontext"
	"food-delivery/module/note/notemodel"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
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

	db = db.Debug()
	log.Println("Connected to DB.", db)
	appCtx := appcontext.New(db)

	router := gin.Default()
	setupRouter(router, appCtx)

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
