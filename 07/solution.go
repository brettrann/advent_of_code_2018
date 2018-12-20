package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func buildGraph(lines []string) (*graph) {
	graph := &graph{heads: make(map[string]*step), steps: make(map[string]*step)}
	for _, line := range lines {
		var beforeID = string(line[5])
		var afterID = string(line[36])
		graph.Add(beforeID, afterID)
	}
	return graph
}

//
// Graph type+funcs
//
type graph struct {
	heads map[string]*step
	steps map[string]*step
}
func (g graph) String() string {
	return fmt.Sprintf("{heads: %v, steps: %v}", g.heads, g.steps)
}
func (g *graph) Add(beforeID string, afterID string) {
	// get or create
	bstep, ok := g.Get(beforeID); if !ok {
		bstep = g.Createstep(beforeID)
		g.heads[bstep.id] = bstep
	}
	astep, ok := g.Get(afterID); if !ok {
		astep = g.Createstep(afterID)
	}

	// if we're an after step we can't be a head step.
	delete(g.heads, astep.id)

	// bstep.children = append(bstep.children, astep)
	bstep.AddChild(astep)
	astep.AddParent(bstep)
}
func (g graph) Get(id string) (step *step, ok bool) {
	step, ok = g.steps[id];
	return step, ok
}
func (g *graph) Createstep(id string) *step {
	n := &step{id: id}
	g.steps[id] = n
	return n
}
// Solve to determine the order of the steps
// graph flows left to right, with parents->children.
// there may be multiple head parents to choose from.
// each step can only be done if its parents are done.
// when multiple steps can be chosen next, choose alphabetically.
func (g *graph) SolveString() string {

	var stepOrder strings.Builder
	candidates := make([]*step, 0, len(g.heads))

	for _, v := range g.heads {
		candidates = append(candidates, v)
	}

	for ; len(candidates) > 0; {
		// from multiple candidates use alphabetical order
		sort.Slice(candidates, func(i, j int) bool {
			return candidates[i].id < candidates[j].id
		})

		// do the step and add children if the child's parents are done.
		var complete *step
		complete, candidates = candidates[0], candidates[1:]
		complete.done = true
		stepOrder.WriteString(complete.id)
		for _, nextStep := range complete.children {
			if nextStep.doable() {
				candidates = append(candidates, nextStep)
			}
		}
	}

	return stepOrder.String()
}

//
// step type+funcs
//
type steps []*step
type step struct {
	parents steps
	children steps
	done bool
	id string
}
func (n step) String() string {
	return fmt.Sprintf("{id:%v, parents:%d, children:%d}", n.id, len(n.parents), len(n.children))
}
func (n *step) AddChild(child *step) {
	n.children = append(n.children, child)
}
func (n *step) AddParent(parent *step) {
	n.parents = append(n.parents, parent)
}
// A step is doable if its parents are done.
func (n *step) doable() bool {
	doable := true
	for _, s := range n.parents {
		if s.done == false {
			doable = false
			break
		}
	}
	return doable
}

func in(path string) ([]string) {
	content, _ := ioutil.ReadFile(path)
	lines := strings.Split(string(content), "\n")
	lines = lines[:len(lines)-1]
	return lines
}

func main() {
	lines := in("input.txt")
	graph := buildGraph(lines)
	fmt.Printf("graph: %v\n", graph)
	fmt.Printf("part1 solution: %v\n", graph.SolveString())
}
