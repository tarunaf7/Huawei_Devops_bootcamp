













FROM python:3.12.8-alpine@sha256:54bec49592c8455de8d5983d984efff76b6417a6af9b5dcc8d0237bf6ad3bd20 AS base

FROM base AS builder

RUN apk update \
    && apk add --no-cache wget g++ linux-headers \
    && rm -rf /var/cache/apk/*

COPY requirements.txt .

RUN pip install --prefix="/install" -r requirements.txt

FROM base

RUN apk update \
    && apk add --no-cache libstdc++ \
    && rm -rf /var/cache/apk/*

WORKDIR /loadgen

COPY --from=builder /install /usr/local


COPY locustfile.py .


ENV GEVENT_SUPPORT=True

ENTRYPOINT locust --host=http://{{FRONTEND_ADDR}} --headless -u {{USERS:-10}} -r {{RATE:-1}} 2>&1
