// Code generated by counterfeiter. DO NOT EDIT.
package storagefakes

import (
	context "context"
	sync "sync"

	storage "github.com/Peripli/service-manager/storage"
)

type FakeStorage struct {
	BrokerStub        func() storage.Broker
	brokerMutex       sync.RWMutex
	brokerArgsForCall []struct {
	}
	brokerReturns struct {
		result1 storage.Broker
	}
	brokerReturnsOnCall map[int]struct {
		result1 storage.Broker
	}
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
	}
	closeReturns struct {
		result1 error
	}
	closeReturnsOnCall map[int]struct {
		result1 error
	}
	CredentialsStub        func() storage.Credentials
	credentialsMutex       sync.RWMutex
	credentialsArgsForCall []struct {
	}
	credentialsReturns struct {
		result1 storage.Credentials
	}
	credentialsReturnsOnCall map[int]struct {
		result1 storage.Credentials
	}
	InTransactionStub        func(context.Context, func(ctx context.Context, storage storage.Warehouse) error) error
	inTransactionMutex       sync.RWMutex
	inTransactionArgsForCall []struct {
		arg1 context.Context
		arg2 func(ctx context.Context, storage storage.Warehouse) error
	}
	inTransactionReturns struct {
		result1 error
	}
	inTransactionReturnsOnCall map[int]struct {
		result1 error
	}
	OpenStub        func(*storage.Settings) error
	openMutex       sync.RWMutex
	openArgsForCall []struct {
		arg1 *storage.Settings
	}
	openReturns struct {
		result1 error
	}
	openReturnsOnCall map[int]struct {
		result1 error
	}
	PingStub        func() error
	pingMutex       sync.RWMutex
	pingArgsForCall []struct {
	}
	pingReturns struct {
		result1 error
	}
	pingReturnsOnCall map[int]struct {
		result1 error
	}
	PlatformStub        func() storage.Platform
	platformMutex       sync.RWMutex
	platformArgsForCall []struct {
	}
	platformReturns struct {
		result1 storage.Platform
	}
	platformReturnsOnCall map[int]struct {
		result1 storage.Platform
	}
	SecurityStub        func() storage.Security
	securityMutex       sync.RWMutex
	securityArgsForCall []struct {
	}
	securityReturns struct {
		result1 storage.Security
	}
	securityReturnsOnCall map[int]struct {
		result1 storage.Security
	}
	ServiceOfferingStub        func() storage.ServiceOffering
	serviceOfferingMutex       sync.RWMutex
	serviceOfferingArgsForCall []struct {
	}
	serviceOfferingReturns struct {
		result1 storage.ServiceOffering
	}
	serviceOfferingReturnsOnCall map[int]struct {
		result1 storage.ServiceOffering
	}
	ServicePlanStub        func() storage.ServicePlan
	servicePlanMutex       sync.RWMutex
	servicePlanArgsForCall []struct {
	}
	servicePlanReturns struct {
		result1 storage.ServicePlan
	}
	servicePlanReturnsOnCall map[int]struct {
		result1 storage.ServicePlan
	}
	VisibilityStub        func() storage.Visibility
	visibilityMutex       sync.RWMutex
	visibilityArgsForCall []struct {
	}
	visibilityReturns struct {
		result1 storage.Visibility
	}
	visibilityReturnsOnCall map[int]struct {
		result1 storage.Visibility
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStorage) Broker() storage.Broker {
	fake.brokerMutex.Lock()
	ret, specificReturn := fake.brokerReturnsOnCall[len(fake.brokerArgsForCall)]
	fake.brokerArgsForCall = append(fake.brokerArgsForCall, struct {
	}{})
	fake.recordInvocation("Broker", []interface{}{})
	fake.brokerMutex.Unlock()
	if fake.BrokerStub != nil {
		return fake.BrokerStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.brokerReturns
	return fakeReturns.result1
}

func (fake *FakeStorage) BrokerCallCount() int {
	fake.brokerMutex.RLock()
	defer fake.brokerMutex.RUnlock()
	return len(fake.brokerArgsForCall)
}

func (fake *FakeStorage) BrokerCalls(stub func() storage.Broker) {
	fake.brokerMutex.Lock()
	defer fake.brokerMutex.Unlock()
	fake.BrokerStub = stub
}

func (fake *FakeStorage) BrokerReturns(result1 storage.Broker) {
	fake.brokerMutex.Lock()
	defer fake.brokerMutex.Unlock()
	fake.BrokerStub = nil
	fake.brokerReturns = struct {
		result1 storage.Broker
	}{result1}
}

func (fake *FakeStorage) BrokerReturnsOnCall(i int, result1 storage.Broker) {
	fake.brokerMutex.Lock()
	defer fake.brokerMutex.Unlock()
	fake.BrokerStub = nil
	if fake.brokerReturnsOnCall == nil {
		fake.brokerReturnsOnCall = make(map[int]struct {
			result1 storage.Broker
		})
	}
	fake.brokerReturnsOnCall[i] = struct {
		result1 storage.Broker
	}{result1}
}

func (fake *FakeStorage) Close() error {
	fake.closeMutex.Lock()
	ret, specificReturn := fake.closeReturnsOnCall[len(fake.closeArgsForCall)]
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
	}{})
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		return fake.CloseStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.closeReturns
	return fakeReturns.result1
}

