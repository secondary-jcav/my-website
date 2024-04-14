package handlers

import (
    "html/template"
    "net/http"
    "github.com/gin-gonic/gin"
    "path/filepath"
)

// Home handler
func Home(c *gin.Context) {
    // Parsing the templates
    tmpl, err := template.ParseFiles(
        "pages/home.html",
        "templates/header.html",
        "templates/work-experience.html",
        "templates/skills.html",
        "templates/certs.html",
        "templates/footer.html",
    )
    if err != nil {
        c.AbortWithError(http.StatusInternalServerError, err)
        return
    }
    
    // Data to pass to the templates
    data := struct {
        Current string
    }{
        Current: "home", // This sets the current page as 'home'
    }
    
    // Execute the template with the data
    err = tmpl.ExecuteTemplate(c.Writer, "home.html", data)
    if err != nil {
        c.AbortWithError(http.StatusInternalServerError, err)
        return
    }
}


// MyProjects handler
func MyProjects(c *gin.Context) {
    templ,err:= template.ParseFiles("pages/project.html","templates/header.html","templates/my-projects.html","templates/footer.html")
    if err != nil {
        c.AbortWithError(http.StatusInternalServerError, err)
        return
    }

        // Data to pass to the templates, now including the 'Current' page identifier
        data := struct {
            Current string // Adding a Current field to manage active state in the navigation
        }{
            Current: "projects", // This sets the current page as 'about'
        } 

    err = templ.ExecuteTemplate(c.Writer,"project.html",data)
    if err != nil {
        c.AbortWithError(http.StatusInternalServerError, err)
        return
    }


}

// AboutMe handler
func AboutMe(c *gin.Context) {
    // Parsing all necessary templates including the header and footer
    templ, err := template.ParseFiles(
        "pages/about.html",
        "templates/header.html",
        "templates/introduction.html",
        "templates/recommendations.html",
        "templates/contact.html",
        "templates/footer.html",
    )
    if err != nil {
        c.AbortWithError(http.StatusInternalServerError, err)
        return
    }

    // Data to pass to the templates, now including the 'Current' page identifier
    data := struct {
        Current string // Adding a Current field to manage active state in the navigation
    }{
        Current: "about", // This sets the current page as 'about'
    } 


    // Execute the template with the data
    err = templ.ExecuteTemplate(c.Writer, "about.html", data)
    if err != nil {
        c.AbortWithError(http.StatusInternalServerError, err)
        return
    }
}


func DownloadCv(c *gin.Context) {
    // Specify the path to the file
    filePath := "download/JulioArellanoCV.pdf"

    // Set the header to inform the browser about the file type and prompt for download
    c.Header("Content-Description", "File Transfer")
    c.Header("Content-Transfer-Encoding", "binary")
    c.Header("Content-Disposition", "attachment; filename="+filepath.Base(filePath))
    c.Header("Content-Type", "application/pdf")

    // Serve the file
    c.File(filePath)
}

