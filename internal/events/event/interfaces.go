package event

type EventRepository interface {
	GetIndicatorsByFilter(filter *EventFilter) (*Indicators, error)
	GetEventsByFilter(filter *EventFilter) ([]EventOut, error)
}
