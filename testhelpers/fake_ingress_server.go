// This Code was generated by counterfeiter, but we had to edit it - see below.
// This is a result of gRPC breaking change referenced here: https://github.com/cloudfoundry/go-loggregator/issues/146
// The command to generate this file is: `counterfeiter -o testhelpers/ vendor/code.cloudfoundry.org/go-loggregator/v9/rpc/loggregator_v2/ingress_grpc.pb.go IngressServer`
// After generating the file, loggregator_v2.UnimplementedIngressServer will need to be embedded in the FakeIngressServer struct.
// If this file needs to be generated again, the above changes will need to be reapplied.
// In order to pass the linter checks the function `(fake *FakeIngressServer) mustEmbedUnimplementedIngressServer()` should be removed.
package testhelpers

import (
	"context"
	"sync"

	"code.cloudfoundry.org/go-loggregator/v9/rpc/loggregator_v2"
)

type FakeIngressServer struct {
	loggregator_v2.UnimplementedIngressServer
	BatchSenderStub        func(loggregator_v2.Ingress_BatchSenderServer) error
	batchSenderMutex       sync.RWMutex
	batchSenderArgsForCall []struct {
		arg1 loggregator_v2.Ingress_BatchSenderServer
	}
	batchSenderReturns struct {
		result1 error
	}
	batchSenderReturnsOnCall map[int]struct {
		result1 error
	}
	SendStub        func(context.Context, *loggregator_v2.EnvelopeBatch) (*loggregator_v2.SendResponse, error)
	sendMutex       sync.RWMutex
	sendArgsForCall []struct {
		arg1 context.Context
		arg2 *loggregator_v2.EnvelopeBatch
	}
	sendReturns struct {
		result1 *loggregator_v2.SendResponse
		result2 error
	}
	sendReturnsOnCall map[int]struct {
		result1 *loggregator_v2.SendResponse
		result2 error
	}
	SenderStub        func(loggregator_v2.Ingress_SenderServer) error
	senderMutex       sync.RWMutex
	senderArgsForCall []struct {
		arg1 loggregator_v2.Ingress_SenderServer
	}
	senderReturns struct {
		result1 error
	}
	senderReturnsOnCall map[int]struct {
		result1 error
	}
	mustEmbedUnimplementedIngressServerStub        func()
	mustEmbedUnimplementedIngressServerMutex       sync.RWMutex
	mustEmbedUnimplementedIngressServerArgsForCall []struct {
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeIngressServer) BatchSender(arg1 loggregator_v2.Ingress_BatchSenderServer) error {
	fake.batchSenderMutex.Lock()
	ret, specificReturn := fake.batchSenderReturnsOnCall[len(fake.batchSenderArgsForCall)]
	fake.batchSenderArgsForCall = append(fake.batchSenderArgsForCall, struct {
		arg1 loggregator_v2.Ingress_BatchSenderServer
	}{arg1})
	stub := fake.BatchSenderStub
	fakeReturns := fake.batchSenderReturns
	fake.recordInvocation("BatchSender", []interface{}{arg1})
	fake.batchSenderMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIngressServer) BatchSenderCallCount() int {
	fake.batchSenderMutex.RLock()
	defer fake.batchSenderMutex.RUnlock()
	return len(fake.batchSenderArgsForCall)
}

func (fake *FakeIngressServer) BatchSenderCalls(stub func(loggregator_v2.Ingress_BatchSenderServer) error) {
	fake.batchSenderMutex.Lock()
	defer fake.batchSenderMutex.Unlock()
	fake.BatchSenderStub = stub
}

func (fake *FakeIngressServer) BatchSenderArgsForCall(i int) loggregator_v2.Ingress_BatchSenderServer {
	fake.batchSenderMutex.RLock()
	defer fake.batchSenderMutex.RUnlock()
	argsForCall := fake.batchSenderArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeIngressServer) BatchSenderReturns(result1 error) {
	fake.batchSenderMutex.Lock()
	defer fake.batchSenderMutex.Unlock()
	fake.BatchSenderStub = nil
	fake.batchSenderReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIngressServer) BatchSenderReturnsOnCall(i int, result1 error) {
	fake.batchSenderMutex.Lock()
	defer fake.batchSenderMutex.Unlock()
	fake.BatchSenderStub = nil
	if fake.batchSenderReturnsOnCall == nil {
		fake.batchSenderReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.batchSenderReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIngressServer) Send(arg1 context.Context, arg2 *loggregator_v2.EnvelopeBatch) (*loggregator_v2.SendResponse, error) {
	fake.sendMutex.Lock()
	ret, specificReturn := fake.sendReturnsOnCall[len(fake.sendArgsForCall)]
	fake.sendArgsForCall = append(fake.sendArgsForCall, struct {
		arg1 context.Context
		arg2 *loggregator_v2.EnvelopeBatch
	}{arg1, arg2})
	stub := fake.SendStub
	fakeReturns := fake.sendReturns
	fake.recordInvocation("Send", []interface{}{arg1, arg2})
	fake.sendMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeIngressServer) SendCallCount() int {
	fake.sendMutex.RLock()
	defer fake.sendMutex.RUnlock()
	return len(fake.sendArgsForCall)
}

func (fake *FakeIngressServer) SendCalls(stub func(context.Context, *loggregator_v2.EnvelopeBatch) (*loggregator_v2.SendResponse, error)) {
	fake.sendMutex.Lock()
	defer fake.sendMutex.Unlock()
	fake.SendStub = stub
}

func (fake *FakeIngressServer) SendArgsForCall(i int) (context.Context, *loggregator_v2.EnvelopeBatch) {
	fake.sendMutex.RLock()
	defer fake.sendMutex.RUnlock()
	argsForCall := fake.sendArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeIngressServer) SendReturns(result1 *loggregator_v2.SendResponse, result2 error) {
	fake.sendMutex.Lock()
	defer fake.sendMutex.Unlock()
	fake.SendStub = nil
	fake.sendReturns = struct {
		result1 *loggregator_v2.SendResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeIngressServer) SendReturnsOnCall(i int, result1 *loggregator_v2.SendResponse, result2 error) {
	fake.sendMutex.Lock()
	defer fake.sendMutex.Unlock()
	fake.SendStub = nil
	if fake.sendReturnsOnCall == nil {
		fake.sendReturnsOnCall = make(map[int]struct {
			result1 *loggregator_v2.SendResponse
			result2 error
		})
	}
	fake.sendReturnsOnCall[i] = struct {
		result1 *loggregator_v2.SendResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeIngressServer) Sender(arg1 loggregator_v2.Ingress_SenderServer) error {
	fake.senderMutex.Lock()
	ret, specificReturn := fake.senderReturnsOnCall[len(fake.senderArgsForCall)]
	fake.senderArgsForCall = append(fake.senderArgsForCall, struct {
		arg1 loggregator_v2.Ingress_SenderServer
	}{arg1})
	stub := fake.SenderStub
	fakeReturns := fake.senderReturns
	fake.recordInvocation("Sender", []interface{}{arg1})
	fake.senderMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIngressServer) SenderCallCount() int {
	fake.senderMutex.RLock()
	defer fake.senderMutex.RUnlock()
	return len(fake.senderArgsForCall)
}

func (fake *FakeIngressServer) SenderCalls(stub func(loggregator_v2.Ingress_SenderServer) error) {
	fake.senderMutex.Lock()
	defer fake.senderMutex.Unlock()
	fake.SenderStub = stub
}

func (fake *FakeIngressServer) SenderArgsForCall(i int) loggregator_v2.Ingress_SenderServer {
	fake.senderMutex.RLock()
	defer fake.senderMutex.RUnlock()
	argsForCall := fake.senderArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeIngressServer) SenderReturns(result1 error) {
	fake.senderMutex.Lock()
	defer fake.senderMutex.Unlock()
	fake.SenderStub = nil
	fake.senderReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIngressServer) SenderReturnsOnCall(i int, result1 error) {
	fake.senderMutex.Lock()
	defer fake.senderMutex.Unlock()
	fake.SenderStub = nil
	if fake.senderReturnsOnCall == nil {
		fake.senderReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.senderReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIngressServer) MustEmbedUnimplementedIngressServerCallCount() int {
	fake.mustEmbedUnimplementedIngressServerMutex.RLock()
	defer fake.mustEmbedUnimplementedIngressServerMutex.RUnlock()
	return len(fake.mustEmbedUnimplementedIngressServerArgsForCall)
}

func (fake *FakeIngressServer) MustEmbedUnimplementedIngressServerCalls(stub func()) {
	fake.mustEmbedUnimplementedIngressServerMutex.Lock()
	defer fake.mustEmbedUnimplementedIngressServerMutex.Unlock()
	fake.mustEmbedUnimplementedIngressServerStub = stub
}

func (fake *FakeIngressServer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.batchSenderMutex.RLock()
	defer fake.batchSenderMutex.RUnlock()
	fake.sendMutex.RLock()
	defer fake.sendMutex.RUnlock()
	fake.senderMutex.RLock()
	defer fake.senderMutex.RUnlock()
	fake.mustEmbedUnimplementedIngressServerMutex.RLock()
	defer fake.mustEmbedUnimplementedIngressServerMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeIngressServer) recordInvocation(key string, args []interface{}) {
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

var _ loggregator_v2.IngressServer = new(FakeIngressServer)
