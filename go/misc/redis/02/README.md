# Requirements and Info

## Links

https://www.influxdata.com
https://grafana.com/grafana/dashboards/2587-k6-load-testing-results/

## Installation

brew install k6
docker pull loadimpact/k6
docker pull influxdb
docker pull grafana/grafana

## Commands

### Run k6 in docker commpose

docker compose run --rm k6 run /scripts/test.js -u5 -d5s

## Run k6 in native environment

k6 run ./scripts/test.js

## Run k6 with influxdb as a time series db storage

k6 run ./scripts/test.js -o influxdb=http://localh
ost:8086/k6