FROM golang:1.18-buster

WORKDIR ./src

COPY ./ ./

RUN go build -o main ./cmd/web
RUN rm ./cmd -r
RUN rm ./internal -r

# install psql
RUN apt-get update
RUN apt-get -y install postgresql-client

# make wait-for-postgres.sh executable
RUN chmod +x wait-for-postgres.sh

CMD ["./main", "-config", "resources/config/docker-app-confit.yaml"]

