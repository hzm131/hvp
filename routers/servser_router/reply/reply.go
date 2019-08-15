package replyManagement

import (
	"com/models/servser_model/comment"
	"github.com/gin-gonic/gin"
	"net/http"
)

func QueryComment(c *gin.Context) {
	limit := c.Request.URL.Query().Get("limit")
	offset := c.Request.URL.Query().Get("offset")
	comment_id := c.Request.URL.Query().Get("comment_id")
	if comment_id == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"error":  nil,
			"data":   "comment_id不存在",
		})
		return
	}
	reply := comment.Reply{}
	value, err := reply.QueryReply(comment_id, limit, offset)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"error":  err,
			"data":   "查询失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"error":  nil,
		"data":   value,
	})
}
