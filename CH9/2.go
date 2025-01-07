package CH9

import "errors"

func deleteIfNecessary(users map[string]userTwo, name string) (deleted bool, err error) {
	elem, ok := users[name]
	if !ok {
		return false, errors.New("not found")
	}
	if elem.scheduledForDeletion != true {
		return false, nil
	}
	delete(users, name)

	return true, nil
}

type userTwo struct {
	name                 string
	number               int
	scheduledForDeletion bool
}
