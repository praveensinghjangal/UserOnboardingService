
FROM golang:1.20-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o Onboarding_Service .

FROM alpine:latest  

WORKDIR /root/

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/Onboarding_Service .

EXPOSE 8080

CMD ["./Onboarding_Service"]
