package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	out, err := primitive("chavy.jpg", "out.png", 100, rotatedEllipse)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)

}

type Mode int

const (
	combo Mode = iota
	triangle
	rect
	ellipse
	circle
	rotatedRect
	beziers
	rotatedEllipse
	polygon
)

func primitive(inputFile, outputFile string, numShapes int, mode Mode) (string, error) {
	argsStr := fmt.Sprintf("-i %s -o %s -n %d -m %d", inputFile, outputFile, numShapes, mode)
	cmd := exec.Command("primitive", strings.Fields(argsStr)...)
	b, err := cmd.CombinedOutput()
	return string(b), err
}
