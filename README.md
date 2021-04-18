# hub

This is part of Maja home project.

Hub - main module in maja ecosystem. Module connect to *mqtt* server and manage requests. Results will be forward to mqtt server and receive to corresponding transport module.

Hub stored inside in database user accounts, device database and so on.

## Compile

go build

## Build docker image

docker build -t . maja/hub:latest

## Run

Recomend run module inside docker-compose
