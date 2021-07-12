test-driver-postgres:
	go test -v -coverprofile=coverage.out -coverpkg=./... ./... -storage-driver 'postgres'

test:
	go test -v -coverprofile=coverage.out -coverpkg=./... ./...

bench:
	go test -test.bench=. -run=^a ./...