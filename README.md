# Cinema-Ticketing
Personal project using Golang for learning purpose

Current Project State : On Development

# Start PostgreSQL container
docker-compose -f docker-compose.db.yml up -d

# Stop PostgreSQL container
docker-compose -f docker-compose.db.yml down

# For Dev Phase : Run with local code mounted
docker run -it --rm \
  -v $(pwd):/app \
  -w /app \
  -p 8080:8080 \
  golang:1.21-alpine \
  sh -c "go mod download && go run main.go"

# For Prod Phase : Build the final production image & Run

docker build -t cinema-app .
docker run -p 8080:8080 cinema-app

