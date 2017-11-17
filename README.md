## mongodb

Simple demo for storing and retrieving students from MongoDB.
Setup to be used with mLabs hosted mongo deployment, or with local db.


## studentdb

Simple REST system for storing/retrieving Students info.

To build do:

   go install github.com/marni/imt2681_studentdb/cmd/studentdb

then run:

   $GOPATH/bin/studentdb


### Running the dockerised version

To run a dockerised variant (including mongo database), add a ``.env`` file to /cmd/studentdb/ containing

``
PORT=8080
DB_HOST=db
``

To run all services, run 

``
docker-compose up
``
in the root directory.

## webhooks

Simple demo for webhooks usage.

