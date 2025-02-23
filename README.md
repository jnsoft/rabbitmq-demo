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

# output messages to file, only error and critical:
go run src/pubsub/receiver/receiver.go "*.error" "*.critical" &> .tmp/logs.log

# terminal 2 - output to screen, all messages:
go run src/pubsub/receiver/receiver.go "#"

# terminal 3 - send logs:
go run src/pubsub/logger/logger.go "kern.critical" "A critical kernel error"
go run src/pubsub/logger/logger.go "cron.error" "An error from cron"
go run src/pubsub/logger/logger.go "kern.info" "This is info from kernel"
```

Demo4 (RPC request reply)  
* The client creates an exclusive callback queue
* For an RPC request, the client sends a message with two properties: reply_to, which is set to the callback queue and correlation_id, which is set to a unique value for every request
* The request is sent to an rpc_queue queue
* The RPC worker (aka: server) is waiting for requests on that queue. When a request appears, it does the job and sends a message with the result back to the Client, using the queue from the reply_to field.
* he client waits for data on the callback queue. When a message appears, it checks the correlation_id property. If it matches the value from the request it returns the response to the application.

You can try running more servers (process more requests) and more clients (produce more requests)  

```
# terminal 1 (server)
go run src/rpc/server/server.go

# terminal 2 (client)
go run src/rpc/client/client.go 30
```

# RabbitMQ

A sender sends messages to an exchange (or in very simple cases, directly to a queue)
A receiver reads messages from a queue  

The relationship between exchange and a queue is called a binding  

Consumer acknowledgements, cover RabbitMQ communication with consumers.  
Publisher confirms cover publisher communication with RabbitMQ.  

### AMQP 0-9-1 protocol 

Has 14 properties predefined that go with a message
* persistent: Marks a message as persistent (wtrue) or transient (false)
* content_type: Used to describe the mime-type of the encoding. For example for the often used JSON encoding it is a good practice to set this property to: application/json.
* reply_to: Commonly used to name a callback queue.
* correlation_id: Useful to correlate RPC responses with requests


### Exchange
Exchange type:  
* fanout: broadcast to every attached queue
* direct: one-to-one or one-to-many, routed to queues based on an exact match between the message's routing key and the queue's binding key.
* topic: one-to-one or one-to-many, routed to queues based on based on pattern matching between the message's routing key and the queue's binding key.
* headers: Messages are routed based on the headers and their values. If you have a headers exchange with a binding rule that requires the headers {"type": "order", "format": "json"}, then only messages with headers that match both criteria will be routed to the corresponding queue.

Topic routing:  
* `*` can substitute for exactly one word
* `#` can substitute for zero or more words
If a routing key is <severity>.<system>.<environment>, then `*.ada.*` will catch all messages from the system ada. `*.*.prod` will catch all messages from environment prod. `error.#` will catch all errors from any system in any environment. A queue bound with `#` will catch all messages. A binding like `info.ada.prod` will behave like a binding to a direct exchange.


