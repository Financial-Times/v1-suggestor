FROM alpine:3.3

ADD *.go /v1-suggestor/

RUN apk add --update bash \
  && apk --update add git bzr \
  && apk --update add go \
  && export GOPATH=/gopath \
  && REPO_PATH="github.com/Financial-Times/publish-availability-monitor" \
  && mkdir -p $GOPATH/src/${REPO_PATH} \
  && mv v1-suggestor/* $GOPATH/src/${REPO_PATH} \
  && cd $GOPATH/src/${REPO_PATH} \
  && go get \
  && go test ./... \
  && go build \
  && mv v1-suggestor /app \
  && apk del go git bzr \
  && rm -rf $GOPATH /var/cache/apk/*

CMD [ "/app" ] 