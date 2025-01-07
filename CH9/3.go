package CH9

func getCounts(messagedUsers []string, validUsers map[string]int) {
	for _, user := range messagedUsers {
		if _, ok := validUsers[user]; !ok {
			continue
		}
		validUsers[user]++
	}
}