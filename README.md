# rabbitmq-go-publisher-consumer

## Makefile

All of the scripts is in the Makefile, check it out for more information

## Publisher

Build the publisher application:

```bash
  go build ./publisher/publisher.go
```

Run the publisher application:

```bash
  RABBIT_HOST=localhost RABBIT_PORT=5672 RABBIT_USERNAME=username RABBIT_PASSWORD=password ./publisher
```

Send the message on publish:

```bash
  curl -X POST -H "Content-Type: application/json" -d "{ \"lang\": \"go\", \"code\": \"fmt.Println(\"Hello there!\")\" }" http://localhost:8081/publish
```

<!--
### Dockerize Publisher

Build the publisher container:

```bash
  docker build . -t  nikolasmelui/rabbitmq-go-publisher:v1.0.0
```

Run the publisher container instance:

``bash
  docker run -it --rm --network rabbitmq -e RABBIT_HOST=localhost -e RABBIT_PORT=5672 -e RABBIT_USERNAME=nikolasmelui -e RABBIT_PASSWORD=password -p 8081:8081 .
```
-->
