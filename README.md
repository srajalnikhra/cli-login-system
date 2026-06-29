# CLI Login System

A secure Command Line Interface (CLI) Login System built using Go, PostgreSQL, Docker, and TOTP-based Multi-Factor Authentication (MFA).

## Features

- User Registration
- Secure Login
- Password Hashing using bcrypt
- PostgreSQL Database
- Multi-Factor Authentication (TOTP)
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

```
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

## Running the Project (Without Docker)

### Clone Repository

```bash
git clone <repository-url>
cd cli-login-system
```

### Install Dependencies

```bash
go mod download
```

### Run PostgreSQL

Create a PostgreSQL database named:

```
cli_login_db
```

Default credentials used:

```
Host: localhost
Port: 5432
User: postgres
Password: password
Database: cli_login_db
```

### Run Project

```bash
go run main.go
```

---

## Running with Docker

Build and start the application:

```bash
docker compose up --build
```

Or run in detached mode:

```bash
docker compose up -d
```

Open the application container:

```bash
docker exec -it cli-login-app sh
```

Run the application:

```bash
./cli-login-system
```

Stop containers:

```bash
docker compose down
```

---

## MFA

The application supports Time-based One-Time Password (TOTP) authentication.

Steps:

1. Login
2. Enable MFA
3. Scan the generated QR Code using Google Authenticator or Microsoft Authenticator.
4. Use the generated OTP during future logins.

---

## Security Features

- Passwords are securely hashed using bcrypt.
- TOTP-based Multi-Factor Authentication.
- Secure PostgreSQL storage.
- Dockerized application environment.

---

## Author

**Srajal Nikhra**