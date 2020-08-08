.PHONY: build_publisher run_publisher build_consumer run_consumer send_message

build_publisher: ;@echo "Building publisher...\n"; \
	cd ./publisher && go build publisher.go

run_publisher: ;@echo "Running publisher...\n"; \
	$1 $2 $3 $4 ./publisher/publisher

build_consumer: ;@echo "Building consumer...\n"; \
	cd ./consumer && go build consumer.go

run_consumer: ;@echo "Running consumer...\n"; \
	$1 $2 $3 $4 ./consumer/consumer

send_message:
	curl -X POST -H "Content-Type: application/json" -d "{ \"lang\": \"${LANG}\", \"code\": \"${CODE}\" }" http://localhost:8081/publish
