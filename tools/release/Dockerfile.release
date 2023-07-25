# Content managed by Project Forge, see [projectforge.md] for details.
FROM golang:alpine

LABEL "org.opencontainers.image.authors"="Kyle U"
LABEL "org.opencontainers.image.source"="https://github.com/kyleu/solitaire"
LABEL "org.opencontainers.image.vendor"="kyleu"
LABEL "org.opencontainers.image.title"="Solitaire"
LABEL "org.opencontainers.image.description"="A solitaire game... details soon"

RUN apk add --update --no-cache ca-certificates libc6-compat

# main http port
EXPOSE 7777
# marketing port
EXPOSE 7778

WORKDIR /

ENTRYPOINT ["/solitaire", "-a", "0.0.0.0"]

COPY solitaire /
