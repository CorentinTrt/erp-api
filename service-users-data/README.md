# Service Users Data _v0.1.0

## Development
You can check on ./DEV.md to check development principles that i'm using on this project

### Run the development environment
  - We are using docker to run this service (api + db)
  - There is a `dev.docker-compose.yml` in `erp-api` folder to run the service
  - There is an hot-reload enabled on this service, so you don't need to rebuild containers when you make changes

### Request the API
  - The exported port for this service is 9090 on localhost

### Visualize database
  - Since we are using mongoDB for this service, we can use tools like Compass to visualize the database
  - Because the service run in docker containers, we have to access with a specific IP adress

To establish a connection with Compass:
  1. Check the IP and the port of the database container : `$ docker inspect database-users-data`
  2. Fill the connection string with the right username / password
