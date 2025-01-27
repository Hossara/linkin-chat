<pre>   __ _       _    _           ___ _           _   
  / /(_)_ __ | | _(_)_ __     / __\ |__   __ _| |_ 
 / / | | '_ \| |/ / | '_ \   / /  | '_ \ / _` | __|
/ /__| | | | |   <| | | | | / /___| | | | (_| | |_ 
\____/_|_| |_|_|\_\_|_| |_| \____/|_| |_|\__,_|\__|
</pre>

Linkin Chat is a TUI (Text User Interface) based chat application that uses HTTP and NATS for communication and broadcasting messages between the client and server.

## Table Of Content
<!-- TOC -->
  * [Table Of Content](#table-of-content)
  * [Tech Stack](#tech-stack)
  * [Prerequisites](#prerequisites)
  * [Development Workflow](#development-workflow)
    * [Start Development Environment](#start-development-environment)
    * [Start Development Environment with Logs](#start-development-environment-with-logs)
    * [Stop Development Environment](#stop-development-environment)
    * [Build Development Containers](#build-development-containers)
    * [View Logs for Development Environment](#view-logs-for-development-environment)
    * [Rebuild and Restart Development Containers](#rebuild-and-restart-development-containers)
  * [Deployment Workflow](#deployment-workflow)
    * [Deploy to Production](#deploy-to-production)
    * [Stop Production Deployment](#stop-production-deployment)
    * [View Logs for Production Deployment](#view-logs-for-production-deployment)
    * [Build Production Containers](#build-production-containers)
  * [Installing the Client](#installing-the-client)
    * [For Ubuntu/Debian](#for-ubuntudebian)
    * [For macOS](#for-macos)
    * [For Windows](#for-windows)
  * [Creating a Release](#creating-a-release)
  * [Additional Notes](#additional-notes)
  * [Features](#features)
  * [Known Issues](#known-issues)
  * [Planned Improvements](#planned-improvements)
  * [Test video](#test-video)
<!-- TOC -->

## Tech Stack
- Server
  - [Fiber (http)](https://gofiber.io)
  - [Nats](https://nats.io)
- Database
  - Postgres
- Tools
  - Docker
  - Makefile
- Cache
  - Redis
- Config
  - [Viper](https://github.com/spf13/viper)
- Client
  - UI: [tview](https://github.com/rivo/tview)
  - Release: [Goreleaser](https://goreleaser.com/) 
  - CLI: [Cobra](https://github.com/spf13/cobra)

## Prerequisites
Before you begin, ensure you have the following tools installed:

- **Docker**: [Install Docker](https://docs.docker.com/get-docker/)
- **Docker Compose**: Included with Docker Desktop or install separately ([Docker Compose Documentation](https://docs.docker.com/compose/))
- **GoReleaser** (optional, for release creation): [Install GoReleaser](https://goreleaser.com/quick-start/)

Make sure your project directory follows the structure defined in the Makefile (e.g., `./build/nats`, `./build/redis`, `./build/postgres`, and `./build/server`).

---

## Development Workflow

### Start Development Environment
To spin up the development environment:
```bash
make dev_up
```
This will:
1. Ensure the Docker network `linkin-chat-network` exists.
2. Start the development versions of NATS, Redis, and PostgreSQL containers defined in the `docker-compose.dev.yaml` files.

### Start Development Environment with Logs
To start the development environment and stream logs to the console:
```bash
make dev_uplog
```

### Stop Development Environment
To stop and remove all running containers in the development environment:
```bash
make dev_down
```

### Build Development Containers
To build or rebuild development containers:
```bash
make dev_build
```

### View Logs for Development Environment
To view logs from all containers in the development environment:
```bash
make dev_logs
```

### Rebuild and Restart Development Containers
To rebuild and restart containers for the development environment:
```bash
make dev_rebuild
```

---

## Deployment Workflow

### Deploy to Production
To deploy the application in production:
```bash
make deploy
```
This will:
1. Ensure the Docker network `linkin-chat-network` exists.
2. Start the production versions of NATS, Redis, PostgreSQL, and the server containers defined in the `docker-compose.prod.yaml` files.

### Stop Production Deployment
To stop and remove all running containers in the production environment:
```bash
make deploy-down
```

### View Logs for Production Deployment
To view logs from all containers in the production environment:
```bash
make deploy-logs
```

To view logs from a specific container in production:
```bash
make deploy-log CONTAINER=<container_name>
```
Replace `<container_name>` with the name of the container.

### Build Production Containers
To build or rebuild production containers:
```bash
make deploy-build
```

---

## Installing the Client

### For Ubuntu/Debian
1. Download the `.deb` file matching your architecture from the [Releases Section](https://github.com/Hossara/linkin-chat/releases):
   - For 64-bit systems: `linkin-chat_version-SNAPSHOT-x_linux_amd64.deb`
   - For ARM64 systems: `linkin-chat_version-SNAPSHOT-x_linux_arm64.deb`
   - For 32-bit systems: `linkin-chat_version-SNAPSHOT-x_linux_386.deb`
2. Install the package using `dpkg`:
   ```bash
   sudo dpkg -i linkin-chat_*.deb
   ```
3. Verify the installation:
   ```bash
   linkin-chat --help
   ```

### For macOS
1. Download the `.tar.gz` file matching your architecture from the [Releases Section](https://github.com/Hossara/linkin-chat/releases):
   - For ARM64: `linkin-chat_Darwin_arm64.tar.gz`
   - For x86_64: `linkin-chat_Darwin_x86_64.tar.gz`
2. Extract the archive:
   ```bash
   tar -xzf linkin-chat_Darwin_*.tar.gz
   ```
3. Move the binary to a directory in your PATH:
   ```bash
   sudo mv linkin-chat /usr/local/bin/linkin-chat
   ```
4. Verify the installation:
   ```bash
   linkin-chat --help
   ```

### For Windows
1. Download the `.zip` file matching your architecture from the [Releases Section](https://github.com/Hossara/linkin-chat/releases):
   - For 64-bit systems: `linkin-chat_Windows_x86_64.zip`
   - For 32-bit systems: `linkin-chat_Windows_i386.zip`
   - For ARM64: `linkin-chat_Windows_arm64.zip`
2. Extract the archive using a tool like WinRAR or 7-Zip.
3. Move the `linkin-chat.exe` file to a directory in your PATH (e.g., `C:\Program Files\LinkinChat`).
4. Add the directory to your system's PATH environment variable if not already set.
5. Verify the installation:
   ```cmd
   linkin-chat.exe --help
   ```

---

## Creating a Release
To create a release snapshot using GoReleaser:
```bash
make create_release
```
This will generate a snapshot release of the project.

---

## Additional Notes
- The Makefile ensures that the Docker network `linkin-chat-network` exists before executing commands. This prevents network-related issues when starting containers.
- For troubleshooting or customization, refer to the `Makefile` for exact command definitions and paths.

## Features
Linkin Chat provides the following core functionalities:
1. **Create User:** Register a new user.
2. **Login User:** Log in to the application with your credentials.
3. **Create Chatroom:** Start a new chatroom for communication.
4. **Delete Chatroom:** Remove an existing chatroom that you own.
5. **Join Chatroom:** Join a chatroom to start chatting with others.
6. **Delete Your Chatroom:** Remove the chatrooms you created.

Users can navigate through the interface using:
- **Tab** and **Shift+Tab** to move between objects.
- **Arrow Keys** (Up/Down) to select items in a list.

## Known Issues
- [ ] **Nats Authentication & Authorization**
- [ ] **Hardcoded Variables:** Some variables in the code are hardcoded, which may cause inflexibility and potential errors.
- [ ] Logging systems for containers, servers, and clients.
- [ ] Integration of the zap logging library.
- [ ] Ban user feature.
- [ ] Cache queries.

This aspect required extensive effort to balance security, performance, and functionality within the limited timeframe.

## Planned Improvements
Once the foundational features are stable, the following improvements and features are planned:
- [ ] Add a fully functional chatroom experience with message persistence.
- [ ] Implement logging systems for better observability.
- [ ] Add security enhancements for user authentication and NATS connections.
- [ ] Introduce the "ban user" feature for moderators to manage chatrooms.
- [ ] Refactor code to remove hardcoded variables and improve configurability.

## Test video
![Vid](./tutorial.gif)