func (fake *FakeStorage) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeStorage) CloseCalls(stub func() error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *FakeStorage) CloseReturns(result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStorage) CloseReturnsOnCall(i int, result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	if fake.closeReturnsOnCall == nil {
		fake.closeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStorage) Credentials() storage.Credentials {
	fake.credentialsMutex.Lock()
	ret, specificReturn := fake.credentialsReturnsOnCall[len(fake.credentialsArgsForCall)]
	fake.credentialsArgsForCall = append(fake.credentialsArgsForCall, struct {
	}{})
	fake.recordInvocation("Credentials", []interface{}{})
	fake.credentialsMutex.Unlock()
	if fake.CredentialsStub != nil {
		return fake.CredentialsStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.credentialsReturns
	return fakeReturns.result1
}

func (fake *FakeStorage) CredentialsCallCount() int {
	fake.credentialsMutex.RLock()
	defer fake.credentialsMutex.RUnlock()
	return len(fake.credentialsArgsForCall)
}

func (fake *FakeStorage) CredentialsCalls(stub func() storage.Credentials) {
	fake.credentialsMutex.Lock()
	defer fake.credentialsMutex.Unlock()
	fake.CredentialsStub = stub
}

func (fake *FakeStorage) CredentialsReturns(result1 storage.Credentials) {
	fake.credentialsMutex.Lock()
	defer fake.credentialsMutex.Unlock()
	fake.CredentialsStub = nil
	fake.credentialsReturns = struct {
		result1 storage.Credentials
	}{result1}
}

func (fake *FakeStorage) CredentialsReturnsOnCall(i int, result1 storage.Credentials) {
	fake.credentialsMutex.Lock()
	defer fake.credentialsMutex.Unlock()
	fake.CredentialsStub = nil
	if fake.credentialsReturnsOnCall == nil {
		fake.credentialsReturnsOnCall = make(map[int]struct {
			result1 storage.Credentials
		})
	}
	fake.credentialsReturnsOnCall[i] = struct {
		result1 storage.Credentials
	}{result1}
}

func (fake *FakeStorage) InTransaction(arg1 context.Context, arg2 func(ctx context.Context, storage storage.Warehouse) error) error {
	fake.inTransactionMutex.Lock()
	ret, specificReturn := fake.inTransactionReturnsOnCall[len(fake.inTransactionArgsForCall)]
	fake.inTransactionArgsForCall = append(fake.inTransactionArgsForCall, struct {
		arg1 context.Context
		arg2 func(ctx context.Context, storage storage.Warehouse) error
	}{arg1, arg2})
	fake.recordInvocation("InTransaction", []interface{}{arg1, arg2})
	fake.inTransactionMutex.Unlock()
	if fake.InTransactionStub != nil {
		return fake.InTransactionStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.inTransactionReturns
	return fakeReturns.result1
}

func (fake *FakeStorage) InTransactionCallCount() int {
	fake.inTransactionMutex.RLock()
	defer fake.inTransactionMutex.RUnlock()
	return len(fake.inTransactionArgsForCall)
}

func (fake *FakeStorage) InTransactionCalls(stub func(context.Context, func(ctx context.Context, storage storage.Warehouse) error) error) {
	fake.inTransactionMutex.Lock()
	defer fake.inTransactionMutex.Unlock()
	fake.InTransactionStub = stub
}

func (fake *FakeStorage) InTransactionArgsForCall(i int) (context.Context, func(ctx context.Context, storage storage.Warehouse) error) {
	fake.inTransactionMutex.RLock()
	defer fake.inTransactionMutex.RUnlock()
	argsForCall := fake.inTransactionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStorage) InTransactionReturns(result1 error) {
	fake.inTransactionMutex.Lock()
	defer fake.inTransactionMutex.Unlock()
	fake.InTransactionStub = nil
	fake.inTransactionReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStorage) InTransactionReturnsOnCall(i int, result1 error) {
	fake.inTransactionMutex.Lock()
	defer fake.inTransactionMutex.Unlock()
	fake.InTransactionStub = nil
	if fake.inTransactionReturnsOnCall == nil {
		fake.inTransactionReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.inTransactionReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStorage) Open(arg1 *storage.Settings) error {
	fake.openMutex.Lock()
	ret, specificReturn := fake.openReturnsOnCall[len(fake.openArgsForCall)]
	fake.openArgsForCall = append(fake.openArgsForCall, struct {
		arg1 *storage.Settings
	}{arg1})
	fake.recordInvocation("Open", []interface{}{arg1})
	fake.openMutex.Unlock()
	if fake.OpenStub != nil {
		return fake.OpenStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.openReturns
	return fakeReturns.result1
}

func (fake *FakeStorage) OpenCallCount() int {
	fake.openMutex.RLock()
	defer fake.openMutex.RUnlock()
	return len(fake.openArgsForCall)
}

func (fake *FakeStorage) OpenCalls(stub func(*storage.Settings) error) {
	fake.openMutex.Lock()
	defer fake.openMutex.Unlock()
	fake.OpenStub = stub
}

func (fake *FakeStorage) OpenArgsForCall(i int) *storage.Settings {
	fake.openMutex.RLock()
	defer fake.openMutex.RUnlock()
	argsForCall := fake.openArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeStorage) OpenReturns(result1 error) {
	fake.openMutex.Lock()
	defer fake.openMutex.Unlock()
	fake.OpenStub = nil
	fake.openReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStorage) OpenReturnsOnCall(i int, result1 error) {
	fake.openMutex.Lock()
	defer fake.openMutex.Unlock()
	fake.OpenStub = nil
	if fake.openReturnsOnCall == nil {
		fake.openReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.openReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStorage) Ping() error {
	fake.pingMutex.Lock()
	ret, specificReturn := fake.pingReturnsOnCall[len(fake.pingArgsForCall)]
	fake.pingArgsForCall = append(fake.pingArgsForCall, struct {
	}{})
	fake.recordInvocation("Ping", []interface{}{})
	fake.pingMutex.Unlock()
	if fake.PingStub != nil {
		return fake.PingStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.pingReturns
	return fakeReturns.result1
}

func (fake *FakeStorage) PingCallCount() int {
	fake.pingMutex.RLock()
	defer fake.pingMutex.RUnlock()
	return len(fake.pingArgsForCall)
}

func (fake *FakeStorage) PingCalls(stub func() error) {
	fake.pingMutex.Lock()
	defer fake.pingMutex.Unlock()
	fake.PingStub = stub
}

func (fake *FakeStorage) PingReturns(result1 error) {
	fake.pingMutex.Lock()
	defer fake.pingMutex.Unlock()
	fake.PingStub = nil
	fake.pingReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStorage) PingReturnsOnCall(i int, result1 error) {
	fake.pingMutex.Lock()
	defer fake.pingMutex.Unlock()
	fake.PingStub = nil
	if fake.pingReturnsOnCall == nil {
		fake.pingReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.pingReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStorage) Platform() storage.Platform {
	fake.platformMutex.Lock()
	ret, specificReturn := fake.platformReturnsOnCall[len(fake.platformArgsForCall)]
	fake.platformArgsForCall = append(fake.platformArgsForCall, struct {
	}{})
	fake.recordInvocation("Platform", []interface{}{})
	fake.platformMutex.Unlock()
	if fake.PlatformStub != nil {
		return fake.PlatformStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.platformReturns
	return fakeReturns.result1
}

func (fake *FakeStorage) PlatformCallCount() int {
	fake.platformMutex.RLock()
	defer fake.platformMutex.RUnlock()
	return len(fake.platformArgsForCall)
}

func (fake *FakeStorage) PlatformCalls(stub func() storage.Platform) {
	fake.platformMutex.Lock()
	defer fake.platformMutex.Unlock()
	fake.PlatformStub = stub
}

func (fake *FakeStorage) PlatformReturns(result1 storage.Platform) {
	fake.platformMutex.Lock()
	defer fake.platformMutex.Unlock()
	fake.PlatformStub = nil
	fake.platformReturns = struct {
		result1 storage.Platform
	}{result1}
}

func (fake *FakeStorage) PlatformReturnsOnCall(i int, result1 storage.Platform) {
	fake.platformMutex.Lock()
	defer fake.platformMutex.Unlock()
	fake.PlatformStub = nil
	if fake.platformReturnsOnCall == nil {
		fake.platformReturnsOnCall = make(map[int]struct {
			result1 storage.Platform
		})
	}
	fake.platformReturnsOnCall[i] = struct {
		result1 storage.Platform
	}{result1}
}

func (fake *FakeStorage) Security() storage.Security {
	fake.securityMutex.Lock()
	ret, specificReturn := fake.securityReturnsOnCall[len(fake.securityArgsForCall)]
	fake.securityArgsForCall = append(fake.securityArgsForCall, struct {
	}{})
	fake.recordInvocation("Security", []interface{}{})
	fake.securityMutex.Unlock()
	if fake.SecurityStub != nil {
		return fake.SecurityStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.securityReturns
	return fakeReturns.result1
}

func (fake *FakeStorage) SecurityCallCount() int {
	fake.securityMutex.RLock()
	defer fake.securityMutex.RUnlock()
	return len(fake.securityArgsForCall)
}

func (fake *FakeStorage) SecurityCalls(stub func() storage.Security) {
	fake.securityMutex.Lock()
	defer fake.securityMutex.Unlock()
	fake.SecurityStub = stub
}

func (fake *FakeStorage) SecurityReturns(result1 storage.Security) {
	fake.securityMutex.Lock()
	defer fake.securityMutex.Unlock()
	fake.SecurityStub = nil
	fake.securityReturns = struct {
		result1 storage.Security
	}{result1}
}

func (fake *FakeStorage) SecurityReturnsOnCall(i int, result1 storage.Security) {
	fake.securityMutex.Lock()
	defer fake.securityMutex.Unlock()
	fake.SecurityStub = nil
	if fake.securityReturnsOnCall == nil {
		fake.securityReturnsOnCall = make(map[int]struct {
			result1 storage.Security
		})
	}
	fake.securityReturnsOnCall[i] = struct {
		result1 storage.Security
	}{result1}
}

func (fake *FakeStorage) ServiceOffering() storage.ServiceOffering {
	fake.serviceOfferingMutex.Lock()
	ret, specificReturn := fake.serviceOfferingReturnsOnCall[len(fake.serviceOfferingArgsForCall)]
	fake.serviceOfferingArgsForCall = append(fake.serviceOfferingArgsForCall, struct {
	}{})
	fake.recordInvocation("ServiceOffering", []interface{}{})
	fake.serviceOfferingMutex.Unlock()
	if fake.ServiceOfferingStub != nil {
		return fake.ServiceOfferingStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.serviceOfferingReturns
	return fakeReturns.result1
}

func (fake *FakeStorage) ServiceOfferingCallCount() int {
	fake.serviceOfferingMutex.RLock()
	defer fake.serviceOfferingMutex.RUnlock()
	return len(fake.serviceOfferingArgsForCall)
}

func (fake *FakeStorage) ServiceOfferingCalls(stub func() storage.ServiceOffering) {
	fake.serviceOfferingMutex.Lock()
	defer fake.serviceOfferingMutex.Unlock()
	fake.ServiceOfferingStub = stub
}

func (fake *FakeStorage) ServiceOfferingReturns(result1 storage.ServiceOffering) {
	fake.serviceOfferingMutex.Lock()
	defer fake.serviceOfferingMutex.Unlock()
	fake.ServiceOfferingStub = nil
	fake.serviceOfferingReturns = struct {
		result1 storage.ServiceOffering
	}{result1}
}

func (fake *FakeStorage) ServiceOfferingReturnsOnCall(i int, result1 storage.ServiceOffering) {
	fake.serviceOfferingMutex.Lock()
	defer fake.serviceOfferingMutex.Unlock()
	fake.ServiceOfferingStub = nil
	if fake.serviceOfferingReturnsOnCall == nil {
		fake.serviceOfferingReturnsOnCall = make(map[int]struct {
			result1 storage.ServiceOffering
		})
	}
	fake.serviceOfferingReturnsOnCall[i] = struct {
		result1 storage.ServiceOffering
	}{result1}
}

func (fake *FakeStorage) ServicePlan() storage.ServicePlan {
	fake.servicePlanMutex.Lock()
	ret, specificReturn := fake.servicePlanReturnsOnCall[len(fake.servicePlanArgsForCall)]
	fake.servicePlanArgsForCall = append(fake.servicePlanArgsForCall, struct {
	}{})
	fake.recordInvocation("ServicePlan", []interface{}{})
	fake.servicePlanMutex.Unlock()
	if fake.ServicePlanStub != nil {
		return fake.ServicePlanStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.servicePlanReturns
	return fakeReturns.result1
}

func (fake *FakeStorage) ServicePlanCallCount() int {
	fake.servicePlanMutex.RLock()
	defer fake.servicePlanMutex.RUnlock()
	return len(fake.servicePlanArgsForCall)
}

func (fake *FakeStorage) ServicePlanCalls(stub func() storage.ServicePlan) {
	fake.servicePlanMutex.Lock()
	defer fake.servicePlanMutex.Unlock()
	fake.ServicePlanStub = stub
}

func (fake *FakeStorage) ServicePlanReturns(result1 storage.ServicePlan) {
	fake.servicePlanMutex.Lock()
	defer fake.servicePlanMutex.Unlock()
	fake.ServicePlanStub = nil
	fake.servicePlanReturns = struct {
		result1 storage.ServicePlan
	}{result1}
}

func (fake *FakeStorage) ServicePlanReturnsOnCall(i int, result1 storage.ServicePlan) {
	fake.servicePlanMutex.Lock()
	defer fake.servicePlanMutex.Unlock()
	fake.ServicePlanStub = nil
	if fake.servicePlanReturnsOnCall == nil {
		fake.servicePlanReturnsOnCall = make(map[int]struct {
			result1 storage.ServicePlan
		})
	}
	fake.servicePlanReturnsOnCall[i] = struct {
		result1 storage.ServicePlan
	}{result1}
}

func (fake *FakeStorage) Visibility() storage.Visibility {
	fake.visibilityMutex.Lock()
	ret, specificReturn := fake.visibilityReturnsOnCall[len(fake.visibilityArgsForCall)]
	fake.visibilityArgsForCall = append(fake.visibilityArgsForCall, struct {
	}{})
	fake.recordInvocation("Visibility", []interface{}{})
	fake.visibilityMutex.Unlock()
	if fake.VisibilityStub != nil {
		return fake.VisibilityStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.visibilityReturns
	return fakeReturns.result1
}

func (fake *FakeStorage) VisibilityCallCount() int {
	fake.visibilityMutex.RLock()
	defer fake.visibilityMutex.RUnlock()
	return len(fake.visibilityArgsForCall)
}

func (fake *FakeStorage) VisibilityCalls(stub func() storage.Visibility) {
	fake.visibilityMutex.Lock()
	defer fake.visibilityMutex.Unlock()
	fake.VisibilityStub = stub
}

func (fake *FakeStorage) VisibilityReturns(result1 storage.Visibility) {
	fake.visibilityMutex.Lock()
	defer fake.visibilityMutex.Unlock()
	fake.VisibilityStub = nil
	fake.visibilityReturns = struct {
		result1 storage.Visibility
	}{result1}
}

func (fake *FakeStorage) VisibilityReturnsOnCall(i int, result1 storage.Visibility) {
	fake.visibilityMutex.Lock()
	defer fake.visibilityMutex.Unlock()
	fake.VisibilityStub = nil
	if fake.visibilityReturnsOnCall == nil {
		fake.visibilityReturnsOnCall = make(map[int]struct {
			result1 storage.Visibility
		})
	}
	fake.visibilityReturnsOnCall[i] = struct {
		result1 storage.Visibility
	}{result1}
}

func (fake *FakeStorage) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.brokerMutex.RLock()
	defer fake.brokerMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.credentialsMutex.RLock()
	defer fake.credentialsMutex.RUnlock()
	fake.inTransactionMutex.RLock()
	defer fake.inTransactionMutex.RUnlock()
	fake.openMutex.RLock()
	defer fake.openMutex.RUnlock()
	fake.pingMutex.RLock()
	defer fake.pingMutex.RUnlock()
	fake.platformMutex.RLock()
	defer fake.platformMutex.RUnlock()
	fake.securityMutex.RLock()
	defer fake.securityMutex.RUnlock()
	fake.serviceOfferingMutex.RLock()
	defer fake.serviceOfferingMutex.RUnlock()
	fake.servicePlanMutex.RLock()
	defer fake.servicePlanMutex.RUnlock()
	fake.visibilityMutex.RLock()
	defer fake.visibilityMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeStorage) recordInvocation(key string, args []interface{}) {
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

var _ storage.Storage = new(FakeStorage)
