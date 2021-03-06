// Code generated by counterfeiter. DO NOT EDIT.
package v6fakes

import (
	sync "sync"

	v3action "code.cloudfoundry.org/cli/actor/v3action"
	v6 "code.cloudfoundry.org/cli/command/v6"
)

type FakeCreateIsolationSegmentActor struct {
	CloudControllerAPIVersionStub        func() string
	cloudControllerAPIVersionMutex       sync.RWMutex
	cloudControllerAPIVersionArgsForCall []struct {
	}
	cloudControllerAPIVersionReturns struct {
		result1 string
	}
	cloudControllerAPIVersionReturnsOnCall map[int]struct {
		result1 string
	}
	CreateIsolationSegmentByNameStub        func(v3action.IsolationSegment) (v3action.Warnings, error)
	createIsolationSegmentByNameMutex       sync.RWMutex
	createIsolationSegmentByNameArgsForCall []struct {
		arg1 v3action.IsolationSegment
	}
	createIsolationSegmentByNameReturns struct {
		result1 v3action.Warnings
		result2 error
	}
	createIsolationSegmentByNameReturnsOnCall map[int]struct {
		result1 v3action.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCreateIsolationSegmentActor) CloudControllerAPIVersion() string {
	fake.cloudControllerAPIVersionMutex.Lock()
	ret, specificReturn := fake.cloudControllerAPIVersionReturnsOnCall[len(fake.cloudControllerAPIVersionArgsForCall)]
	fake.cloudControllerAPIVersionArgsForCall = append(fake.cloudControllerAPIVersionArgsForCall, struct {
	}{})
	fake.recordInvocation("CloudControllerAPIVersion", []interface{}{})
	fake.cloudControllerAPIVersionMutex.Unlock()
	if fake.CloudControllerAPIVersionStub != nil {
		return fake.CloudControllerAPIVersionStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.cloudControllerAPIVersionReturns
	return fakeReturns.result1
}

func (fake *FakeCreateIsolationSegmentActor) CloudControllerAPIVersionCallCount() int {
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	return len(fake.cloudControllerAPIVersionArgsForCall)
}

func (fake *FakeCreateIsolationSegmentActor) CloudControllerAPIVersionCalls(stub func() string) {
	fake.cloudControllerAPIVersionMutex.Lock()
	defer fake.cloudControllerAPIVersionMutex.Unlock()
	fake.CloudControllerAPIVersionStub = stub
}

func (fake *FakeCreateIsolationSegmentActor) CloudControllerAPIVersionReturns(result1 string) {
	fake.cloudControllerAPIVersionMutex.Lock()
	defer fake.cloudControllerAPIVersionMutex.Unlock()
	fake.CloudControllerAPIVersionStub = nil
	fake.cloudControllerAPIVersionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCreateIsolationSegmentActor) CloudControllerAPIVersionReturnsOnCall(i int, result1 string) {
	fake.cloudControllerAPIVersionMutex.Lock()
	defer fake.cloudControllerAPIVersionMutex.Unlock()
	fake.CloudControllerAPIVersionStub = nil
	if fake.cloudControllerAPIVersionReturnsOnCall == nil {
		fake.cloudControllerAPIVersionReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.cloudControllerAPIVersionReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeCreateIsolationSegmentActor) CreateIsolationSegmentByName(arg1 v3action.IsolationSegment) (v3action.Warnings, error) {
	fake.createIsolationSegmentByNameMutex.Lock()
	ret, specificReturn := fake.createIsolationSegmentByNameReturnsOnCall[len(fake.createIsolationSegmentByNameArgsForCall)]
	fake.createIsolationSegmentByNameArgsForCall = append(fake.createIsolationSegmentByNameArgsForCall, struct {
		arg1 v3action.IsolationSegment
	}{arg1})
	fake.recordInvocation("CreateIsolationSegmentByName", []interface{}{arg1})
	fake.createIsolationSegmentByNameMutex.Unlock()
	if fake.CreateIsolationSegmentByNameStub != nil {
		return fake.CreateIsolationSegmentByNameStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createIsolationSegmentByNameReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCreateIsolationSegmentActor) CreateIsolationSegmentByNameCallCount() int {
	fake.createIsolationSegmentByNameMutex.RLock()
	defer fake.createIsolationSegmentByNameMutex.RUnlock()
	return len(fake.createIsolationSegmentByNameArgsForCall)
}

func (fake *FakeCreateIsolationSegmentActor) CreateIsolationSegmentByNameCalls(stub func(v3action.IsolationSegment) (v3action.Warnings, error)) {
	fake.createIsolationSegmentByNameMutex.Lock()
	defer fake.createIsolationSegmentByNameMutex.Unlock()
	fake.CreateIsolationSegmentByNameStub = stub
}

func (fake *FakeCreateIsolationSegmentActor) CreateIsolationSegmentByNameArgsForCall(i int) v3action.IsolationSegment {
	fake.createIsolationSegmentByNameMutex.RLock()
	defer fake.createIsolationSegmentByNameMutex.RUnlock()
	argsForCall := fake.createIsolationSegmentByNameArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCreateIsolationSegmentActor) CreateIsolationSegmentByNameReturns(result1 v3action.Warnings, result2 error) {
	fake.createIsolationSegmentByNameMutex.Lock()
	defer fake.createIsolationSegmentByNameMutex.Unlock()
	fake.CreateIsolationSegmentByNameStub = nil
	fake.createIsolationSegmentByNameReturns = struct {
		result1 v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeCreateIsolationSegmentActor) CreateIsolationSegmentByNameReturnsOnCall(i int, result1 v3action.Warnings, result2 error) {
	fake.createIsolationSegmentByNameMutex.Lock()
	defer fake.createIsolationSegmentByNameMutex.Unlock()
	fake.CreateIsolationSegmentByNameStub = nil
	if fake.createIsolationSegmentByNameReturnsOnCall == nil {
		fake.createIsolationSegmentByNameReturnsOnCall = make(map[int]struct {
			result1 v3action.Warnings
			result2 error
		})
	}
	fake.createIsolationSegmentByNameReturnsOnCall[i] = struct {
		result1 v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeCreateIsolationSegmentActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	fake.createIsolationSegmentByNameMutex.RLock()
	defer fake.createIsolationSegmentByNameMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCreateIsolationSegmentActor) recordInvocation(key string, args []interface{}) {
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

var _ v6.CreateIsolationSegmentActor = new(FakeCreateIsolationSegmentActor)
