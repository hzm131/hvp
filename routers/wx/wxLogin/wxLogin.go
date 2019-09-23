package wxLogin

import (
	"com/models/wx/user"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)


type Code struct {
	Code string `json:"code"`
	Id int `json:"id"`
}

func Login(c *gin.Context) {
	value, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return
	}
	data := Code{}
	json.Unmarshal(value,&data)
	appid := "wx3103bc03ab579a65"
	secret := "147e54ad0fe7931331ce6310833eee72"
	r,err := getFile("https://api.weixin.qq.com/sns/jscode2session?appid=" + appid + "&secret=" + secret +"&js_code="+ data.Code + "&grant_type=authorization_code")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"error":  err,
			"data":  "失败",
		})
		return
	}

	sky := user.WxUser{}
	json.Unmarshal(r,&sky)
	str := md5V(*sky.SessionKey + " " + *sky.OpenId)
	sky.SessionId = &str
	wxUser,err := sky.CreateData()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"error":  err,
			"data":   "创建失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"error":  nil,
		"data":   wxUser,
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

func md5V(str string) string  {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}