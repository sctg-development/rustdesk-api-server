FROM golang:1.22 as gobuilder
RUN apt update && apt install -y build-essential libsqlite3-dev
RUN mkdir /app
COPY . /app/
RUN cd /app/ && go mod tidy && CGO_ENABLED=1 go build -v -o rustdesk-api-server

FROM debian:bookworm-slim
LABEL maintainer="ronan <ronan@sctg.eu.org>"
RUN mkdir /app

# Copy the configuration file
COPY conf /app/conf

# Copy the master file
COPY --from=gobuilder /app/rustdesk-api-server /app/rustdesk-api-server
WORKDIR /app
ENTRYPOINT ["/app/rustdesk-api-server"]

# Export port numbers
EXPOSE 21114
