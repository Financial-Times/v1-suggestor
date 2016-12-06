[![Coverage Status](https://coveralls.io/repos/github/Financial-Times/v1-suggestor/badge.svg?branch=master)](https://coveralls.io/github/Financial-Times/v1-suggestor?branch=master)

# V1 suggestor
* Reads V1 metadata from the source kafka queue
* filters and transforms it to UP standard json representation
* puts the result on the destination kafka queue

# How to run

```
export|set SRC_ADDR=http://kafka:8080
export|set SRC_GROUP=FooGroup
export|set SRC_TOPIC=FooBarEvents
export|set SRC_QUEUE=kafka
export|set SRC_CONCURRENT_PROCESSING=true
export|set DEST_ADDRESS=http://kafka:8080
export|set DEST_TOPIC=DestTopic
export|set DEST_QUEUE=kafka

./v1-suggestor[.exe]
```

Note: Brigthcove video metadata is the same pipeline as metadata for Methode articles, so brands are added in the same way.