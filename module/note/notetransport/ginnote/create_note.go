package ginnote

import (
	"food-delivery/common"
	"food-delivery/module/note/notebusiness"
	"food-delivery/module/note/notemodel"
	"food-delivery/module/note/notestorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateNote(context common.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		db := context.GetMainDBConnection()
		var data notemodel.CreateNote

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(400, common.ErrInvalidRequest(err))
			return
		}

		store := notestorage.NewSQLStore(db)
		biz := notebusiness.NewCreateNoteBiz(store)
		data.UserId = requester.GetUserId()

		if err := biz.CreateNote(c.Request.Context(), &data); err != nil {
			c.JSON(401, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}
