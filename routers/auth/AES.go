package auth

import (
	"bytes"
	"com/models/wx/wxuser"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AesEncrypt(orig string, key string) (str string,err error) {
	// 转成字节数组
	origData := []byte(orig)
	k := []byte(key)
	// 分组秘钥
	// NewCipher该函数限制了输入k的长度必须为16, 24或者32
	block, err := aes.NewCipher(k)
	if err != nil {
		return
	}
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 补全码
	origData = PKCS7Padding(origData, blockSize)
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
	// 创建数组
	cryted := make([]byte, len(origData))
	// 加密
	blockMode.CryptBlocks(cryted, origData)

	str = base64.StdEncoding.EncodeToString(cryted)
	return
}

func AesDecrypt(cryted string, key string) (str string,err error) {
	// 转成字节数组
	crytedByte, err := base64.StdEncoding.DecodeString(cryted)
	if err != nil {
		return
	}
	k := []byte(key)
	// 分组秘钥
	block, err := aes.NewCipher(k)
	if err != nil {
		return
	}
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 加密模式
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	// 创建数组
	orig := make([]byte, len(crytedByte))
	// 解密
	blockMode.CryptBlocks(orig, crytedByte)
	// 去补全码
	orig = PKCS7UnPadding(orig)

	str = string(orig)
	return
}
//补码
//AES加密数据块分组长度必须为128bit(byte[16])，密钥长度可以是128bit(byte[16])、192bit(byte[24])、256bit(byte[32])中的任意一个。
func PKCS7Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
//去码
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}


func ParseAES(c *gin.Context) {
	h := c.Request.Header.Get("sessionId")
	if h == "" {
		fmt.Println("sessionId不能为空")
		c.JSON(http.StatusOK, gin.H{
			"status": 401,
			"error":  nil,
			"data":  "sessionId不存在",
		})
		c.Abort()
		return
	}
	openId,err := AesDecrypt(h,Key)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 401,
			"error":  err,
			"data":   "解密失败",
		})
		c.Abort()
		return
	}
	wxUser := wxuser.WxUser{
		OpenId:openId,
	}
	wu,err := wxUser.FindOpenId()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 401,
			"error":  err,
			"data":"查询出错",
		})
		c.Abort()
		return
	}
	if wu.ID <= 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": 401,
			"error":  nil,
			"data":"不存在此用户",
		})
		c.Abort()
		return
	}
	c.Set("openId", wu.OpenId)
	c.Next()
	return
}
