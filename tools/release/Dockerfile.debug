# Content managed by Project Forge, see [projectforge.md] for details.
FROM golang:alpine

LABEL "org.opencontainers.image.authors"="Kyle U"
LABEL "org.opencontainers.image.source"="https://github.com/kyleu/solitaire"
LABEL "org.opencontainers.image.vendor"="kyleu"
LABEL "org.opencontainers.image.title"="Solitaire"
LABEL "org.opencontainers.image.description"="A solitaire game... details soon"

RUN apk add --update --no-cache ca-certificates tzdata bash curl htop libc6-compat

RUN apk add --no-cache ca-certificates dpkg gcc git musl-dev \
    && mkdir -p "$GOPATH/src" "$GOPATH/bin" \
    && chmod -R 777 "$GOPATH"

RUN go install github.com/go-delve/delve/cmd/dlv@latest

SHELL ["/bin/bash", "-c"]

# main http port
EXPOSE 7777
# marketing port
EXPOSE 7778

WORKDIR /

ENTRYPOINT ["/solitaire", "-a", "0.0.0.0"]

COPY solitaire /
