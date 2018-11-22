.PHONY: all build certs clean

all: build
ifeq (,$(wildcard certs))
	$(MAKE) certs
endif

build:
	go build -o client src/client/main.go
	go build -o server src/server/main.go

certs: SHELL:=/bin/bash
certs:
	@mkdir -p certs

	# crate CA
	openssl req -x509 -nodes -days 3650 -newkey ec:<(openssl ecparam -name prime256v1) -keyout certs/CA.key -out certs/CA.pem -subj "/CN=example.com"

	# create server certs
	openssl req -new -nodes -newkey ec:<(openssl ecparam -name prime256v1) -keyout certs/server.key -out certs/server.csr -subj "/CN=server.example.com"
	openssl x509 -days 3065 -sha256 -req -in certs/server.csr -CA certs/CA.pem -CAkey certs/CA.key -set_serial 01 -out certs/server.crt

	# create client certs
	openssl req -new -nodes -newkey ec:<(openssl ecparam -name prime256v1) -keyout certs/client.key -out certs/client.csr -subj "/CN=client.example.com"
	openssl x509 -days 3065 -sha256 -req -in certs/client.csr -CA certs/CA.pem -CAkey certs/CA.key -set_serial 01 -out certs/client.crt


clean:
	@rm -rf certs client server
