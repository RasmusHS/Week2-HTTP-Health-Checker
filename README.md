# HTTP Health Checker
This project is the 2nd in my 12 weeks 12 projects plan. [Overview](https://github.com/RasmusHS/12-Weeks-of-Coding) \
A concurrent command-line tool written in Go that continuously monitors the health of a list of URLs. It performs HTTP GET requests in parallel, reports status codes, response times, and up/down status in a formatted table, and repeats on a configurable interval.

## Features

- **Concurrent checks** — URLs are checked in parallel using goroutines, with a concurrency limit of 10 to avoid overwhelming the system.
- **Configurable interval** — Set the check interval via the `-interval` flag (default: 60 seconds).
- **Live countdown** — Displays a real-time countdown until the next check cycle.
- **Formatted output** — Results are displayed in an aligned table showing URL, status code, response time, and up/down status.
- **Graceful exit** — Press Enter at any time during the countdown to stop the program.
- **Status classification** — Differentiates between UP (2xx), REDIRECT (3xx), client errors (4xx), server errors (5xx), and connection failures.

## Project Structure

| File         | Purpose                                    |
|--------------|--------------------------------------------|
| `main.go`    | Entry point, orchestration, and check loop |
| `checker.go` | HTTP checking logic (5s timeout per URL)   |
| `config.go`  | JSON config file loading and parsing       |
| `result.go`  | Result struct and table output formatting  |
| `urls.json`  | List of URLs to monitor                    |

## Getting Started

### Prerequisites

- [Go](https://go.dev/dl/) 1.26+

### Installation

```bash
git clone <repository-url>
cd "HTTP Health Checker"
go build -o health-checker .
```

### Configuration
Edit urls.json to specify the URLs you want to monitor:
```json
"urls": [
    "https://www.example.com",
    "https://www.google.com",
    "https://www.github.com"
  ]
```

### Usage

```sh
# Run with default 60-second interval
go run .

# Run with a custom interval (e.g., 30 seconds)
go run . -interval 30
```

## Example Output
<!-- |---|---|---|---| -->
```bash
URL                            | Status | Response Time | Up/Down Status
https://www.google.com         | 200    | 150ms         | UP
https://www.example.com        | 200    | 85ms          | UP
https://www.badurl.invalid     | ---    | ---           | DOWN (error: ...)

Next check in 60 seconds... Press Enter to exit.
Time remaining: 47 seconds
```
