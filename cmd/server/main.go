package main

import (
    "github.com/gin-gonic/gin"
    "github.com/secondary-jcav/my-website/handlers"
)

func main() {
    r := gin.Default()

    r.LoadHTMLGlob("pages/*")

    // Define endpoints
    r.GET("/", handlers.Home)
    r.GET("/projects", handlers.MyProjects)
    r.GET("/about", handlers.AboutMe)
    r.GET("/download-file", handlers.DownloadCv)

    // Serve static content
    r.Static("/static", "./static")



    // Start the server
    r.Run(":8080") // Listen and serve on localhost:8080
}
