sliceperf: main.go
	go build -o sliceperf

.PHONY: go-format
go-format:
	go fmt $(shell sh -c "go list ./...")

.PHONY: clean
clean:
	rm -f sliceperf
 
