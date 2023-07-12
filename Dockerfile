# Stage 1: Build
FROM golang:1.19.5 AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o product

# Stage 2: Run
FROM debian:buster-slim
WORKDIR /app
COPY --from=build /app/product .
COPY view/*.html view/
EXPOSE 3000
CMD ["./product"]
