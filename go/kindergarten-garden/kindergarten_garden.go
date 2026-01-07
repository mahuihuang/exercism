package kindergarten

import (
	"errors"
	"slices"
	"strings"
)

// Define the Garden type here.

// The diagram argument starts each row with a '\n'.  This allows Go's
// raw string literals to present diagrams in source code nicely as two
// rows flush left, for example,
//
//	diagram := `
//	VVCCGG
//	VVCCGG`
type Garden struct {
	FirstRow  string
	SecondRow string
	Children  []string
}

func NewGarden(diagram string, children []string) (*Garden, error) {
	if !strings.HasPrefix(diagram, "\n") {
		return nil, errors.New("diagram isn't start with \n")
	}
	rows := formatDiagram(diagram)
	if rows == nil {
		return nil, errors.New("invalid diagram")
	}
	if !validChildren(children) {
		return nil, errors.New("invalid children list")
	}
	return &Garden{
		FirstRow:  rows[0],
		SecondRow: rows[1],
		Children:  children,
	}, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	if !slices.Contains(g.Children, child) {
		return nil, false
	}

	plants := []string{}

	childrenIndex := slices.Index(slices.Sorted(slices.Values(g.Children)), child)
	if childrenIndex == -1 {
		return nil, false
	}

	diagramIndex := childrenIndex * 2
	codes := g.FirstRow[diagramIndex:diagramIndex+2] + g.SecondRow[diagramIndex:diagramIndex+2]
	for _, code := range codes {
		plants = append(plants, translateDiagramCode(code))
	}
	return plants, true
}

func formatDiagram(diagram string) []string {
	rows := strings.Split(strings.TrimPrefix(diagram, "\n"), "\n")
	if len(rows[0]) != len(rows[1]) {
		return nil
	}
	for _, row := range rows {
		if len(row)%2 != 0 {
			return nil
		}
		for _, seed := range row {
			if translateDiagramCode(seed) == "" {
				return nil
			}
		}
	}
	return rows
}

func translateDiagramCode(code rune) string {
	switch code {
	case 'G':
		return "grass"
	case 'C':
		return "clover"
	case 'R':
		return "radishes"
	case 'V':
		return "violets"
	default:
		return ""
	}
}
func validChildren(children []string) bool {
	childrenCount := map[string]int{}
	for _, c := range children {
		if _, ok := childrenCount[c]; ok {
			return false
		} else {
			childrenCount[c]++
		}
	}
	return true
}
