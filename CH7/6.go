package CH7

/*
Assignment
Complete the countConnections function that takes an integer groupSize representing the number of people in the group chat and
returns an integer representing the number of connections between them. For each additional person in the group, the number of new connections continues to grow.
Use a for loop to accumulate the number of connections instead of directly using a mathematical formula.
*/
func countConnections(groupSize int) int {
	var connections int
	for i := 1; i < groupSize; i++ {
		connections += i
	}

	return connections
}