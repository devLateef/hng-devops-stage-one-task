# Go API Project

## Description
Simple Go API with 3 endpoints deployed on AWS EC2 using Nginx reverse proxy.

## Run Locally
```bash
go run main.go

## Endpoints
- GET /
- GET /health
- GET /me

## Response:
```bash
{
"message": "API is running"
}
```bash
{
"message": "healthy"
}
```bash
{
"name": "Your Full Name",
"email": "you@example.com
",
"github": "https://github.com/yourusername
"
}
