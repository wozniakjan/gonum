package tsp

import (
	"time"

	"gonum.org/v1/gonum/graph"
	//"gonum.org/v1/gonum/graph/traverse"
)

type Algorithm interface {
	Run(TSPIn)
}

type Result struct {
	path []graph.WeightedEdge
	cost float64
}

type TSPOut interface {
	GetResult() (Result, error)
	Finish()
}

type TSPIn interface {
	SendResult(Result)
	SendError(error)
}

type TSPConfig struct {
	Timeout time.Time
	Async   bool
	Algo    Algorithm
}

func TravellingSalesman(cfg TSPConfig) TSPOut {
	in, out := init(cfg)
	cfg.Algo.Run(in)
	return out
}
