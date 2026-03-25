FROM golang:1.22-alpine AS build
WORKDIR /app
COPY . .
RUN go build -o app

FROM alpine
WORKDIR /app
COPY --from=build /app/app .
EXPOSE 8080
CMD [./app]


