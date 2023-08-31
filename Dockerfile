FROM golang:1.19.3-alpine3.16 as builder
WORKDIR /app
COPY . .

RUN cd cmd/ && CGO_ENABLED=0 GOOS=linux go build -v -o ../smoketest

FROM alpine@sha256:21a3deaa0d32a8057914f36584b5288d2e5ecc984380bc0118285c70fa8c9300
COPY --from=builder /app/smoketest .

ENTRYPOINT ["/smoketest"]
