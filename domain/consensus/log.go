package consensus

import (
	"github.com/stokesnetwork/stokes/infrastructure/logger"
	"github.com/stokesnetwork/stokes/util/panics"
)

var log = logger.RegisterSubSystem("BDAG")
var spawn = panics.GoroutineWrapperFunc(log)
