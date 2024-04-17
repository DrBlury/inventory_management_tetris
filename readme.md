# Inventory Management System API

This is a simple inventory management system API. It is built using golang and uses PostgreSQL as the database.

## Features (Work in progress) 🚧

-   [x] API documentation using openapi
-   [x] API test page (Swagger UI alternative)
-   [x] CRUD operations for items and inventories
-   [x] Move items between inventories (e.g. from a player inventory to a chest and vice versa)
-   Currency system (add/remove currency from inventories/players)
-   Add custom attributes to inventories/items (e.g. weight, capacity, etc.)
-   Add custom actions to items/inventories using scripting (golang/js)
-   Add custom checks for inventory placement using scripting (e.g. check if item can only be obtained once and if already obtained, do not allow to be placed in inventory)
-   Create loot tables and temporary inventories (e.g. for looting/lootboxes)
-   Feature toggles for different types of inventories/items (e.g. disable trading for certain items) or disable features completely like trading, mail system or shaped inventories/items
-   Sort items in inventory by category and custom sorting algorithms
-   Access control for items and inventories
-   Trading system (trade items between inventories and players)
-   Mail system (send items asynchronously between inventories and players)
-   Websocket support for real-time updates

## Required Dependencies

### Running the application

-   [Docker](https://docs.docker.com/get-docker/)
-   [Docker Compose](https://docs.docker.com/compose/install/)
-   [Taskfile](https://taskfile.dev/#/installation) (optional, you can run the commands manually)

### Development

-   [Golang](https://golang.org/dl/) (or only run it inside a container)
-   [oapi-codegen](https://github.com/deepmap/oapi-codegen) (to generate the API Types and server code from the openapi specification)
-   [Node.js](https://nodejs.org/en/download/) & [NPM](https://www.npmjs.com/get-npm) (to install the formatter/plugins)

Alternative: Use `task format-in-docker`

## Getting started

1. Clone the repository
2. Install the required software
3. Run `task gen-env-file` to generate the `.env` file
4. Change the environment variables in the generated `.env` file according to your needs
5. Run `task start` to start the application

Useful commands:

-   Run `task --list-all` to see all available task commands

## Changing the database schema

To change the database schema, adjust the `database/schema.dbml` file and run `task gen-sql` to generate the SQL schema file. Write a corresponding migration file and place it into the `database/migrations` folder. On the next start, the migration will be applied.

## API Documentation

The API documentation is generated from the openapi specification file `api-spec/bundle.yml`. To regenerate the API types and server code, run `task gen-api`.

### Prettier

This project uses prettier to format the code.
To format the code, run `task format`.
