# Golang REST API with GORM, Echo, PostgreSQL, JWT Authentication, and Roles

This is a RESTful API built with Golang, GORM, and Echo. It includes JWT-based authentication, role-based access control (admin, editor, user), and supports CRUD operations for both posts and comments. Swagger is used for API documentation.

## Features

- **User Authentication**: JWT-based authentication with secure password hashing using `bcrypt`.
- **Role-Based Authorization**:
  - Admin: Full access (CRUD) for both posts and comments.
  - Editor: Can create, update, and delete posts and comments.
  - User: Can view posts and comments.
- **Posts and Comments**: Users can create, update, delete posts and their associated comments. Certain routes are protected based on the user's role.
- **Swagger Documentation**: Auto-generated API documentation using Swagger.
- **Database**: PostgreSQL is used with GORM for ORM-based interactions.

## Requirements

- [Go](https://golang.org/dl/) 1.18+
- [PostgreSQL](https://www.postgresql.org/download/)
- [Swagger](https://swagger.io/)
- [GORM](https://gorm.io/)
- [Echo](https://echo.labstack.com/)

## Setup

1. Clone the repository:

    ```bash
    git clone https://github.com/repoleved08/Blog_Post.git
    cd Blog_Post
    ```

2. Install dependencies:

    ```bash
    go mod tidy
    ```

3. Set up environment variables:

    Create a `.env` file in the root directory with the following structure:

    ```env
    DB_DSN=host=localhost user=postgres password=Your@Paswword dbname=blog_post_db port=5432 sslmode=disable
    JWT_SECRET=thisismysecretjwttokenfromtechxtrasol
    APP_DOMAIN=localhost:8080
    ```

4. Run database migrations (if applicable):

    You can run migrations manually by executing:

    ```bash
    go run main.go migrate
    ```

5. Run the application:

    ```bash
    go run main.go
    ```

    The server should now be running at `http://localhost:8080`.

6. View Swagger Documentation:

    Once the server is running, you can access the Swagger documentation at:

    ```
    http://localhost:8080/swagger/index.html
    ```

## Folder Structure

```bash
.
├── config             # Configuration files (e.g., database configuration)
│   └── database.go
├── docs               # Swagger documentation files
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── handlers           # HTTP request handlers
│   ├── comment.go
│   ├── post.go
│   └── user.go
├── main.go            # Application entry point
├── middleware         # Middleware for JWT and role-based access
│   └── middleware.go
├── models             # Database models and DTOs
│   ├── comment.go
│   ├── dto
│   │   ├── logindto.go
│   │   └── registerdto.go
│   ├── post.go
│   └── user.go
├── routes             # API route definitions
│   └── routes.go
├── validators         # Input validation logic
│   └── validator.go
└── go.mod             # Go module definition
└── go.sum             # Go module checksum

```
## API Endpoints

### Authentication

- **Login**: `POST /auth/login`  
  Allows users to log in and receive a JWT token.

- **Register**: `POST /auth/register`  
  Allows new users to register.

### Post Endpoints

- **Create Post**: `POST /posts`  
  **Protected** (Admin, Editor)  
  Allows authenticated users to create a new post.

- **Get Posts**: `GET /posts`  
  **Public**  
  Retrieves a list of posts.

- **Update Post**: `PUT /posts/:id`  
  **Protected** (Admin, Editor)  
  Updates an existing post by ID.

- **Delete Post**: `DELETE /posts/:id`  
  **Protected** (Admin, Editor)  
  Deletes a post by ID.

### Comment Endpoints

- **Create Comment**: `POST /posts/:id/comments`  
  **Protected** (Admin, Editor)  
  Adds a new comment to a post.

- **Get Comments**: `GET /posts/:id/comments`  
  **Public**  
  Retrieves comments for a specific post.

- **Update Comment**: `PUT /posts/:id/comments/:comment_id`  
  **Protected** (Admin, Editor)  
  Updates a comment by ID.

- **Delete Comment**: `DELETE /posts/:id/comments/:comment_id`  
  **Protected** (Admin, Editor)  
  Deletes a comment by ID.

## Role-Based Access

- **Admin**: Full access to all resources (posts and comments).
- **Editor**: Can create, update, and delete posts and comments.
- **User**: Can view posts and comments.

## Middleware

- **JWT Authentication**: Protects create, update, and delete routes for posts and comments.
- **Role-Based Middleware**: Ensures that only users with the appropriate roles (admin, editor) can access certain routes.

## Swagger Documentation

To generate or update the Swagger documentation, use the `swag` package:

```bash
swag init
```
The generated documentation files will be placed in the docs folder. To view the documentation in your browser, visit:

```
http://localhost:8080/swagger/index.html
```

## License

This project is open-source and available under the [MIT License](LICENSE).
```


 ## Contact
- [LinkedIn](https://www.linkedin.com/in/norman-bii-87382722a)
- [Twitter](https://x.com/normangeek1)
- [GitHub](https://github.com/repoleved08)
- [Email](mailto:techxtrasol.design@gmail.com)
- [Website](https://techxtrasol.com)

