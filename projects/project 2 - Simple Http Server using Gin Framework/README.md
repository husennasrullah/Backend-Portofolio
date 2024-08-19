# CRUD PROJECT

This project contains an app with several endpoints, including:
1. API for Inserting a Product
2. API for Getting a List of Products
3. API for Viewing a Product

This app also implements Docker for containerization and Redis for temporary data caching. Error handling in this project is centralized through a single function.

# PROJECT ARCHITECTURE

This project uses an architecture with the following structure: `attribute`, `config`, `dto`, `endpoint`, `model`, `router`, `service`, `util`, `main.go`. The project employs a layered architecture, where each file with different functions is encapsulated in a separate package. This design facilitates easier function lookup and centralized error handling, while also avoiding import cycles.

- **attribute**: Contains functions for initiating DB and Redis connections used during project runtime.
- **config**: Contains configurations used in the app.
- **dto**: Contains structs for request payloads when making API requests.
- **model**: Contains structs representing product models or error models.
- **router**: Contains routing configurations for endpoints.
- **service**: Contains functions for data retrieval from the database.
- **util**: Contains utility functions required by other services.

# RUNNING THE PROJECT

1. **Manual Setup**
   - Create a new database named `erajaya`.
   - Configure the DB and Redis settings in the `config.json` file.
   - Add an environment variable in Windows or your IDE with the following env: `projectconfig=./config`.
   - Run the application.
   - Happy testing!

2. **Running with Docker**
   - Execute the following command in the terminal:
     ```bash
     docker-compose up
     ```
   - Adjust the environment settings in the `docker-compose.yml` file.
