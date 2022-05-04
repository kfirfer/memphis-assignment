### Simple producer-consumer queue example using NATS Jetstream

* Web server: Client is consumer, Server is producer
* Persistent included
* Docker-compose for NATS streaming server


Run docker-compose (nats streaming server):
```bash
docker-compose up
```

Stop and remove docker-compose volume:
```bash
docker-compose down -v
```

CURL to the server (producer):
```bash
curl --location --request POST 'http://localhost:1323/sendMessage' \
--header 'Content-Type: application/json' \
--data-raw '{
    "message": "from messages subject"
}'
```