# File Storage

A modern file storage service built with Go and React that enables secure file uploads and storage using AWS S3.

## Features

- File upload and storage to AWS S3
- CloudFront CDN integration for fast file delivery
- SQLite database for metadata management
- RESTful API with Gin framework
- React frontend with TypeScript and Vite
- Docker support for easy deployment
- Graceful shutdown handling

## Tech Stack

### Backend
- **Go 1.25.6** - Backend language
- **Gin** - HTTP web framework
- **GORM** - ORM for database operations
- **SQLite** - Database
- **AWS SDK v2** - S3 integration
- **OpenAI Go SDK** - AI integration

### Frontend
- **React 19** - UI library
- **TypeScript** - Type safety
- **Vite** - Build tool and dev server

### Infrastructure
- **Docker** - Containerization
- **AWS S3** - File storage
- **CloudFront** - CDN (optional)

## Prerequisites

- Go 1.25.6 or higher
- Node.js 18+ and npm
- Docker and Docker Compose (for containerized deployment)
- AWS account with S3 bucket configured
- AWS credentials (Access Key ID and Secret Access Key)

## Installation

### Clone the repository

```bash
git clone https://github.com/NikitaYurchyk/file-storage.git
cd file-storage
```

### Backend Setup

1. Install Go dependencies:
```bash
go mod download
```

2. Copy the example environment file:
```bash
cp .env.example .env
```

3. Configure your environment variables in `.env`:
```env
DB_PATH=./data/app.db
PORT=8091
S3_BUCKET=your-bucket-name
S3_REGION=us-west-1
S3_CF_DISTRO=your-cloudfront-id
CloudFrontDistributionDomain=your-cloudfront-domain
AWS_ACCESS_KEY_ID=your-access-key
AWS_SECRET_ACCESS_KEY=your-secret-key
```

### Frontend Setup

```bash
cd app
npm install
```

## Running the Application

### Option 1: Using Docker Compose (Recommended)

```bash
docker-compose up --build
```

The API will be available at `http://localhost:8091`

### Option 2: Running Locally

#### Start the Backend

```bash
# From project root
go run cmd/api/main.go
```

#### Start the Frontend

```bash
cd app
npm run dev
```

### Option 3: Build and Run the Binary

```bash
# Build the binary
go build -o api cmd/api/main.go

# Run the binary
./api
```

## Configuration

### Environment Variables

| Variable | Description | Required |
|----------|-------------|----------|
| `DB_PATH` | Path to SQLite database file | Yes |
| `PORT` | Server port (default: 8091) | Yes |
| `S3_BUCKET` | AWS S3 bucket name | Yes |
| `S3_REGION` | AWS S3 region | Yes |
| `S3_CF_DISTRO` | CloudFront distribution ID | No |
| `CloudFrontDistributionDomain` | CloudFront domain | No |
| `AWS_ACCESS_KEY_ID` | AWS access key | Yes |
| `AWS_SECRET_ACCESS_KEY` | AWS secret key | Yes |

### AWS S3 Setup

1. Create an S3 bucket in your AWS account
2. Configure bucket permissions for public/private access
3. (Optional) Set up CloudFront distribution for CDN
4. Generate AWS access credentials with S3 permissions

## Project Structure

```
.
├── cmd/
│   └── api/
│       └── main.go              # Application entry point
├── internal/
│   ├── DTO/                     # Data Transfer Objects
│   ├── apperror/                # Error handling
│   ├── config/                  # Configuration management
│   ├── handler/                 # HTTP handlers
│   ├── infrastructure/          # External services (S3)
│   ├── repository/              # Database layer
│   ├── router/                  # Route definitions
│   └── service/                 # Business logic
├── pkg/
│   └── database/                # Database utilities
├── app/                         # React frontend
│   ├── src/
│   └── package.json
├── data/                        # SQLite database storage
├── docker-compose.yml           # Docker compose configuration
├── Dockerfile                   # Docker image definition
├── .env.example                 # Example environment variables
└── go.mod                       # Go dependencies

```

## API Endpoints

### File Operations

- `POST /api/files/upload` - Upload a file
- `GET /api/files/:id` - Get file metadata
- `DELETE /api/files/:id` - Delete a file

## Development

### Run Tests

```bash
go test ./...
```

### Lint Frontend

```bash
cd app
npm run lint
```

### Build Frontend

```bash
cd app
npm run build
```

## Docker

### Build Docker Image

```bash
docker build -t file-storage .
```

### Run Docker Container

```bash
docker run -p 8091:8091 --env-file .env file-storage
```

## License

MIT

## Author

NikitaYurchyk
