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
- PostgreSQL
- Docker (optional, for running PostgreSQL in a container)

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

Alternatively, you can use Docker to run a PostgreSQL container:

```bash
    docker run --name postgres -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres
```

1. **Create the Database**
Connect to your PostgreSQL instance and create a new database:

```bash
    CREATE DATABASE tidetracker;
```

3. **Run Migrations**
Use `goose` for database migrations:

- Install `goose`:
  
```bash
    go install github.com/pressly/goose/v3/cmd/goose@latest
```

- Run migrations:
```bash
    goose -dir sql/migrations postgres <postgres_connection_url> up
```

### Generate Database Code with sqlc (Development Only)

The `sqlc` tool is used to generate Go code from SQL queries and schema definitions. This step is primarily relevant during the development phase, particularly when introducing new SQL queries or modifying existing ones. If you have not made any changes to your SQL queries or schema that would impact the generated code, you do not need to run `sqlc generate` again.

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

### TODO:
- Set up the database inside a docker container (maybe share network?)
- Finish dockerizing