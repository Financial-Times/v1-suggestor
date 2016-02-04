FROM alpine

ADD *.go /v1-suggestor/
ADD config.json.template /v1-suggestor/config.json
ADD startup.sh /

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
  && mv config.json /config.json \
  && apk del go git bzr \
  && rm -rf $GOPATH /var/cache/apk/*

ENTRYPOINT [ "/bin/sh", "-c" ]
CMD [ "/startup.sh" ] 
