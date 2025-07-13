# TCP-Chat

A simple chat application using TCP socket with transport layer security(TLS).

## Getting Started

**1. Prerequisites:**
  
- [Go](https://go.dev/doc/install) installed
- You also need a [tcp-chat server](https://github.com/AnhTTx13/tcp-chat-server) running somewhere to connect to.

**2. Install the repository:**
  
  ```sh
    go install github.com/AnhTTx13/tcp-chat
  ```

## Usage
  
- First start the chat client:
  
  ```sh
    tcp-chat [flag]
  ```

- To send something to the server:
  
  ```sh
    > [/cmd] [message]
  ```

**Command:**

- `/name`: Change your nick name.
- `/join`: Join a room or create if it doesn't exists.
- `/msg`: Send a message to a room.
- `/rooms`: Get all existing rooms.

**Example:**

```sh
> /name xbro        # Rename to xbro
> /join room1       # Join room1
> /msg Hello there  # Send message to room1
> /join room2       # Join room2
> /msg Hello room2  # Send message to room2
```

**Avaiable flags:**

- `--host [string]`: Specify the host name/ip address of the server (default localhost).
- `--port [int]`: Specify the port number of the server (default 8080).
