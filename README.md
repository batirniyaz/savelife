# SaveLife Backend

SaveLife Backend is a Golang application designed to facilitate ambulance requests through a mobile application. The backend provides RESTful APIs for user authentication, ambulance request creation, and retrieval of user requests.

## Features

- **User Authentication:** Users can authenticate using their username and password. Authentication is handled securely using JWT tokens.
- **Ambulance Request Creation:** Users can create ambulance requests by providing details such as full name, date of birth, illness, location, and phone number.
- **Request Retrieval:** Users can retrieve their ambulance requests to view their status and details.

## Technologies Used

- **Golang:** Backend server development is done in Golang, a fast and efficient programming language.
- **SQLite:** Database storage is implemented using SQLite, a lightweight and self-contained SQL database engine.
- **Gorilla Mux:** HTTP router used for routing incoming requests to the appropriate handlers.
- **JWT-Go:** Library for JSON Web Token (JWT) authentication.
- **WebSocket:** Initially planned for real-time communication between users and ambulance cars, later replaced with REST API.

## Installation and Setup

1. Clone the repository:
   ```sh
   git clone https://github.com/torexanovich/savelife.git
2. Navigate to the backend directory:
    ```sh
    cd savelife/backend
3. Install dependencies:
    ```sh
    go mod download
4. Build and run the server:
    ```sh
    go run cmd/server/main.go

## API Endpoints
### Authentication:
- **POST /login:**  Authenticate user and obtain JWT token.
### Ambulance Requests
- **POST /request:** Create a new ambulance request
- **GET /requests:**  Retrieve all requests for the authenticated user.


## License
This project is licensed under the MIT License - see the LICENSE file for details.

