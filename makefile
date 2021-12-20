build:
	@echo "Building stuntman..."
	go build stuntman.go 
	

install:
	@echo "Installing stuntman..."
	@echo "Compiling..."
	go build stuntman.go
	@echo "Copying to /usr/local/bin..."
	cp stuntman /usr/local/bin
uninstall:
	@echo "Uninstaling stuntman..."
	rm /usr/local/bin/stuntman
	@echo "stuntman sucessfully removed from /usr/local/bin and uninstalled..."