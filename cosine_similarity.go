package main

import (
	"errors"
	"math"
)

type Vector []float64

func Matching(vectorA, vectorB Vector) (float64, error) {
	if vectorA == nil || len(vectorA) < 0 || vectorB == nil || len(vectorB) < 0 {
		return 0.0, errors.New("向量参数异常")
	}

	if len(vectorA) != len(vectorB) {
		return 0.0, errors.New("空间向量 A, B 维度不一致")
	}

	return dotProduct(vectorA, vectorB) / (norm(vectorA) * norm(vectorB)), nil
}

/**
 * 计算向量 A, B 点积
 */
func dotProduct(vectorA, vectorB Vector) float64 {
	dotProduct := 0.0
	for i := 0; i < len(vectorA); i++ {
		dotProduct += vectorA[i] * vectorB[i]
	}
	return dotProduct
}

/**
 * 计算向量模
 */
func norm(vector Vector) float64 {
	return math.Sqrt(dotProduct(vector, vector))
}
