# Deploy locally
You can deploy **Photon** locally among two ways. Using _Docker Compose_ to start up all the required services through docker images or 
you could install all the required services _manually_ within your environment.

## Docker
After you got installed _Docker_ and _Docker Compose_ into your environment, add the required .env files or the required environment
variables. Once the previous process are fulfilled, you must get inside **Photon**'s root folder and run the following command.

```console
root@host:~$ docker-compose up
```

Stopping containers

```console
root@host:~$ docker-compose down
```

## Manually
In order to make **Photon** work correctly within your environment, you must install the following services.
* Redis
* PostgreSQL
* RabbitMQ

After completing the previous process, you must add the required .env files or the environment variables to every microservice.
Once the environment variables are all set, you must run the main SQL script called 'photon.sql' located at data/root.

You may run the microservices with the following command using _Golang_'s CLI.

```console
root@host:~$ go run main.go
```

Or you could run the binaries, compiling the _Golang_'s microservices first.

```console
root@host:~$ go build -o PROJECT_NAME .
root@host:~$ ./PROJECT_NAME
```
