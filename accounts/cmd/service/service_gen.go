// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package service

import (
	endpoint1 "github.com/go-kit/kit/endpoint"
	log "github.com/go-kit/kit/log"
	prometheus "github.com/go-kit/kit/metrics/prometheus"
	opentracing "github.com/go-kit/kit/tracing/opentracing"
	grpc "github.com/go-kit/kit/transport/grpc"
	group "github.com/oklog/oklog/pkg/group"
	opentracinggo "github.com/opentracing/opentracing-go"
	endpoint "github.com/tryu-fullerton-edu/AMSE_CPSC546_Cartelize/accounts/pkg/endpoint"
	service "github.com/tryu-fullerton-edu/AMSE_CPSC546_Cartelize/accounts/pkg/service"
)

func createService(endpoints endpoint.Endpoints) (g *group.Group) {
	g = &group.Group{}
	initGRPCHandler(endpoints, g)
	return g
}
func defaultGRPCOptions(logger log.Logger, tracer opentracinggo.Tracer) map[string][]grpc.ServerOption {
	options := map[string][]grpc.ServerOption{
		"CreateAdmin":     {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "CreateAdmin", logger))},
		"GetAdmin":        {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "GetAdmin", logger))},
		"GetAdminByEmail": {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "GetAdminByEmail", logger))},
		"ListAdmins":      {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "ListAdmins", logger))},
	}
	return options
}
func addDefaultEndpointMiddleware(logger log.Logger, duration *prometheus.Summary, mw map[string][]endpoint1.Middleware) {
	mw["ListAdmins"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "ListAdmins")), endpoint.InstrumentingMiddleware(duration.With("method", "ListAdmins"))}
	mw["GetAdmin"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "GetAdmin")), endpoint.InstrumentingMiddleware(duration.With("method", "GetAdmin"))}
	mw["GetAdminByEmail"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "GetAdminByEmail")), endpoint.InstrumentingMiddleware(duration.With("method", "GetAdminByEmail"))}
	mw["CreateAdmin"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "CreateAdmin")), endpoint.InstrumentingMiddleware(duration.With("method", "CreateAdmin"))}
}
func addDefaultServiceMiddleware(logger log.Logger, mw []service.Middleware) []service.Middleware {
	return append(mw, service.LoggingMiddleware(logger))
}
func addEndpointMiddlewareToAllMethods(mw map[string][]endpoint1.Middleware, m endpoint1.Middleware) {
	methods := []string{"ListAdmins", "GetAdmin", "GetAdminByEmail", "CreateAdmin"}
	for _, v := range methods {
		mw[v] = append(mw[v], m)
	}
}
