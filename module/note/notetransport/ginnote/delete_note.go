package ginnote

import (
	"food-delivery/common"
	"food-delivery/module/note/notebusiness"
	"food-delivery/module/note/notestorage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DeleteNote(context common.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("note-id"))
		db := context.GetMainDBConnection()

		store := notestorage.NewSQLStore(db)
		biz := notebusiness.NewDeleteNoteBiz(store)

		if err := biz.DeleteNote(id); err != nil {
			c.JSON(401, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
