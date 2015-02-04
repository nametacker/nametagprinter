package nametagprinter

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
	"time"
)

type PageController struct {
	cacheLifetTime time.Duration
	pageTemplates  map[string]*template.Template
	cache          bool
}

func NewPageController() (c *PageController) {
	c = new(PageController)
	c.cacheLifetTime = time.Minute * 30
	c.pageTemplates = make(map[string]*template.Template)
	c.cache = false
	return
}

func (c *PageController) PageHandler(w http.ResponseWriter, r *http.Request, matches []string) {
	w.Header().Add("X-Click-Counter-Iframe-Version", VERSION)
	if r.Method != "GET" {
		w.WriteHeader(400)
		return
	}

	w.Header().Add("Content-Type", "text/html; charset=UTF-8")
	w.Header().Add("Cache-Control", fmt.Sprintf("public, s-maxage=%d", int64(c.cacheLifetTime/time.Second)))
	w.Header().Add("Expires", time.Now().Add(c.cacheLifetTime).Format(http.TimeFormat))
	// w.Header().Add("Last-Modified", domain.Updated.Format(http.TimeFormat))

	var tpl *template.Template
	var data map[string]string
	data = make(map[string]string)
	data["Version"] = VERSION

	tpl, err := c.getPageTemplate(matches[0])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("failed to load template")
		log.Println(err.Error())
		return
	}

	err = tpl.Execute(w, data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("failed to parse template")
		log.Println(err.Error())
		return
	}
}

func (c *PageController) IndexHandler(w http.ResponseWriter, r *http.Request, matches []string) {
	c.PageHandler(w, r, []string{"index"})
}

func (c *PageController) getPageTemplate(name string) (tpl *template.Template, err error) {
	var ok bool
	if tpl, ok = c.pageTemplates[name]; !ok || !c.cache {
		tpl, err = loadTemplate(name, "./templates/"+name+".html")
		if err != nil {
			return
		}
	}
	c.pageTemplates[name] = tpl
	return
}

func loadTemplate(ident string, filename string) (tpl *template.Template, err error) {
	tplSource, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	tpl = template.Must(template.New(ident).Parse(string(tplSource)))
	return
}
