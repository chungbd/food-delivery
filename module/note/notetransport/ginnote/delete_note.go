package ginnote

import (
	"food-delivery/common"
	"food-delivery/module/note/notebusiness"
	"food-delivery/module/note/notestorage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteNote(context common.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("note-id"))
		db := context.GetMainDBConnection()

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		store := notestorage.NewSQLStore(db)
		biz := notebusiness.NewDeleteNoteBiz(store, requester)

		if err := biz.DeleteNote(c.Request.Context(), id); err != nil {
			c.JSON(401, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
