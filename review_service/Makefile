run:
	go run cmd/main.go

run_script:
	./script/gen-proto.sh

migrate_up:
	migrate -path migrations/ -database postgres://postgres:123@localhost:5432/ratingdb up

migrate_down:
	migrate -path migrations/ -database postgres://postgres:123@localhost:5432/ratingdb down

migrate_force:
	migrate -path migrations/ -database postgres://postgres:123@localhost:5432/ratingdb force 2