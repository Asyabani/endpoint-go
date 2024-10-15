# endpoint-go

**endpoint-go** is a RESTful API developed in Go (Golang) for managing various endpoints. This project aims to provide a simple and efficient API structure, allowing users to create, read, update, and delete (CRUD) data through a clean and easy-to-use interface.

## Features

- **RESTful API**: Follows REST architecture principles for standard HTTP methods.
- **CRUD Operations**: Supports Create, Read, Update, and Delete operations for managing data.
- **JSON Format**: Data is exchanged in JSON format, making it lightweight and easy to handle.
- **SQLite Support**: Uses SQLite as the database for simple setup and ease of use.
- **Middleware Support**: Includes middleware for logging and error handling.

## Tech Stack

- **Go (Golang)**: The programming language used for developing the API.
- **Gorilla Mux**: A powerful URL router and dispatcher for matching incoming requests to their respective handler.
- **SQLite**: A lightweight, file-based database for data storage.

## Prerequisites

To run this project, ensure you have the following installed:
- [Go](https://golang.org/dl/)
- [SQLite](https://www.sqlite.org/download.html) (optional; can use embedded SQLite)

## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/Asyabani/endpoint-go.git
    ```

2. Navigate to the project directory:
    ```bash
    cd endpoint-go
    ```

3. Install the necessary Go modules:
    ```bash
    go mod tidy
    ```

4. Set up the SQLite database:
   - The application will automatically create an SQLite database file (`database.db`) if it does not exist.
   - You may need to modify the database connection settings in the code if necessary.

5. Start the server:
    ```bash
    go run main.go
    ```

   The server will run on [http://localhost:8080](http://localhost:8080) by default.

## API Endpoints

The following RESTful API endpoints are available:

http://localhost:1323/todos/2
- **GET** `/todos`: Retrieve a list of all items.
- **GET** `/todos/:id`: Retrieve a single item by ID.
- **POST** `/todos`: Create a new item.
- **PUT** `/todos/:id`: Update an existing item.
- **DELETE** `/todos/:id`: Delete an item.

## Directory Structure

```plaintext
endpoint-go/
│
├── db/                        # SQLite database file
│   └── database.db            # SQLite database
│
├── main.go                    # Entry point for the application
│
├── go.mod                     # Go module file
│
└── go.sum                     # Go module dependencies
```

## Usage

- **Start the server**: Run the server using `go run main.go`.
- **Test the API**: Use tools like [Postman](https://www.postman.com/) or [curl](https://curl.se/) to test the API endpoints.

### Example Requests

- **Get all items**:
    ```http
    GET http://localhost:1323/todos
    ```

- **Add a new item**:
    ```http
    POST http://localhost:1323/todos
    Content-Type: application/json

    {
        "name": "New Item",
        "description": "This is a new item."
    }
    ```

## Future Improvements

- **Authentication**: Implement user authentication to secure the API.
- **Pagination**: Add pagination support for item retrieval.
- **Unit Tests**: Implement unit tests for handlers and middleware.
- **Docker Support**: Provide Docker configuration for easier deployment.

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/Asyabani/endpoint-go/blob/main/LICENSE) file for details.

Feel free to fork, clone, and contribute to this project. If you have any suggestions or issues, please create an issue in the repository.