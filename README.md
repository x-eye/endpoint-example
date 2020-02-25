# endpoint-example
Homework for Skoltech

## Prerequisites
- [Docker](https://docs.docker.com/get-started/)
- Login to docker and do [what's necessary](https://docs.docker.com/install/linux/linux-postinstall/) 
- [docker-compose](https://docs.docker.com/compose/install/)

## Installation and deployment

1. Clone the code from https://github.com/x-eye/endpoint-example
1. Run `docker-compose up --build -d`

## Observing the work
1. To run sample: `cd scripts` then `sh good.sh`
1. To see what gets to the partner:`docker attach endpointexample_nc_1`
1. To see result in kafka: `docker exec -it endpointexample_kafka_1 bash` then 
`kafka-console-consumer --bootstrap-server kafka:9092 --topic Example` 
1. To run tests: `docker exec -it endpointexample_example_1 bash` then `make test`