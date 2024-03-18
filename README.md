# Golang REST API Personalities
This is a RESTful API built with Go. It provides CRUD operations for managing "personalidades".

## Getting Started
These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites
- Go 1.21.3 or later
- Docker
### Installing
1. Clone the repository:
```bash
git clone https://github.com/helioLJ/golang-rest-api-personalities.git
```
2. Navigate to the project directory:
```bash
cd golang-rest-api-personalities
```
3. Install the dependencies:
```bash
go mod download
```
4. Run the application:
```bash
go run main.go
```
The server will start on `localhost:8080`.

## API Endpoints
- `GET /api/personalidades`: Fetch all personalidades
- `GET /api/personalidades/{id}`: Fetch a single personalidade by ID
- `POST /api/personalidades`: Create a new personalidade
- `DELETE /api/personalidades/{id}`: Delete a personalidade by ID
- `PUT /api/personalidades/{id}`: Update a personalidade by ID
## Built With
- Go - The programming language used
- Gorilla Mux - The HTTP router and URL matcher for building Go web servers
- PostgreSQL - The database used
