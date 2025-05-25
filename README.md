# FollMe Comment Service

FollMe Comment Service is a backend service built with Go. It provides APIs for managing comments and WebSocket-based real-time interactions for posts. This project is designed to support scalable and efficient server-side applications.

## Features

- **Comment Management**: Create, retrieve, and manage comments for posts.
- **WebSocket Support**: Real-time updates for comments and interactions.
- **Database Integration**: Uses PostgreSQL for data storage.
- **Middleware**: Includes authentication and error-handling middleware.

## Project Structure

The project is organized into the following main directories:

- **cmd/**: Entry point for the application.
- **internal/**: Contains the core business logic and services.
  - **comment_service/**: Manages comment-related operations.
  - **story_with_you/**: Handles story-related operations.
- **pkg/**: Contains reusable packages such as middleware, configuration, and database utilities.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/FollMe-Comment-Service.git
   cd FollMe-Comment-Service
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Set up environment variables:
   - Create a `.env` file in the root directory.
   - Add the following variables:
     ```env
     # Database
     DB_HOST=<your-database-host>
     DB_NAME=<your-database-name>
     DB_USERNAME=<your-database-username>
     DB_PASSWORD=<your-database-password>
     DB_PORT=<your-database-port>

     # WebSocket Token
     WS_TOKEN=<your-websocket-token>
     ```

4. Run the application:
   ```bash
   go run cmd/comment-service/main.go
   ```

## API Documentation

### Comment Service
- **GET /api/comments/{postId}**: Retrieve comments for a specific post.
- **POST /api/comments**: Create a new comment for a post (protected).
- **POST /api/comments/count**: Get the number of comments for multiple posts.
- **GET /ws**: Establish a WebSocket connection.

### Story With You Service
- **GET /api/commit-date/{id}**: Retrieve commit date information.
- **POST /api/commit-date/{id}**: Update commit date information.

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Commit your changes with clear and descriptive messages.
4. Push your changes to your fork.
5. Submit a pull request.

## License

This project is authored by [Sum Duong](https://github.com/sumsv50). All rights reserved. You are free to use, modify, and distribute this software as long as proper credit is given to the author.