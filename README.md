# AnonTalk — Anonymous Real-Time Chat with Rooms

**AnonTalk** is a lightweight real-time chat application built in Go. It allows users to create chat rooms and communicate anonymously over WebSocket, either with a chosen nickname or fully incognito.


---

## 📸 UI Preview

Frontend is included and served via HTTP — simple HTML + JS interface for quick testing and use.  
Messages are styled and updated in real time.

---

## 📦 Building

### 🔧 Requirements

- [Go](https://golang.org/dl/) ≥ 1.21
- [Task](https://taskfile.dev/) — for build automation

Install `task`:

```bash
go install github.com/go-task/task/v3/cmd/task@latest
```
Clone the project
```bash
git clone https://github.com/ASsssker/AnonTalk.git
cd AnonTalk
```

Build for both Linux and Windows:
```bash
task build:all
```
The final binaries will be in the bin/ directory.