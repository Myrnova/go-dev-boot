package CH9

import "errors"

func getUserMap(names []string, phoneNumbers []int) (map[string]userOne, error) {
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("invalid sizes")
	}

	userMap := make(map[string]userOne)

	for i, name := range names {
		userMap[name] = userOne{
			name,
			phoneNumbers[i],
		}
	}

	return userMap, nil
}

type userOne struct {
	name        string
	phoneNumber int
}
