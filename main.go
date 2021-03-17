package main

import (
	"fmt"
	"food-delivery/common"
	"food-delivery/component/appcontext"
	"food-delivery/module/note/notemodel"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

// LoginData is login body
type LoginData struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func main() {
	secretKey := os.Getenv("SYSTEM_SECRET")

	// Create the token
	token := jwt.New(jwt.GetSigningMethod("HS256"))

	// Set some claims
	token.Claims = jwt.MapClaims{
		"foo": "bar",
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(secretKey))
	fmt.Println(tokenString, err)

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
		SQLModel: common.SQLModel{Id: 1, Status: 1},
		Title:    "",
		Content:  "",
	}, nil
}

func (fakeDeleteNoteStore) Delete(id int) error {
	return nil
}
