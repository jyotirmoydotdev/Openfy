<p align="center">
<a href="" target="">
<img src="https://raw.githubusercontent.com/jyotirmoydotdev/Openfy/main/src/OpenfyLogo.svg" alt="Openfy logo">
</a>
</p>

> ⚠️ Note: This project is currently under development and may not be fully functional. Feel free to explore the code, but be aware that some features may not work as intended.

# Openfy

Openfy is a self-hosted e-commerce platform built with Go and Gin, providing a flexible and customizable solution for managing products and handling user authentication.

## Features

- **Product Management**: Easily create, update, and delete products.
- **User Authentication**: Secure user authentication with JWT tokens.
- **Admin Panel**: Admin-specific routes for managing products.
- **Flexible and Extendable**: Built with Go and Gin for flexibility and extensibility.

## Prerequisites

Before running Openfy, ensure you have the following installed:

- Go (1.15 or higher)

## Getting Started

1. Clone the repository:

   ```bash
   git clone https://github.com/jyotirmoydotdev/openfy.git
   cd openfy
   ```

2. Install dependencies:

   ```bash
   go mod download
   ```

3. Set up your configuration:

   Copy the `.env.example` file to `.env` and configure your environment variables, such as database connection details and secret keys.

4. Run the application:

   ```bash
   go run main.go
   ```

   The application will be accessible at `http://localhost:8080`.

## Routes

- **User Routes**:
  - `POST /signup`: Register a new user.
  - `POST /login`: User login.
  - `GET /products`: Retrieve details of all active products.

- **Admin Routes**:
  - `POST /admin/signup`: Register a new admin user.
  - `POST /admin/login`: Admin login.

- **Authenticated User Routes** (Protected by JWT):
  - `GET /user/ping`: Test endpoint for authenticated users.

- **Authenticated Admin Routes** (Protected by JWT):
  - `GET /admin/product`: Retrieve details of a specific product.
  - `GET /admin/products`: Retrieve details of all products.
  - `POST /admin/product/new`: Create a new product.
  - `PUT /admin/product`: Update an existing product.
  - `DELETE /admin/product`: Delete a product.
  - `DELETE /admin/variant`: Delete a product variant.


## Authentication

- User authentication is implemented using JWT tokens.
- To authenticate, include the JWT token in the "Authorization" header using the "Bearer" scheme.

## Preview

To experience a sneak peek of our project, please visit the following link: [Project Preview](https://www.figma.com/embed?embed_host=share&url=https%3A%2F%2Fwww.figma.com%2Fproto%2FDdnZ03JxOvicLeXQuts4gA%2FOpenfy%3Fpage-id%3D0%253A1%26type%3Ddesign%26node-id%3D33-20%26viewport%3D404%252C559%252C0.94%26t%3D98564Tf0lb67O8Ca-1%26scaling%3Dmin-zoom%26starting-point-node-id%3D33%253A20%26mode%3Ddesign).

Explore the features, test functionalities, and provide us with valuable feedback. Your insights are crucial in shaping the final release, and we appreciate your participation in this preview phase.

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.
