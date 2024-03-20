# Inventory Management System API

This is a simple inventory management system API. It is built using golang and uses PostgreSQL as the database.

## Required Dependencies

### Running the application

- [Docker](https://docs.docker.com/get-docker/) (preferably)
- [Docker Compose](https://docs.docker.com/compose/install/) (preferably)
- [Taskfile](https://taskfile.dev/#/installation) (optional, you can run the commands manually)

### Development

- [Golang](https://golang.org/dl/) (or only run it inside a container)
- [Node.js](https://nodejs.org/en/download/) (to run the pnpm command)
- [pnpm](https://pnpm.io/installation) (to install the required node modules)
- [DBML cli](https://www.dbml.org/docs/#installation) (to change the database schema)
- [oapi-codegen](https://github.com/deepmap/oapi-codegen) (to generate the API Types and server code from the openapi specification)

## Getting started

1. Clone the repository
2. Install the required software
3. Run `task generate-env-file` to generate the environment file
4. Change the environment variables in the generated `.env` file according to your needs
5. Run `task start` to start the application

Useful commands:

- Run `task --list-all` to see all available task commands

## Changing the database schema

To change the database schema, adjust the `database/schema.dbml` file and run `task gen-sql` to generate the SQL schema file. Write a corresponding migration file and place it into the `database/migrations` folder. On the next start, the migration will be applied.

## API Documentation

The API documentation is generated from the openapi specification file `api/openapi.yaml`. To regenerate the API types and server code, run `task gen-api`.

### Prettier

This project uses prettier to format the code.
To format the code, run `task format`.
