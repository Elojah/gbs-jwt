# builder
FROM golang:1.15.2 AS builder
ARG GITHUB_TOKEN
WORKDIR /secrets
COPY . .
RUN echo "prod" > .version
RUN echo "machine github.com\n\tlogin elojah\n\tpassword ${GITHUB_TOKEN}" >> ~/.netrc
ENV GOPRIVATE=github.com/elojah/*
# RUN git config --global url."ssh://git@github.com/".insteadOf https://github.com/
RUN make tidy
RUN env GOARCH=amd64 GOOS=linux CGO_ENABLED=0 make secrets

# runner
FROM alpine
WORKDIR /app
COPY --from=builder /secrets/bin/gbs-jwt_secrets .
COPY --from=builder /secrets/config/secrets/prod.json .
RUN chmod +x /app/gbs-jwt_secrets
ENTRYPOINT ["/app/gbs-jwt_secrets", "/app/prod.json"]
