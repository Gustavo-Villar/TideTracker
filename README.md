# TideTracker

TideTracker is an RSS feed aggregator application that allows users to subscribe to, manage, and view updates from their favorite feeds. Built with Go, it utilizes PostgreSQL for data persistence and offers a REST API for feed management and retrieval.

## Features

- User registration and management
- RSS feed subscription and management
- Feed update fetching and viewing
- Authentication via API keys

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go (version 1.21 or later)
- Goose (To run the migrations)
- PostgreSQL
- Docker (optional, for running the database and/or the app on a container)

### Setting Up the Project

1. **Clone the Repository**

```bash
   git clone https://github.com/Gustavo-Villar/TideTracker.git
   cd TideTracker
```

2. **Set Up Environment Variables**
Copy the .env.example file to a new file named .env and adjust the environment variables to match your local setup.

```bash
    cp .env.example .env
```
Edit the .env file with your PostgreSQL connection details and any other configuration you need.

3. **Install Dependencies**
Run the following command to install the Go dependencies:

```bash
    go mod tidy
```

### Setting Up the PostgreSQL Database

1. **Install PostgreSQL**
Follow the official documentation to install PostgreSQL on your system.

Alternatively, you can use Docker to run a PostgreSQL container as listed on the section below:

2. **Run Migrations**
Use `goose` for database migrations:

- Install `goose`:
  
```bash
    go install github.com/pressly/goose/v3/cmd/goose@latest
```

- Run migrations:
```bash
    goose -dir sql/schema postgres <postgres_connection_url> up
```

### Generate Database Code with sqlc (Development Only)

The `sqlc` tool is used to generate Go code from SQL queries and schema definitions. This step is primarily relevant during the development phase, particularly when introducing new SQL queries or modifying existing ones. If you have not made any changes to your SQL queries or schema that would impact the generated code, you do not need to run `sqlc generate` again.

To install `sqlc`:

```bash
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

To generate or update the database code after making changes, run:

```bash
    sqlc generate
```

This command should be executed in the project root.

Note for Contributors
If you're contributing to the TideTracker project and make changes affecting database interactions, please ensure to run sqlc generate and include the updated code in your pull requests.

### Building the Application
To compile the TideTracker application into an executable, run:

```bash
    go build
```
This command generates an executable file named TideTracker in the current directory.

### Running the Application
After building, you can start the TideTracker application by running:

```bash
    ./TideTracker
```
This command executes the compiled binary of the TideTracker application, starting the server on the port specified in your .env file.

### Making API Calls
Use a tool like `Postman` or `Thunder Client` to make API calls to the application. There are some example requests on the [thunder-collection.json](thunder-collection_TideTracker.json) file:

### Understanding the Vendor Folder
The vendor folder is part of Go's dependency management system. It is populated when you run go mod vendor and includes the exact versions of external packages your project is using. This folder is crucial for ensuring reproducible builds and dependency availability. It's only necessary to re-run go mod vendor if you've updated your dependencies.


# TideTracker Dockerization Guide

This guide covers the Docker setup for the TideTracker project. It includes instructions for running the application and PostgreSQL database using Docker and Docker Compose.

Before you start using the TideTracker application, it's essential to prepare the database structure by running the migrations explained above.

## Prerequisites

- Docker
- Docker Compose

## Development Database Setup

To run only the PostgreSQL database for development purposes, you can use a custom `docker-compose-db-only.yml` file.

### Up the Development Database

1. Ensure you have `docker-compose-db-only.yml` with the following content:

    ```yaml
    version: '3.8'

    services:
      db:
        image: postgres
        environment:
          POSTGRES_DB: tidetracker
          POSTGRES_USER: postgres
          POSTGRES_PASSWORD: postgres
        ports:
          - "5432:5432"
        volumes:
          - postgres_data_dev:/var/lib/postgresql/data

    volumes:
      postgres_data_dev:
    ```

2. Run the following command in the terminal:

    ```bash
    docker-compose -f docker-compose-db-only.yml up -d
    ```

This command starts a PostgreSQL container with the development configuration.

## Running the Application and Database Together

To run both the application and the database in containers, use Docker Compose with the main `docker-compose.yml` file.

### Up Both Services

1. Ensure your `docker-compose.yml` is configured correctly for both the `app` and `db` services.

2. Start the services by running:

    ```bash
    docker-compose up -d --build
    ```

This command builds and starts both the application and the database containers based on the configurations provided in `docker-compose.yml`.

## Dockerizing Only the Application

If you need to Dockerize only the application (e.g., for deployment or testing in isolation), follow these steps:

### Build the Application Image

1. Build the Docker image for the application:

    ```bash
    docker build . -t tidetracker:latest
    ```

This command builds a Docker image from the `Dockerfile` in the current directory and tags it as `tidetracker:latest`.

### Run the Application Container

2. Run the application container:

    ```bash
    docker run -p 8080:8080 tidetracker
    ```

This command runs the `tidetracker` Docker image as a container and maps port 8080 from the container to port 8080 on the host machine, allowing you to access the application at `http://localhost:8080`.

## Notes

- Make sure to adjust environment variables and ports as needed for your specific development or production environments.
- Regularly backup your database data, especially when running in production environments.
- Always test your Docker configurations in a development environment before deploying to production.
- When pushing to DockerHub (or other similar service) remember to set the correct tags so that you can have both a numbered version and the latest version actually up to date
```bash
docker build -t username/imagename:0.0.0 -t username/imagename:latest .
docker push username/imagename --all-tags
```