AUTH_SVC=auth
PUBLISHER_SVC=publisher
VERSION=v0.0.1

AUTH_BIN=$(PWD)/$(AUTH_SVC)/bin/$(AUTH_SVC)
AUTH_SRC=$(PWD)/$(AUTH_SVC)/*.go

PUBLISHER_BIN=$(PWD)/$(PUBLISHER_SVC)/bin/$(PUBLISHER_SVC)
PUBLISHER_SRC=$(PWD)/$(PUBLISHER_SVC)/*.go
PROTO_DIR=$(PWD)/$(PUBLISHER_SVC)/proto
PROTO_SRC=$(PROTO_DIR)/*.proto

COMPOSE=manifests/docker-compose.yml

AUTH_LDFLAGS='-extldflags "static" -X main.svcVersion=$(VERSION) -X main.svcName=$(AUTH_SVC)'
PUBLISHER_LDFLAGS='-extldflags "static" -X main.svcVersion=$(VERSION) -X main.svcName=$(PUBLISHER_SVC)'

clean c:
	@echo "[clean] Cleaning files..."
	@rm $(AUTH_BIN)
	@rm $(PUBLISHER_BIN)

auth_build:
	@echo "Build auth service"
	@GOOS=linux go build -o $(AUTH_BIN) -i -ldflags=$(AUTH_LDFLAGS) $(AUTH_SRC)

proto p:
	@protoc --proto_path=$(GOPATH)/src:$(PROTO_DIR)/ --go_out=plugins=grpc:$(GOPATH)/src $(PROTO_SRC)

publisher: proto
	@echo "Build publisher service"
	@GOOS=linux go build -o $(PUBLISHER_BIN) -i -ldflags=$(PUBLISHER_LDFLAGS) $(PUBLISHER_SRC)

run r: auth_build publisher
	docker-compose -f $(COMPOSE) up --build --force-recreate