package main

import (
	"github.com/stokesnetwork/stokes/infrastructure/logger"
	"github.com/stokesnetwork/stokes/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("MNJS")
	spawn      = panics.GoroutineWrapperFunc(log)
)
