MIGRATE_CMD=migrate
MIGRATE_DIR=./migrations
DB_DSN=postgres://root:root@127.0.0.1:5438/job_portal?sslmode=disable
DATE=$(shell date +%Y%m%d_%H%M%S)

# Generates mocks for interfaces
INTERFACES_GO_FILES := $(shell find internal -name "interfaces.go")
INTERFACES_GEN_GO_FILES := $(INTERFACES_GO_FILES:%.go=%.mock.gen.go)

generate_mocks: $(INTERFACES_GEN_GO_FILES)
$(INTERFACES_GEN_GO_FILES): %.mock.gen.go: %.go
	@echo "Generating mocks $@ for $<"
	mockgen -source=$< -destination=$@ -package=$(shell basename $(dir $<))

# generate protobuf
PROTOBUF_FILES := $(shell find api/proto -name "*.proto")
PROTOBUF_GEN_FILES := $(PROTOBUF_FILES:api/proto/%.proto=generated/proto/%.pb.go)
generate_protobuf: $(PROTOBUF_GEN_FILES)
	@echo "Generating protobuf files"

$(PROTOBUF_GEN_FILES): generated/proto/%.pb.go: api/proto/%.proto
	@echo "Generating protobuf files $@ for $<"
	buf generate --path $<

generate: api/api.yml generate_mocks generate_protobuf
	mkdir -p generated/api
	oapi-codegen --package api -generate types $< > generated/api/api-types.gen.go

force:
	@$(MIGRATE_CMD) -path migrations -database=$(DB_DSN) force 20250108084709

create:
	@$(MIGRATE_CMD) create -ext sql -dir $(MIGRATE_DIR) $(NAME)

up:
	@$(MIGRATE_CMD) -source file://$(MIGRATE_DIR) -database=$(DB_DSN) up

reset:
	@$(MIGRATE_CMD) reset -dir $(MIGRATE_DIR)

refresh: reset up

down:
	@$(MIGRATE_CMD) -source file://$(MIGRATE_DIR) -database=$(DB_DSN) down
status:
	@$(MIGRATE_CMD) status -dir $(MIGRATE_DIR)

clean:
	find . -name "*.mock.gen.go" -type f -delete
	find . -name "*.out" -type f -delete

CONFIG_FILE=/etc/kafka/kafka-config/config.properties
BOOTSTRAP_SERVER=localhost:8003
TOPIC_NAME=jobpostetl
REPLICATION_FACTOR=1
PARTITIONS=1
create-topic:
	docker exec -it kafka_1 kafka-topics --create \
		--bootstrap-server $(BOOTSTRAP_SERVER) \
		--replication-factor $(REPLICATION_FACTOR) \
		--partitions $(PARTITIONS) \
		--topic $(TOPIC_NAME) \
		--command-config $(CONFIG_FILE)

install_dependency:
	./bin/install_dependency.sh

docker_compose:
	docker-compose up -d

unit_test:
	go test -coverprofile=coverage.out ./...


init: install_dependency docker_compose generate up create-topic

preview_open_api:
	redocly preview-docs api/api.yml

run_rest_api:
	go run ./cmd rest-api

run_consumer:
	go run ./cmd consumer-etl-job-post