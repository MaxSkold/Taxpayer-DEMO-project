# RESTful API for Managing Taxpayers

## Description
This project is a RESTful API developed in Go, designed for managing taxpayer data. It provides endpoints for CRUD operations, ensuring seamless interaction between the backend and frontend. The system includes features for validating input, handling errors, and securely managing data using PostgreSQL.

---

## Features

- Full CRUD support for taxpayers.
- Validation of taxpayer data (e.g., unique TIN, positive income).
- Easy-to-navigate frontend built with HTML5 and Bootstrap.
- RESTful API powered by Go and the Gin framework.
- PostgreSQL as the relational database.

---

## Technology Stack

- **Backend**: Go, Gin, Gorm
- **Frontend**: HTML5, Bootstrap, JavaScript, Axios
- **Database**: PostgreSQL

---

## Installation and Setup

### Prerequisites

1. Go (latest version)
2. PostgreSQL

### Steps

1. **Clone the repository:**
   ```bash
   git clone https://github.com/your-username/taxpayer-api.git
   cd taxpayer-api
   ```

2. **Install dependencies:**
   ```bash
   go mod tidy
   ```

3. **Set up the database:**
   - Create a PostgreSQL database.
   - Update the database connection details in the code (e.g., environment variables or config file).

4. **Run the application:**
   ```bash
   go run main.go
   ```

5. **Access the application:**
   - API: `http://localhost:8080`
   - Frontend: `http://localhost:8080/static`

---

## API Endpoints

### Taxpayers

- **Get all taxpayers**
  - `GET /taxpayers`

- **Add a taxpayer**
  - `POST /taxpayer`
  - Request body:
    ```json
    {
        "name": "Ivan Ivanov",
        "tin": "13-28-087306-07",
        "address": "Moscow, Lenina St. 1",
        "income": 50000
    }
    ```

- **Update a taxpayer**
  - `PUT /taxpayer/:id`

- **Delete a taxpayer**
  - `DELETE /taxpayer/:id`

---

## Screenshots

### Frontend UI
![Frontend UI](path/to/frontend-screenshot.png)

### API Request Example
![API Request](path/to/api-request-screenshot.png)

---

## Known Issues

- Large payloads may cause slower response times.
- Error messages could be more descriptive for complex validation cases.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
