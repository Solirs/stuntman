build:
	@echo "Building stuntman..."
	@go build stuntman.go colors.go

install:
	@echo "Installing stuntman..."
	@go build stuntman.go colors.go
	@cp stuntman /usr/local/bin