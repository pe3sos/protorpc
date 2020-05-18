package protorpc

import (
	"context"
	"fmt"
	"reflect"

	"google.golang.org/protobuf/proto"
)

type MethodInfo struct {
	Name           string
	IsClientStream bool
	IsServerStream bool
}

type ServiceInfo struct {
	Methods []MethodInfo
	About   interface{}
}

type MethodMap struct {
	Name    string
	Handler methodHandler
}

type methodHandler func(service interface{}, ctx context.Context, payload proto.Message) (proto.Message, error)

type ServiceDescriptor struct {
	Name        string
	HandlerType interface{}
	Methods     []MethodMap
	About       interface{}
}

func BuildService(sd *ServiceDescriptor, implementation interface{}) (*RPCDispatcher, error) {
	requiredType := reflect.TypeOf(sd.HandlerType).Elem()
	givenType := reflect.TypeOf(implementation)

	if !givenType.Implements(requiredType) {
		return nil, fmt.Errorf("wsrpc: Server.RegisterService found the handler of type %v that does not satisfy %v", givenType, requiredType)
	}
	srv := &RPCDispatcher{
		descriptor:     sd,
		implementation: implementation,
		methods:        make(map[string]*MethodMap),
		about:          sd.About,
	}
	for i := range sd.Methods {
		method := &sd.Methods[i]
		srv.methods[method.Name] = method
	}
	return srv, nil
}

type RPCDispatcher struct {
	descriptor     *ServiceDescriptor
	implementation interface{}
	methods        map[string]*MethodMap
	about          interface{}
}

func (i RPCDispatcher) Name() string {
	if i.descriptor == nil {
		return ""
	}
	return i.descriptor.Name
}

func (i RPCDispatcher) RPC(methodName string, ctx context.Context, in proto.Message) (proto.Message, error) {
	implementation, handler, err := i.getRPCHandler(methodName)
	if err != nil {
		return nil, err
	}
	return handler(implementation, ctx, in)
}

func (i RPCDispatcher) getRPCHandler(methodName string) (interface{}, methodHandler, error) {
	mm, ok := i.methods[methodName]
	if !ok {
		return nil, nil, fmt.Errorf("Unknown method %s", methodName)
	}
	return i.implementation, mm.Handler, nil
}
