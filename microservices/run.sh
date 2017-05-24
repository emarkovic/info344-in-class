#!/usr/bin/env bash

#TODO: 
# - create private network,
# - run hellosvc in network with no published ports
# - run gateway in network publishing port 443
#   and using volumes to give it access to your
#   cert and key files in `gateway/tls`
#   and setting environment variables
#    - CERTPATH = path to cert file in container
#    - KEYPATH = path to private key file in container
#    - HELLOADDR = net address of hellosvc container

docker network create microservice-net 

# node stuff
docker run -d \
--name nodehellosvc1 \
--network microservice-net \
em42/hellosvc

docker run -d \
--name nodehellosvc2 \
--network microservice-net \
em42/hellosvc

# go 
docker run -d \
--name gogateway \
--network microservice-net \
-p 443:443 \
-v $(pwd)/gateway/tls:/tls:ro \
-e CERTPATH=/tls/fullchain.pem \
-e KEYPATH=/tls/privkey.pem \
-e HELLOSVCADDR=nodehellosvc1,nodehellosvc2 \
em42/gateway