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

Demo 1 (simple sender and receiver)
```
go run src/send_receive/sender/send.go
go run src/send_receive/receiver/receive.go
```

Demo2 (Two workers fetching tasks from task queue. A work queue is that each task is delivered to exactly one worker.)
```
# run in shell 1:
go run src/task_queue/worker/worker.go

# run in shell 2:
go run src/task_queue/worker/worker.go

# run in shell 3:
go run src/task_queue/task/task.go First message.
go run src/task_queue/task/task.go Second message..
go run src/task_queue/task/task.go Third  message...
go run src/task_queue/task/task.go Fourth message....
go run src/task_queue/task/task.go Fifth message.....
```

Demo3 (publish / subscribe, deliver a message to multiple consumers)
```

```
