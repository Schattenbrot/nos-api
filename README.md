# Nos API
An API for finding several items in NosTale with the goal of creating damage calculator as frontend to shed some light on actual itemupgrades.

## Installation

For now it's pretty simple.
Either use the docker-compose with `docker-compose up`. Or already have database then ... good luck in finding the database-connection string in the code (it's at the start of the main function in cmd/api/main.go).

## Usage

Use it *kappa*

## Roadmap

1. Full CRUD for weapons:
  - Create (implemented)
  - Get all weapons (implemented)
  - Get all weapons based on profession (implemented)
  - Get a specific weapon by id
  - Update a weapon by id
  - Delete a weapon by id
2. Full CRUD for fairies:
  - Create (implemented)
  - Get all fairies (implemented)
  - Get all fairies by element
  - Get all fairies by maxlevel
  - Get a specific fairy by id
  - Update a fairy by id
  - Delete a fairy by id
3. Savety (login required)
  - Thinking about GraphQL O.o worth or not worth?
4. Full CRUD for armor:
  - Same as weapons ... just with armor
5. Full CRUD for accessories:
  - Create
  - Get all accessories
  - Get all accessories by slot
  - Get a specific accessory by id
  - Update an accessory by id
  - Delete an accessory by id
6. Full CRUD for costumes:
  - Create
  - Get all costumes
  - Get all costumes by slot
  - Get a specific costume by id
  - Update a costume by id
  - Delete a costume by id

## License

This project is licensed under the [MIT](LICENSE) License.