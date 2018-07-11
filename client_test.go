package diego_logging_client_test

import (
	"time"

	client "code.cloudfoundry.org/diego-logging-client"
	"code.cloudfoundry.org/go-loggregator"
	"code.cloudfoundry.org/go-loggregator/rpc/loggregator_v2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DiegoLoggingClient", func() {
	var (
		s       *spyLogClient
		c       client.IngressClient
		gauge   *loggregator_v2.Envelope
		counter *loggregator_v2.Envelope
	)

	BeforeEach(func() {
		s = newSpyLogClient()
		c = client.WrapClient(s, "some-source-id", "some-instance-id")
		gauge = &loggregator_v2.Envelope{
			Timestamp: time.Now().UnixNano(),
			Message: &loggregator_v2.Envelope_Gauge{
				Gauge: &loggregator_v2.Gauge{
					Metrics: make(map[string]*loggregator_v2.GaugeValue),
				},
			},
			Tags: make(map[string]string),
		}
		counter = &loggregator_v2.Envelope{
			Timestamp: time.Now().UnixNano(),
			Message: &loggregator_v2.Envelope_Counter{
				Counter: &loggregator_v2.Counter{},
			},
			Tags: make(map[string]string),
		}
	})

	Describe("SendDuration", func() {
		It("sets app info", func() {
			c.SendDuration("time", 18*time.Second)

			for _, o := range s.gaugeOpts {
				o(gauge)
			}

			Expect(gauge.GetSourceId()).To(Equal("some-source-id"))
			Expect(gauge.GetInstanceId()).To(Equal("some-instance-id"))
		})
	})

	Describe("SendMebiBytes", func() {
		It("sets app info", func() {
			c.SendMebiBytes("disk-free", 23)

			for _, o := range s.gaugeOpts {
				o(gauge)
			}

			Expect(gauge.GetSourceId()).To(Equal("some-source-id"))
			Expect(gauge.GetInstanceId()).To(Equal("some-instance-id"))
		})
	})

	Describe("SendMetric", func() {
		It("sets app info", func() {
			c.SendMetric("errors", 3)

			for _, o := range s.gaugeOpts {
				o(gauge)
			}

			Expect(gauge.GetSourceId()).To(Equal("some-source-id"))
			Expect(gauge.GetInstanceId()).To(Equal("some-instance-id"))
		})
	})

	Describe("SendBytesPerSecond", func() {
		It("sets app info", func() {
			c.SendBytesPerSecond("speed", 3)

			for _, o := range s.gaugeOpts {
				o(gauge)
			}

			Expect(gauge.GetSourceId()).To(Equal("some-source-id"))
			Expect(gauge.GetInstanceId()).To(Equal("some-instance-id"))
		})
	})

	Describe("SendRequestsPerSecond", func() {
		It("sets app info", func() {
			c.SendRequestsPerSecond("homepage", 37)

			for _, o := range s.gaugeOpts {
				o(gauge)
			}

			Expect(gauge.GetSourceId()).To(Equal("some-source-id"))
			Expect(gauge.GetInstanceId()).To(Equal("some-instance-id"))
		})
	})

	Describe("IncrementCounter", func() {
		It("sets app info", func() {
			c.IncrementCounter("its")

			for _, o := range s.counterOpts {
				o(counter)
			}

			Expect(counter.GetSourceId()).To(Equal("some-source-id"))
			Expect(counter.GetInstanceId()).To(Equal("some-instance-id"))
		})
	})

	Describe("IncrementCounterWithDelta", func() {
		It("sets app info", func() {
			c.IncrementCounterWithDelta("its", 5)

			for _, o := range s.counterOpts {
				o(counter)
			}

			Expect(counter.GetSourceId()).To(Equal("some-source-id"))
			Expect(counter.GetInstanceId()).To(Equal("some-instance-id"))
		})
	})

	Describe("SendComponentMetric", func() {
		It("sets app info", func() {
			c.SendComponentMetric("memory", 37, "GB")

			for _, o := range s.gaugeOpts {
				o(gauge)
			}

			Expect(gauge.GetSourceId()).To(Equal("some-source-id"))
			Expect(gauge.GetInstanceId()).To(Equal("some-instance-id"))
		})
	})
})

type spyLogClient struct {
	gaugeOpts   []loggregator.EmitGaugeOption
	counterOpts []loggregator.EmitCounterOption
}

func newSpyLogClient() *spyLogClient {
	return &spyLogClient{}
}

func (c *spyLogClient) EmitLog(msg string, opts ...loggregator.EmitLogOption) {}

func (c *spyLogClient) EmitGauge(opts ...loggregator.EmitGaugeOption) {
	c.gaugeOpts = opts
}

func (c *spyLogClient) EmitCounter(name string, opts ...loggregator.EmitCounterOption) {
	c.counterOpts = opts
}
