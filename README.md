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
