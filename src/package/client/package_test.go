package client

import (
	"go_learning/src/package/series"
	"testing"
)

func TestPackage(t *testing.T) {
	t.Log(series.GetFibonacciSeries(5))
	//t.Log(series.square(1))
}
