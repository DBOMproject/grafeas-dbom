FROM golang:1.16 as builder

WORKDIR /workspace/src
COPY src/go.mod src/go.sum ./ 
RUN go mod download
COPY src/ ./

RUN CGO_ENABLED=0 go build -o grafeas-server .

FROM alpine:latest
WORKDIR /
COPY --from=builder /workspace/src/grafeas-server /grafeas-server
EXPOSE 8080
ENTRYPOINT ["/grafeas-server"]
