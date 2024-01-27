## Braz E-commerce
This project is a small e-commerce done with good practices and the best of golang!

### First steps
As a new member of the engineering team the first step that you'll need to do is to execute the project setup command:
```sh
make setup
```

This project is developer friendly and will watch your changes and automatically reload the microservices. In order to achieve this result, you'll need to execute:
```sh
make start
```
The command above will start all microservices and infrastructure resources in docker containers and start to scan for file changes. To stop the running containers you can easily execute:
```sh
make stop
```
