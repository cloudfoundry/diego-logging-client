// Code generated by counterfeiter. DO NOT EDIT.
package testhelpers

import (
	"sync"
	"time"

	diego_logging_client "code.cloudfoundry.org/diego-logging-client"
	loggregator "code.cloudfoundry.org/go-loggregator/v8"
)

type FakeIngressClient struct {
	IncrementCounterStub        func(string) error
	incrementCounterMutex       sync.RWMutex
	incrementCounterArgsForCall []struct {
		arg1 string
	}
	incrementCounterReturns struct {
		result1 error
	}
	incrementCounterReturnsOnCall map[int]struct {
		result1 error
	}
	IncrementCounterWithDeltaStub        func(string, uint64) error
	incrementCounterWithDeltaMutex       sync.RWMutex
	incrementCounterWithDeltaArgsForCall []struct {
		arg1 string
		arg2 uint64
	}
	incrementCounterWithDeltaReturns struct {
		result1 error
	}
	incrementCounterWithDeltaReturnsOnCall map[int]struct {
		result1 error
	}
	SendAppErrorLogStub        func(string, string, map[string]string) error
	sendAppErrorLogMutex       sync.RWMutex
	sendAppErrorLogArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 map[string]string
	}
	sendAppErrorLogReturns struct {
		result1 error
	}
	sendAppErrorLogReturnsOnCall map[int]struct {
		result1 error
	}
	SendAppLogStub        func(string, string, map[string]string) error
	sendAppLogMutex       sync.RWMutex
	sendAppLogArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 map[string]string
	}
	sendAppLogReturns struct {
		result1 error
	}
	sendAppLogReturnsOnCall map[int]struct {
		result1 error
	}
	SendAppLogRateStub        func(float64, float64, map[string]string) error
	sendAppLogRateMutex       sync.RWMutex
	sendAppLogRateArgsForCall []struct {
		arg1 float64
		arg2 float64
		arg3 map[string]string
	}
	sendAppLogRateReturns struct {
		result1 error
	}
	sendAppLogRateReturnsOnCall map[int]struct {
		result1 error
	}
	SendAppMetricsStub        func(diego_logging_client.ContainerMetric) error
	sendAppMetricsMutex       sync.RWMutex
	sendAppMetricsArgsForCall []struct {
		arg1 diego_logging_client.ContainerMetric
	}
	sendAppMetricsReturns struct {
		result1 error
	}
	sendAppMetricsReturnsOnCall map[int]struct {
		result1 error
	}
	SendBytesPerSecondStub        func(string, float64) error
	sendBytesPerSecondMutex       sync.RWMutex
	sendBytesPerSecondArgsForCall []struct {
		arg1 string
		arg2 float64
	}
	sendBytesPerSecondReturns struct {
		result1 error
	}
	sendBytesPerSecondReturnsOnCall map[int]struct {
		result1 error
	}
	SendComponentMetricStub        func(string, float64, string) error
	sendComponentMetricMutex       sync.RWMutex
	sendComponentMetricArgsForCall []struct {
		arg1 string
		arg2 float64
		arg3 string
	}
	sendComponentMetricReturns struct {
		result1 error
	}
	sendComponentMetricReturnsOnCall map[int]struct {
		result1 error
	}
	SendDurationStub        func(string, time.Duration, ...loggregator.EmitGaugeOption) error
	sendDurationMutex       sync.RWMutex
	sendDurationArgsForCall []struct {
		arg1 string
		arg2 time.Duration
		arg3 []loggregator.EmitGaugeOption
	}
	sendDurationReturns struct {
		result1 error
	}
	sendDurationReturnsOnCall map[int]struct {
		result1 error
	}
	SendMebiBytesStub        func(string, int, ...loggregator.EmitGaugeOption) error
	sendMebiBytesMutex       sync.RWMutex
	sendMebiBytesArgsForCall []struct {
		arg1 string
		arg2 int
		arg3 []loggregator.EmitGaugeOption
	}
	sendMebiBytesReturns struct {
		result1 error
	}
	sendMebiBytesReturnsOnCall map[int]struct {
		result1 error
	}
	SendMetricStub        func(string, int, ...loggregator.EmitGaugeOption) error
	sendMetricMutex       sync.RWMutex
	sendMetricArgsForCall []struct {
		arg1 string
		arg2 int
		arg3 []loggregator.EmitGaugeOption
	}
	sendMetricReturns struct {
		result1 error
	}
	sendMetricReturnsOnCall map[int]struct {
		result1 error
	}
	SendRequestsPerSecondStub        func(string, float64) error
	sendRequestsPerSecondMutex       sync.RWMutex
	sendRequestsPerSecondArgsForCall []struct {
		arg1 string
		arg2 float64
	}
	sendRequestsPerSecondReturns struct {
		result1 error
	}
	sendRequestsPerSecondReturnsOnCall map[int]struct {
		result1 error
	}
	SendSpikeMetricsStub        func(diego_logging_client.SpikeMetric) error
	sendSpikeMetricsMutex       sync.RWMutex
	sendSpikeMetricsArgsForCall []struct {
		arg1 diego_logging_client.SpikeMetric
	}
	sendSpikeMetricsReturns struct {
		result1 error
	}
	sendSpikeMetricsReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeIngressClient) IncrementCounter(arg1 string) error {
	fake.incrementCounterMutex.Lock()
	ret, specificReturn := fake.incrementCounterReturnsOnCall[len(fake.incrementCounterArgsForCall)]
	fake.incrementCounterArgsForCall = append(fake.incrementCounterArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.IncrementCounterStub
	fakeReturns := fake.incrementCounterReturns
	fake.recordInvocation("IncrementCounter", []interface{}{arg1})
	fake.incrementCounterMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIngressClient) IncrementCounterCallCount() int {
	fake.incrementCounterMutex.RLock()
	defer fake.incrementCounterMutex.RUnlock()
	return len(fake.incrementCounterArgsForCall)
}

func (fake *FakeIngressClient) IncrementCounterCalls(stub func(string) error) {
	fake.incrementCounterMutex.Lock()
	defer fake.incrementCounterMutex.Unlock()
	fake.IncrementCounterStub = stub
}

func (fake *FakeIngressClient) IncrementCounterArgsForCall(i int) string {
	fake.incrementCounterMutex.RLock()
	defer fake.incrementCounterMutex.RUnlock()
	argsForCall := fake.incrementCounterArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeIngressClient) IncrementCounterReturns(result1 error) {
	fake.incrementCounterMutex.Lock()
	defer fake.incrementCounterMutex.Unlock()
	fake.IncrementCounterStub = nil
	fake.incrementCounterReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIngressClient) IncrementCounterReturnsOnCall(i int, result1 error) {
	fake.incrementCounterMutex.Lock()
	defer fake.incrementCounterMutex.Unlock()
	fake.IncrementCounterStub = nil
	if fake.incrementCounterReturnsOnCall == nil {
		fake.incrementCounterReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.incrementCounterReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIngressClient) IncrementCounterWithDelta(arg1 string, arg2 uint64) error {
	fake.incrementCounterWithDeltaMutex.Lock()
	ret, specificReturn := fake.incrementCounterWithDeltaReturnsOnCall[len(fake.incrementCounterWithDeltaArgsForCall)]
	fake.incrementCounterWithDeltaArgsForCall = append(fake.incrementCounterWithDeltaArgsForCall, struct {
		arg1 string
		arg2 uint64
	}{arg1, arg2})
	stub := fake.IncrementCounterWithDeltaStub
	fakeReturns := fake.incrementCounterWithDeltaReturns
	fake.recordInvocation("IncrementCounterWithDelta", []interface{}{arg1, arg2})
	fake.incrementCounterWithDeltaMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIngressClient) IncrementCounterWithDeltaCallCount() int {
	fake.incrementCounterWithDeltaMutex.RLock()
	defer fake.incrementCounterWithDeltaMutex.RUnlock()
	return len(fake.incrementCounterWithDeltaArgsForCall)
}

func (fake *FakeIngressClient) IncrementCounterWithDeltaCalls(stub func(string, uint64) error) {
	fake.incrementCounterWithDeltaMutex.Lock()
	defer fake.incrementCounterWithDeltaMutex.Unlock()
	fake.IncrementCounterWithDeltaStub = stub
}

func (fake *FakeIngressClient) IncrementCounterWithDeltaArgsForCall(i int) (string, uint64) {
	fake.incrementCounterWithDeltaMutex.RLock()
	defer fake.incrementCounterWithDeltaMutex.RUnlock()
	argsForCall := fake.incrementCounterWithDeltaArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeIngressClient) IncrementCounterWithDeltaReturns(result1 error) {
	fake.incrementCounterWithDeltaMutex.Lock()
	defer fake.incrementCounterWithDeltaMutex.Unlock()
	fake.IncrementCounterWithDeltaStub = nil
	fake.incrementCounterWithDeltaReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIngressClient) IncrementCounterWithDeltaReturnsOnCall(i int, result1 error) {
	fake.incrementCounterWithDeltaMutex.Lock()
	defer fake.incrementCounterWithDeltaMutex.Unlock()
	fake.IncrementCounterWithDeltaStub = nil
	if fake.incrementCounterWithDeltaReturnsOnCall == nil {
		fake.incrementCounterWithDeltaReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.incrementCounterWithDeltaReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIngressClient) SendAppErrorLog(arg1 string, arg2 string, arg3 map[string]string) error {
	fake.sendAppErrorLogMutex.Lock()
	ret, specificReturn := fake.sendAppErrorLogReturnsOnCall[len(fake.sendAppErrorLogArgsForCall)]
	fake.sendAppErrorLogArgsForCall = append(fake.sendAppErrorLogArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 map[string]string
	}{arg1, arg2, arg3})
	stub := fake.SendAppErrorLogStub
	fakeReturns := fake.sendAppErrorLogReturns
	fake.recordInvocation("SendAppErrorLog", []interface{}{arg1, arg2, arg3})
	fake.sendAppErrorLogMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIngressClient) SendAppErrorLogCallCount() int {
	fake.sendAppErrorLogMutex.RLock()
	defer fake.sendAppErrorLogMutex.RUnlock()
	return len(fake.sendAppErrorLogArgsForCall)
}

func (fake *FakeIngressClient) SendAppErrorLogCalls(stub func(string, string, map[string]string) error) {
	fake.sendAppErrorLogMutex.Lock()
	defer fake.sendAppErrorLogMutex.Unlock()
	fake.SendAppErrorLogStub = stub
}

func (fake *FakeIngressClient) SendAppErrorLogArgsForCall(i int) (string, string, map[string]string) {
	fake.sendAppErrorLogMutex.RLock()
	defer fake.sendAppErrorLogMutex.RUnlock()
	argsForCall := fake.sendAppErrorLogArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeIngressClient) SendAppErrorLogReturns(result1 error) {
	fake.sendAppErrorLogMutex.Lock()
	defer fake.sendAppErrorLogMutex.Unlock()
	fake.SendAppErrorLogStub = nil
	fake.sendAppErrorLogReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIngressClient) SendAppErrorLogReturnsOnCall(i int, result1 error) {
	fake.sendAppErrorLogMutex.Lock()
	defer fake.sendAppErrorLogMutex.Unlock()
	fake.SendAppErrorLogStub = nil
	if fake.sendAppErrorLogReturnsOnCall == nil {
		fake.sendAppErrorLogReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.sendAppErrorLogReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIngressClient) SendAppLog(arg1 string, arg2 string, arg3 map[string]string) error {
	fake.sendAppLogMutex.Lock()
	ret, specificReturn := fake.sendAppLogReturnsOnCall[len(fake.sendAppLogArgsForCall)]
	fake.sendAppLogArgsForCall = append(fake.sendAppLogArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 map[string]string
	}{arg1, arg2, arg3})
	stub := fake.SendAppLogStub
	fakeReturns := fake.sendAppLogReturns
	fake.recordInvocation("SendAppLog", []interface{}{arg1, arg2, arg3})
	fake.sendAppLogMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIngressClient) SendAppLogCallCount() int {
	fake.sendAppLogMutex.RLock()
	defer fake.sendAppLogMutex.RUnlock()
	return len(fake.sendAppLogArgsForCall)
}

func (fake *FakeIngressClient) SendAppLogCalls(stub func(string, string, map[string]string) error) {
	fake.sendAppLogMutex.Lock()
	defer fake.sendAppLogMutex.Unlock()
	fake.SendAppLogStub = stub
}

func (fake *FakeIngressClient) SendAppLogArgsForCall(i int) (string, string, map[string]string) {
	fake.sendAppLogMutex.RLock()
	defer fake.sendAppLogMutex.RUnlock()
	argsForCall := fake.sendAppLogArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeIngressClient) SendAppLogReturns(result1 error) {
	fake.sendAppLogMutex.Lock()
	defer fake.sendAppLogMutex.Unlock()
	fake.SendAppLogStub = nil
	fake.sendAppLogReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIngressClient) SendAppLogReturnsOnCall(i int, result1 error) {
	fake.sendAppLogMutex.Lock()
	defer fake.sendAppLogMutex.Unlock()
	fake.SendAppLogStub = nil
	if fake.sendAppLogReturnsOnCall == nil {
		fake.sendAppLogReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.sendAppLogReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIngressClient) SendAppLogRate(arg1 float64, arg2 float64, arg3 map[string]string) error {
	fake.sendAppLogRateMutex.Lock()
	ret, specificReturn := fake.sendAppLogRateReturnsOnCall[len(fake.sendAppLogRateArgsForCall)]
	fake.sendAppLogRateArgsForCall = append(fake.sendAppLogRateArgsForCall, struct {
		arg1 float64
		arg2 float64
		arg3 map[string]string
	}{arg1, arg2, arg3})
	stub := fake.SendAppLogRateStub
	fakeReturns := fake.sendAppLogRateReturns
	fake.recordInvocation("SendAppLogRate", []interface{}{arg1, arg2, arg3})
	fake.sendAppLogRateMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIngressClient) SendAppLogRateCallCount() int {
	fake.sendAppLogRateMutex.RLock()
	defer fake.sendAppLogRateMutex.RUnlock()
	return len(fake.sendAppLogRateArgsForCall)
}

func (fake *FakeIngressClient) SendAppLogRateCalls(stub func(float64, float64, map[string]string) error) {
	fake.sendAppLogRateMutex.Lock()
	defer fake.sendAppLogRateMutex.Unlock()
	fake.SendAppLogRateStub = stub
}

func (fake *FakeIngressClient) SendAppLogRateArgsForCall(i int) (float64, float64, map[string]string) {
	fake.sendAppLogRateMutex.RLock()
	defer fake.sendAppLogRateMutex.RUnlock()
	argsForCall := fake.sendAppLogRateArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeIngressClient) SendAppLogRateReturns(result1 error) {
	fake.sendAppLogRateMutex.Lock()
	defer fake.sendAppLogRateMutex.Unlock()
	fake.SendAppLogRateStub = nil
	fake.sendAppLogRateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIngressClient) SendAppLogRateReturnsOnCall(i int, result1 error) {
	fake.sendAppLogRateMutex.Lock()
	defer fake.sendAppLogRateMutex.Unlock()
	fake.SendAppLogRateStub = nil
	if fake.sendAppLogRateReturnsOnCall == nil {
		fake.sendAppLogRateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.sendAppLogRateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIngressClient) SendAppMetrics(arg1 diego_logging_client.ContainerMetric) error {
	fake.sendAppMetricsMutex.Lock()
	ret, specificReturn := fake.sendAppMetricsReturnsOnCall[len(fake.sendAppMetricsArgsForCall)]
	fake.sendAppMetricsArgsForCall = append(fake.sendAppMetricsArgsForCall, struct {
		arg1 diego_logging_client.ContainerMetric
	}{arg1})
	stub := fake.SendAppMetricsStub
	fakeReturns := fake.sendAppMetricsReturns
	fake.recordInvocation("SendAppMetrics", []interface{}{arg1})
	fake.sendAppMetricsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIngressClient) SendAppMetricsCallCount() int {
	fake.sendAppMetricsMutex.RLock()
	defer fake.sendAppMetricsMutex.RUnlock()
	return len(fake.sendAppMetricsArgsForCall)
}

func (fake *FakeIngressClient) SendAppMetricsCalls(stub func(diego_logging_client.ContainerMetric) error) {
	fake.sendAppMetricsMutex.Lock()
	defer fake.sendAppMetricsMutex.Unlock()
	fake.SendAppMetricsStub = stub
}

func (fake *FakeIngressClient) SendAppMetricsArgsForCall(i int) diego_logging_client.ContainerMetric {
	fake.sendAppMetricsMutex.RLock()
	defer fake.sendAppMetricsMutex.RUnlock()
	argsForCall := fake.sendAppMetricsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeIngressClient) SendAppMetricsReturns(result1 error) {
	fake.sendAppMetricsMutex.Lock()
	defer fake.sendAppMetricsMutex.Unlock()
	fake.SendAppMetricsStub = nil
	fake.sendAppMetricsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIngressClient) SendAppMetricsReturnsOnCall(i int, result1 error) {
	fake.sendAppMetricsMutex.Lock()
	defer fake.sendAppMetricsMutex.Unlock()
	fake.SendAppMetricsStub = nil
	if fake.sendAppMetricsReturnsOnCall == nil {
		fake.sendAppMetricsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.sendAppMetricsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIngressClient) SendBytesPerSecond(arg1 string, arg2 float64) error {
	fake.sendBytesPerSecondMutex.Lock()
	ret, specificReturn := fake.sendBytesPerSecondReturnsOnCall[len(fake.sendBytesPerSecondArgsForCall)]
	fake.sendBytesPerSecondArgsForCall = append(fake.sendBytesPerSecondArgsForCall, struct {
		arg1 string
		arg2 float64
	}{arg1, arg2})
	stub := fake.SendBytesPerSecondStub
	fakeReturns := fake.sendBytesPerSecondReturns
	fake.recordInvocation("SendBytesPerSecond", []interface{}{arg1, arg2})
	fake.sendBytesPerSecondMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIngressClient) SendBytesPerSecondCallCount() int {
	fake.sendBytesPerSecondMutex.RLock()
	defer fake.sendBytesPerSecondMutex.RUnlock()
	return len(fake.sendBytesPerSecondArgsForCall)
}

func (fake *FakeIngressClient) SendBytesPerSecondCalls(stub func(string, float64) error) {
	fake.sendBytesPerSecondMutex.Lock()
	defer fake.sendBytesPerSecondMutex.Unlock()
	fake.SendBytesPerSecondStub = stub
}

func (fake *FakeIngressClient) SendBytesPerSecondArgsForCall(i int) (string, float64) {
	fake.sendBytesPerSecondMutex.RLock()
	defer fake.sendBytesPerSecondMutex.RUnlock()
	argsForCall := fake.sendBytesPerSecondArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeIngressClient) SendBytesPerSecondReturns(result1 error) {
	fake.sendBytesPerSecondMutex.Lock()
	defer fake.sendBytesPerSecondMutex.Unlock()
	fake.SendBytesPerSecondStub = nil
	fake.sendBytesPerSecondReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIngressClient) SendBytesPerSecondReturnsOnCall(i int, result1 error) {
	fake.sendBytesPerSecondMutex.Lock()
	defer fake.sendBytesPerSecondMutex.Unlock()
	fake.SendBytesPerSecondStub = nil
	if fake.sendBytesPerSecondReturnsOnCall == nil {
		fake.sendBytesPerSecondReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.sendBytesPerSecondReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIngressClient) SendComponentMetric(arg1 string, arg2 float64, arg3 string) error {
	fake.sendComponentMetricMutex.Lock()
	ret, specificReturn := fake.sendComponentMetricReturnsOnCall[len(fake.sendComponentMetricArgsForCall)]
	fake.sendComponentMetricArgsForCall = append(fake.sendComponentMetricArgsForCall, struct {
		arg1 string
		arg2 float64
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.SendComponentMetricStub
	fakeReturns := fake.sendComponentMetricReturns
	fake.recordInvocation("SendComponentMetric", []interface{}{arg1, arg2, arg3})
	fake.sendComponentMetricMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIngressClient) SendComponentMetricCallCount() int {
	fake.sendComponentMetricMutex.RLock()
	defer fake.sendComponentMetricMutex.RUnlock()
	return len(fake.sendComponentMetricArgsForCall)
}

func (fake *FakeIngressClient) SendComponentMetricCalls(stub func(string, float64, string) error) {
	fake.sendComponentMetricMutex.Lock()
	defer fake.sendComponentMetricMutex.Unlock()
	fake.SendComponentMetricStub = stub
}

func (fake *FakeIngressClient) SendComponentMetricArgsForCall(i int) (string, float64, string) {
	fake.sendComponentMetricMutex.RLock()
	defer fake.sendComponentMetricMutex.RUnlock()
	argsForCall := fake.sendComponentMetricArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeIngressClient) SendComponentMetricReturns(result1 error) {
	fake.sendComponentMetricMutex.Lock()
	defer fake.sendComponentMetricMutex.Unlock()
	fake.SendComponentMetricStub = nil
	fake.sendComponentMetricReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIngressClient) SendComponentMetricReturnsOnCall(i int, result1 error) {
	fake.sendComponentMetricMutex.Lock()
	defer fake.sendComponentMetricMutex.Unlock()
	fake.SendComponentMetricStub = nil
	if fake.sendComponentMetricReturnsOnCall == nil {
		fake.sendComponentMetricReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.sendComponentMetricReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIngressClient) SendDuration(arg1 string, arg2 time.Duration, arg3 ...loggregator.EmitGaugeOption) error {
	fake.sendDurationMutex.Lock()
	ret, specificReturn := fake.sendDurationReturnsOnCall[len(fake.sendDurationArgsForCall)]
	fake.sendDurationArgsForCall = append(fake.sendDurationArgsForCall, struct {
		arg1 string
		arg2 time.Duration
		arg3 []loggregator.EmitGaugeOption
	}{arg1, arg2, arg3})
	stub := fake.SendDurationStub
	fakeReturns := fake.sendDurationReturns
	fake.recordInvocation("SendDuration", []interface{}{arg1, arg2, arg3})
	fake.sendDurationMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIngressClient) SendDurationCallCount() int {
	fake.sendDurationMutex.RLock()
	defer fake.sendDurationMutex.RUnlock()
	return len(fake.sendDurationArgsForCall)
}

func (fake *FakeIngressClient) SendDurationCalls(stub func(string, time.Duration, ...loggregator.EmitGaugeOption) error) {
	fake.sendDurationMutex.Lock()
	defer fake.sendDurationMutex.Unlock()
	fake.SendDurationStub = stub
}

func (fake *FakeIngressClient) SendDurationArgsForCall(i int) (string, time.Duration, []loggregator.EmitGaugeOption) {
	fake.sendDurationMutex.RLock()
	defer fake.sendDurationMutex.RUnlock()
	argsForCall := fake.sendDurationArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeIngressClient) SendDurationReturns(result1 error) {
	fake.sendDurationMutex.Lock()
	defer fake.sendDurationMutex.Unlock()
	fake.SendDurationStub = nil
	fake.sendDurationReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIngressClient) SendDurationReturnsOnCall(i int, result1 error) {
	fake.sendDurationMutex.Lock()
	defer fake.sendDurationMutex.Unlock()
	fake.SendDurationStub = nil
	if fake.sendDurationReturnsOnCall == nil {
		fake.sendDurationReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.sendDurationReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIngressClient) SendMebiBytes(arg1 string, arg2 int, arg3 ...loggregator.EmitGaugeOption) error {
	fake.sendMebiBytesMutex.Lock()
	ret, specificReturn := fake.sendMebiBytesReturnsOnCall[len(fake.sendMebiBytesArgsForCall)]
	fake.sendMebiBytesArgsForCall = append(fake.sendMebiBytesArgsForCall, struct {
		arg1 string
		arg2 int
		arg3 []loggregator.EmitGaugeOption
	}{arg1, arg2, arg3})
	stub := fake.SendMebiBytesStub
	fakeReturns := fake.sendMebiBytesReturns
	fake.recordInvocation("SendMebiBytes", []interface{}{arg1, arg2, arg3})
	fake.sendMebiBytesMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIngressClient) SendMebiBytesCallCount() int {
	fake.sendMebiBytesMutex.RLock()
	defer fake.sendMebiBytesMutex.RUnlock()
	return len(fake.sendMebiBytesArgsForCall)
}

func (fake *FakeIngressClient) SendMebiBytesCalls(stub func(string, int, ...loggregator.EmitGaugeOption) error) {
	fake.sendMebiBytesMutex.Lock()
	defer fake.sendMebiBytesMutex.Unlock()
	fake.SendMebiBytesStub = stub
}

func (fake *FakeIngressClient) SendMebiBytesArgsForCall(i int) (string, int, []loggregator.EmitGaugeOption) {
	fake.sendMebiBytesMutex.RLock()
	defer fake.sendMebiBytesMutex.RUnlock()
	argsForCall := fake.sendMebiBytesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeIngressClient) SendMebiBytesReturns(result1 error) {
	fake.sendMebiBytesMutex.Lock()
	defer fake.sendMebiBytesMutex.Unlock()
	fake.SendMebiBytesStub = nil
	fake.sendMebiBytesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIngressClient) SendMebiBytesReturnsOnCall(i int, result1 error) {
	fake.sendMebiBytesMutex.Lock()
	defer fake.sendMebiBytesMutex.Unlock()
	fake.SendMebiBytesStub = nil
	if fake.sendMebiBytesReturnsOnCall == nil {
		fake.sendMebiBytesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.sendMebiBytesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIngressClient) SendMetric(arg1 string, arg2 int, arg3 ...loggregator.EmitGaugeOption) error {
	fake.sendMetricMutex.Lock()
	ret, specificReturn := fake.sendMetricReturnsOnCall[len(fake.sendMetricArgsForCall)]
	fake.sendMetricArgsForCall = append(fake.sendMetricArgsForCall, struct {
		arg1 string
		arg2 int
		arg3 []loggregator.EmitGaugeOption
	}{arg1, arg2, arg3})
	stub := fake.SendMetricStub
	fakeReturns := fake.sendMetricReturns
	fake.recordInvocation("SendMetric", []interface{}{arg1, arg2, arg3})
	fake.sendMetricMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIngressClient) SendMetricCallCount() int {
	fake.sendMetricMutex.RLock()
	defer fake.sendMetricMutex.RUnlock()
	return len(fake.sendMetricArgsForCall)
}

func (fake *FakeIngressClient) SendMetricCalls(stub func(string, int, ...loggregator.EmitGaugeOption) error) {
	fake.sendMetricMutex.Lock()
	defer fake.sendMetricMutex.Unlock()
	fake.SendMetricStub = stub
}

func (fake *FakeIngressClient) SendMetricArgsForCall(i int) (string, int, []loggregator.EmitGaugeOption) {
	fake.sendMetricMutex.RLock()
	defer fake.sendMetricMutex.RUnlock()
	argsForCall := fake.sendMetricArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeIngressClient) SendMetricReturns(result1 error) {
	fake.sendMetricMutex.Lock()
	defer fake.sendMetricMutex.Unlock()
	fake.SendMetricStub = nil
	fake.sendMetricReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIngressClient) SendMetricReturnsOnCall(i int, result1 error) {
	fake.sendMetricMutex.Lock()
	defer fake.sendMetricMutex.Unlock()
	fake.SendMetricStub = nil
	if fake.sendMetricReturnsOnCall == nil {
		fake.sendMetricReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.sendMetricReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIngressClient) SendRequestsPerSecond(arg1 string, arg2 float64) error {
	fake.sendRequestsPerSecondMutex.Lock()
	ret, specificReturn := fake.sendRequestsPerSecondReturnsOnCall[len(fake.sendRequestsPerSecondArgsForCall)]
	fake.sendRequestsPerSecondArgsForCall = append(fake.sendRequestsPerSecondArgsForCall, struct {
		arg1 string
		arg2 float64
	}{arg1, arg2})
	stub := fake.SendRequestsPerSecondStub
	fakeReturns := fake.sendRequestsPerSecondReturns
	fake.recordInvocation("SendRequestsPerSecond", []interface{}{arg1, arg2})
	fake.sendRequestsPerSecondMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIngressClient) SendRequestsPerSecondCallCount() int {
	fake.sendRequestsPerSecondMutex.RLock()
	defer fake.sendRequestsPerSecondMutex.RUnlock()
	return len(fake.sendRequestsPerSecondArgsForCall)
}

func (fake *FakeIngressClient) SendRequestsPerSecondCalls(stub func(string, float64) error) {
	fake.sendRequestsPerSecondMutex.Lock()
	defer fake.sendRequestsPerSecondMutex.Unlock()
	fake.SendRequestsPerSecondStub = stub
}

func (fake *FakeIngressClient) SendRequestsPerSecondArgsForCall(i int) (string, float64) {
	fake.sendRequestsPerSecondMutex.RLock()
	defer fake.sendRequestsPerSecondMutex.RUnlock()
	argsForCall := fake.sendRequestsPerSecondArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeIngressClient) SendRequestsPerSecondReturns(result1 error) {
	fake.sendRequestsPerSecondMutex.Lock()
	defer fake.sendRequestsPerSecondMutex.Unlock()
	fake.SendRequestsPerSecondStub = nil
	fake.sendRequestsPerSecondReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIngressClient) SendRequestsPerSecondReturnsOnCall(i int, result1 error) {
	fake.sendRequestsPerSecondMutex.Lock()
	defer fake.sendRequestsPerSecondMutex.Unlock()
	fake.SendRequestsPerSecondStub = nil
	if fake.sendRequestsPerSecondReturnsOnCall == nil {
		fake.sendRequestsPerSecondReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.sendRequestsPerSecondReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIngressClient) SendSpikeMetrics(arg1 diego_logging_client.SpikeMetric) error {
	fake.sendSpikeMetricsMutex.Lock()
	ret, specificReturn := fake.sendSpikeMetricsReturnsOnCall[len(fake.sendSpikeMetricsArgsForCall)]
	fake.sendSpikeMetricsArgsForCall = append(fake.sendSpikeMetricsArgsForCall, struct {
		arg1 diego_logging_client.SpikeMetric
	}{arg1})
	stub := fake.SendSpikeMetricsStub
	fakeReturns := fake.sendSpikeMetricsReturns
	fake.recordInvocation("SendSpikeMetrics", []interface{}{arg1})
	fake.sendSpikeMetricsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIngressClient) SendSpikeMetricsCallCount() int {
	fake.sendSpikeMetricsMutex.RLock()
	defer fake.sendSpikeMetricsMutex.RUnlock()
	return len(fake.sendSpikeMetricsArgsForCall)
}

func (fake *FakeIngressClient) SendSpikeMetricsCalls(stub func(diego_logging_client.SpikeMetric) error) {
	fake.sendSpikeMetricsMutex.Lock()
	defer fake.sendSpikeMetricsMutex.Unlock()
	fake.SendSpikeMetricsStub = stub
}

func (fake *FakeIngressClient) SendSpikeMetricsArgsForCall(i int) diego_logging_client.SpikeMetric {
	fake.sendSpikeMetricsMutex.RLock()
	defer fake.sendSpikeMetricsMutex.RUnlock()
	argsForCall := fake.sendSpikeMetricsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeIngressClient) SendSpikeMetricsReturns(result1 error) {
	fake.sendSpikeMetricsMutex.Lock()
	defer fake.sendSpikeMetricsMutex.Unlock()
	fake.SendSpikeMetricsStub = nil
	fake.sendSpikeMetricsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIngressClient) SendSpikeMetricsReturnsOnCall(i int, result1 error) {
	fake.sendSpikeMetricsMutex.Lock()
	defer fake.sendSpikeMetricsMutex.Unlock()
	fake.SendSpikeMetricsStub = nil
	if fake.sendSpikeMetricsReturnsOnCall == nil {
		fake.sendSpikeMetricsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.sendSpikeMetricsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIngressClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.incrementCounterMutex.RLock()
	defer fake.incrementCounterMutex.RUnlock()
	fake.incrementCounterWithDeltaMutex.RLock()
	defer fake.incrementCounterWithDeltaMutex.RUnlock()
	fake.sendAppErrorLogMutex.RLock()
	defer fake.sendAppErrorLogMutex.RUnlock()
	fake.sendAppLogMutex.RLock()
	defer fake.sendAppLogMutex.RUnlock()
	fake.sendAppLogRateMutex.RLock()
	defer fake.sendAppLogRateMutex.RUnlock()
	fake.sendAppMetricsMutex.RLock()
	defer fake.sendAppMetricsMutex.RUnlock()
	fake.sendBytesPerSecondMutex.RLock()
	defer fake.sendBytesPerSecondMutex.RUnlock()
	fake.sendComponentMetricMutex.RLock()
	defer fake.sendComponentMetricMutex.RUnlock()
	fake.sendDurationMutex.RLock()
	defer fake.sendDurationMutex.RUnlock()
	fake.sendMebiBytesMutex.RLock()
	defer fake.sendMebiBytesMutex.RUnlock()
	fake.sendMetricMutex.RLock()
	defer fake.sendMetricMutex.RUnlock()
	fake.sendRequestsPerSecondMutex.RLock()
	defer fake.sendRequestsPerSecondMutex.RUnlock()
	fake.sendSpikeMetricsMutex.RLock()
	defer fake.sendSpikeMetricsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeIngressClient) recordInvocation(key string, args []interface{}) {
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

var _ diego_logging_client.IngressClient = new(FakeIngressClient)
