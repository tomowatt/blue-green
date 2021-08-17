FROM golang:1.16-alpine as builder

WORKDIR /build
COPY ./go.mod ./main.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o blue-green .

FROM scratch
WORKDIR /app
COPY --from=builder /build/blue-green blue-green
COPY ./template/ ./template

CMD [ "/app/blue-green" ]
