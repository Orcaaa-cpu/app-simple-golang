Product Application
This is a web application for managing product data. It allows users to log in, register, and perform CRUD operations on product items.

Features
~User authentication (login, register, and logout)
~Google OAuth2 login integration
~Product item management (create, read, update, and delete)
~Input validation using Go Validator
~PostgreSQL database integration using go-pg library
~Docker containerization for easy deployment

Prerequisites
~Go 1.19.5 or later
~PostgreSQL 13 or later
~Docker (optional)

Installation

1. Clone the repository:

shell
Copy code
git clone https://github.com/your-username/product-app.git

2. Navigate to the project directory:

shell
Copy code
cd product-app

3. Install the necessary Go dependencies:

shell
Copy code
go mod download

4. Set up the PostgreSQL database with the following environment variables:

shell
Copy code
export DB_DRIVER=postgres
export DB_HOST=db
export DB_PORT=5432
export DB_NAME=go_product
export DB_USERNAME=root
export DB_PASSWORD=secret
Note: Adjust the values according to your database configuration.

5. Build the Go application:

shell
Copy code
go build -o product

6. Start the application:

shell
Copy code
./product
The application should now be running on http://localhost:3000.

Usage
1. Open a web browser and navigate to http://localhost:3000.
2. You will be redirected to the login page. If you don't have an account, click on the "Register" link  to create a new account.
3. After logging in or registering, you will be redirected to the product item management page.
4. On the product item management page, you can perform the following actions:
~View a list of existing product items.
~Add a new product item by clicking on the "Add Product" button and filling out the form.
~Edit an existing product item by clicking on the "Edit" button next to the item.
~Delete a product item by clicking on the "Delete" button next to the item.
5. You can also log out by clicking on the "Logout" link in the navigation bar.

Docker Deployment
To deploy the application using Docker, follow these steps:

1. Make sure you have Docker installed and running on your system.

2. Build the Docker image:

shell
Copy code
docker build -t product-app .

3. Start the Docker containers:

shell
Copy code
docker-compose up -d
The application will be available at http://localhost:3000.

Contributing
Contributions are welcome! If you find any issues or would like to add new features, please submit an issue or a pull request.

License
This project is licensed under the MIT License.

Feel free to update and modify the README.md file according to your needs.