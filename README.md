# diego-logging-client

**Note**: This package must be imported as `code.cloudfoundry.org/diego-logging-client`.

The Diego Logging Client provides a generic client for
[Diego](https://github.com/cloudfoundry/diego-release) to
Cloud Foundry's logging subsystem,
[Loggregator](https://github.com/cloudfoundry/loggregator).

The client wraps the [go-loggregator](https://github.com/cloudfoundry/go-loggregator) library
to provide a tailored interface for Diego components.
