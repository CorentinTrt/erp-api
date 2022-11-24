FROM golang:1.18-alpine3.15 AS builder

ENV GO111MODULE=on 

WORKDIR /service-users-data
COPY . .

ENV TZ America/Chicago

#dependencies
RUN go mod tidy 

WORKDIR /service-users-data/cmd/service-users-data
 
RUN GOOS=linux go build -a -installsuffix cgo -o app .


#Second stage for a smaller image
FROM alpine:3.15
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /service-users-data/cmd/service-users-data/app .
# COPY --from=builder /go/src/github.com/javiersoto15/skeleton-tutorial/cmd/tutorial/app .
CMD ["./app"]
