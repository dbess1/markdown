package main


/* To make sure the application from our last chapter will work on Heroku, 
we will need to make a few changes. Heroku gives us a PORT environment 
variable and expects our web application to bind to it. Let's start by 
importing the "os" package so we can grab that PORT environment variable: */

import (
	"net/http"
	"os"
	"github.com/russross/blackfriday"
)

func main() {

/* Next, we need to grab the PORT environment variable, check if it is set, 
	and if it is we should bind to that instead of our hardcoded port (8080).*/

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

/* Lastly, we want to bind to that port in our http.ListenAndServe call: */
	
	http.HandleFunc("/markdown", generateMarkdown)
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":" + port, nil)
}

func generateMarkdown(rw http.ResponseWriter, r *http.Request) {
	markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
	rw.Write(markdown)
}
	
