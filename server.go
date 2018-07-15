package main

import (
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

var (
	static = http.FileServer(http.Dir("static"))

	pathOld = regexp.MustCompile(`^/([0-9]{1,10})(/.*)?$`)
	pathNew = regexp.MustCompile(`^/([0-9a-f]{16})(/.*)?$`)
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
		w.Write(paste)
	case "/json":
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
	if m := pathNew.FindStringSubmatch(r.URL.Path); m != nil {
		src, err := hex.DecodeString(m[1])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		dst := make([]byte, 8)
		cipher.Decrypt(dst, src)

		id := binary.BigEndian.Uint64(dst)
		servePaste(w, id, m[2])
	} else if m := pathOld.FindStringSubmatch(r.URL.Path); m != nil {
		id, err := strconv.ParseUint(m[1], 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if id >= 256 {
			id = (id * 0x7FFFFFFF) % 0x100000000
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

		src := make([]byte, 8)
		binary.BigEndian.PutUint64(src, uint64(id))

		dst := make([]byte, 8)
		cipher.Encrypt(dst, src)

		http.Redirect(w, r, hex.EncodeToString(dst), http.StatusSeeOther)
	} else {
		static.ServeHTTP(w, r)
	}
}

func main() {
	http.HandleFunc("/", handle)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
