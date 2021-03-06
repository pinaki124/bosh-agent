// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry/yagnats"
)

type FakeNATSClient struct {
	PingStub        func() bool
	pingMutex       sync.RWMutex
	pingArgsForCall []struct{}
	pingReturns     struct {
		result1 bool
	}
	pingReturnsOnCall map[int]struct {
		result1 bool
	}
	ConnectStub        func(connectionProvider yagnats.ConnectionProvider) error
	connectMutex       sync.RWMutex
	connectArgsForCall []struct {
		connectionProvider yagnats.ConnectionProvider
	}
	connectReturns struct {
		result1 error
	}
	connectReturnsOnCall map[int]struct {
		result1 error
	}
	DisconnectStub        func()
	disconnectMutex       sync.RWMutex
	disconnectArgsForCall []struct{}
	PublishStub           func(subject string, payload []byte) error
	publishMutex          sync.RWMutex
	publishArgsForCall    []struct {
		subject string
		payload []byte
	}
	publishReturns struct {
		result1 error
	}
	publishReturnsOnCall map[int]struct {
		result1 error
	}
	PublishWithReplyToStub        func(subject, reply string, payload []byte) error
	publishWithReplyToMutex       sync.RWMutex
	publishWithReplyToArgsForCall []struct {
		subject string
		reply   string
		payload []byte
	}
	publishWithReplyToReturns struct {
		result1 error
	}
	publishWithReplyToReturnsOnCall map[int]struct {
		result1 error
	}
	SubscribeStub        func(subject string, callback yagnats.Callback) (int64, error)
	subscribeMutex       sync.RWMutex
	subscribeArgsForCall []struct {
		subject  string
		callback yagnats.Callback
	}
	subscribeReturns struct {
		result1 int64
		result2 error
	}
	subscribeReturnsOnCall map[int]struct {
		result1 int64
		result2 error
	}
	SubscribeWithQueueStub        func(subject, queue string, callback yagnats.Callback) (int64, error)
	subscribeWithQueueMutex       sync.RWMutex
	subscribeWithQueueArgsForCall []struct {
		subject  string
		queue    string
		callback yagnats.Callback
	}
	subscribeWithQueueReturns struct {
		result1 int64
		result2 error
	}
	subscribeWithQueueReturnsOnCall map[int]struct {
		result1 int64
		result2 error
	}
	UnsubscribeStub        func(subscription int64) error
	unsubscribeMutex       sync.RWMutex
	unsubscribeArgsForCall []struct {
		subscription int64
	}
	unsubscribeReturns struct {
		result1 error
	}
	unsubscribeReturnsOnCall map[int]struct {
		result1 error
	}
	UnsubscribeAllStub        func(subject string)
	unsubscribeAllMutex       sync.RWMutex
	unsubscribeAllArgsForCall []struct {
		subject string
	}
	BeforeConnectCallbackStub        func(callback func())
	beforeConnectCallbackMutex       sync.RWMutex
	beforeConnectCallbackArgsForCall []struct {
		callback func()
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeNATSClient) Ping() bool {
	fake.pingMutex.Lock()
	ret, specificReturn := fake.pingReturnsOnCall[len(fake.pingArgsForCall)]
	fake.pingArgsForCall = append(fake.pingArgsForCall, struct{}{})
	fake.recordInvocation("Ping", []interface{}{})
	fake.pingMutex.Unlock()
	if fake.PingStub != nil {
		return fake.PingStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.pingReturns.result1
}

func (fake *FakeNATSClient) PingCallCount() int {
	fake.pingMutex.RLock()
	defer fake.pingMutex.RUnlock()
	return len(fake.pingArgsForCall)
}

func (fake *FakeNATSClient) PingReturns(result1 bool) {
	fake.PingStub = nil
	fake.pingReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeNATSClient) PingReturnsOnCall(i int, result1 bool) {
	fake.PingStub = nil
	if fake.pingReturnsOnCall == nil {
		fake.pingReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.pingReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeNATSClient) Connect(connectionProvider yagnats.ConnectionProvider) error {
	fake.connectMutex.Lock()
	ret, specificReturn := fake.connectReturnsOnCall[len(fake.connectArgsForCall)]
	fake.connectArgsForCall = append(fake.connectArgsForCall, struct {
		connectionProvider yagnats.ConnectionProvider
	}{connectionProvider})
	fake.recordInvocation("Connect", []interface{}{connectionProvider})
	fake.connectMutex.Unlock()
	if fake.ConnectStub != nil {
		return fake.ConnectStub(connectionProvider)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.connectReturns.result1
}

func (fake *FakeNATSClient) ConnectCallCount() int {
	fake.connectMutex.RLock()
	defer fake.connectMutex.RUnlock()
	return len(fake.connectArgsForCall)
}

func (fake *FakeNATSClient) ConnectArgsForCall(i int) yagnats.ConnectionProvider {
	fake.connectMutex.RLock()
	defer fake.connectMutex.RUnlock()
	return fake.connectArgsForCall[i].connectionProvider
}

func (fake *FakeNATSClient) ConnectReturns(result1 error) {
	fake.ConnectStub = nil
	fake.connectReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeNATSClient) ConnectReturnsOnCall(i int, result1 error) {
	fake.ConnectStub = nil
	if fake.connectReturnsOnCall == nil {
		fake.connectReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.connectReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeNATSClient) Disconnect() {
	fake.disconnectMutex.Lock()
	fake.disconnectArgsForCall = append(fake.disconnectArgsForCall, struct{}{})
	fake.recordInvocation("Disconnect", []interface{}{})
	fake.disconnectMutex.Unlock()
	if fake.DisconnectStub != nil {
		fake.DisconnectStub()
	}
}

func (fake *FakeNATSClient) DisconnectCallCount() int {
	fake.disconnectMutex.RLock()
	defer fake.disconnectMutex.RUnlock()
	return len(fake.disconnectArgsForCall)
}

func (fake *FakeNATSClient) Publish(subject string, payload []byte) error {
	var payloadCopy []byte
	if payload != nil {
		payloadCopy = make([]byte, len(payload))
		copy(payloadCopy, payload)
	}
	fake.publishMutex.Lock()
	ret, specificReturn := fake.publishReturnsOnCall[len(fake.publishArgsForCall)]
	fake.publishArgsForCall = append(fake.publishArgsForCall, struct {
		subject string
		payload []byte
	}{subject, payloadCopy})
	fake.recordInvocation("Publish", []interface{}{subject, payloadCopy})
	fake.publishMutex.Unlock()
	if fake.PublishStub != nil {
		return fake.PublishStub(subject, payload)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.publishReturns.result1
}

func (fake *FakeNATSClient) PublishCallCount() int {
	fake.publishMutex.RLock()
	defer fake.publishMutex.RUnlock()
	return len(fake.publishArgsForCall)
}

func (fake *FakeNATSClient) PublishArgsForCall(i int) (string, []byte) {
	fake.publishMutex.RLock()
	defer fake.publishMutex.RUnlock()
	return fake.publishArgsForCall[i].subject, fake.publishArgsForCall[i].payload
}

func (fake *FakeNATSClient) PublishReturns(result1 error) {
	fake.PublishStub = nil
	fake.publishReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeNATSClient) PublishReturnsOnCall(i int, result1 error) {
	fake.PublishStub = nil
	if fake.publishReturnsOnCall == nil {
		fake.publishReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.publishReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeNATSClient) PublishWithReplyTo(subject string, reply string, payload []byte) error {
	var payloadCopy []byte
	if payload != nil {
		payloadCopy = make([]byte, len(payload))
		copy(payloadCopy, payload)
	}
	fake.publishWithReplyToMutex.Lock()
	ret, specificReturn := fake.publishWithReplyToReturnsOnCall[len(fake.publishWithReplyToArgsForCall)]
	fake.publishWithReplyToArgsForCall = append(fake.publishWithReplyToArgsForCall, struct {
		subject string
		reply   string
		payload []byte
	}{subject, reply, payloadCopy})
	fake.recordInvocation("PublishWithReplyTo", []interface{}{subject, reply, payloadCopy})
	fake.publishWithReplyToMutex.Unlock()
	if fake.PublishWithReplyToStub != nil {
		return fake.PublishWithReplyToStub(subject, reply, payload)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.publishWithReplyToReturns.result1
}

func (fake *FakeNATSClient) PublishWithReplyToCallCount() int {
	fake.publishWithReplyToMutex.RLock()
	defer fake.publishWithReplyToMutex.RUnlock()
	return len(fake.publishWithReplyToArgsForCall)
}

func (fake *FakeNATSClient) PublishWithReplyToArgsForCall(i int) (string, string, []byte) {
	fake.publishWithReplyToMutex.RLock()
	defer fake.publishWithReplyToMutex.RUnlock()
	return fake.publishWithReplyToArgsForCall[i].subject, fake.publishWithReplyToArgsForCall[i].reply, fake.publishWithReplyToArgsForCall[i].payload
}

func (fake *FakeNATSClient) PublishWithReplyToReturns(result1 error) {
	fake.PublishWithReplyToStub = nil
	fake.publishWithReplyToReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeNATSClient) PublishWithReplyToReturnsOnCall(i int, result1 error) {
	fake.PublishWithReplyToStub = nil
	if fake.publishWithReplyToReturnsOnCall == nil {
		fake.publishWithReplyToReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.publishWithReplyToReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeNATSClient) Subscribe(subject string, callback yagnats.Callback) (int64, error) {
	fake.subscribeMutex.Lock()
	ret, specificReturn := fake.subscribeReturnsOnCall[len(fake.subscribeArgsForCall)]
	fake.subscribeArgsForCall = append(fake.subscribeArgsForCall, struct {
		subject  string
		callback yagnats.Callback
	}{subject, callback})
	fake.recordInvocation("Subscribe", []interface{}{subject, callback})
	fake.subscribeMutex.Unlock()
	if fake.SubscribeStub != nil {
		return fake.SubscribeStub(subject, callback)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.subscribeReturns.result1, fake.subscribeReturns.result2
}

func (fake *FakeNATSClient) SubscribeCallCount() int {
	fake.subscribeMutex.RLock()
	defer fake.subscribeMutex.RUnlock()
	return len(fake.subscribeArgsForCall)
}

func (fake *FakeNATSClient) SubscribeArgsForCall(i int) (string, yagnats.Callback) {
	fake.subscribeMutex.RLock()
	defer fake.subscribeMutex.RUnlock()
	return fake.subscribeArgsForCall[i].subject, fake.subscribeArgsForCall[i].callback
}

func (fake *FakeNATSClient) SubscribeReturns(result1 int64, result2 error) {
	fake.SubscribeStub = nil
	fake.subscribeReturns = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeNATSClient) SubscribeReturnsOnCall(i int, result1 int64, result2 error) {
	fake.SubscribeStub = nil
	if fake.subscribeReturnsOnCall == nil {
		fake.subscribeReturnsOnCall = make(map[int]struct {
			result1 int64
			result2 error
		})
	}
	fake.subscribeReturnsOnCall[i] = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeNATSClient) SubscribeWithQueue(subject string, queue string, callback yagnats.Callback) (int64, error) {
	fake.subscribeWithQueueMutex.Lock()
	ret, specificReturn := fake.subscribeWithQueueReturnsOnCall[len(fake.subscribeWithQueueArgsForCall)]
	fake.subscribeWithQueueArgsForCall = append(fake.subscribeWithQueueArgsForCall, struct {
		subject  string
		queue    string
		callback yagnats.Callback
	}{subject, queue, callback})
	fake.recordInvocation("SubscribeWithQueue", []interface{}{subject, queue, callback})
	fake.subscribeWithQueueMutex.Unlock()
	if fake.SubscribeWithQueueStub != nil {
		return fake.SubscribeWithQueueStub(subject, queue, callback)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.subscribeWithQueueReturns.result1, fake.subscribeWithQueueReturns.result2
}

func (fake *FakeNATSClient) SubscribeWithQueueCallCount() int {
	fake.subscribeWithQueueMutex.RLock()
	defer fake.subscribeWithQueueMutex.RUnlock()
	return len(fake.subscribeWithQueueArgsForCall)
}

func (fake *FakeNATSClient) SubscribeWithQueueArgsForCall(i int) (string, string, yagnats.Callback) {
	fake.subscribeWithQueueMutex.RLock()
	defer fake.subscribeWithQueueMutex.RUnlock()
	return fake.subscribeWithQueueArgsForCall[i].subject, fake.subscribeWithQueueArgsForCall[i].queue, fake.subscribeWithQueueArgsForCall[i].callback
}

func (fake *FakeNATSClient) SubscribeWithQueueReturns(result1 int64, result2 error) {
	fake.SubscribeWithQueueStub = nil
	fake.subscribeWithQueueReturns = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeNATSClient) SubscribeWithQueueReturnsOnCall(i int, result1 int64, result2 error) {
	fake.SubscribeWithQueueStub = nil
	if fake.subscribeWithQueueReturnsOnCall == nil {
		fake.subscribeWithQueueReturnsOnCall = make(map[int]struct {
			result1 int64
			result2 error
		})
	}
	fake.subscribeWithQueueReturnsOnCall[i] = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeNATSClient) Unsubscribe(subscription int64) error {
	fake.unsubscribeMutex.Lock()
	ret, specificReturn := fake.unsubscribeReturnsOnCall[len(fake.unsubscribeArgsForCall)]
	fake.unsubscribeArgsForCall = append(fake.unsubscribeArgsForCall, struct {
		subscription int64
	}{subscription})
	fake.recordInvocation("Unsubscribe", []interface{}{subscription})
	fake.unsubscribeMutex.Unlock()
	if fake.UnsubscribeStub != nil {
		return fake.UnsubscribeStub(subscription)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.unsubscribeReturns.result1
}

func (fake *FakeNATSClient) UnsubscribeCallCount() int {
	fake.unsubscribeMutex.RLock()
	defer fake.unsubscribeMutex.RUnlock()
	return len(fake.unsubscribeArgsForCall)
}

func (fake *FakeNATSClient) UnsubscribeArgsForCall(i int) int64 {
	fake.unsubscribeMutex.RLock()
	defer fake.unsubscribeMutex.RUnlock()
	return fake.unsubscribeArgsForCall[i].subscription
}

func (fake *FakeNATSClient) UnsubscribeReturns(result1 error) {
	fake.UnsubscribeStub = nil
	fake.unsubscribeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeNATSClient) UnsubscribeReturnsOnCall(i int, result1 error) {
	fake.UnsubscribeStub = nil
	if fake.unsubscribeReturnsOnCall == nil {
		fake.unsubscribeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.unsubscribeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeNATSClient) UnsubscribeAll(subject string) {
	fake.unsubscribeAllMutex.Lock()
	fake.unsubscribeAllArgsForCall = append(fake.unsubscribeAllArgsForCall, struct {
		subject string
	}{subject})
	fake.recordInvocation("UnsubscribeAll", []interface{}{subject})
	fake.unsubscribeAllMutex.Unlock()
	if fake.UnsubscribeAllStub != nil {
		fake.UnsubscribeAllStub(subject)
	}
}

func (fake *FakeNATSClient) UnsubscribeAllCallCount() int {
	fake.unsubscribeAllMutex.RLock()
	defer fake.unsubscribeAllMutex.RUnlock()
	return len(fake.unsubscribeAllArgsForCall)
}

func (fake *FakeNATSClient) UnsubscribeAllArgsForCall(i int) string {
	fake.unsubscribeAllMutex.RLock()
	defer fake.unsubscribeAllMutex.RUnlock()
	return fake.unsubscribeAllArgsForCall[i].subject
}

func (fake *FakeNATSClient) BeforeConnectCallback(callback func()) {
	fake.beforeConnectCallbackMutex.Lock()
	fake.beforeConnectCallbackArgsForCall = append(fake.beforeConnectCallbackArgsForCall, struct {
		callback func()
	}{callback})
	fake.recordInvocation("BeforeConnectCallback", []interface{}{callback})
	fake.beforeConnectCallbackMutex.Unlock()
	if fake.BeforeConnectCallbackStub != nil {
		fake.BeforeConnectCallbackStub(callback)
	}
}

func (fake *FakeNATSClient) BeforeConnectCallbackCallCount() int {
	fake.beforeConnectCallbackMutex.RLock()
	defer fake.beforeConnectCallbackMutex.RUnlock()
	return len(fake.beforeConnectCallbackArgsForCall)
}

func (fake *FakeNATSClient) BeforeConnectCallbackArgsForCall(i int) func() {
	fake.beforeConnectCallbackMutex.RLock()
	defer fake.beforeConnectCallbackMutex.RUnlock()
	return fake.beforeConnectCallbackArgsForCall[i].callback
}

func (fake *FakeNATSClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.pingMutex.RLock()
	defer fake.pingMutex.RUnlock()
	fake.connectMutex.RLock()
	defer fake.connectMutex.RUnlock()
	fake.disconnectMutex.RLock()
	defer fake.disconnectMutex.RUnlock()
	fake.publishMutex.RLock()
	defer fake.publishMutex.RUnlock()
	fake.publishWithReplyToMutex.RLock()
	defer fake.publishWithReplyToMutex.RUnlock()
	fake.subscribeMutex.RLock()
	defer fake.subscribeMutex.RUnlock()
	fake.subscribeWithQueueMutex.RLock()
	defer fake.subscribeWithQueueMutex.RUnlock()
	fake.unsubscribeMutex.RLock()
	defer fake.unsubscribeMutex.RUnlock()
	fake.unsubscribeAllMutex.RLock()
	defer fake.unsubscribeAllMutex.RUnlock()
	fake.beforeConnectCallbackMutex.RLock()
	defer fake.beforeConnectCallbackMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeNATSClient) recordInvocation(key string, args []interface{}) {
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

var _ yagnats.NATSClient = new(FakeNATSClient)
