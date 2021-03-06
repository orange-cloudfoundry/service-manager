// Code generated by counterfeiter. DO NOT EDIT.
package securityfakes

import (
	context "context"
	sync "sync"

	security "github.com/Peripli/service-manager/pkg/security"
)

type FakeEncrypter struct {
	DecryptStub        func(context.Context, []byte) ([]byte, error)
	decryptMutex       sync.RWMutex
	decryptArgsForCall []struct {
		arg1 context.Context
		arg2 []byte
	}
	decryptReturns struct {
		result1 []byte
		result2 error
	}
	decryptReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	EncryptStub        func(context.Context, []byte) ([]byte, error)
	encryptMutex       sync.RWMutex
	encryptArgsForCall []struct {
		arg1 context.Context
		arg2 []byte
	}
	encryptReturns struct {
		result1 []byte
		result2 error
	}
	encryptReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEncrypter) Decrypt(arg1 context.Context, arg2 []byte) ([]byte, error) {
	var arg2Copy []byte
	if arg2 != nil {
		arg2Copy = make([]byte, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.decryptMutex.Lock()
	ret, specificReturn := fake.decryptReturnsOnCall[len(fake.decryptArgsForCall)]
	fake.decryptArgsForCall = append(fake.decryptArgsForCall, struct {
		arg1 context.Context
		arg2 []byte
	}{arg1, arg2Copy})
	fake.recordInvocation("Decrypt", []interface{}{arg1, arg2Copy})
	fake.decryptMutex.Unlock()
	if fake.DecryptStub != nil {
		return fake.DecryptStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.decryptReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeEncrypter) DecryptCallCount() int {
	fake.decryptMutex.RLock()
	defer fake.decryptMutex.RUnlock()
	return len(fake.decryptArgsForCall)
}

func (fake *FakeEncrypter) DecryptCalls(stub func(context.Context, []byte) ([]byte, error)) {
	fake.decryptMutex.Lock()
	defer fake.decryptMutex.Unlock()
	fake.DecryptStub = stub
}

func (fake *FakeEncrypter) DecryptArgsForCall(i int) (context.Context, []byte) {
	fake.decryptMutex.RLock()
	defer fake.decryptMutex.RUnlock()
	argsForCall := fake.decryptArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeEncrypter) DecryptReturns(result1 []byte, result2 error) {
	fake.decryptMutex.Lock()
	defer fake.decryptMutex.Unlock()
	fake.DecryptStub = nil
	fake.decryptReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeEncrypter) DecryptReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.decryptMutex.Lock()
	defer fake.decryptMutex.Unlock()
	fake.DecryptStub = nil
	if fake.decryptReturnsOnCall == nil {
		fake.decryptReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.decryptReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeEncrypter) Encrypt(arg1 context.Context, arg2 []byte) ([]byte, error) {
	var arg2Copy []byte
	if arg2 != nil {
		arg2Copy = make([]byte, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.encryptMutex.Lock()
	ret, specificReturn := fake.encryptReturnsOnCall[len(fake.encryptArgsForCall)]
	fake.encryptArgsForCall = append(fake.encryptArgsForCall, struct {
		arg1 context.Context
		arg2 []byte
	}{arg1, arg2Copy})
	fake.recordInvocation("Encrypt", []interface{}{arg1, arg2Copy})
	fake.encryptMutex.Unlock()
	if fake.EncryptStub != nil {
		return fake.EncryptStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.encryptReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeEncrypter) EncryptCallCount() int {
	fake.encryptMutex.RLock()
	defer fake.encryptMutex.RUnlock()
	return len(fake.encryptArgsForCall)
}

func (fake *FakeEncrypter) EncryptCalls(stub func(context.Context, []byte) ([]byte, error)) {
	fake.encryptMutex.Lock()
	defer fake.encryptMutex.Unlock()
	fake.EncryptStub = stub
}

func (fake *FakeEncrypter) EncryptArgsForCall(i int) (context.Context, []byte) {
	fake.encryptMutex.RLock()
	defer fake.encryptMutex.RUnlock()
	argsForCall := fake.encryptArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeEncrypter) EncryptReturns(result1 []byte, result2 error) {
	fake.encryptMutex.Lock()
	defer fake.encryptMutex.Unlock()
	fake.EncryptStub = nil
	fake.encryptReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeEncrypter) EncryptReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.encryptMutex.Lock()
	defer fake.encryptMutex.Unlock()
	fake.EncryptStub = nil
	if fake.encryptReturnsOnCall == nil {
		fake.encryptReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.encryptReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeEncrypter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.decryptMutex.RLock()
	defer fake.decryptMutex.RUnlock()
	fake.encryptMutex.RLock()
	defer fake.encryptMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeEncrypter) recordInvocation(key string, args []interface{}) {
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

var _ security.Encrypter = new(FakeEncrypter)
