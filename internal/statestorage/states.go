package statestorage

var initState = &UserState{
	State:    "start",
	TempData: make(map[string]interface{}),
}
