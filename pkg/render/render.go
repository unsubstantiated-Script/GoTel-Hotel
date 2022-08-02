package render

import (
	"WebGo/pkg/config"
	"WebGo/pkg/models"
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

//// This renders our template
//func RenderTemplateTest(w http.ResponseWriter, tmpl string) {
//	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
//	err := parsedTemplate.Execute(w, nil)
//	if err != nil {
//		fmt.Println("error parsing template:", err)
//	}
//}

//Template caching below

//var templateCache = make(map[string]*template.Template)
//
//func RenderTemplate(w http.ResponseWriter, t string) {
//	var tmpl *template.Template
//	var err error
//
//	// check to see if we already have the template in the cache
//	_, inMap := templateCache[t]
//
//	if !inMap {
//		// Need to create the template
//		log.Println("creating template and adding to cache")
//		err = createTemplateCache(t)
//		if err != nil {
//			log.Println(err)
//		}
//	} else {
//		log.Println("Using cached template")
//	}
//
//	tmpl = templateCache[t]
//
//	err = tmpl.Execute(w, nil)
//	if err != nil {
//		log.Println(err)
//	}
//}
//
//func createTemplateCache(t string) error {
//	templates := []string{
//		fmt.Sprintf("./templates/%s", t),
//		"./templates/base.layout.tmpl",
//	}
//
//	//parse template
//
//	tmpl, err := template.ParseFiles(templates...)
//
//	// checking for errors
//	if err != nil {
//		return err
//	}
//
//	//add template to cache
//
//	templateCache[t] = tmpl
//
//	return nil
//}

var app *config.AppConfig

//NewTemplates sets the config for the config package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var templateCache map[string]*template.Template
	//Using caching on development vs production to allow for hot reloads
	if app.UseCache {
		// get the template cache from the app config

		templateCache = app.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCache()
	}

	// get requested template from cache
	t, ok := templateCache[tmpl]
	if !ok {
		log.Fatal("Could not create template from Template Cache")
	}

	// For finer grained error checking
	buf := new(bytes.Buffer)

	td = AddDefaultData(td)
	
	_ = t.Execute(buf, td)

	// render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all of the files that stat with *.page.tmpl

	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return myCache, err
	}

	//	range through all of the templates
	for _, page := range pages {
		// returns the last part of a file path to give us a page name
		name := filepath.Base(page)
		templateSet, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// adds to the template set if indeed it is a layout
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		// combines all the things together
		if len(matches) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		// Adds final template to this map
		myCache[name] = templateSet
	}
	return myCache, nil
}
