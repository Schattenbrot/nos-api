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
  - Get a specific weapon by id (implemented)
  - Update a weapon by id (implemented)
  - Delete a weapon by id (implemented)
2. Full CRUD for fairies:
  - Create (implemented)
  - Get all fairies (implemented)
  - Get all fairies by element (implemented)
  - Get all fairies by maxlevel (implemented)
  - Get a specific fairy by id (implemented)
  - Update a fairy by id (implemented)
  - Delete a fairy by id (implemented)
3. Miscellaneous
  - Move handlers into handlers-module for better readability. (done, but no receiver anymore, moved config into config-module)
  - httprouter to chi (done)
  - Fix response messages to be more schemey. (done)
  - validation check
  - Savety (login required)
  - Thinking about GraphQL O.o worth or not worth?
  - Pagination?!
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

## Response format

Basically using the format described [HERE](https://github.com/cryptlex/rest-api-response-format).
POST it will return the inserted ObjectID instead of the message.
UPDATE it will return the number of updated entries.
DELETE will return the number of deleted entries.

## Contributing

I have no clue how I can make this cmd/api thing more readable since I'm using receivers to pass around the app configuration which I personally like.

If someone got an idea for that I would be happy to hear about it.

I did it ... but without receivers. If someone got an idea for a better solution I'm even more happy to hear about it. :3

## License

This project is licensed under the [MIT](LICENSE) License.