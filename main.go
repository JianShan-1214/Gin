package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var userData = map[string]string{
	"admin": "admin",
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("./template/*")
	r.Static("/static", "./template/")
	r.GET("/", loginPage)
	r.GET("/register", registerPage)
	r.POST("/", loginAuth)
	r.POST("/register", registerAuth)
	r.Run()
}

func loginPage(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

func registerPage(c *gin.Context) {
	c.HTML(200, "register.html", nil)
}

func loginAuth(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if userData[username] == password {
		c.HTML(200, "login.html", gin.H{
			"success": "登入成功",
		})
	} else {
		c.HTML(403, "login.html", gin.H{
			"error": "帳號或密碼錯誤",
		})
	}
}

func registerAuth(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	secendPassword := c.PostForm("secondPassword")
	fmt.Println(username, password, secendPassword)
	if password != secendPassword {
		c.HTML(403, "register.html", gin.H{
			"error": "密碼不一致",
		})
	} else if userData[username] != "" {
		c.HTML(403, "register.html", gin.H{
			"error": "帳號已被使用",
		})
	} else {
		userData[username] = password
		c.HTML(200, "register.html", gin.H{
			"success": "註冊成功",
		})
	}
}
