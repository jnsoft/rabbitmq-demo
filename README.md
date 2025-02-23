# rabbitmq-demo
Test and demo of RabbitMQ


```
go mod init github.com/jnsoft/rabbitmqdemo 
go get github.com/rabbitmq/amqp091-go
```

# Init
```
podman run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 docker.io/rabbitmq:4.0-management
podman run -it --rm --name rabbitmq --network host -p 5672:5672 -p 15672:15672 docker.io/rabbitmq:4.0-management
```

# Run
```
go run src/sender/send.go
```
