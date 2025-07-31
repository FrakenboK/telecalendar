package statestorage

var (
	initState = &UserState{
		State:    startState,
		TempData: make(map[string]interface{}),
	}
)

const (
	startState = "start"
)
