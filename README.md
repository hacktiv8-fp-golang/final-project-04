# Toko Belanja
![Project Status](https://img.shields.io/badge/status-in%20progress-yellow)
![Contributors](https://img.shields.io/badge/contributors-3-blue)
![License](https://img.shields.io/badge/license-MIT-green)

Toko Belanja is an e-commerce application that offers various features to facilitate users in comfortably and securely shopping for products. This project is part of the Final Project 04 for the Magang dan Studi Independen Bersertifikat (MSIB) program at Hacktiv8.

## Team Members

- Ahmad Arsy (GLNG-KS-08-05)
- Eki Alfani (GLNG-KS08-012)
- Fabianus Jericho Harapan Jaya Sipayung (GLNG-KS08-015)

## Technologies

The project is built using the following technologies:

- Programming Language: Golang
- Framework: Gin-Gonic
- Database: PostgreSQL
- HTTP Server: Gin HTTP Server
- API Docs: Swagger

## Installation

1. Clone the repo
```sh
  git clone https://github.com/hacktiv8-fp-golang/final-project-04.git
```
2. Install the required dependencies
```sh
  go get
```
3. Navigate to the root directory
```sh
  cd final-project-04
```
4. Create a .env file in the root directory
```sh
  touch .env
  # or
  echo > .env
```
5. Add environment variable
```sh
  DB_HOST=Database host (example: localhost)
  DB_USER=Database username (example: postgres)
  DB_PASSWORD=Database user's password (example: postgres)
  DB_PORT=Database port (example: 5432)
  DB_NAME=Name of the database to be used (example: postgres)
  DB_SSLMODE=Database SSL mode (example: disable)
  PORT=Port for the server to run on (example: 8080)
```
6. Run the project
```sh
  go run main.go
```
7. The project will be available at http://localhost:8080 by default.

## API Documentation

The API documentation can be found [here](https://final-project-04-production.up.railway.app/swagger/index.html).

## Deployment

The API has been deployed and can be accessed [here](https://final-project-04-production.up.railway.app/).

## Contributing

We welcome contributions from developers to improve this project. Please check the [ISSUES](https://github.com/hacktiv8-fp-golang/final-project-04/issues) for a list of tasks that need to be done.

## License

This project is licensed under the [MIT License](LICENSE).
