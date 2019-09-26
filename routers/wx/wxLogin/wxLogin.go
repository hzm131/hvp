package wxLogin

import (
	"com/models/wx/wxuser"
	"com/routers/auth"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)



func Login(c *gin.Context) {
	value, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return
	}
	user := wxuser.WxUser{}
	err = json.Unmarshal(value,&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"error":  err,
			"data":  "code有问题",
		})
		return
	}
	appid := "wx3103bc03ab579a65"
	secret := "147e54ad0fe7931331ce6310833eee72"
	r,err := getFile("https://api.weixin.qq.com/sns/jscode2session?appid=" + appid + "&secret=" + secret +"&js_code="+ user.Code + "&grant_type=authorization_code")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"error":  err,
			"data":  "获取openid是失败",
		})
		return
	}
	err = json.Unmarshal(r,&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"error":  err,
			"data":  "Unmarshal openid失败",
		})
		return
	}
	wxUser,err := user.CreateData()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"error":  err,
			"data":  "失败",
		})
		return
	}

	encrypt,err := auth.AesEncrypt(wxUser.OpenId,auth.Key)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"error":  err,
			"data":   "openid加密失败",
		})
		return
	}
	wxUser.OpenId = encrypt
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"error":  err,
		"data":   wxUser,
		"sessionId":   encrypt,
	})
}

func getFile(url string) (oids []byte,err error) {
	oid, err := http.Get(url)
	if err != nil {
		return
	}
	defer oid.Body.Close()
	oids, err = ioutil.ReadAll(oid.Body)
	if err != nil {
		return
	}
	return
}



