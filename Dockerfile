FROM gcr.io/test-servers-256610/ronin_ronin-base-image:base-image-0c47779be as builder

ENV GOPRIVATE="github.com/axieinfinity"

WORKDIR /opt

RUN --mount=type=secret,id=github_token git config --global url."https://x-access-token:$(cat /run/secrets/github_token)@github.com".insteadOf "https://github.com"

COPY . /opt/bridge
RUN cd bridge && make bridge

FROM debian:buster

RUN apt-get update  && apt-get install -y --no-install-recommends ca-certificates
RUN update-ca-certificates

WORKDIR "/opt"

ENV CONFIG_PATH ''
ENV RONIN_RPC ''
ENV RONIN_VALIDATOR_KEY ''
ENV RONIN_RELAYER_KEY ''
ENV RONIN_TASK_INTERVAL ''
ENV RONIN_TRANSACTION_CHECK_PERIOD ''
ENV RONIN_MAX_PROCESSING_TASKS ''

ENV RONIN_MAX_TASK_QUERY ''

ENV ETHEREUM_RPC ''
ENV ETHEREUM_VALIDATOR_KEY ''
ENV ETHEREUM_RELAYER_KEY ''
ENV ETHEREUM_GET_LOGS_BATCH_SIZE ''
ENV VERBOSITY 3
ENV NUMBER_OF_WORKERS 1024

ENV DB_HOST ''
ENV DB_PORT ''
ENV DB_NAME ''
ENV DB_USERNAME ''
ENV DB_PASSWORD ''
ENV DB_CONN_MAX_LIFE_TIME ''
ENV DB_MAX_IDLE_CONNS ''
ENV DB_MAX_OPEN_CONNS ''

COPY --from=builder /go/bin/bridge /usr/local/bin/bridge
COPY --from=builder /opt/bridge/config/ ./
COPY --from=builder /opt/bridge/docker/entrypoint.sh ./

ENTRYPOINT ["./entrypoint.sh"]