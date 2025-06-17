# 📡 Broadcast Server (CLI WebSocket Chat)

A CLI-based broadcast server in Go that allows multiple clients to connect, send messages, and receive real-time broadcasts from others via WebSockets.

🔗 Listed on [roadmap.sh Projects](https://roadmap.sh/projects/broadcast-server)

---

## 🧠 Features

- Start a WebSocket server using `broadcast-server start`
- Connect clients using `broadcast-server connect`
- Real-time message broadcast to all connected clients
- Graceful handling of connections and disconnections
- CLI-based messaging
- Simple, extensible structure

---

## ⚙️ Usage

### 🔧 1. Build the project

```bash
go build -o broadcast-server main.go