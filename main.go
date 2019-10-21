package main

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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
		a, err := generateImage(file, ext, 100, primitive.ModeTriangle)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		file.Seek(0, 0)
		b, err := generateImage(file, ext, 100, primitive.ModeRect)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		file.Seek(0, 0)
		c, err := generateImage(file, ext, 100, primitive.ModeEllipse)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		file.Seek(0, 0)
		d, err := generateImage(file, ext, 100, primitive.ModeCircle)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		file.Seek(0, 0)
		e, err := generateImage(file, ext, 100, primitive.ModeRotatedRect)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		file.Seek(0, 0)
		f, err := generateImage(file, ext, 100, primitive.ModeBeziers)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		file.Seek(0, 0)
		g, err := generateImage(file, ext, 100, primitive.ModeRotatedEllipse)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		file.Seek(0, 0)
		h, err := generateImage(file, ext, 100, primitive.ModePolygon)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		file.Seek(0, 0)
		i, err := generateImage(file, ext, 100, primitive.ModeCombo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		html :=
			`<html><body>
			{{range .}}
			<img src="/{{.}}">
			{{end}}
		</body></html>`
		templte := template.Must(template.New("").Parse(html))
		images := []string{a, b, c, d, e, f, g, h, i}
		templte.Execute(w, images)
	})

	fileServer := http.FileServer(http.Dir("./img/"))
	mux.Handle("/img/", http.StripPrefix("/img", fileServer))
	log.Fatal(http.ListenAndServe(":3000", mux))
}

func tempFile(prefix, ext string) (*os.File, error) {
	in, err := ioutil.TempFile("./img/", prefix)
	if err != nil {
		return nil, err
	}
	defer os.Remove(in.Name())
	return os.Create(fmt.Sprintf("%s.%s", in.Name(), ext))
}

func generateImage(read io.Reader, ext string, numShapes int, mode primitive.Mode) (string, error) {
	out, err := primitive.Transform(read, ext, numShapes, primitive.WithMode(mode))
	if err != nil {
		return "", err
	}
	outputFile, err := tempFile("", ext)
	if err != nil {
		return "", err
	}
	defer outputFile.Close()
	io.Copy(outputFile, out)
	return outputFile.Name(), nil
}
