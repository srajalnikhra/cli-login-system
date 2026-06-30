# CLI Login System

A secure Command Line Interface (CLI) Login System built using Go, PostgreSQL, Docker, and Time-based One-Time Password (TOTP) Multi-Factor Authentication (MFA).

## Features

- User Registration
- Secure Login
- Password Hashing using bcrypt
- PostgreSQL Database
- Time-based One-Time Password (TOTP) MFA
- QR Code Generation for MFA Setup
- User Profile
- Logout
- Help Menu
- Docker Support
- Docker Compose Support

## Tech Stack

- Go
- PostgreSQL
- Docker
- Docker Compose
- bcrypt
- pquerna/otp

## Project Structure

```text
cli-login-system/
├── internal/
│   ├── auth/
│   ├── cli/
│   ├── database/
│   ├── mfa/
│   ├── models/
│   ├── repository/
│   └── session/
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
├── main.go
└── README.md
```

---

# Running the Project (Without Docker)

## Clone the Repository

```bash
git clone https://github.com/srajalnikhra/cli-login-system.git
cd cli-login-system
```

## Install Dependencies

```bash
go mod download
```

## Configure PostgreSQL

Create a PostgreSQL database named:

```text
cli_login_db
```

Default database configuration:

```text
Host: localhost
Port: 5432
User: postgres
Password: password
Database: cli_login_db
```

## Run the Application

```bash
go run main.go
```

---

# Running with Docker

## Build the image and start the containers

Run this command the first time you use the project or whenever you modify the Go source code or Dockerfile.

```bash
docker compose up -d --build
```

## Launch the CLI application

```bash
docker exec -it cli-login-app ./cli-login-system
```

## Start existing containers (without rebuilding)

If the containers have been stopped and no code changes were made:

```bash
docker compose up -d
```

Then launch the CLI again:

```bash
docker exec -it cli-login-app ./cli-login-system
```

## Stop the containers

```bash
docker compose down
```

> **Note:** The CLI application runs inside the Docker container. Once the containers are running, use the `docker exec` command above whenever you want to launch the interactive CLI.

---

# Multi-Factor Authentication (MFA)

The application supports Time-based One-Time Password (TOTP) authentication.

### Steps

1. Register or Login.
2. Enable MFA.
3. Scan the generated QR Code using Google Authenticator or Microsoft Authenticator.
4. Enter the generated OTP during future logins.

---

# Security Features

- Passwords are securely hashed using bcrypt.
- Time-based One-Time Password (TOTP) authentication.
- Secure PostgreSQL storage.
- Dockerized application environment.

---

# Author

**Srajal Nikhra**
