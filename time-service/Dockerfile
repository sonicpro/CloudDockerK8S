# Загружаем образ для билда
FROM golang:1.13 as builder

# Set convenient workdir
WORKDIR /opt/time-service

# Copy the source file into container's file system
COPY main.go .

# Coберем программу в статический бинарник с названием time-service
RUN CGO_ENABLED=0 GOOS=linux go build -a -o time-service main.go

# The second stage - load linux Alpine image to create a runtime container
FROM alpine:3.10

# Номер порта, используемого контейнером
EXPOSE 8080
# Map builder container workdir to the same named runtime container workdir
WORKDIR /opt/time-service

# Copy the built binary from the first stage to the new workdir
COPY --from=builder /opt/time-service/time-service .

CMD [ "/opt/time-service/time-service" ]
