package internal

import (
	"github.com/grahamchill/blog_aggregator/internal/config"
	"github.com/grahamchill/blog_aggregator/internal/database"
)

type State struct {
	Db  *database.Queries
	Cfg *config.Config
}
