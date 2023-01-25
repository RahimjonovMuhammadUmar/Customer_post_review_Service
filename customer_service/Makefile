run:
	go run cmd/main.go

run_script:
	./script/gen-proto.sh

migrate_up:
	migrate -path migrations/ -database postgres://postgres:123@db-postgres-1/customerdb up

migrate_down:
	migrate -path migrations/ -database postgres://postgres:123@db-postgres-1/customerdb down

migrate_force:
	migrate -path migrations/ -database postgres://postgres:123@db-postgres-1/customerdb force 1