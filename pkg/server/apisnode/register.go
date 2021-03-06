package apisnode

import (
	"github.com/emicklei/go-restful"

	"github.com/baidu/ote-stack/pkg/server/apisnode/resource"
	"github.com/baidu/ote-stack/pkg/server/handler"
	"github.com/baidu/ote-stack/pkg/util"
)

const (
	ServePath = "/apis/node.k8s.io/v1beta1"
)

func NewServiceHandler(ctx *handler.HandlerContext) *handler.ServiceHandler {
	serverHandler := map[string]handler.Handler{}

	// TODO add needed server handler here.
	serverHandler[util.ResourceRuntimeClass] = resource.NewRuntimeClassHandler(ctx)

	return &handler.ServiceHandler{
		ServerHandler: serverHandler,
		Ctx:           ctx,
	}
}

func NewWebsService(ctx *handler.HandlerContext) *restful.WebService {
	serviceHandler := NewServiceHandler(ctx)

	ws := new(restful.WebService)
	ws.Path(ServePath)
	// Register create handler
	serviceHandler.RegisterCreateHandler(ws)
	serviceHandler.RegisterDeleteHandler(ws)
	serviceHandler.RegisterListHandler(ws)
	serviceHandler.RegisterUpdateHandler(ws)

	return ws
}
