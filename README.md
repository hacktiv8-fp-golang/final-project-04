# Toko Belanja

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

## Deployment
The API has been deployed and can be accessed [here](https://final-project-04-production.up.railway.app/).
