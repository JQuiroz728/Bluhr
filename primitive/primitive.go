package primitive

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

// Mode - User defined shapes for transforming images
type Mode int

// Supported modes by primitive
const (
	ModeCombo Mode = iota
	ModeTriangle
	ModeRect
	ModeEllipse
	ModeCircle
	ModeRotatedRect
	ModeBeziers
	ModeRotatedEllipse
	ModePolygon
)

// WithMode will provide option for transform func that will define desired mode. Default mode = triangle
func WithMode(mode Mode) func() []string {
	return func() []string {
		return []string{"-m", fmt.Sprintf("%d", mode)}
	}
}

// Transform will take provided image and apply transformation, returns reader to the result
func Transform(image io.Reader, ext string, numShapes int, options ...func() []string) (io.Reader, error) {
	var args []string
	for _, option := range options {
		args = append(args, option()...)
	}
	// input
	in, err := tempFile("in_", ext)
	if err != nil {
		return nil, err
	}
	defer os.Remove(in.Name())
	// output
	out, err := tempFile("out_", ext)
	if err != nil {
		return nil, err
	}
	defer os.Remove(out.Name())
	// read image into file
	_, err = io.Copy(in, image)
	if err != nil {
		return nil, err
	}
	stdCombo, err := primitive(in.Name(), out.Name(), numShapes, args...)
	if err != nil {
		return nil, err
	}
	fmt.Println(stdCombo)

	b := bytes.NewBuffer(nil)
	_, err = io.Copy(b, out)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func tempFile(prefix, ext string) (*os.File, error) {
	in, err := ioutil.TempFile("", prefix)
	if err != nil {
		return nil, err
	}
	defer os.Remove(in.Name())
	return os.Create(fmt.Sprintf("%s.%s", in.Name(), ext))
}

func primitive(inputFile, outputFile string, numShapes int, args ...string) (string, error) {
	argsStr := fmt.Sprintf("-i %s -o %s -n %d", inputFile, outputFile, numShapes)
	args = append(strings.Fields(argsStr), args...)
	cmd := exec.Command("primitive", args...)
	b, err := cmd.CombinedOutput()
	return string(b), err
}
