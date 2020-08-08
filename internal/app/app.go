package app

import (
	"xsolla-sdk-backend/internal/cache"
	"xsolla-sdk-backend/internal/config"
	"xsolla-sdk-backend/internal/store"

	log "github.com/sirupsen/logrus"
)

type Application struct {
	Store  store.Store
	Cache  cache.Cache
	Config *config.Config
	Log    *log.Entry
}

func NewApplication(storeInstance store.Store,
	cacheInstance cache.Cache,
	configInstance *config.Config,
	logger *log.Entry) Application {
	return Application{
		Store:  storeInstance,
		Cache:  cacheInstance,
		Config: configInstance,
		Log:    logger,
	}
}
