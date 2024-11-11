# DiceDB TCP Server

This repository implements a basic synchronous TCP server that allows clients to connect, send commands, and receive responses. The server logs each client connection and command and can handle multiple concurrent clients. It also includes configurable options for host and port settings through command-line flags.

## Project Structure

```
.
├── main.go             # Entry point for the server
├── config/
│   └── main.go         # Configuration settings for host and port
└── server/
    └── tcp.go          # TCP server implementation
```

### Files Overview

- **`main.go`**: Initializes the server with host and port settings from command-line flags and starts the synchronous TCP server.
- **`config/main.go`**: Holds default configurations for the server's `Host` and `Port`.
- **`server/tcp.go`**: Contains the implementation of the synchronous TCP server, including command handling and response logic.

## Features

- **Configurable Host and Port**: Set the host and port for the server through command-line flags.
- **Synchronous TCP Server**: Manages client connections and allows command handling.
- **Client Connection Logging**: Logs each client’s IP address and the current number of connected clients.
- **Command Echoing**: Reads commands from clients and echoes them back.

## Usage

### Prerequisites

- [Go](https://golang.org/doc/install) (version 1.16 or higher recommended)

### Setup and Run

1. **Clone the repository:**
   ```bash
   git clone https://github.com/yourusername/dicedb-tcp-server.git
   cd dicedb-tcp-server
   ```

2. **Run the server** with default settings:
   ```bash
   go run main.go
   ```

   You can also set the `Host` and `Port` using command-line flags:
   ```bash
   go run main.go -host 127.0.0.1 -port 7379
   ```

   This command will start the TCP server on the specified host and port.

### Interacting with the Server

After starting the server, you can connect to it via a TCP client (like `telnet` or `nc`) on the specified host and port.

For example, using `nc`:
```bash
nc 127.0.0.1 7379
```

Type any command, and the server will echo it back.

## Code Overview

### `main.go`

- Defines command-line flags for setting the `Host` and `Port`.
- Initializes the TCP server by calling `server.RunSyncTCPServer()`.

### `config/main.go`

- Provides default configuration values for `Host` (`0.0.0.0`) and `Port` (`7379`).
- The host and port can be overridden by command-line flags in `main.go`.

### `server/tcp.go`

- **`RunSyncTCPServer`**: Starts a synchronous TCP server, listening on the specified host and port.
- **Client Handling**:
  - Logs each new connection and the number of concurrent clients.
  - Reads commands from each client, logs them, and sends back responses.
  - Closes connections and updates the concurrent client count on client disconnect.

## Example Commands

```bash
# Start server with default settings
go run main.go

# Start server with custom host and port
go run main.go -host 127.0.0.1 -port 8080
```

## License

This project is licensed under the MIT License.