GO=go build

all: vigenere-analyze vigenere-decrypt vigenere-encrypt

vigenere-analyze: vigenere-analyze.go
	$(GO) vigenere-analyze.go

vigenere-decrypt: vigenere-decrypt.go
	$(GO) vigenere-decrypt.go

vigenere-encrypt: vigenere-encrypt.go
	$(GO) vigenere-encrypt.go

clean:
	rm vigenere-analyze vigenere-decrypt vigenere-encrypt

