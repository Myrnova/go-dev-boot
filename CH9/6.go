package CH9

import (
	"slices"
)

func findSuggestedFriends(username string, friendships map[string][]string) []string {
	var suggestedFriends []string
	userCurrentFriends := friendships[username]
	for _, userCurrentFriend := range userCurrentFriends {
		friendFriendships := friendships[userCurrentFriend]
		for _, friend := range friendFriendships {
			if friend != username &&
				!slices.Contains(userCurrentFriends, friend) &&
				!slices.Contains(suggestedFriends, friend) {
				suggestedFriends = append(suggestedFriends, friend)
			}
		}
	}
	return suggestedFriends
}

//Complete the findSuggestedFriends function. It takes a username string, and a friendships
//map as inputs. The map keys are usernames, and the values are slices of strings representing the direct friends of that user.
//
//Example friendship map:
//
//friendships := map[string][]string{
//"Alice":   {"Bob", "Charlie"},
//"Bob":     {"Alice", "Charlie", "David"},
//"Charlie": {"Alice", "Bob", "David", "Eve"},
//"David":   {"Bob", "Charlie"},
//"Eve":     {"Charlie"},
//}
//Copy icon
//The function should return a slice of strings containing the user's suggested friends.
//A suggested friend is someone who is not a direct friend of the user but is a direct friend of one or more of the user's direct friends.
//Each suggested friend should appear only once in the slice, even if they are found through multiple direct friends.
//
//If there are no suggested friends, return a nil slice.
//
//For example using the friendships map above:
//
//suggestedFriends := findSuggestedFriends("Alice", friendships)
//// suggestedFriends = [David, Eve]
