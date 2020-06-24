run:
	docker-compose up
migration:
	cd migrator && docker build -t migrator . && docker run --network host migrator -path=/migrations/ -database "postgresql://postgres:mypassword@localhost:5432/time_tracker?sslmode=disable" up
