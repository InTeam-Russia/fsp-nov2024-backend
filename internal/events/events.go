package events

import (
	"github.com/InTeam-Russia/go-backend-template/internal/events/api"
	"github.com/InTeam-Russia/go-backend-template/internal/events/event"
)

type EventRepository = event.EventRepository

var NewMockEventRepository = event.NewMockEventRepository
var SetupRoutes = api.SetupRoutes
var MockEvents = event.MockEvents
