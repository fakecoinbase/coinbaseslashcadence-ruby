// Code generated by thriftrw-plugin-yarpc
// @generated

package workflowservicefx

import (
	workflowserviceclient "github.com/coinbase/cadence-ruby/proxy/gen/cadence/workflowserviceclient"
	fx "go.uber.org/fx"
	yarpc "go.uber.org/yarpc"
	thrift "go.uber.org/yarpc/encoding/thrift"
)

// Params defines the dependencies for the WorkflowService client.
type Params struct {
	fx.In

	Provider yarpc.ClientConfig
}

// Result defines the output of the WorkflowService client module. It provides a
// WorkflowService client to an Fx application.
type Result struct {
	fx.Out

	Client workflowserviceclient.Interface

	// We are using an fx.Out struct here instead of just returning a client
	// so that we can add more values or add named versions of the client in
	// the future without breaking any existing code.
}

// Client provides a WorkflowService client to an Fx application using the given name
// for routing.
//
// 	fx.Provide(
// 		workflowservicefx.Client("..."),
// 		newHandler,
// 	)
func Client(name string, opts ...thrift.ClientOption) interface{} {
	return func(p Params) Result {
		client := workflowserviceclient.New(p.Provider.ClientConfig(name), opts...)
		return Result{Client: client}
	}
}