package ginnote

import (
	"food-delivery/common"
	"food-delivery/module/note/notebusiness"
	"food-delivery/module/note/notestorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListNote(context common.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		db := context.GetMainDBConnection()

		paging := common.Paging{}
		paging.Fulfill()
		store := notestorage.NewSQLStore(db)
		biz := notebusiness.NewListNoteBiz(store)

		result, err := biz.ListNote(&paging)
		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": result})
	}
}
