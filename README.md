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
mkdir .tmp
# output messages (warning and error) to file:
go run src/pubsub/receiver/receiver.go warning error &> .tmp/logs.log
# terminal 2 - output to screen (all messages, info warning error):
go run src/pubsub/receiver/receiver.go info warning error
# terminal 3 - send logs:
go run src/pubsub/logger/logger.go error "This is an error!"
go run src/pubsub/logger/logger.go warning "This is a warning!"
go run src/pubsub/logger/logger.go info "This is info!"

```

# RabbitMQ

A sender sends messages to an exchange (or in very simple cases, directly to a queue)
A receiver reads messages from a queue  
That relationship between exchange and a queue is called a binding  



### Exchange
Exchange type:  
* fanout: broadcast to every attached queue
* direct: one-to-one or one-to-many, routed to queues based on an exact match between the message's routing key and the queue's binding key.
* topic: one-to-one or one-to-many, routed to queues based on based on pattern matching between the message's routing key and the queue's binding key.
* headers: Messages are routed based on the headers and their values. If you have a headers exchange with a binding rule that requires the headers {"type": "order", "format": "json"}, then only messages with headers that match both criteria will be routed to the corresponding queue.
