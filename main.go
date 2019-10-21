package main

import (
	"io"
	"os"

	"github.com/JQuiroz728/imageTransform/primitive"
)

func main() {
	file, err := os.Open("chavy.jpg")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	out, err := primitive.Transform(file, 50)
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, out)

}
