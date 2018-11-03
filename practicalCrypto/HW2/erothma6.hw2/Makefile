GO=go build

all: encrypt-auth decrypt-attack

encrypt-auth: encrypt-auth.go
	$(GO) encrypt-auth.go

decrypt-attack: decrypt-attack.go
	$(GO) decrypt-attack.go

clean:
	rm encrypt-auth decrypt-attack

