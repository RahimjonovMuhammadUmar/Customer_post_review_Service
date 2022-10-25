run:
	go run cmd/main.go

run_script:
	./script/gen-proto.sh

migrate_up:
	migrate -path migrations/ -database postgres://postgres:123@localhost:5432/postdb up

migrate_down:
	migrate -path migrations/ -database postgres://postgres:123@localhost:5432/postdb down

migrate_force:
	migrate -path migrations/ -database postgres://postgres:123@localhost:5432/postdb force 2