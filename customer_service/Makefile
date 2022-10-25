run:
	go run cmd/main.go

run_script:
	./script/gen-proto.sh

migrate_up:
	migrate -path migrations/ -database postgres://postgres:123@localhost:5432/customerdb up

migrate_down:
	migrate -path migrations/ -database postgres://postgres:123@localhost:5432/customerdb down

migrate_force:
	migrate -path migrations/ -database postgres://postgres:123@localhost:5432/customerdb force 2