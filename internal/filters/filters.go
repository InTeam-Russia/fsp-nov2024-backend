package filters

import (
	"github.com/InTeam-Russia/go-backend-template/internal/filters/api"
	"github.com/InTeam-Russia/go-backend-template/internal/filters/filter"
)

type FilterRepository = filter.FilterRepository

var SetupRoutes = api.SetupRoutes

func NewMockFilterRepository() FilterRepository {
	return &filter.MockFilterRepository{}
}

var NewPgFilterRepository = filter.NewPgFilterRepository
