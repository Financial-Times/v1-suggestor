# V1 suggestor
* Reads V1 metadata from the source kafka queue
* filters and transforms it to UP standard json representation
* puts the result on the destination kafka queue

# How to run

`./v1-suggestor[.exe] -config=example-config.json`