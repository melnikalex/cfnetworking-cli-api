// Code generated by counterfeiter. DO NOT EDIT.
package cfnetv1fakes

import (
	"sync"

	"code.cloudfoundry.org/cfnetworking-cli-api"
	"code.cloudfoundry.org/cfnetworking-cli-api/cfnetv1"
)

type FakeConnectionWrapper struct {
	MakeStub        func(request *cfnetworking.Request, passedResponse *cfnetworking.Response) error
	makeMutex       sync.RWMutex
	makeArgsForCall []struct {
		request        *cfnetworking.Request
		passedResponse *cfnetworking.Response
	}
	makeReturns struct {
		result1 error
	}
	makeReturnsOnCall map[int]struct {
		result1 error
	}
	WrapStub        func(innerconnection cfnetworking.Connection) cfnetworking.Connection
	wrapMutex       sync.RWMutex
	wrapArgsForCall []struct {
		innerconnection cfnetworking.Connection
	}
	wrapReturns struct {
		result1 cfnetworking.Connection
	}
	wrapReturnsOnCall map[int]struct {
		result1 cfnetworking.Connection
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeConnectionWrapper) Make(request *cfnetworking.Request, passedResponse *cfnetworking.Response) error {
	fake.makeMutex.Lock()
	ret, specificReturn := fake.makeReturnsOnCall[len(fake.makeArgsForCall)]
	fake.makeArgsForCall = append(fake.makeArgsForCall, struct {
		request        *cfnetworking.Request
		passedResponse *cfnetworking.Response
	}{request, passedResponse})
	fake.recordInvocation("Make", []interface{}{request, passedResponse})
	fake.makeMutex.Unlock()
	if fake.MakeStub != nil {
		return fake.MakeStub(request, passedResponse)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.makeReturns.result1
}

func (fake *FakeConnectionWrapper) MakeCallCount() int {
	fake.makeMutex.RLock()
	defer fake.makeMutex.RUnlock()
	return len(fake.makeArgsForCall)
}

func (fake *FakeConnectionWrapper) MakeArgsForCall(i int) (*cfnetworking.Request, *cfnetworking.Response) {
	fake.makeMutex.RLock()
	defer fake.makeMutex.RUnlock()
	return fake.makeArgsForCall[i].request, fake.makeArgsForCall[i].passedResponse
}

func (fake *FakeConnectionWrapper) MakeReturns(result1 error) {
	fake.MakeStub = nil
	fake.makeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeConnectionWrapper) MakeReturnsOnCall(i int, result1 error) {
	fake.MakeStub = nil
	if fake.makeReturnsOnCall == nil {
		fake.makeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.makeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeConnectionWrapper) Wrap(innerconnection cfnetworking.Connection) cfnetworking.Connection {
	fake.wrapMutex.Lock()
	ret, specificReturn := fake.wrapReturnsOnCall[len(fake.wrapArgsForCall)]
	fake.wrapArgsForCall = append(fake.wrapArgsForCall, struct {
		innerconnection cfnetworking.Connection
	}{innerconnection})
	fake.recordInvocation("Wrap", []interface{}{innerconnection})
	fake.wrapMutex.Unlock()
	if fake.WrapStub != nil {
		return fake.WrapStub(innerconnection)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.wrapReturns.result1
}

func (fake *FakeConnectionWrapper) WrapCallCount() int {
	fake.wrapMutex.RLock()
	defer fake.wrapMutex.RUnlock()
	return len(fake.wrapArgsForCall)
}

func (fake *FakeConnectionWrapper) WrapArgsForCall(i int) cfnetworking.Connection {
	fake.wrapMutex.RLock()
	defer fake.wrapMutex.RUnlock()
	return fake.wrapArgsForCall[i].innerconnection
}

func (fake *FakeConnectionWrapper) WrapReturns(result1 cfnetworking.Connection) {
	fake.WrapStub = nil
	fake.wrapReturns = struct {
		result1 cfnetworking.Connection
	}{result1}
}

func (fake *FakeConnectionWrapper) WrapReturnsOnCall(i int, result1 cfnetworking.Connection) {
	fake.WrapStub = nil
	if fake.wrapReturnsOnCall == nil {
		fake.wrapReturnsOnCall = make(map[int]struct {
			result1 cfnetworking.Connection
		})
	}
	fake.wrapReturnsOnCall[i] = struct {
		result1 cfnetworking.Connection
	}{result1}
}

func (fake *FakeConnectionWrapper) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.makeMutex.RLock()
	defer fake.makeMutex.RUnlock()
	fake.wrapMutex.RLock()
	defer fake.wrapMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeConnectionWrapper) recordInvocation(key string, args []interface{}) {
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

var _ cfnetv1.ConnectionWrapper = new(FakeConnectionWrapper)
