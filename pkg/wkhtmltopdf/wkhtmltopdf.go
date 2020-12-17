package wkhtmltopdf

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	wkhtmltopdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/rs/xid"
)

// Generate ...
func Generate(location, res string) (err error) {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		fmt.Println(1)
		return err
	}

	f, err := os.Open(location)
	if f != nil {
		defer f.Close()
	}
	if err != nil {
		return errors.New("HTML file not found")
	}

	pdfg.AddPage(wkhtmltopdf.NewPageReader(f))
	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)
	pdfg.Dpi.Set(300)

	err = pdfg.Create()
	if err != nil {
		return err
	}

	err = pdfg.WriteFile(res)
	if err != nil {
		return errors.New("Save pdf")
	}

	return err
}

// GenerateOld ...
func GenerateOld() (err error) {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return err
	}

	htmlFile, err := ioutil.ReadFile("../html_template/1.html")
	if err != nil {
		return err
	}

	// Replace value
	html := string(htmlFile)
	html = strings.Replace(html, "Lorem", "HI", 1)

	newPath := "../html_template/" + xid.New().String() + ".html"
	err = ioutil.WriteFile(newPath, []byte(html), 0644)
	if err != nil {
		return err
	}

	f, err := os.Open(newPath)
	if f != nil {
		defer f.Close()
	}
	if err != nil {
		return errors.New("HTML file not found")
	}

	pdfg.AddPage(wkhtmltopdf.NewPageReader(f))

	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)
	pdfg.Dpi.Set(300)

	err = pdfg.Create()
	if err != nil {
		return err
	}

	err = pdfg.WriteFile("../static/output.pdf")
	if err != nil {
		return errors.New("Save pdf")
	}

	os.Remove(newPath)

	return err
}
