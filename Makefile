# Build the Golang backend binary.
build-backend:
	cd ./cmd/packcalculator && go build -o ../../bin/packcalculator

# Run the Golang backend application.
run-backend:
	cd ./cmd/packcalculator && go run main.go

# Build the Vue.js frontend application.
build-frontend:
	cd ./ui && npm install && npm run build

# Run the Vue.js frontend application.
run-frontend:
	cd ./ui && npm run serve -- --port 8081
