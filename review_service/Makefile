run:
	go run cmd/main.go

run_script:
	./script/gen-proto.sh

migrate_up:
	migrate -path migrations/ -database postgres://muhammadumar:123@database-1.c9lxq3r1itbt.us-east-1.rds.amazonaws.com/ratingdb up

migrate_down:
	migrate -path migrations/ -database postgres://muhammadumar:123@database-1.c9lxq3r1itbt.us-east-1.rds.amazonaws.com/ratingdb down

migrate_force:
	migrate -path migrations/ -database postgres://muhammadumar:123@database-1.c9lxq3r1itbt.us-east-1.rds.amazonaws.com/ratingdb force 2