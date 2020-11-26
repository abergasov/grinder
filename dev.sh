#!/usr/bin/env bash

clean() {
  echo "stop containers";
  docker container stop grnd.mf
  echo "drop containers"
  docker rm -v grnd.mf
}

clean


serviceList="grinderSql"
echo "RUNNING SERVICES: $serviceList"
echo "RUN docker-compose-dev.yml "
docker-compose -f docker-compose-dev.yml pull
docker-compose -f docker-compose-dev.yml up --build $serviceList