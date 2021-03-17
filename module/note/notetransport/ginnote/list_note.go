package ginnote

import (
	"food-delivery/common"
	"food-delivery/module/note/notebusiness"
	"food-delivery/module/note/notestorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func demoCrash() {
	var list []int
	_ = list[0]
}

func ListNote(context common.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		paging := common.Paging{}

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := context.GetMainDBConnection()

		paging.Fulfill()
		store := notestorage.NewSQLStore(db)
		biz := notebusiness.NewListNoteBiz(store)

		result, err := biz.ListNote(c.Request.Context(), &paging)
		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].GenUID(common.DbTypeNote)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, nil))
	}
}
