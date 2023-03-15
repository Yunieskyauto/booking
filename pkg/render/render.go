package render

import (
	"bytes"
	"html/template"
	"log"
	"github.com/yunieskyauto/booking/pkg/config"
	"github.com/yunieskyauto/booking/pkg/models"
	"net/http"
	"path/filepath"
)

var app *config.AppConfig

// Newtemplate sets the config for the template
func NewTemplate(a *config.AppConfig) {
	app = a
}

// Rendertemplate renders a template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTempalteCache()
	}

	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTempalteCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all of the files named *.page.html from ./templates
	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	// range through all files ending with *.page.html
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		// until here I gat the name of the file I l currently at exp(home.page.html)
		if err != nil {
			return myCache, err
		}

		maches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(maches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}
	return myCache, nil
}

/*

var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	//checl  to see if we already have the template in owr cache
	_, inMap := tc[t]
	if !inMap {
		// need to create a new template
		log.Println("creating template and adding to chach")
		err = createTempalteCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		// we have the template in the cache
		log.Println("using cached template")
	}

	tmpl = tc[t]

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func createTempalteCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.html",
	}

	//parse template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	//add template
	tc[t] = tmpl
	return nil
}
*/
