package main

import (
    "net/http"
	//"text/template"

	//"github.com/forrestly/gwiki/handlers"
	"github.com/gin-gonic/gin"
)

//Index is the test handler
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "install.tmpl", gin.H{"msg": "Main website"})
}

func main() {
	gin.SetMode(gin.DebugMode)
	g := gin.Default()
	g.StaticFile("/favicon.ico", "./favicon.ico")
	g.Static("/assets", "./assets")
	g.LoadHTMLGlob("templates/*")
	g.GET("/install", Index)
    
	g.Run(":5678")
	// t := template.New("ss")
	// t.Execute(nil, nil)

}


