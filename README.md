# authorization-backend

Backend authentication service written in Go with Gin. Handles user
registration, login, JWT token management, and session storage. Built with
security and scalability in mind.

## Stack

- **[Go](https://go.dev/)** and **[Gin](https://gin-gonic.com/)** - HTTP server
- **[PostgreSQL](https://www.postgresql.org/)** - user and session storage
- **[Redis](https://redis.io/)** - JWT caching
- **[JWT](https://jwt.io/)** - token-based authentication
- **[bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)** - password hashing
- **[Docker](https://www.docker.com/)** - multi-stage build

## Endpoints

| Method | Path                    | Description              | Auth Required |
| ------ | ----------------------- | ------------------------ | ------------- |
| `POST` | `/api/v1/auth/register` | Register new user        | ❌            |
| `POST` | `/api/v1/auth/login`    | Login and get tokens     | ❌            |
| `POST` | `/api/v1/token/refresh` | Refresh access token     | ❌            |
| `POST` | `/api/v1/auth/logout`   | Logout and revoke tokens | ✅            |
| `GET`  | `/api/v1/profile`       | Get user profile         | ✅            |

## Running locally

Requirements: Go, PostgreSQL, Redis

1. Clone the repo and copy the example env file:

   ```bash
   git clone https://github.com/whymewaomi/authorization-backend.git
   cd auth-service
   cp .env.example .env
   ```

## Running with Docker

```bash
docker compose up -d --build
```
