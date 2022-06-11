package graph_test

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"github.com/qulia/go-qulia/lib/graph"
)

var dotExe string

func dotToImageGraphviz(fileName string, format string, dot []byte) (string, error) {
	if fileName == "" {
		return "", errors.New("empty file name")
	}
	if dotExe == "" {
		dot, err := exec.LookPath("dot")
		if err != nil {
			return "", err
		}
		dotExe = dot
	}

	img := fmt.Sprintf("%s.%s", fileName, format)
	cmd := exec.Command(dotExe, fmt.Sprintf("-T%s", format), "-o", img)
	cmd.Stdin = bytes.NewReader(dot)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return "", err
	}
	return img, nil
}

// Return Dot string of the graph, can be used with Graphviz(https://graphviz.org/) to visualize
func graphToDot[T comparable](g graph.Graph[T]) string {
	sb := strings.Builder{}

	sb.WriteString(`strict digraph {`)
	sb.WriteString("\n")

	nodes := g.GetNodes().ToSlice()
	for _, n := range nodes {
		WriteDot(n, g, &sb)
	}

	sb.WriteString(`}`)

	return sb.String()
}

func WriteDot[T comparable](n T, g graph.Graph[T], sb *strings.Builder) {
	sb.WriteString(fmt.Sprintf("%v", n))
	sb.WriteString("\n")
	for target := range g.Adjacencies(n) {
		sb.WriteString(fmt.Sprintf("%v -> %v\n", n, target))
	}
}
