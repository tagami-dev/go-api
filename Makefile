.PHONY: up setup-db run-server test-article cleanup-db

up:
	docker-compose up -d

setup-db:
	mysql -h 127.0.0.1 -u docker sampledb -p < repositories/testdata/setupDB.sql

run-server:
	DB_USER="docker" DB_PASSWORD="docker" DB_NAME="sampledb" go run main.go

# Make a test HTTP request to the server
test-article:
	curl http://localhost:8080/article/1

cleanup-db:
	mysql -h 127.0.0.1 -u docker sampledb -p -e "source repositories/testdata/cleanupDB.sql"