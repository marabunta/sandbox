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
	openssl req -x509 -newkey rsa:2048 -sha256 -nodes -keyout certs/CA.key -out certs/CA.pem -subj "/CN=example.com" -days 365
	# create server certs
	openssl req -newkey rsa:2048 -sha256 -nodes -keyout certs/server.key -out certs/server.csr -subj "/CN=server.example.com"
	openssl x509 -days 3065 -sha256 -req -in certs/server.csr -CA certs/CA.pem -CAkey certs/CA.key -set_serial 01 -out certs/server.crt -extfile <(printf "subjectAltName = DNS:localhost,DNS:server.example.com,IP:127.0.0.1,IP:0.0.0.0")
	# create client certs
	openssl req -newkey rsa:2048 -sha256 -nodes -keyout certs/client.key -out certs/client.csr -subj "/CN=client.example.com"
	openssl x509 -days 3065 -sha256 -req -in certs/client.csr -CA certs/CA.pem -CAkey certs/CA.key -set_serial 01 -out certs/client.crt -extfile <(printf "subjectAltName = DNS:localhost,DNS:client.example.com,IP:127.0.0.1,IP:0.0.0.0")

clean:
	@rm -rf certs client server
