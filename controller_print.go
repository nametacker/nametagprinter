package nametagprinter

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

type PrintController struct {
	template string
}

func NewPrintController(template string) (c *PrintController) {
	c = new(PrintController)
	c.template = template
	return
}

func (c *PrintController) PrintTagHandler(w http.ResponseWriter, r *http.Request, matches []string) {
	if r.Method != "POST" {
		HttpProblem(w, http.StatusBadRequest, "Expected POST got "+r.Method)
		return
	}
	if !strings.Contains(r.Header.Get("Content-Type"), "application/json") {
		HttpProblem(w, http.StatusBadRequest, "Expected application/json got "+r.Header.Get("Content-Type"))
		return
	}

	tplSource, err := ioutil.ReadFile(c.template)
	if err != nil {
		HttpProblem(w, http.StatusInternalServerError, "Failed to read template")
		log.Println(err.Error())
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		HttpProblem(w, http.StatusBadRequest, "Failed to read body: "+err.Error())
		log.Println(err.Error())
		return
	}

	var data map[string]interface{}
	unmarshalErr := json.Unmarshal(b, &data)
	if unmarshalErr != nil {
		HttpProblem(w, http.StatusBadRequest, "Failed to parse request: "+bytes.NewBuffer(b).String())
		return
	}
	// TODO: read this from nametag template
	fields := []string{"firstname", "lastname", "twitter", "tag1", "tag2", "tag3"}
	for _, field := range fields {
		if data[field] == nil {
			data[field] = ""
		}
	}

	// Write template
	f, err := ioutil.TempFile("", "nametag-")
	if err != nil {
		HttpProblem(w, http.StatusInternalServerError, "Failed to open temp file")
		log.Println(err.Error())
		return
	}
	log.Println("Writing to " + f.Name())
	defer f.Close()
	filename := f.Name()

	tpl := template.Must(template.New("nametag").Parse(string(tplSource)))
	err = tpl.Execute(f, data)
	if err != nil {
		HttpProblem(w, http.StatusInternalServerError, "Failed execute template")
		log.Println(err.Error())
		return
	}
	f.Close()
	os.Rename(filename, filename+".svg")
	svgfile := filename + ".svg"
	pdffile := filename + ".pdf"

	// Convert to pdf
	toPdfCmd := exec.Command("inkscape", "--export-pdf="+pdffile, svgfile)
	err = toPdfCmd.Start()
	if err != nil {
		HttpProblem(w, http.StatusInternalServerError, "Failed to execute inkscape")
		log.Println(err.Error())
	}
	err = toPdfCmd.Wait()
	if err != nil {
		HttpProblem(w, http.StatusInternalServerError, "Failed convert to PDF")
		log.Println(err.Error())
	}
	log.Println(pdffile + " written")

	// Print
	printCmd := exec.Command("lpr", "-P", "Brother_QL-720NW_USB", pdffile)
	err = printCmd.Start()
	if err != nil {
		HttpProblem(w, http.StatusInternalServerError, "Failed send job to printer")
		log.Println(err.Error())
	}
	err = printCmd.Wait()
	if err != nil {
		HttpProblem(w, http.StatusInternalServerError, "Printing failed")
		log.Println(err.Error())
	}
	log.Println(pdffile + " printed")

	w.Header().Add("X-Nametagprinter-Version", VERSION)
}
