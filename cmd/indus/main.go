package main

import (
	"fmt"
	"log"

	"framagit.org/andinus/indus/clipboard"
	"framagit.org/andinus/indus/fetch"
	"framagit.org/andinus/indus/summarize"
)

func main() {
	// Get the primary clipboard selection.
	sel, err := clipboard.GetSel()
	if err != nil {
		log.Fatal(err)
	}

	// If the selection is an empty string then the program should
	// exit because otherwise Notify function will fail with
	// non-zero exit code.
	if len(sel) == 0 {
		err = fmt.Errorf("clipboard: primary clipboard empty")
		log.Fatal(err)
	}

	notifyWiki(sel)
}

func notifyWiki(sel string) {
	// Fetch the summary from wikipedia.
	w, err := fetch.Wikipedia(sel)
	if err != nil {
		log.Fatal(err)
	}

	// Get notification information.
	n, err := summarize.Wikipedia(w)
	if err != nil {
		log.Fatal(err)
	}

	// Send a desktop notification to the user.
	err = n.Notify()
	if err != nil {
		log.Fatal(err)
	}
}
