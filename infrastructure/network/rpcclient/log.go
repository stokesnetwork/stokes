package rpcclient

import (
	"github.com/Sam-Stokes/stokes/infrastructure/logger"
	"github.com/Sam-Stokes/stokes/util/panics"
)

var log = logger.RegisterSubSystem("RPCC")
var spawn = panics.GoroutineWrapperFunc(log)
