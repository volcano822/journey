package routes

import (
	"github.com/dimfeld/httptreemux"
	"github.com/volcano822/journey/common/filenames"
	"github.com/volcano822/journey/common/helpers"
	"github.com/volcano822/journey/common/templates"
	"github.com/volcano822/journey/common/structure"
	"net/http"
	"path/filepath"
	"strings"
	"time"
)

func pagesHandler(w http.ResponseWriter, r *http.Request, params map[string]string) {
	path := filepath.Join(filenames.PagesFilepath, params["filepath"])
	// If the path points to a directory, add a trailing slash to the path (needed if the page loads relative assets).
	if helpers.IsDirectory(path) && !strings.HasSuffix(r.RequestURI, "/") {
		http.Redirect(w, r, r.RequestURI+"/", 301)
		return
	}
	http.ServeFile(w, r, path)
	return
}

func aboutHandler(w http.ResponseWriter, r *http.Request, params map[string]string) {
	//http.ServeFile(w, r, filenames.PagesFilepath+"/about.html")
	var t time.Time
	t = time.Now()
	content := `
		<h2> Mr. Coffee</h2>

		<ul>
			<li><p>p1</p></li>
			<li><p>p2</p></li>
		</ul>
	`
	templates.ShowPostTemplateByContent(w, r, &structure.Post{
		//Markdown: []byte(content),
		Title: []byte("About"),
		Id:    int64(-1),
		Uuid:  []byte("uuid for about"),
		Slug:  "about",
		Html:  []byte(content),
		//IsFeatured      bool
		//IsPage          bool
		//IsPublished     bool
		Date: &t,
		//Tags            []Tag
		//Author          *User
		//MetaDescription []byte
		//Image           []byte
	})
	return
}

func InitializePages(router *httptreemux.TreeMux) {
	router.GET("/about", aboutHandler)
	// For serving standalone projects or pages saved in in content/pages
	router.GET("/pages/*filepath", pagesHandler)
}
