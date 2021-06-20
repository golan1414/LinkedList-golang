package main

import (
	"errors"
	"fmt"
	. "linkedList/LinkedList"
)

func main() {
	var g Graph
	for i := 0; i < 7; i++ {
		g.AddVertex()
	}
	g.AddEdge(0, 1)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(1, 4)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	g.AddEdge(4, 5)
	g.ShortestPathsBFS(0)
	fmt.Println(g.vertices[0].distances)
	g.ShortestPathsBFS(1)
	fmt.Println(g.vertices[1].distances)
	g.ShortestPathsBFS(2)
	fmt.Println(g.vertices[2].distances)
	g.ShortestPathsBFS(4)
	fmt.Println(g.vertices[4].distances)
	g.ShortestPathsBFS(6)
	fmt.Println(g.vertices[6].distances)

}

type Vertex struct {
	neighbors []int
	distances map[int]int
	id        int
}

type Graph struct{ vertices []Vertex }

func (g *Graph) AddVertex() {
	g.vertices = append(g.vertices, Vertex{
		neighbors: []int{},
		distances: make(map[int]int),
		id:        len(g.vertices),
	})
}

func (g *Graph) AddEdge(first, second int) error {
	if first < 0 || second < 0 || first >= len(g.vertices) || second >= len(g.vertices) {
		return errors.New("either first or second are invalid vertex id")
	}
	if contains(g.vertices[first].neighbors, second) || contains(g.vertices[second].neighbors, first) {
		return errors.New("cant create the same edge twice")
	}
	g.vertices[first].neighbors = append(g.vertices[first].neighbors, second)
	g.vertices[second].neighbors = append(g.vertices[second].neighbors, first)
	return nil
}

func (g *Graph) ShortestPathsBFS(source int) error {
	if source < 0 || source > len(g.vertices) {
		return errors.New(fmt.Sprintf("vertex number %d does not exist", source))
	}
	previous := make(map[int]int)
	v := g.vertices[source]
	for i := 0; i < len(g.vertices); i++ {
		v.distances[i] = -1
		previous[i] = -1
	}
	v.distances[source] = 0
	queue := NewLinkedList()
	curVertex := v
	var tmp interface{}
	var err error
	var ok bool
	queue.PushBack(v)
	for queue.Len() > 0 {
		tmp, err = queue.PopFront()
		if err != nil {
			return errors.New("error while popping")
		}
		curVertex, ok = tmp.(Vertex)
		if !ok {
			return errors.New("error while converting to vertex")
		}
		for _, u := range curVertex.neighbors {
			if v.distances[u] == -1 {
				queue.PushBack(g.vertices[u])
				v.distances[u] = 1 + v.distances[curVertex.id]
			}
		}
	}
	return nil
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
