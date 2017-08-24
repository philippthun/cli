// Code generated by counterfeiter. DO NOT EDIT.
package v3fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v3action"
	"code.cloudfoundry.org/cli/command/v3"
)

type FakeV3RestartAppInstanceActor struct {
	DeleteInstanceByApplicationNameSpaceProcessTypeAndIndexStub        func(appName string, spaceGUID string, processType string, instanceIndex int) (v3action.Warnings, error)
	deleteInstanceByApplicationNameSpaceProcessTypeAndIndexMutex       sync.RWMutex
	deleteInstanceByApplicationNameSpaceProcessTypeAndIndexArgsForCall []struct {
		appName       string
		spaceGUID     string
		processType   string
		instanceIndex int
	}
	deleteInstanceByApplicationNameSpaceProcessTypeAndIndexReturns struct {
		result1 v3action.Warnings
		result2 error
	}
	deleteInstanceByApplicationNameSpaceProcessTypeAndIndexReturnsOnCall map[int]struct {
		result1 v3action.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeV3RestartAppInstanceActor) DeleteInstanceByApplicationNameSpaceProcessTypeAndIndex(appName string, spaceGUID string, processType string, instanceIndex int) (v3action.Warnings, error) {
	fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexMutex.Lock()
	ret, specificReturn := fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexReturnsOnCall[len(fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexArgsForCall)]
	fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexArgsForCall = append(fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexArgsForCall, struct {
		appName       string
		spaceGUID     string
		processType   string
		instanceIndex int
	}{appName, spaceGUID, processType, instanceIndex})
	fake.recordInvocation("DeleteInstanceByApplicationNameSpaceProcessTypeAndIndex", []interface{}{appName, spaceGUID, processType, instanceIndex})
	fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexMutex.Unlock()
	if fake.DeleteInstanceByApplicationNameSpaceProcessTypeAndIndexStub != nil {
		return fake.DeleteInstanceByApplicationNameSpaceProcessTypeAndIndexStub(appName, spaceGUID, processType, instanceIndex)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexReturns.result1, fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexReturns.result2
}

func (fake *FakeV3RestartAppInstanceActor) DeleteInstanceByApplicationNameSpaceProcessTypeAndIndexCallCount() int {
	fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexMutex.RLock()
	defer fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexMutex.RUnlock()
	return len(fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexArgsForCall)
}

func (fake *FakeV3RestartAppInstanceActor) DeleteInstanceByApplicationNameSpaceProcessTypeAndIndexArgsForCall(i int) (string, string, string, int) {
	fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexMutex.RLock()
	defer fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexMutex.RUnlock()
	return fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexArgsForCall[i].appName, fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexArgsForCall[i].spaceGUID, fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexArgsForCall[i].processType, fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexArgsForCall[i].instanceIndex
}

func (fake *FakeV3RestartAppInstanceActor) DeleteInstanceByApplicationNameSpaceProcessTypeAndIndexReturns(result1 v3action.Warnings, result2 error) {
	fake.DeleteInstanceByApplicationNameSpaceProcessTypeAndIndexStub = nil
	fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexReturns = struct {
		result1 v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeV3RestartAppInstanceActor) DeleteInstanceByApplicationNameSpaceProcessTypeAndIndexReturnsOnCall(i int, result1 v3action.Warnings, result2 error) {
	fake.DeleteInstanceByApplicationNameSpaceProcessTypeAndIndexStub = nil
	if fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexReturnsOnCall == nil {
		fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexReturnsOnCall = make(map[int]struct {
			result1 v3action.Warnings
			result2 error
		})
	}
	fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexReturnsOnCall[i] = struct {
		result1 v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeV3RestartAppInstanceActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexMutex.RLock()
	defer fake.deleteInstanceByApplicationNameSpaceProcessTypeAndIndexMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeV3RestartAppInstanceActor) recordInvocation(key string, args []interface{}) {
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

var _ v3.V3RestartAppInstanceActor = new(FakeV3RestartAppInstanceActor)