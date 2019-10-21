package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"path/filepath"

	"github.com/JQuiroz728/imageTransform/primitive"
)

func main() {
	// Set up web server
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		html :=
			`<html>
			<body>
				<form action="/upload" method="post" enctype="multipart/form-data">
				<input type="file" name="image">
				<button type="submit">Upload Image</button>
				</form>
			</body>
		</html>`
		fmt.Fprint(w, html)
	})

	mux.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		file, header, err := r.FormFile("image")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		ext := filepath.Ext(header.Filename)[1:]
		out, err := primitive.Transform(file, ext, 50)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		// Validate proper image formats
		switch ext {
		case "jpg":
			fallthrough
		case "jpeg":
			w.Header().Set("Content-Type", "image/jpeg")
		case "png":
			w.Header().Set("Content-Type", "image/png")
		default:
			http.Error(w, "Invalid Image Format (Please use .png or .jpg)", http.StatusInternalServerError)
			return
		}
		io.Copy(w, out)
	})
	fileServer := http.FileServer(http.Dir("./img/"))
	mux.Handle("/img/", http.StripPrefix("/img", fileServer))
	log.Fatal(http.ListenAndServe(":3000", mux))
}
