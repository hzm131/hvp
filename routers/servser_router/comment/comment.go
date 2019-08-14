package commentManagement

import (
	"com/models/servser_model/comment"
	"github.com/gin-gonic/gin"
	"net/http"
)

func QueryComment(c *gin.Context) {
	limit := c.Request.URL.Query().Get("limit")
	offset := c.Request.URL.Query().Get("offset")
	condition := c.Request.URL.Query().Get("condition")
	orderBy := c.Request.URL.Query().Get("order_by")

	comment := comment.Comment{}
	value,err := comment.QueryComment(condition, orderBy, limit, offset)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"error": err,
			"data": "查询失败",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"error": nil,
		"data": value,
	})
}

func DeleteComment(c *gin.Context){
	commentId := c.Param("id")
	if commentId == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"data":   "没有id怎么查",
		})
		return
	}
	comment := comment.Comment{}
	err := comment.DeleteComment(commentId)
	if err != nil{
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"error":err,
			"data":   "删除失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"error":nil,
		"data":   "删除成功",
	})
}