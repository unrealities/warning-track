package main

import "fmt"

type gameInfo struct {
	Id     string
	Status string
	Li     float64
}

func (g gameInfo) String() string {
	return fmt.Sprintf("%s: %-1.1f", g.Id, g.Li)
}

type ByLi []gameInfo

func (g ByLi) Len() int           { return len(g) }
func (g ByLi) Swap(i, j int)      { g[i], g[j] = g[j], g[i] }
func (g ByLi) Less(i, j int) bool { return g[i].Li > g[j].Li }
