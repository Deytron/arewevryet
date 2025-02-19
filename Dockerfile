FROM golang:1.23.4-alpine3.21 AS builder

# Stage 1 : dependencies download
WORKDIR /app
COPY src/go.mod src/go.sum ./
RUN go mod download
COPY  src .
ENV GOCACHE=/home/user-apps/.cache/go-build
RUN --mount=type=cache,target="/home/user-apps/.cache/go-build" CGO_ENABLED=0 GOOS=linux go build -p 8 -o /arewevryet

FROM alpine:3.21
# Stage 2 : Simply deploy the built bin

WORKDIR /app
COPY  src .
COPY --from=builder /arewevryet .
EXPOSE 9445

CMD ["/app/arewevryet"]