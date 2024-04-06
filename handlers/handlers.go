package handlers

import (
    "html/template"
    "net/http"
    "github.com/gin-gonic/gin"
)

// Home handler
func Home(c *gin.Context) {

        tmpl, err := template.ParseFiles("pages/home.html","templates/header.html","templates/work-experience.html",  "templates/skills.html", "templates/certs.html","templates/footer.html")
        if err != nil {
            c.AbortWithError(http.StatusInternalServerError, err)
            return
        }
    
        // Data to pass to the templates (static at the moment)
        data := struct {
            // Define any data
        }{
            // Populate data here
        }
        
        err = tmpl.ExecuteTemplate(c.Writer, "home.html", data)
        if err != nil {
            c.AbortWithError(http.StatusInternalServerError, err)
            return
        }
    
    }

// MyProjects handler
func MyProjects(c *gin.Context) {
    templ,err:= template.ParseFiles("pages/project.html","templates/header.html","templates/footer.html")
    if err != nil {
        c.AbortWithError(http.StatusInternalServerError, err)
        return
    }
    // Data to pass to the templates (static at the moment)
    data := struct {
        // Define any data
    }{
        // Populate data here
    }
    err = templ.ExecuteTemplate(c.Writer,"project.html",data)
    if err != nil {
        c.AbortWithError(http.StatusInternalServerError, err)
        return
    }


}

// AboutMe handler
func AboutMe(c *gin.Context) {

    templ, err:= template.ParseFiles("pages/about.html","templates/header.html","templates/footer.html")
    if err != nil {
        c.AbortWithError(http.StatusInternalServerError, err)
        return
    }

      // Data to pass to the templates (static at the moment)
      data := struct {
        // Define any data
    }{
        // Populate data here
    }

    err = templ.ExecuteTemplate(c.Writer, "about.html", data)
        if err != nil {
            c.AbortWithError(http.StatusInternalServerError, err)
            return
        }
        
}

