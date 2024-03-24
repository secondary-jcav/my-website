package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// Home handler
func Home(c *gin.Context) {
    c.HTML(http.StatusOK, "home.html", gin.H{
        "title": "Home Page",
    })
}

// MyProjects handler
func MyProjects(c *gin.Context) {
    c.HTML(http.StatusOK, "project.html", gin.H{
        "title": "My Projects",
    })
}

// AboutMe handler
func AboutMe(c *gin.Context) {
    c.HTML(http.StatusOK, "about.html", gin.H{
        "title": "About Me",
    })
}

