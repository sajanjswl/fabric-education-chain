
docker-compose -f TLS-CA/tls-ca-server.yml down
sleep 5s

docker-compose -f org2/org2CA.yml down
sleep 3s

docker-compose -f org1/org1CA.yml down
sleep 3s

docker-compose -f org0/ordererCA.yml down

