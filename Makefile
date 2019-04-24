all:
	@godep go build -a -installsuffix cgo -o supernova-mgtw

test:
	@godep go test -v ./cmd