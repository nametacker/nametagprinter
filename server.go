package nametagprinter

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
)

type route struct {
	re      *regexp.Regexp
	handler func(http.ResponseWriter, *http.Request, []string)
}

type RegexpHandler struct {
	routes []*route
}

func (h *RegexpHandler) AddRoute(re string, handler func(http.ResponseWriter, *http.Request, []string)) {
	r := &route{regexp.MustCompile(re), handler}
	h.routes = append(h.routes, r)
}

func (h *RegexpHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	for _, route := range h.routes {
		matches := route.re.FindStringSubmatch(r.URL.Path)
		if matches != nil {
			route.handler(rw, r, matches)
			break
		}
	}
}

func Serve(c *Config) (err error) {
	log.Println(fmt.Sprintf("Starting server %s on at %s:%d ...", VERSION, c.Server.Address, c.Server.Port))

	pageCntrl := NewPageController()
	tagCntrl := NewTagController()
	printCntrl := NewPrintController(c.Tag.Template)

	reHandler := new(RegexpHandler)
	reHandler.AddRoute("^/api/print$", printCntrl.PrintTagHandler)
	reHandler.AddRoute("^/new$", tagCntrl.NewTagHandler)
	reHandler.AddRoute("^/([^/\\.]+)$", pageCntrl.PageHandler)
	reHandler.AddRoute("^/$", pageCntrl.IndexHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", c.Server.Address, c.Server.Port), reHandler))
	return
}
