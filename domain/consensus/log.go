package consensus

import (
	"github.com/Sam-Stokes/stokes/infrastructure/logger"
	"github.com/Sam-Stokes/stokes/util/panics"
)

var log = logger.RegisterSubSystem("BDAG")
var spawn = panics.GoroutineWrapperFunc(log)
