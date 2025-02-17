// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package grpc

import (
	grpc "github.com/go-kit/kit/transport/grpc"
	endpoint "github.com/tryu-fullerton-edu/AMSE_CPSC546_Cartelize/accounts/pkg/endpoint"
	pb "github.com/tryu-fullerton-edu/AMSE_CPSC546_Cartelize/accounts/pkg/grpc/pb"
)

// NewGRPCServer makes a set of endpoints available as a gRPC AddServer
type grpcServer struct {
	listAdmins      grpc.Handler
	getAdmin        grpc.Handler
	getAdminByEmail grpc.Handler
	createAdmin     grpc.Handler
}

func NewGRPCServer(endpoints endpoint.Endpoints, options map[string][]grpc.ServerOption) pb.AccountsServer {
	return &grpcServer{
		createAdmin:     makeCreateAdminHandler(endpoints, options["CreateAdmin"]),
		getAdmin:        makeGetAdminHandler(endpoints, options["GetAdmin"]),
		getAdminByEmail: makeGetAdminByEmailHandler(endpoints, options["GetAdminByEmail"]),
		listAdmins:      makeListAdminsHandler(endpoints, options["ListAdmins"]),
	}
}
