FROM alpine:3.21.3

ARG S3_BUCKET
ARG MINIO_ROOT_USER
ARG MINIO_ROOT_PASSWORD
ARG MINIO_REGION

COPY docker-entrypoint.sh .

RUN apk update && apk add --no-cache aws-cli bash curl  \
    && curl -sS -o minio https://dl.min.io/server/minio/release/linux-amd64/minio \
    && chmod +x minio \
    && mv minio /usr/local/bin/ \
    && curl -sS -o mc https://dl.min.io/client/mc/release/linux-amd64/mc \
    && chmod +x mc \
    && mv mc /usr/local/bin/

CMD ./docker-entrypoint.sh & minio server /data/minio --console-address ':9090'
