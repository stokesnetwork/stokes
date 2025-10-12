package rpchandlers

import (
	"github.com/Sam-Stokes/stokes/app/appmessage"
	"github.com/Sam-Stokes/stokes/app/rpc/rpccontext"
	"github.com/Sam-Stokes/stokes/infrastructure/network/netadapter/router"
)

// HandleGetCurrentNetwork handles the respectively named RPC command
func HandleGetCurrentNetwork(context *rpccontext.Context, _ *router.Router, _ appmessage.Message) (appmessage.Message, error) {
	response := appmessage.NewGetCurrentNetworkResponseMessage(context.Config.ActiveNetParams.Net.String())
	return response, nil
}
