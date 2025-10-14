package main

import (
	"github.com/stokesnetwork/stokes/infrastructure/logger"
	"github.com/stokesnetwork/stokes/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("RORG")
	spawn      = panics.GoroutineWrapperFunc(log)
)
