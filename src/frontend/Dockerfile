













FROM  golang:1.23.4-alpine@sha256:c23339199a08b0e12032856908589a6d41a0dab141b8b3b21f156fc571a3f1d3 AS builder

WORKDIR /src


COPY go.mod go.sum ./
RUN go mod download
COPY . .


RUN  CGO_ENABLED=0 go build -o /go/bin/frontend .

FROM scratch
WORKDIR /src
COPY --from=builder /go/bin/frontend /src/server
COPY ./templates ./templates
COPY ./static ./static




ENV GOTRACEBACK=single

EXPOSE 8080
ENTRYPOINT ["/src/server"]
