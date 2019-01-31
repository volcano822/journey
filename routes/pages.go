package routes

import (
	"github.com/dimfeld/httptreemux"
	"github.com/volcano822/journey/common/filenames"
	"github.com/volcano822/journey/common/helpers"
	"net/http"
	"path/filepath"
	"strings"
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
	http.ServeFile(w, r, filenames.PagesFilepath+"/about.html")
	return
}

func InitializePages(router *httptreemux.TreeMux) {
	router.GET("/about", aboutHandler)
	// For serving standalone projects or pages saved in in content/pages
	router.GET("/pages/*filepath", pagesHandler)
}
