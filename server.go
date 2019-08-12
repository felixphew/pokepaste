package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"regexp"
)

var (
	assets = http.FileServer(http.Dir("assets"))

	path    = regexp.MustCompile(`^/([0-9a-f]{16})(/.*)?$`)
	pathOld = regexp.MustCompile(`^/([0-9]{1,10})(/.*)?$`)
)

func servePaste(w http.ResponseWriter, id uint64, p string) {
	paste, title, author, notes, err := getPaste(id)
	if err != nil {
		http.NotFound(w, nil)
		return
	}

	switch p {
	case "", "/":
		renderPaste(w, paste, title, author, notes)
	case "/raw":
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write(paste)
	case "/json":
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"paste":  string(paste),
			"title":  string(title),
			"author": string(author),
			"notes":  string(notes),
		})
	default:
		http.NotFound(w, nil)
	}
}

func handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Strict-Transport-Security", "max-age=31536000")
	w.Header().Set("Content-Security-Policy", "default-src 'none'; font-src 'self'; img-src 'self'; script-src 'self'; style-src 'self'; base-uri 'none'; form-action 'self'; frame-ancestors 'none'")

	if m := path.FindStringSubmatch(r.URL.Path); m != nil {
		id, err := decodeID(m[1])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		servePaste(w, id, m[2])
	} else if m := pathOld.FindStringSubmatch(r.URL.Path); m != nil {
		id, err := decodeOldID(m[1])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if id >= 1000 {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		servePaste(w, id, m[2])
	} else if r.URL.Path == "/create" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		paste := r.Form.Get("paste")
		if len(paste) == 0 {
			http.Error(w, "No (or Invalid) Paste", http.StatusBadRequest)
			return
		}

		title := r.Form.Get("title")
		author := r.Form.Get("author")
		notes := r.Form.Get("notes")

		id, err := postPaste(&paste, &title, &author, &notes)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/"+encodeID(id), http.StatusSeeOther)
	} else {
		assets.ServeHTTP(w, r)
	}
}

func main() {
	http.HandleFunc("/", handle)

	log.Fatal(http.ListenAndServe(os.Getenv("ADDR"), nil))
}
