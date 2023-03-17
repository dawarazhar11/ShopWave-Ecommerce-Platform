# E-Commerce Platform

This repository contains the source code for a basic e-commerce platform built using Golang for the backend, Angular for the frontend, and PostgreSQL for the database. The platform allows users to browse products, add items to their cart, and complete the checkout process. It also includes user accounts, order history, and payment integration features.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Installation](#installation)
  - [Backend](#backend)
  - [Frontend](#frontend)
- [Usage](#usage)
  - [Running the Backend](#running-the-backend)
  - [Running the Frontend](#running-the-frontend)
- [Disclaimer](#disclaimer)

## Prerequisites

Before you begin, ensure you have the following installed on your local machine:

- [Go](https://golang.org/doc/install) (version 1.16 or later)
- [Node.js](https://nodejs.org/en/download/) (version 14.x or later)
- [Angular CLI](https://angular.io/cli) (version 12.x or later)
- [PostgreSQL](https://www.postgresql.org/download/)

## Installation

### Backend

1. Clone the repository:

git clone https://github.com/yourusername/ecommerce-platform.git


2. Change to the `backend` directory:

cd ecommerce-platform/backend


3. Download the required Go dependencies:

go mod download


4. Update the `main.go` file with your PostgreSQL credentials:

go
db, err := gorm.Open("postgres", "host=localhost dbname=ecommerce sslmode=disable user=youruser password=yourpassword")"

Replace youruser and yourpassword with your PostgreSQL username and password.

### Frontend

1. Change to the frontend directory:

cd ecommerce-platform/frontend

2. Install the required npm packages:

npm install


## Usage

### Running the Backend

1. Change to the backend directory:

cd ecommerce-platform/backend

2. Run the backend server:

go run main.go

The backend server will start on http://localhost:8080.

### Running the Frontend

1. Change to the frontend directory:

cd ecommerce-platform/frontend

2. Run the frontend development server:

ng serve


The frontend server will start on http://localhost:4200.

## Disclaimer

Please note that this e-commerce platform is intended for educational and demonstration purposes only. The platform has not been thoroughly tested, and some security features may not be implemented. For production use, consider making the necessary changes and improvements to ensure the application is secure and efficient.


Make sure to replace `yourusername` with your actual GitHub username or any other namespace that you prefer.


## License
This project is licensed under the [MIT License](https://opensource.org/licenses/MIT). 



