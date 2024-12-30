package CH7

func bulkSend(numMessages int) float64 {
	var cost float64
	for i := 0; i < numMessages; i++ {
		cost += 1.0 + float64(i)/100
	}

	return cost
}
