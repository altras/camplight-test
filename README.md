# User Management Application

This is a monorepo containing a user management application with a Go backend and a Next.js frontend.

## Prerequisites

- Docker
- Docker Compose

## Running the Application

1. Clone the repository:
   ```
   git clone https://github.com/altras/camplight-test.git
   cd camplight-test
   ```

2. Create a `.env` file in the root directory with the following content:
   ```
   DB_USER=postgres
   DB_PASSWORD=yourpassword
   DB_NAME=usermanagement
   ```

3. Build and run the application using Docker Compose:
   ```
   docker-compose up --build
   ```

4. Access the application:
   - Frontend: http://localhost:3000
   - Backend API: http://localhost:8080

## Project Structure

```
backend/
├── backend/
│   ├── cmd/
│   ├── internal/
│   │   ├── domain/
│   │   ├── application/
│   │   ├── infrastructure/
│   │   └── interfaces/
│   └── tests/
├── frontend/
│   ├── pages/
│   ├── components/
│   ├── services/
│   └── styles/
└── docker-compose.yml
```

## Development

- To run the backend tests: `cd backend && go test ./...`
- To run the frontend tests: `cd frontend && npm test`

## API Endpoints

- GET /api/users - List users
- POST /api/users - Create a new user
- DELETE /api/users/:id - Delete a user

## Technologies Used

- Backend: Go, PostgreSQL
- Frontend: Next.js, React Query
- Containerization: Docker, Docker Compose

## Contributing

Please read CONTRIBUTING.md for details on our code of conduct and the process for submitting pull requests.

## License

This project is licensed under the MIT License - see the LICENSE.md file for details.
