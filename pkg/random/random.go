/*
 * @Descripttion:
 * @version:
 * @Author: yyeiei
 * @Date: 2019-08-12 14:22:56
 * @LastEditors: yyeiei
 * @LastEditTime: 2019-08-16 16:00:07
 */
package random

import (
	"math/rand"
	"sort"
)

func String(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(Int(65, 90)) // A-Z
	}
	return string(bytes)
}

func Int(min, max int) int {
	return min + rand.Intn(max-min)
}

func Float64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func Float64WithWeight(weights []float64) float64 {
	if weights == nil || len(weights) == 0 {
		return 0
	}
	if len(weights) == 1 {
		return weights[0]
	}

	var sum float64
	for _, v := range weights {
		sum += v
	}

	sort.Float64s(weights)
	rand := Float64(0, sum)
	index := 0
	for i, v := range weights {
		if rand <= v {
			index = i
			break
		}
		rand -= v
	}
	return weights[index]
}
