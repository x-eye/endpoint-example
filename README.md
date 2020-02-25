# endpoint-example
Homework for Skoltech

## Prerequisites
- [Docker](https://docs.docker.com/get-started/)
- Login to docker and do [what's necessary](https://docs.docker.com/install/linux/linux-postinstall/) 
- [docker-compose](https://docs.docker.com/compose/install/)

## Installation and deployment

1. Clone the code from https://github.com/x-eye/endpoint-example
1. Go into it `cd endpoint-example`
1. Run `docker-compose up -d`

## Configuration
Example service gets configured by passing parameters with environment variables:

Variable | Meaning | Default
--- | --- | ---
PORT | Port to listen on | 9003
ID_REGEXP | Regular expression to match ID fields against | `^([0-9A-Fa-f]{2}[-]){5}([0-9A-Fa-f]{2})$`
PARTNER_URL | URL of partner service to send packets to | n/a
PARTNER_QUEUE_MAX_LENGTH | Maximum capacity of partner packets queue | 1000
KAFKA_ADDR | Kafka broker address to produce events | n/a
KAFKA_TOPIC | Kafka topic to produce events | n/a

## Observing the work
1. To run sample: `cd scripts` then `sh good.sh`
1. To see what gets to the partner:`docker attach endpointexample_nc_1`
1. To see result in kafka: `docker exec -it endpointexample_kafka_1 kafka-console-consumer --bootstrap-server kafka:9092 --topic Example` 
1. To run tests: `docker exec -it endpointexample_example_1 make test`