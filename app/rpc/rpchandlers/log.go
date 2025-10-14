package rpchandlers

import (
	"github.com/stokesnetwork/stokes/infrastructure/logger"
	"github.com/stokesnetwork/stokes/util/panics"
)

var log = logger.RegisterSubSystem("RPCS")
var spawn = panics.GoroutineWrapperFunc(log)
