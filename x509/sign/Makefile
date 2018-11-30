.PHONY: all

all: SHELL:=/bin/bash
all:
	@mkdir -p certs
	# crate CA
	# openssl req -new -x509 -key private-key.pem -out server.pem -days 730
	openssl req -x509 -nodes -days 3650 -newkey ec:<(openssl ecparam -name prime256v1) -keyout certs/CA.pem -out certs/CA.crt -subj "/CN=marabunta"
	openssl ec -in certs/CA.pem -outform DER -out certs/CA.key
