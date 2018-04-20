// Code generated by counterfeiter. DO NOT EDIT.
package serverfakes

import (
	"sync"

	"github.com/Peripli/service-manager/server"
)

type FakeEnvironment struct {
	LoadStub        func() error
	loadMutex       sync.RWMutex
	loadArgsForCall []struct{}
	loadReturns     struct {
		result1 error
	}
	loadReturnsOnCall map[int]struct {
		result1 error
	}
	GetStub        func(key string) interface{}
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		key string
	}
	getReturns struct {
		result1 interface{}
	}
	getReturnsOnCall map[int]struct {
		result1 interface{}
	}
	SetStub        func(key string, value interface{})
	setMutex       sync.RWMutex
	setArgsForCall []struct {
		key   string
		value interface{}
	}
	UnmarshalStub        func(value interface{}) error
	unmarshalMutex       sync.RWMutex
	unmarshalArgsForCall []struct {
		value interface{}
	}
	unmarshalReturns struct {
		result1 error
	}
	unmarshalReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEnvironment) Load() error {
	fake.loadMutex.Lock()
	ret, specificReturn := fake.loadReturnsOnCall[len(fake.loadArgsForCall)]
	fake.loadArgsForCall = append(fake.loadArgsForCall, struct{}{})
	fake.recordInvocation("Load", []interface{}{})
	fake.loadMutex.Unlock()
	if fake.LoadStub != nil {
		return fake.LoadStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.loadReturns.result1
}

func (fake *FakeEnvironment) LoadCallCount() int {
	fake.loadMutex.RLock()
	defer fake.loadMutex.RUnlock()
	return len(fake.loadArgsForCall)
}

func (fake *FakeEnvironment) LoadReturns(result1 error) {
	fake.LoadStub = nil
	fake.loadReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeEnvironment) LoadReturnsOnCall(i int, result1 error) {
	fake.LoadStub = nil
	if fake.loadReturnsOnCall == nil {
		fake.loadReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.loadReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeEnvironment) Get(key string) interface{} {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		key string
	}{key})
	fake.recordInvocation("Get", []interface{}{key})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(key)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getReturns.result1
}

func (fake *FakeEnvironment) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeEnvironment) GetArgsForCall(i int) string {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return fake.getArgsForCall[i].key
}

func (fake *FakeEnvironment) GetReturns(result1 interface{}) {
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 interface{}
	}{result1}
}

func (fake *FakeEnvironment) GetReturnsOnCall(i int, result1 interface{}) {
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 interface{}
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 interface{}
	}{result1}
}

func (fake *FakeEnvironment) Set(key string, value interface{}) {
	fake.setMutex.Lock()
	fake.setArgsForCall = append(fake.setArgsForCall, struct {
		key   string
		value interface{}
	}{key, value})
	fake.recordInvocation("Set", []interface{}{key, value})
	fake.setMutex.Unlock()
	if fake.SetStub != nil {
		fake.SetStub(key, value)
	}
}

func (fake *FakeEnvironment) SetCallCount() int {
	fake.setMutex.RLock()
	defer fake.setMutex.RUnlock()
	return len(fake.setArgsForCall)
}

func (fake *FakeEnvironment) SetArgsForCall(i int) (string, interface{}) {
	fake.setMutex.RLock()
	defer fake.setMutex.RUnlock()
	return fake.setArgsForCall[i].key, fake.setArgsForCall[i].value
}

func (fake *FakeEnvironment) Unmarshal(value interface{}) error {
	fake.unmarshalMutex.Lock()
	ret, specificReturn := fake.unmarshalReturnsOnCall[len(fake.unmarshalArgsForCall)]
	fake.unmarshalArgsForCall = append(fake.unmarshalArgsForCall, struct {
		value interface{}
	}{value})
	fake.recordInvocation("Unmarshal", []interface{}{value})
	fake.unmarshalMutex.Unlock()
	if fake.UnmarshalStub != nil {
		return fake.UnmarshalStub(value)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.unmarshalReturns.result1
}

func (fake *FakeEnvironment) UnmarshalCallCount() int {
	fake.unmarshalMutex.RLock()
	defer fake.unmarshalMutex.RUnlock()
	return len(fake.unmarshalArgsForCall)
}

func (fake *FakeEnvironment) UnmarshalArgsForCall(i int) interface{} {
	fake.unmarshalMutex.RLock()
	defer fake.unmarshalMutex.RUnlock()
	return fake.unmarshalArgsForCall[i].value
}

func (fake *FakeEnvironment) UnmarshalReturns(result1 error) {
	fake.UnmarshalStub = nil
	fake.unmarshalReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeEnvironment) UnmarshalReturnsOnCall(i int, result1 error) {
	fake.UnmarshalStub = nil
	if fake.unmarshalReturnsOnCall == nil {
		fake.unmarshalReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.unmarshalReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeEnvironment) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.loadMutex.RLock()
	defer fake.loadMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.setMutex.RLock()
	defer fake.setMutex.RUnlock()
	fake.unmarshalMutex.RLock()
	defer fake.unmarshalMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeEnvironment) recordInvocation(key string, args []interface{}) {
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

var _ server.Environment = new(FakeEnvironment)
