setup:
	go get -u github.com/google/wire/cmd/wire
	
start:
	docker-compose up development

build:
	docker build --target=production -t app:latest .

wire:
	cd pkg && go generate

test:
	go test -v ./pkg/...

# .PHONY is used for reserving tasks words
.PHONY: setup start build wire