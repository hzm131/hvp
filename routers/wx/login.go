package wx

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)


type WxLogin struct {
	Code string `json:"code"`
	Appid string `json:"appid"`
}

func Login(c *gin.Context) {
	value, err := ioutil.ReadAll(c.Request.Body)
	fmt.Println("value", value)
	if err != nil {
		return
	}
	appid := "wx3103bc03ab579a65"
	secret := "147e54ad0fe7931331ce6310833eee72"
	resp,err := http.Get("https://api.weixin.qq.com/sns/jscode2session?appid="+appid+"&grant_type=authorization_code&secret="+secret+"&js_code="+ string(value) )
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println("resp",resp)
}
