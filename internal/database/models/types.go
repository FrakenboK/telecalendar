package models

type EventType string

var (
	Disposable  EventType = "DISPOSABLE"
	EveryDay    EventType = "EVERY_DAY"
	EveryWeek   EventType = "EVERY_WEEK"
	EveryMounth EventType = "EVERY_MOUNTH"
	EveryYear   EventType = "EVERY_YEAR"
)
