// Code generated by counterfeiter. DO NOT EDIT.
package v6fakes

import (
	sync "sync"

	v3action "code.cloudfoundry.org/cli/actor/v3action"
	v6 "code.cloudfoundry.org/cli/command/v6"
)

type FakeV3GetHealthCheckActor struct {
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
	GetApplicationProcessHealthChecksByNameAndSpaceStub        func(string, string) ([]v3action.ProcessHealthCheck, v3action.Warnings, error)
	getApplicationProcessHealthChecksByNameAndSpaceMutex       sync.RWMutex
	getApplicationProcessHealthChecksByNameAndSpaceArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getApplicationProcessHealthChecksByNameAndSpaceReturns struct {
		result1 []v3action.ProcessHealthCheck
		result2 v3action.Warnings
		result3 error
	}
	getApplicationProcessHealthChecksByNameAndSpaceReturnsOnCall map[int]struct {
		result1 []v3action.ProcessHealthCheck
		result2 v3action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeV3GetHealthCheckActor) CloudControllerAPIVersion() string {
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

func (fake *FakeV3GetHealthCheckActor) CloudControllerAPIVersionCallCount() int {
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	return len(fake.cloudControllerAPIVersionArgsForCall)
}

func (fake *FakeV3GetHealthCheckActor) CloudControllerAPIVersionCalls(stub func() string) {
	fake.cloudControllerAPIVersionMutex.Lock()
	defer fake.cloudControllerAPIVersionMutex.Unlock()
	fake.CloudControllerAPIVersionStub = stub
}

func (fake *FakeV3GetHealthCheckActor) CloudControllerAPIVersionReturns(result1 string) {
	fake.cloudControllerAPIVersionMutex.Lock()
	defer fake.cloudControllerAPIVersionMutex.Unlock()
	fake.CloudControllerAPIVersionStub = nil
	fake.cloudControllerAPIVersionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeV3GetHealthCheckActor) CloudControllerAPIVersionReturnsOnCall(i int, result1 string) {
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

func (fake *FakeV3GetHealthCheckActor) GetApplicationProcessHealthChecksByNameAndSpace(arg1 string, arg2 string) ([]v3action.ProcessHealthCheck, v3action.Warnings, error) {
	fake.getApplicationProcessHealthChecksByNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.getApplicationProcessHealthChecksByNameAndSpaceReturnsOnCall[len(fake.getApplicationProcessHealthChecksByNameAndSpaceArgsForCall)]
	fake.getApplicationProcessHealthChecksByNameAndSpaceArgsForCall = append(fake.getApplicationProcessHealthChecksByNameAndSpaceArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetApplicationProcessHealthChecksByNameAndSpace", []interface{}{arg1, arg2})
	fake.getApplicationProcessHealthChecksByNameAndSpaceMutex.Unlock()
	if fake.GetApplicationProcessHealthChecksByNameAndSpaceStub != nil {
		return fake.GetApplicationProcessHealthChecksByNameAndSpaceStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getApplicationProcessHealthChecksByNameAndSpaceReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeV3GetHealthCheckActor) GetApplicationProcessHealthChecksByNameAndSpaceCallCount() int {
	fake.getApplicationProcessHealthChecksByNameAndSpaceMutex.RLock()
	defer fake.getApplicationProcessHealthChecksByNameAndSpaceMutex.RUnlock()
	return len(fake.getApplicationProcessHealthChecksByNameAndSpaceArgsForCall)
}

func (fake *FakeV3GetHealthCheckActor) GetApplicationProcessHealthChecksByNameAndSpaceCalls(stub func(string, string) ([]v3action.ProcessHealthCheck, v3action.Warnings, error)) {
	fake.getApplicationProcessHealthChecksByNameAndSpaceMutex.Lock()
	defer fake.getApplicationProcessHealthChecksByNameAndSpaceMutex.Unlock()
	fake.GetApplicationProcessHealthChecksByNameAndSpaceStub = stub
}

func (fake *FakeV3GetHealthCheckActor) GetApplicationProcessHealthChecksByNameAndSpaceArgsForCall(i int) (string, string) {
	fake.getApplicationProcessHealthChecksByNameAndSpaceMutex.RLock()
	defer fake.getApplicationProcessHealthChecksByNameAndSpaceMutex.RUnlock()
	argsForCall := fake.getApplicationProcessHealthChecksByNameAndSpaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeV3GetHealthCheckActor) GetApplicationProcessHealthChecksByNameAndSpaceReturns(result1 []v3action.ProcessHealthCheck, result2 v3action.Warnings, result3 error) {
	fake.getApplicationProcessHealthChecksByNameAndSpaceMutex.Lock()
	defer fake.getApplicationProcessHealthChecksByNameAndSpaceMutex.Unlock()
	fake.GetApplicationProcessHealthChecksByNameAndSpaceStub = nil
	fake.getApplicationProcessHealthChecksByNameAndSpaceReturns = struct {
		result1 []v3action.ProcessHealthCheck
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3GetHealthCheckActor) GetApplicationProcessHealthChecksByNameAndSpaceReturnsOnCall(i int, result1 []v3action.ProcessHealthCheck, result2 v3action.Warnings, result3 error) {
	fake.getApplicationProcessHealthChecksByNameAndSpaceMutex.Lock()
	defer fake.getApplicationProcessHealthChecksByNameAndSpaceMutex.Unlock()
	fake.GetApplicationProcessHealthChecksByNameAndSpaceStub = nil
	if fake.getApplicationProcessHealthChecksByNameAndSpaceReturnsOnCall == nil {
		fake.getApplicationProcessHealthChecksByNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 []v3action.ProcessHealthCheck
			result2 v3action.Warnings
			result3 error
		})
	}
	fake.getApplicationProcessHealthChecksByNameAndSpaceReturnsOnCall[i] = struct {
		result1 []v3action.ProcessHealthCheck
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3GetHealthCheckActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	fake.getApplicationProcessHealthChecksByNameAndSpaceMutex.RLock()
	defer fake.getApplicationProcessHealthChecksByNameAndSpaceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeV3GetHealthCheckActor) recordInvocation(key string, args []interface{}) {
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

var _ v6.V3GetHealthCheckActor = new(FakeV3GetHealthCheckActor)
