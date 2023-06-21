# Zerolog-GCP-Integration

Zerolog-GCP-Integration is a module for integrating Zerolog with Google Cloud Platform (GCP). It enables Golang applications to push log events to GCP's Cloud Logging, enhancing monitoring and troubleshooting capabilities.

## Features

- Seamlessly integrates with Zerolog
- Log events are directly pushed to Google Cloud's Cloud Logging
- Offers different log levels for enhanced debugging and information tracking
- Compatible with OpenTelemetry's trace API

## Installation

```bash
go get github.com/program-world-labs/pwlogger
```

## Usage

Import the package into your project and use it to create a new logger:

```go
import (
    "github.com/program-world-labs/pwlogger"
)

func main() {
    log := logger.NewProductionLogger("<your_project_id>")
    log.Info().Msg("This is an info message")

    // or for development
    log := logger.NewDevelopmentLogger("<your_project_id>")
    log.Debug().Msg("This is a debug message")
}
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT](https://choosealicense.com/licenses/mit/)
