#!/bin/bash

curl --request POST -i \
     --url 'http://localhost:9003/v1/example'\
     -d "@../internal/validators/test_data/good.json"