[![Circle CI](https://circleci.com/gh/Financial-Times/v1-suggestor.svg?style=shield)](https://circleci.com/gh/Financial-Times/v1-suggestor)[![Go Report Card](https://goreportcard.com/badge/github.com/Financial-Times/v1-suggestor)](https://goreportcard.com/report/github.com/Financial-Times/v1-suggestor) [![Coverage Status](https://coveralls.io/repos/github/Financial-Times/v1-suggestor/badge.svg)](https://coveralls.io/github/Financial-Times/v1-suggestor)

# V1 suggestor
Processes metadata about content that comes from QMI system - aka V1 annotations.  

* Reads V1 metadata for an article from the  kafka source topic _NativeCmsMetadataPublicationEvents_
* Filters and transforms it to UP standard json representation
* Puts the result onto the kafka destination topic _ConceptSuggestions_

v1-suggestor service communicates with kafka via http-rest-proxy. It polls kafka-rest-proxy for messages and POSTs transfirmed messages to kafka-rest-proxy.  
This service is deployed in the Delivery clusters.
## Installation

* `go get -u github.com/Financial-Times/v1-sugestor`
* `cd $GOPATH/src/github.com/Financial-Times/v1-sugestor`
* `go install`

## Startup parameters

| **Parameter** | **Value in prod** | **Explained** |
|---|---|---|
| **SRC_ADDR** |_http://localhost:8080_ | Url of the _http-rest-proxy_ host to connect to in order to **receive** messages from kafka. |
| **SRC_GROUP** | _v1Suggestor_ | The consumer group for receiving messages from kafka. |
| **SRC_TOPIC** | _NativeCmsMetadataPublicationEvents_ | kafka topic to consume messages from. |
| **SRC_QUEUE** | _kafka_ |  Used by _Vulkan_ to route http requests based on _Host_ header. In docker cluster all hosts are at _http://localhost:8080_. This http header is supplied to distinguish one service from another.  Host header _kafka_ points to _http-rest-proxy_. |
| **SRC_CONCURRENT_PROCESSING** | _false_ | Should the consumer process messages concurrently or sequentially. |
| **DEST_ADDRESS** | _http://localhost:8080_| Url of the _http-rest-proxy_ host to connect to in order to **send** messages to kafka. In prod env this is typically the same address as the SRC_ADDR. |
| **DEST_TOPIC** | _ConceptSuggestions_ | kafka topic to **send** messages to.  |
| **DEST_QUEUE** | _kafka_ |  Used by _Vulkan_ to route http requests based on _Host_ header. In prod docker cluster it is the same as SRC_QUEUE. |


## Prerequisites
1.  In order to set LDFLAGS to provide correct versionning information to */build-info* endpoint we need to enable access to remote GIT repo via https (not ssh) because Docker does not have your ssh keys.                                                                                                                                                                     
    When we run the service locally (not in Docker) this is not required but you would need to execute manually the whole section of the Dockerfile that deals with setting LDFLAGS parameters.
````
git config remote.origin.url https://github.com/Financial-Times/v1-suggestor.git
````
When docker build finished do not forget to set remote.origin.url to use ssh
````
git config remote.origin.url git@github.com:Financial-Times/v1-suggestor.git
````
2. In order to run v1-suggestor you would need at least kafka/zookeeper and kafka-rest-proxy to be accessible somewhere
and you would need to provide the host and the port to connect to them as startup parameters.

## Run locally
````
   export|set SRC_ADDR=http://kafkahost:8080
   export|set SRC_GROUP=FooGroup
   export|set SRC_TOPIC=FooBarEvents
   export|set SRC_QUEUE=kafka
   export|set SRC_CONCURRENT_PROCESSING=true
   export|set DEST_ADDRESS=http://kafkahost:8080
   export|set DEST_TOPIC=DestTopic
   export|set DEST_QUEUE=kafka
   export|set ENVIRONMENT=coco-semantic
   export|set DOCKER_APP_VERSION=latest
````

````
./v1-suggestor[.exe]
````
## Run in Docker
````
docker build -t coco/v1-suggestor:$DOCKER_APP_VERSION .
````

````
docker run --name v1-suggestor -p 8080 \
--env "SRC_ADDR=http://kafka:8080" \
	--env "SRC_GROUP=v1Suggestor" \
	--env "SRC_TOPIC=NativeCmsMetadataPublicationEvents" \
	--env "SRC_QUEUE=kafka" \
	--env "SRC_CONCURRENT_PROCESSING=false" \
	--env "DEST_ADDRESS=http://kafka:8080" \
	--env "DEST_TOPIC=ConceptSuggestions" \
	--env "DEST_QUEUE=kafka" \
	--env "ENVIRONMENT=coco-$ENVIRONMENT_TAG" \
	coco/v1-suggestor:$DOCKER_APP_VERSION
````
 
## Admin Endpoints

|Endpoint        | Explained |
|---|---|
| /__health      | checks that v1-sugestor can communicate to kafka via http-rest-proxy|
|/__ping         | _response status_: **200**  _body_:**"pong"** |
|/ping           | the same as above for compatibility with Dropwizard java apps |
|/__gtg          | _response status_: **200** when "good to go" or **503** when not "good to go"|
|/__build-info   | consisting of _**version** (release tag), git **repository** url, **revision** (git commit-id), deployment **datetime**, **builder** (go or java or ...)_ 
|/build-info     | the same as above for compatibility with Dropwizard java apps |


Note: Brigthcove video metadata is the same pipeline as metadata for Methode articles, so brands are added in the same way.

_TODO: We should consider documenting the format of the in/out messages to kafka._