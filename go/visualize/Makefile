vis: main.go
	go build -o vis

.PHONY: clean
clean:
	rm -f vis
 
.PHONY: go-format
go-format:
	go fmt $(shell sh -c "go list ./...")
