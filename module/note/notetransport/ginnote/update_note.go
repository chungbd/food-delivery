package ginnote

import (
	"food-delivery/common"
	"food-delivery/module/note/notebusiness"
	"food-delivery/module/note/notemodel"
	"food-delivery/module/note/notestorage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UpdateNote(context common.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data notemodel.UpdateNote
		id, _ := strconv.Atoi(c.Param("note-id"))

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(400, common.ErrInvalidRequest(err))
		}

		data.Id = id

		db := context.GetMainDBConnection()

		store := notestorage.NewSQLStore(db)
		biz := notebusiness.NewUpdateNoteBiz(store)

		if err := biz.UpdateNote(c.Request.Context(), data); err != nil {
			c.JSON(401, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
