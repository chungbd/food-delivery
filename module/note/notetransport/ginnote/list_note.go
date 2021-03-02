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
		paging := common.Paging{}

		if err := c.ShouldBind(&paging); err != nil {
			c.AbortWithStatusJSON(400, gin.H{"message": err})
			return
		}

		db := context.GetMainDBConnection()

		paging.Fulfill()
		store := notestorage.NewSQLStore(db)
		biz := notebusiness.NewListNoteBiz(store)

		result, err := biz.ListNote(&paging)
		if err != nil {
			c.JSON(401, err)
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, nil))
	}
}
