# TCP-Chat

A simple chat application to communicate with other devices on your local network via TCP socket. Protected with transport layer security (TLS).

## Getting Started

**1. Prerequisites:**
  
- [Go](https://go.dev/doc/install) installed
- You also need a [tcp-chat server](https://github.com/anhtr13/tcp-chat-server) running somewhere on your local network to connect to.

**2. Install the repository:**
  
  ```sh
    go install github.com/anhtr13/tcp-chat@latest
  ```

## Usage
  
- First start the chat client:
  
  ```sh
    tcp-chat --host=[server host/ip] --port=[server port number]
    # default: --host=localhost
    #          --port=8080
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
