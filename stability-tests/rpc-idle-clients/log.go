package main

import (
	"github.com/Sam-Stokes/stokes/infrastructure/logger"
	"github.com/Sam-Stokes/stokes/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("RPIC")
	spawn      = panics.GoroutineWrapperFunc(log)
)
