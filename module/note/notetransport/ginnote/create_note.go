package ginnote

import (
	"food-delivery/common"
	"food-delivery/module/note/notebusiness"
	"food-delivery/module/note/notemodel"
	"food-delivery/module/note/notestorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateNote(context common.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		db := context.GetMainDBConnection()
		var data notemodel.CreateNote

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(400, common.ErrInvalidRequest(err))
		}

		store := notestorage.NewSQLStore(db)
		biz := notebusiness.NewCreateNoteBiz(store)

		if err := biz.CreateNote(c.Request.Context(), &data); err != nil {
			c.JSON(401, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.ID))
	}
}
