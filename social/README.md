# Social Media API

A modern, RESTful social media backend API built with Go, featuring user management and post creation functionality.

## 🚀 Features

- **User Management**: Create and manage user accounts with secure authentication
- **Post Creation**: Users can create posts with titles, content, and tags
- **Post Retrieval**: Fetch individual posts by ID
- **Health Monitoring**: Built-in health check endpoint
- **Database Migrations**: Automated database schema management
- **Docker Support**: Easy local development with Docker Compose
- **Hot Reloading**: Development environment with Air for automatic recompilation

## 🏗️ Architecture

This project follows Go best practices with a clean architecture:

```
social/
├── cmd/
│   ├── api/          # HTTP API handlers and main application
│   └── migrate/      # Database migration files
├── internal/
│   ├── db/          # Database connection utilities
│   ├── env/         # Environment variable management
│   └── store/       # Data access layer (repositories)
├── bin/             # Compiled binaries
├── docs/            # Documentation
└── scripts/         # Build and deployment scripts
```

## 🛠️ Tech Stack

- **Language**: Go 1.23.4
- **Database**: PostgreSQL 16.3
- **Router**: Chi v5 (lightweight, fast HTTP router)
- **Database Driver**: lib/pq (PostgreSQL driver)
- **Migrations**: golang-migrate
- **Development**: Air (hot reloading)
- **Containerization**: Docker & Docker Compose

## 📋 Prerequisites

- Go 1.23.4 or later
- PostgreSQL 16+ (or use Docker Compose)
- Docker & Docker Compose (optional, for local development)

## 🚀 Quick Start

### Using Docker Compose (Recommended)

1. **Clone and navigate to the project**:
   ```bash
   git clone <your-repo-url>
   cd social
   ```

2. **Start the database**:
   ```bash
   docker-compose up -d db
   ```

3. **Run database migrations**:
   ```bash
   # On Windows
   migrate.exe -path cmd/migrate/migrations -database "postgres://admin:adminpassword@localhost:5432/social?sslmode=disable" up
   
   # On Linux/Mac
   migrate -path cmd/migrate/migrations -database "postgres://admin:adminpassword@localhost:5432/social?sslmode=disable" up
   ```

4. **Start the API server**:
   ```bash
   go run cmd/api/*.go
   ```

   Or with hot reloading (requires Air):
   ```bash
   air
   ```

The API will be available at `http://localhost:8080`

### Manual Setup

1. **Set up PostgreSQL database**:
   ```sql
   CREATE DATABASE social;
   CREATE USER admin WITH PASSWORD 'adminpassword';
   GRANT ALL PRIVILEGES ON DATABASE social TO admin;
   ```

2. **Set environment variables** (optional):
   ```bash
   export DB_ADDR="postgres://admin:adminpassword@localhost:5432/social?sslmode=disable"
   export ADDR=":8080"
   export ENV="development"
   ```

3. **Run migrations and start the server** (same as steps 3-4 above)

## 📡 API Endpoints

### Health Check
```http
GET /v1/health
```
**Response**:
```json
{
  "status": "ok",
  "env": "development",
  "version": "0.0.1"
}
```

### Posts

#### Create a Post
```http
POST /v1/posts
Content-Type: application/json

{
  "title": "My First Post",
  "content": "This is the content of my post",
  "tags": ["golang", "api", "social"]
}
```

**Response**:
```json
{
  "id": 1,
  "title": "My First Post",
  "content": "This is the content of my post",
  "user_id": 1,
  "tags": ["golang", "api", "social"],
  "created_at": "2024-01-15T10:30:00Z",
  "updated_at": "2024-01-15T10:30:00Z"
}
```

#### Get a Post
```http
GET /v1/posts/{postID}
```

**Response**:
```json
{
  "id": 1,
  "title": "My First Post",
  "content": "This is the content of my post",
  "user_id": 1,
  "tags": ["golang", "api", "social"],
  "created_at": "2024-01-15T10:30:00Z",
  "updated_at": "2024-01-15T10:30:00Z"
}
```

## 🗄️ Database Schema

### Users Table
```sql
CREATE TABLE users (
    id bigserial PRIMARY KEY,
    email citext UNIQUE NOT NULL,
    username varchar(255) UNIQUE NOT NULL,
    password bytea NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now()
);
```

### Posts Table
```sql
CREATE TABLE posts (
    id bigserial PRIMARY KEY,
    title text NOT NULL,
    user_id bigint NOT NULL REFERENCES users(id),
    content text NOT NULL,
    tags varchar(100)[],
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now()
);
```

## ⚙️ Configuration

The application can be configured using environment variables:

| Variable | Default | Description |
|----------|---------|-------------|
| `ADDR` | `:8080` | Server address and port |
| `DB_ADDR` | `postgres://admin:adminpassword@localhost/social?sslmode=disable` | Database connection string |
| `DB_MAX_OPEN_CONNS` | `30` | Maximum open database connections |
| `DB_MAX_IDLE_CONNS` | `30` | Maximum idle database connections |
| `DB_MAX_IDLE_TIME` | `15m` | Maximum connection idle time |
| `ENV` | `development` | Environment (development/production) |

## 🔄 Database Migrations

Migrations are located in `cmd/migrate/migrations/` and follow the naming convention:
```
{version}_{description}.{up|down}.sql
```

**Run migrations**:
```bash
migrate -path cmd/migrate/migrations -database $DB_ADDR up
```

**Rollback migrations**:
```bash
migrate -path cmd/migrate/migrations -database $DB_ADDR down 1
```

## 🧪 Development

### Hot Reloading with Air

Install Air for automatic recompilation during development:
```bash
go install github.com/cosmtrek/air@latest
```

Start development server:
```bash
air
```

### Project Structure Details

- **cmd/api/**: Main application entry point and HTTP handlers
  - `main.go`: Application bootstrap and configuration
  - `api.go`: Route definitions and middleware setup
  - `posts.go`: Post-related HTTP handlers
  - `health.go`: Health check handler
  - `errors.go`: Error handling utilities
  - `json.go`: JSON encoding/decoding utilities

- **internal/**: Private application code
  - `db/`: Database connection management
  - `env/`: Environment variable utilities
  - `store/`: Data access layer with repository pattern

## 🚧 Roadmap

- [ ] User authentication and authorization
- [ ] User registration and login endpoints
- [ ] Post comments and likes
- [ ] User following system
- [ ] File upload for post images
- [ ] Full-text search for posts
- [ ] Rate limiting and API key management
- [ ] Comprehensive test suite
- [ ] API documentation with Swagger

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 📞 Support

If you have any questions or need help with setup, please open an issue on GitHub. 