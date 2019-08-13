package jwt

import (
	"com/models/servser_model/users"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

/*var str = []byte("hezhimin")
func GenerateJWT()(string,error){
	// 创建一个新的token 参数是使用的签名 New方法返回Token结构体
	token := jwt.New(jwt.SigningMethodHS256)
	// 令牌的第二部分   Token结构体包含Claims接口
	//type MapClaims map[string]interface{}
	//Claims是接口  MapClaims实现了此接口的Valid()方法 等于创建了一个映射
	//创建一个jwt提供的映射，进行json解码的数据类型
	claims := make(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims

	//完整的签名令牌  使用token结构的SignedString方法
	tokenString,err := token.SignedString(str)
	if err != nil {
		fmt.Errorf(err.Error())
		return "",err
	}
	return tokenString,nil
}*/

/*
jwt的组成部分
Header：{
	typ: "JWT", typ是默认的一种标识.标识这条信息采用JWT规范.
	alg: "HS256"  alg表示签名使用的加密算法.通常有ES256,ES512,RS256等等
	}

payload：有效载荷
用来承载要传递的数据，它的一个属性对被称为claim，这样的标准成为claims标准，同样是将其用Base64Url编码
	Audience  string `json:"aud,omitempty"`
	ExpiresAt int64  `json:"exp,omitempty"`
	Id        string `json:"jti,omitempty"`
	IssuedAt  int64  `json:"iat,omitempty"`
	Issuer    string `json:"iss,omitempty"`
	NotBefore int64  `json:"nbf,omitempty"`
	Subject   string `json:"sub,omitempty"`

这段结构体是我们在golang中使用到的字段. 可以在这个的基础上进行组合,定义新的Claims部分.

1. aud 标识token的接收者.
2. exp 过期时间.通常与Unix UTC时间做对比过期后token无效
3. jti 是自定义的id号
4. iat 签名发行时间.
5. iss 是签名的发行者.
6. nbf 这条token信息生效时间.这个值可以不设置,但是设定后,一定要大于当前Unix UTC,否则token将会延迟生效.
7. sub 签名面向的用户

可以随意加

注意，不要在JWT的payload或header中放置敏感信息，除非它们是加密的。

Signature ：（签名）
组成部分
将Header与Claims信息拼接起来[base64(header)+"."+base64(claims)],采用Header中指定的加密算法进行加密,得到Signature部分.
token组成部分
base64(header) + “.” + base64(claims) + “.” + 加密签名


是用 header + payload + secret组合起来加密的,公式是:

HMACSHA256(
  base64UrlEncode(header) + "." +
  base64UrlEncode(payload),
  secret)
这里 secret就是自己定义的一个随机字符串,这一个过程只能发生在 server 端,会随机生成一个 hash 值

这样组合起来之后就是一个完整的 jwt 了:


例：
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjI5NTg5MTA4MDYsImlzcyI6InRlc3QiLCJuYmYiOjE0Nzk0NTczMTZ9.57gqtlk1nNezXSa0VgWBOwu2b2FCDJ6wXizuJF6IY10

如上边的token,由2个点好分割.第一部分是header的base64编码,第二部分是claims的base64编码,第三部分是加密签名信息.
*/
/*type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}*/
var mySigningKey = []byte("hzwy23")

//创建令牌
func CreateJWT(user users.Users) (ss string, err error) {

	/*// 创建声明 就是设置token的一些东西
	claims := &jwt.StandardClaims{
		NotBefore: int64(time.Now().Unix() - 1000), //token信息生效时间.这个值可以不设置,但是设定后,一定要大于当前Unix UTC,否则token将会延迟生效.
		ExpiresAt: int64(time.Now().Unix() + 1000), //过期时间.通常与Unix UTC时间做对比过期后token无效
		Issuer:    "test",                          //是签名的发行者
		Id:        string(id),
	}*/
	claims := jwt.MapClaims{
		"user": user,
		"exp":  int64(time.Now().Unix() + 1000),
		"iss":  "hzm",
	}

	//创建签名 第一个参数是算法 第二个参数是配置
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//获取完整的签名令牌
	ss, err = token.SignedString(mySigningKey)
	if err != nil {
		fmt.Errorf("签名失败")
		return
	}

	return

}

//解析令牌
func ParseToken(c *gin.Context, str string) {
	//解析，验证并返回令牌,如果token 信息parse后与签名信息不一致,则会爆出异常.
	token, err := jwt.Parse(str, func(*jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err == nil {
		if token.Valid {
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				fmt.Println("HS256的token解析错误，err:", err)
				c.JSON(http.StatusOK, gin.H{
					"status": 401,
					"error":  "token解析错误",
				})
				return
			}
			fmt.Println(claims)
			fmt.Println(&claims)
			user := claims["user"]
			c.Set("user", user)
			c.Next()
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": 401,
				"data":   "token无效",
			})
			c.Abort()
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": 401,
			"data":   "未经授权访问此资源",
		})
		c.Abort()
	}
}
