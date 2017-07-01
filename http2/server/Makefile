all: cert
	go build

cert:
	openssl req -x509 -newkey rsa:4096 -sha256 -nodes -keyout server.key -out server.crt -subj "/CN=example.com" -days 3650

clean:
	@rm -f server*
