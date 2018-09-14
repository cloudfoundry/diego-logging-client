package diego_logging_client_test

import (
	"os"
	"path"
	"time"

	client "code.cloudfoundry.org/diego-logging-client"
	"code.cloudfoundry.org/diego-logging-client/testhelpers"
	"code.cloudfoundry.org/go-loggregator/rpc/loggregator_v2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	fixturesPath         = path.Join(os.Getenv("GOPATH"), "src/code.cloudfoundry.org/diego-logging-client/fixtures")
	metronCAFile         = path.Join(fixturesPath, "metron", "CA.crt")
	metronServerCertFile = path.Join(fixturesPath, "metron", "metron.crt")
	metronServerKeyFile  = path.Join(fixturesPath, "metron", "metron.key")
	metronClientCertFile = path.Join(fixturesPath, "metron", "client.crt")
	metronClientKeyFile  = path.Join(fixturesPath, "metron", "client.key")
)

var _ = Describe("DiegoLoggingClient", func() {
	var (
		c client.IngressClient
	)

	Context("when the v2 api is used", func() {
		var (
			testIngressServer *testhelpers.TestIngressServer
			metricsPort       int
		)

		BeforeEach(func() {
			var err error

			testIngressServer, err = testhelpers.NewTestIngressServer(metronServerCertFile, metronServerKeyFile, metronCAFile)
			Expect(err).NotTo(HaveOccurred())

			Expect(testIngressServer.Start()).To(Succeed())

			metricsPort, err = testIngressServer.Port()
			Expect(err).NotTo(HaveOccurred())
		})

		Context("and the loggregator agent isn't up", func() {
			BeforeEach(func() {
				testIngressServer.Stop()
			})

			It("returns an error when constructing the loggregator client", func() {
				metricsPort := 8080

				_, err := client.NewIngressClient(client.Config{
					SourceID:           "some-source-id",
					InstanceID:         "some-instance-id",
					BatchFlushInterval: 10 * time.Millisecond,
					BatchMaxSize:       1,
					UseV2API:           true,
					APIPort:            metricsPort,
					CACertPath:         metronCAFile,
					KeyPath:            metronClientKeyFile,
					CertPath:           metronClientCertFile,
				})
				Expect(err).To(HaveOccurred())
			})
		})

		Context("and the loggregator agent is up", func() {
			BeforeEach(func() {
				var err error
				c, err = client.NewIngressClient(client.Config{
					SourceID:           "some-source-id",
					InstanceID:         "some-instance-id",
					BatchFlushInterval: 10 * time.Millisecond,
					BatchMaxSize:       1,
					UseV2API:           true,
					APIPort:            metricsPort,
					CACertPath:         metronCAFile,
					KeyPath:            metronClientKeyFile,
					CertPath:           metronClientCertFile,
				})
				Expect(err).NotTo(HaveOccurred())
			})

			getEnvelopeBatch := func() *loggregator_v2.EnvelopeBatch {
				var sender loggregator_v2.Ingress_BatchSenderServer
				Eventually(testIngressServer.Receivers()).Should(Receive(&sender))

				batch, err := sender.Recv()
				Expect(err).NotTo(HaveOccurred())
				return batch
			}

			assertEnvelopeSourceAndInstanceIDAreCorrect := func(batch *loggregator_v2.EnvelopeBatch) {
				Expect(batch.Batch).To(HaveLen(1))
				Expect(batch.Batch[0].GetSourceId()).To(Equal("some-source-id"))
				Expect(batch.Batch[0].GetInstanceId()).To(Equal("some-instance-id"))
			}

			Describe("SendDuration", func() {
				It("sets app info", func() {
					Expect(c.SendDuration("time", 18*time.Second)).To(Succeed())

					assertEnvelopeSourceAndInstanceIDAreCorrect(getEnvelopeBatch())
				})
			})

			Describe("SendMebiBytes", func() {
				It("sets app info", func() {
					Expect(c.SendMebiBytes("disk-free", 23)).To(Succeed())

					assertEnvelopeSourceAndInstanceIDAreCorrect(getEnvelopeBatch())
				})
			})

			Describe("SendMetric", func() {
				It("sets app info", func() {
					Expect(c.SendMetric("errors", 3)).To(Succeed())

					assertEnvelopeSourceAndInstanceIDAreCorrect(getEnvelopeBatch())
				})
			})

			Describe("SendBytesPerSecond", func() {
				It("sets app info", func() {
					Expect(c.SendBytesPerSecond("speed", 3)).To(Succeed())

					assertEnvelopeSourceAndInstanceIDAreCorrect(getEnvelopeBatch())
				})
			})

			Describe("SendRequestsPerSecond", func() {
				It("sets app info", func() {
					Expect(c.SendRequestsPerSecond("homepage", 37)).To(Succeed())

					assertEnvelopeSourceAndInstanceIDAreCorrect(getEnvelopeBatch())
				})
			})

			Describe("IncrementCounter", func() {
				It("sets app info", func() {
					Expect(c.IncrementCounter("its")).To(Succeed())

					assertEnvelopeSourceAndInstanceIDAreCorrect(getEnvelopeBatch())
				})
			})

			Describe("IncrementCounterWithDelta", func() {
				It("sets app info", func() {
					Expect(c.IncrementCounterWithDelta("its", 5)).To(Succeed())

					assertEnvelopeSourceAndInstanceIDAreCorrect(getEnvelopeBatch())
				})
			})

			Describe("SendComponentMetric", func() {
				It("sets app info", func() {
					Expect(c.SendComponentMetric("memory", 37, "GB")).To(Succeed())

					assertEnvelopeSourceAndInstanceIDAreCorrect(getEnvelopeBatch())
				})
			})

			Describe("SendCPUUsage", func() {
				var batch *loggregator_v2.EnvelopeBatch

				JustBeforeEach(func() {
					Expect(c.SendCPUUsage("some-source-id", 345, 1, 2, 3)).To(Succeed())
					batch = getEnvelopeBatch()
				})

				It("sets app info", func() {
					Expect(batch.Batch).To(HaveLen(1))
					Expect(batch.Batch[0].GetSourceId()).To(Equal("some-source-id"))
					Expect(batch.Batch[0].GetInstanceId()).To(Equal("345"))
				})

				It("sends cpu usage", func() {
					metrics := batch.Batch[0].GetGauge().GetMetrics()
					Expect(metrics["absolute_usage"].GetValue()).To(Equal(float64(1)))
					Expect(metrics["absolute_usage"].GetUnit()).To(Equal("nanoseconds"))

					Expect(metrics["absolute_entitlement"].GetValue()).To(Equal(float64(2)))
					Expect(metrics["absolute_entitlement"].GetUnit()).To(Equal("nanoseconds"))

					Expect(metrics["container_age"].GetValue()).To(Equal(float64(3)))
					Expect(metrics["container_age"].GetUnit()).To(Equal("nanoseconds"))
				})
			})
		})
	})
})
