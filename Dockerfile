FROM golang:alpine as builder

LABEL maintainer="Mohammad Fikri <fikri.mohammad30@gmail.com>"

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download 

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ficree-api .

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/ficree-api .
COPY --from=builder /app/app.env .       

EXPOSE 3000

CMD ["./ficree-api"]