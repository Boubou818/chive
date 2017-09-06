package handlers

import (
	"github.com/GeertJohan/go.rice"
	"github.com/go-chi/chi"
	"html/template"
	"log"
	"net/http"
	//	"github.com/go-chi/chi/middleware"
	"gopkg.in/mgo.v2"
)

type siteMetas struct {
	SiteName        string
	Title           string // Title of this asset
	Description     string
	ImageURL        string
	URL             string
	TwitterUsername string // including the @
	Type            string // website or article
}

var metas = siteMetas{
	SiteName:    "Chive",
	Title:       "Welcome to Chive",
	Description: "A full stack website boilerplate written in Go and Vue.",
	ImageURL:    "https://i.imgur.com/PComu8U.jpg",
	URL:         "http://localhost:3333/",
	Type:        "website",
}

// Index serves as the anchor for all the handlers based on top-level routes
type Index struct {
	RespFormat  string
	TemplateBox *rice.Box
	Db          *mgo.Database
}

// Routes creates a REST router for the index resource
func (rs Index) Routes() chi.Router {
	r := chi.NewRouter()
	// r.Use() // some middleware..
	//r.Use(middleware.WithValue("respFormat", rs.RespFormat))

	r.Get("/", rs.Home) // GET /

	return r
}

// Home grabs the home page
func (rs Index) Home(w http.ResponseWriter, r *http.Request) {

	// get file contents as string
	indexString, err := rs.TemplateBox.String("index.tpl")
	headerString, err := rs.TemplateBox.String("header.tpl")
	footerString, err := rs.TemplateBox.String("footer.tpl")

	// parse and execute the template
	tmpl, err := template.New("index").Parse(headerString)
	tmpl2, err := template.Must(tmpl.Clone()).Parse(footerString)
	tmpl3, err := template.Must(tmpl2.Clone()).Parse(indexString)
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl3.ExecuteTemplate(w, "index", metas)
	if err != nil {
		// TODO: return an HTTP ERROR internal server error
		log.Println(err)
	}

}
