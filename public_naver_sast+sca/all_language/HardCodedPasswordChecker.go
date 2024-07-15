package HardCodedPasswordChecker

import (
	models "Go-Gin/model"
	"github.com/gin-gonic/gin"
	"log"
)

func D_1(c *gin.Context) {
	var pass = "1q2w3e!@#"      //@violation
	const PaSsWD = "1q2w3e!@#"  //@violation
	password := "1q2w3e!@#"     //@violation
	eeeepassword := "1q2w3e!@#" //@violation
	passwordeeee := "1q2w3e!@#"
	var (
		pwd    = "1q2w3e!@#" //@violation
		secret = "1q2w3e!@#" //@violation
	)
	const (
		token        = "1q2w3e!@#" //@violation
		thisIsPasswd = "1q2w3e!@#" //@violation
		aptpwd //@violation
	)

	tokentest := "1q2w3e!@#"
	fuckingtoken := "1q2w3e!@#" //@violation
	nopwdddd := "1q2w3e!@#"
	secretcheck := "1q2w3e!@#"
	log.Println(pwd, pass, passwd, password, secret, token, thisIsPasswd, nopwdddd, secretcheck, eeeepassword, passwordeeee, tokentest, fuckingtoken)
	c.String(200, pwd)
}

func S_1(c *gin.Context) {
	password := models.Strings
	pass := models.Strings
	passwd := models.Strings
	pwd := models.Strings
	secret := models.Strings
	token := models.Strings
	log.Println(pwd, pass, passwd, password, secret, token)
	c.String(200, pwd)
}
