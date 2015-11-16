package nametagprinter

import (
	"net/http"
	"text/template"
	"time"
)

type TagField struct {
	Name        string
	Label       string
	Placeholder string
	Required    bool
}

type TagController struct {
	PageController
}

func NewTagController() (c *TagController) {
	c = new(TagController)
	c.cacheLifetTime = time.Minute * 30
	c.pageTemplates = make(map[string]*template.Template)
	c.cache = false
	return
}

func (c *PageController) NewTagHandler(w http.ResponseWriter, r *http.Request, matches []string) {
	fields := make([]TagField, 0)
	// TODO: read this from nametag template
	fields = append(fields, TagField{"firstname", "Name", "John Doe", true})
	fields = append(fields, TagField{"lastname", "Title", "Founder", false})
	fields = append(fields, TagField{"twitter", "Twitter", "@johndoe", false})
	fields = append(fields, TagField{"tag1", "1. Tag", "#Startups", false})
	fields = append(fields, TagField{"tag2", "2. Tag", "#RheinMainRocks", false})
	fields = append(fields, TagField{"tag3", "3. Tag", "#foobar", false})
	data := make(map[string]interface{})
	data["fields"] = fields
	c.RenderPage(w, r, "new", data)
}
