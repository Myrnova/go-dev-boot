package CH8

import (
	"sort"
)

type cost struct {
	day   int
	value float64
}

func getCostsByDay(costs []cost) []float64 {
	sort.Slice(costs, func(i, j int) bool {
		return costs[i].day < costs[j].day
	})
	lastValue := costs[len(costs)-1]
	costsByDay := make([]float64, lastValue.day+1)

	for i := 0; i < len(costs); i++ {
		costsByDay[costs[i].day] += costs[i].value
	}
	return costsByDay
}
