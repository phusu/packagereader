package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"packagereaderproject/packagereader"
)

var pr *packagereader.PackageInfoReader
var htmltemplates = template.Must(template.ParseFiles("htmltemplates/index.html", "htmltemplates/package.html"))

// Handles all HTML requests. Parses the URL and based on that
// decides if it serves the index.html template or package.html template.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Path[len("/"):]
	if len(page) > 1 {
		packages := pr.Packages()
		elem, ok := packages[page]
		if !ok {
			fmt.Println("Package doesn't exist, redirecting to index")
			http.Redirect(w, r, "/", http.StatusFound)
			fmt.Println("Request served.")
			return
		}
		err := htmltemplates.ExecuteTemplate(w, "package.html", elem)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		fmt.Println("Request served.")
		return
	}
	err := htmltemplates.ExecuteTemplate(w, "index.html", pr.Packages())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Println("Request served.")
}

func main() {
	// Parse command line arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: packagereaderproject <filename>")
		fmt.Println("Filename: location of the package status file, e.g. /var/lib/dpkg/status.")
		return
	}
	fileName := os.Args[1]

	// Read the packages
	fmt.Println("Starting.")
	pr = packagereader.NewPackageInfoReader()
	pr.ParseFile(fileName)
	fmt.Println("File parsed.")

	// Register handler
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("./css"))))
	http.HandleFunc("/", indexHandler)

	// Start web server
	fmt.Println("Starting web server.")
	http.ListenAndServe(":8080", nil)
}
