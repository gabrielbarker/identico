// svg/svg.go

package svg

import (
	"fmt"
	"math/rand"
	"strings"
)

type SVG struct {
	header        string
	innerElements []string
}

func New() *SVG {
	return &SVG{header: defaultHeader(), innerElements: []string{}}
}

func (s *SVG) ToString() string {
	inner := strings.Join(s.innerElements, "\n\t")
	tags := []string{s.header, inner, closingTag()}
	return strings.Join(tags, "\n")
}

func (s *SVG) AddElement(element string) {
	s.innerElements = append(s.innerElements, element)
}

func defaultHeader() string {
	return `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 15 16" shape-rendering="crispEdges">`
}

func closingTag() string {
	return `</svg>`
}

func NewRandomSquare(x int, y int) string {
	return newRectElement(randomHexColor(), x, y, 1, 1)
}

func newRectElement(fillHex string, x int, y int, width int, height int) string {
	return fmt.Sprintf(`<rect fill="%v" x="%v" y="%v" width="%v" height="%v" />`, fillHex, x, y, width, height)
}

func randomHexColor() string {
	return fmt.Sprintf("#%x%x%x", rand.Intn(256), rand.Intn(256), rand.Intn(256))
}
