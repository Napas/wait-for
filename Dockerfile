FROM        golang:1.13-alpine AS builder

RUN         mkdir /app

WORKDIR     /app

COPY        ./ /app

RUN         go build -v -o wait-for . && \
            chmod +x wait-for

FROM        golang:1.13-alpine

COPY        --from=builder /app/wait-for /usr/local/bin/wait-for

ENTRYPOINT  ["wait-for"]

CMD         ["-h"]