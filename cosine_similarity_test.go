package main

import "testing"

func TestMatching(t *testing.T) {
	vectorA := Vector{1, 2, 3}
	vectorB := Vector{1, 2, 3}
	if value, error := Matching(vectorA, vectorB); error == nil && value == 1.0 {
		t.Log("test successful")
	}else{
		t.Error("test failed")
	}
}
