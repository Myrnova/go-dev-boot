package CH7

func maxMessages(thresh int) int {
	auxThresh := thresh
	var count int
	for i := 0; ; i++ {
		auxThresh -= 100 + i
		if auxThresh < 0 {
			break
		}
		count++
	}
	return count
}
