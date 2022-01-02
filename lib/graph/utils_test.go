package graph_test

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

var dotExe string

func dotToImageGraphviz(fileName string, format string, dot []byte) (string, error) {
	if fileName == "" {
		log.Fatal("Provide fileName for the image")
	}
	if dotExe == "" {
		dot, err := exec.LookPath("dot")
		if err != nil {
			log.Fatalln("unable to find program 'dot', please install it or check your PATH")
		}
		dotExe = dot
	}

	img := fmt.Sprintf("%s.%s", fileName, format)
	cmd := exec.Command(dotExe, fmt.Sprintf("-T%s", format), "-o", img)
	cmd.Stdin = bytes.NewReader(dot)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("command '%v': %v\n%v", cmd, err, stderr.String())
	}
	return img, nil
}
