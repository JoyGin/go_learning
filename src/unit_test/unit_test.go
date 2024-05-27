package unit_test

import "testing"

func TestError(t *testing.T) {
	t.Log("start")
	t.Error("error")
	t.Log("end")
}

func TestFatal(t *testing.T) {
	t.Log("start")
	t.Fatal("fatal")
	t.Log("end")
}
