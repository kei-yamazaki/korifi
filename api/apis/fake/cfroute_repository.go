// Code generated by counterfeiter. DO NOT EDIT.
package fake

import (
	"context"
	"sync"

	"code.cloudfoundry.org/cf-k8s-controllers/api/apis"
	"code.cloudfoundry.org/cf-k8s-controllers/api/authorization"
	"code.cloudfoundry.org/cf-k8s-controllers/api/repositories"
)

type CFRouteRepository struct {
	AddDestinationsToRouteStub        func(context.Context, authorization.Info, repositories.AddDestinationsToRouteMessage) (repositories.RouteRecord, error)
	addDestinationsToRouteMutex       sync.RWMutex
	addDestinationsToRouteArgsForCall []struct {
		arg1 context.Context
		arg2 authorization.Info
		arg3 repositories.AddDestinationsToRouteMessage
	}
	addDestinationsToRouteReturns struct {
		result1 repositories.RouteRecord
		result2 error
	}
	addDestinationsToRouteReturnsOnCall map[int]struct {
		result1 repositories.RouteRecord
		result2 error
	}
	CreateRouteStub        func(context.Context, authorization.Info, repositories.CreateRouteMessage) (repositories.RouteRecord, error)
	createRouteMutex       sync.RWMutex
	createRouteArgsForCall []struct {
		arg1 context.Context
		arg2 authorization.Info
		arg3 repositories.CreateRouteMessage
	}
	createRouteReturns struct {
		result1 repositories.RouteRecord
		result2 error
	}
	createRouteReturnsOnCall map[int]struct {
		result1 repositories.RouteRecord
		result2 error
	}
	GetRouteStub        func(context.Context, authorization.Info, string) (repositories.RouteRecord, error)
	getRouteMutex       sync.RWMutex
	getRouteArgsForCall []struct {
		arg1 context.Context
		arg2 authorization.Info
		arg3 string
	}
	getRouteReturns struct {
		result1 repositories.RouteRecord
		result2 error
	}
	getRouteReturnsOnCall map[int]struct {
		result1 repositories.RouteRecord
		result2 error
	}
	ListRoutesStub        func(context.Context, authorization.Info, repositories.ListRoutesMessage) ([]repositories.RouteRecord, error)
	listRoutesMutex       sync.RWMutex
	listRoutesArgsForCall []struct {
		arg1 context.Context
		arg2 authorization.Info
		arg3 repositories.ListRoutesMessage
	}
	listRoutesReturns struct {
		result1 []repositories.RouteRecord
		result2 error
	}
	listRoutesReturnsOnCall map[int]struct {
		result1 []repositories.RouteRecord
		result2 error
	}
	ListRoutesForAppStub        func(context.Context, authorization.Info, string, string) ([]repositories.RouteRecord, error)
	listRoutesForAppMutex       sync.RWMutex
	listRoutesForAppArgsForCall []struct {
		arg1 context.Context
		arg2 authorization.Info
		arg3 string
		arg4 string
	}
	listRoutesForAppReturns struct {
		result1 []repositories.RouteRecord
		result2 error
	}
	listRoutesForAppReturnsOnCall map[int]struct {
		result1 []repositories.RouteRecord
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *CFRouteRepository) AddDestinationsToRoute(arg1 context.Context, arg2 authorization.Info, arg3 repositories.AddDestinationsToRouteMessage) (repositories.RouteRecord, error) {
	fake.addDestinationsToRouteMutex.Lock()
	ret, specificReturn := fake.addDestinationsToRouteReturnsOnCall[len(fake.addDestinationsToRouteArgsForCall)]
	fake.addDestinationsToRouteArgsForCall = append(fake.addDestinationsToRouteArgsForCall, struct {
		arg1 context.Context
		arg2 authorization.Info
		arg3 repositories.AddDestinationsToRouteMessage
	}{arg1, arg2, arg3})
	stub := fake.AddDestinationsToRouteStub
	fakeReturns := fake.addDestinationsToRouteReturns
	fake.recordInvocation("AddDestinationsToRoute", []interface{}{arg1, arg2, arg3})
	fake.addDestinationsToRouteMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CFRouteRepository) AddDestinationsToRouteCallCount() int {
	fake.addDestinationsToRouteMutex.RLock()
	defer fake.addDestinationsToRouteMutex.RUnlock()
	return len(fake.addDestinationsToRouteArgsForCall)
}

func (fake *CFRouteRepository) AddDestinationsToRouteCalls(stub func(context.Context, authorization.Info, repositories.AddDestinationsToRouteMessage) (repositories.RouteRecord, error)) {
	fake.addDestinationsToRouteMutex.Lock()
	defer fake.addDestinationsToRouteMutex.Unlock()
	fake.AddDestinationsToRouteStub = stub
}

func (fake *CFRouteRepository) AddDestinationsToRouteArgsForCall(i int) (context.Context, authorization.Info, repositories.AddDestinationsToRouteMessage) {
	fake.addDestinationsToRouteMutex.RLock()
	defer fake.addDestinationsToRouteMutex.RUnlock()
	argsForCall := fake.addDestinationsToRouteArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *CFRouteRepository) AddDestinationsToRouteReturns(result1 repositories.RouteRecord, result2 error) {
	fake.addDestinationsToRouteMutex.Lock()
	defer fake.addDestinationsToRouteMutex.Unlock()
	fake.AddDestinationsToRouteStub = nil
	fake.addDestinationsToRouteReturns = struct {
		result1 repositories.RouteRecord
		result2 error
	}{result1, result2}
}

func (fake *CFRouteRepository) AddDestinationsToRouteReturnsOnCall(i int, result1 repositories.RouteRecord, result2 error) {
	fake.addDestinationsToRouteMutex.Lock()
	defer fake.addDestinationsToRouteMutex.Unlock()
	fake.AddDestinationsToRouteStub = nil
	if fake.addDestinationsToRouteReturnsOnCall == nil {
		fake.addDestinationsToRouteReturnsOnCall = make(map[int]struct {
			result1 repositories.RouteRecord
			result2 error
		})
	}
	fake.addDestinationsToRouteReturnsOnCall[i] = struct {
		result1 repositories.RouteRecord
		result2 error
	}{result1, result2}
}

func (fake *CFRouteRepository) CreateRoute(arg1 context.Context, arg2 authorization.Info, arg3 repositories.CreateRouteMessage) (repositories.RouteRecord, error) {
	fake.createRouteMutex.Lock()
	ret, specificReturn := fake.createRouteReturnsOnCall[len(fake.createRouteArgsForCall)]
	fake.createRouteArgsForCall = append(fake.createRouteArgsForCall, struct {
		arg1 context.Context
		arg2 authorization.Info
		arg3 repositories.CreateRouteMessage
	}{arg1, arg2, arg3})
	stub := fake.CreateRouteStub
	fakeReturns := fake.createRouteReturns
	fake.recordInvocation("CreateRoute", []interface{}{arg1, arg2, arg3})
	fake.createRouteMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CFRouteRepository) CreateRouteCallCount() int {
	fake.createRouteMutex.RLock()
	defer fake.createRouteMutex.RUnlock()
	return len(fake.createRouteArgsForCall)
}

func (fake *CFRouteRepository) CreateRouteCalls(stub func(context.Context, authorization.Info, repositories.CreateRouteMessage) (repositories.RouteRecord, error)) {
	fake.createRouteMutex.Lock()
	defer fake.createRouteMutex.Unlock()
	fake.CreateRouteStub = stub
}

func (fake *CFRouteRepository) CreateRouteArgsForCall(i int) (context.Context, authorization.Info, repositories.CreateRouteMessage) {
	fake.createRouteMutex.RLock()
	defer fake.createRouteMutex.RUnlock()
	argsForCall := fake.createRouteArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *CFRouteRepository) CreateRouteReturns(result1 repositories.RouteRecord, result2 error) {
	fake.createRouteMutex.Lock()
	defer fake.createRouteMutex.Unlock()
	fake.CreateRouteStub = nil
	fake.createRouteReturns = struct {
		result1 repositories.RouteRecord
		result2 error
	}{result1, result2}
}

func (fake *CFRouteRepository) CreateRouteReturnsOnCall(i int, result1 repositories.RouteRecord, result2 error) {
	fake.createRouteMutex.Lock()
	defer fake.createRouteMutex.Unlock()
	fake.CreateRouteStub = nil
	if fake.createRouteReturnsOnCall == nil {
		fake.createRouteReturnsOnCall = make(map[int]struct {
			result1 repositories.RouteRecord
			result2 error
		})
	}
	fake.createRouteReturnsOnCall[i] = struct {
		result1 repositories.RouteRecord
		result2 error
	}{result1, result2}
}

func (fake *CFRouteRepository) GetRoute(arg1 context.Context, arg2 authorization.Info, arg3 string) (repositories.RouteRecord, error) {
	fake.getRouteMutex.Lock()
	ret, specificReturn := fake.getRouteReturnsOnCall[len(fake.getRouteArgsForCall)]
	fake.getRouteArgsForCall = append(fake.getRouteArgsForCall, struct {
		arg1 context.Context
		arg2 authorization.Info
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.GetRouteStub
	fakeReturns := fake.getRouteReturns
	fake.recordInvocation("GetRoute", []interface{}{arg1, arg2, arg3})
	fake.getRouteMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CFRouteRepository) GetRouteCallCount() int {
	fake.getRouteMutex.RLock()
	defer fake.getRouteMutex.RUnlock()
	return len(fake.getRouteArgsForCall)
}

func (fake *CFRouteRepository) GetRouteCalls(stub func(context.Context, authorization.Info, string) (repositories.RouteRecord, error)) {
	fake.getRouteMutex.Lock()
	defer fake.getRouteMutex.Unlock()
	fake.GetRouteStub = stub
}

func (fake *CFRouteRepository) GetRouteArgsForCall(i int) (context.Context, authorization.Info, string) {
	fake.getRouteMutex.RLock()
	defer fake.getRouteMutex.RUnlock()
	argsForCall := fake.getRouteArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *CFRouteRepository) GetRouteReturns(result1 repositories.RouteRecord, result2 error) {
	fake.getRouteMutex.Lock()
	defer fake.getRouteMutex.Unlock()
	fake.GetRouteStub = nil
	fake.getRouteReturns = struct {
		result1 repositories.RouteRecord
		result2 error
	}{result1, result2}
}

func (fake *CFRouteRepository) GetRouteReturnsOnCall(i int, result1 repositories.RouteRecord, result2 error) {
	fake.getRouteMutex.Lock()
	defer fake.getRouteMutex.Unlock()
	fake.GetRouteStub = nil
	if fake.getRouteReturnsOnCall == nil {
		fake.getRouteReturnsOnCall = make(map[int]struct {
			result1 repositories.RouteRecord
			result2 error
		})
	}
	fake.getRouteReturnsOnCall[i] = struct {
		result1 repositories.RouteRecord
		result2 error
	}{result1, result2}
}

func (fake *CFRouteRepository) ListRoutes(arg1 context.Context, arg2 authorization.Info, arg3 repositories.ListRoutesMessage) ([]repositories.RouteRecord, error) {
	fake.listRoutesMutex.Lock()
	ret, specificReturn := fake.listRoutesReturnsOnCall[len(fake.listRoutesArgsForCall)]
	fake.listRoutesArgsForCall = append(fake.listRoutesArgsForCall, struct {
		arg1 context.Context
		arg2 authorization.Info
		arg3 repositories.ListRoutesMessage
	}{arg1, arg2, arg3})
	stub := fake.ListRoutesStub
	fakeReturns := fake.listRoutesReturns
	fake.recordInvocation("ListRoutes", []interface{}{arg1, arg2, arg3})
	fake.listRoutesMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CFRouteRepository) ListRoutesCallCount() int {
	fake.listRoutesMutex.RLock()
	defer fake.listRoutesMutex.RUnlock()
	return len(fake.listRoutesArgsForCall)
}

func (fake *CFRouteRepository) ListRoutesCalls(stub func(context.Context, authorization.Info, repositories.ListRoutesMessage) ([]repositories.RouteRecord, error)) {
	fake.listRoutesMutex.Lock()
	defer fake.listRoutesMutex.Unlock()
	fake.ListRoutesStub = stub
}

func (fake *CFRouteRepository) ListRoutesArgsForCall(i int) (context.Context, authorization.Info, repositories.ListRoutesMessage) {
	fake.listRoutesMutex.RLock()
	defer fake.listRoutesMutex.RUnlock()
	argsForCall := fake.listRoutesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *CFRouteRepository) ListRoutesReturns(result1 []repositories.RouteRecord, result2 error) {
	fake.listRoutesMutex.Lock()
	defer fake.listRoutesMutex.Unlock()
	fake.ListRoutesStub = nil
	fake.listRoutesReturns = struct {
		result1 []repositories.RouteRecord
		result2 error
	}{result1, result2}
}

func (fake *CFRouteRepository) ListRoutesReturnsOnCall(i int, result1 []repositories.RouteRecord, result2 error) {
	fake.listRoutesMutex.Lock()
	defer fake.listRoutesMutex.Unlock()
	fake.ListRoutesStub = nil
	if fake.listRoutesReturnsOnCall == nil {
		fake.listRoutesReturnsOnCall = make(map[int]struct {
			result1 []repositories.RouteRecord
			result2 error
		})
	}
	fake.listRoutesReturnsOnCall[i] = struct {
		result1 []repositories.RouteRecord
		result2 error
	}{result1, result2}
}

func (fake *CFRouteRepository) ListRoutesForApp(arg1 context.Context, arg2 authorization.Info, arg3 string, arg4 string) ([]repositories.RouteRecord, error) {
	fake.listRoutesForAppMutex.Lock()
	ret, specificReturn := fake.listRoutesForAppReturnsOnCall[len(fake.listRoutesForAppArgsForCall)]
	fake.listRoutesForAppArgsForCall = append(fake.listRoutesForAppArgsForCall, struct {
		arg1 context.Context
		arg2 authorization.Info
		arg3 string
		arg4 string
	}{arg1, arg2, arg3, arg4})
	stub := fake.ListRoutesForAppStub
	fakeReturns := fake.listRoutesForAppReturns
	fake.recordInvocation("ListRoutesForApp", []interface{}{arg1, arg2, arg3, arg4})
	fake.listRoutesForAppMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CFRouteRepository) ListRoutesForAppCallCount() int {
	fake.listRoutesForAppMutex.RLock()
	defer fake.listRoutesForAppMutex.RUnlock()
	return len(fake.listRoutesForAppArgsForCall)
}

func (fake *CFRouteRepository) ListRoutesForAppCalls(stub func(context.Context, authorization.Info, string, string) ([]repositories.RouteRecord, error)) {
	fake.listRoutesForAppMutex.Lock()
	defer fake.listRoutesForAppMutex.Unlock()
	fake.ListRoutesForAppStub = stub
}

func (fake *CFRouteRepository) ListRoutesForAppArgsForCall(i int) (context.Context, authorization.Info, string, string) {
	fake.listRoutesForAppMutex.RLock()
	defer fake.listRoutesForAppMutex.RUnlock()
	argsForCall := fake.listRoutesForAppArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *CFRouteRepository) ListRoutesForAppReturns(result1 []repositories.RouteRecord, result2 error) {
	fake.listRoutesForAppMutex.Lock()
	defer fake.listRoutesForAppMutex.Unlock()
	fake.ListRoutesForAppStub = nil
	fake.listRoutesForAppReturns = struct {
		result1 []repositories.RouteRecord
		result2 error
	}{result1, result2}
}

func (fake *CFRouteRepository) ListRoutesForAppReturnsOnCall(i int, result1 []repositories.RouteRecord, result2 error) {
	fake.listRoutesForAppMutex.Lock()
	defer fake.listRoutesForAppMutex.Unlock()
	fake.ListRoutesForAppStub = nil
	if fake.listRoutesForAppReturnsOnCall == nil {
		fake.listRoutesForAppReturnsOnCall = make(map[int]struct {
			result1 []repositories.RouteRecord
			result2 error
		})
	}
	fake.listRoutesForAppReturnsOnCall[i] = struct {
		result1 []repositories.RouteRecord
		result2 error
	}{result1, result2}
}

func (fake *CFRouteRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addDestinationsToRouteMutex.RLock()
	defer fake.addDestinationsToRouteMutex.RUnlock()
	fake.createRouteMutex.RLock()
	defer fake.createRouteMutex.RUnlock()
	fake.getRouteMutex.RLock()
	defer fake.getRouteMutex.RUnlock()
	fake.listRoutesMutex.RLock()
	defer fake.listRoutesMutex.RUnlock()
	fake.listRoutesForAppMutex.RLock()
	defer fake.listRoutesForAppMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *CFRouteRepository) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ apis.CFRouteRepository = new(CFRouteRepository)
