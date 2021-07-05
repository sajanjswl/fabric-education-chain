docker-compose -f TLS-CA/tls-ca-server.yml up -d
sleep 3s

docker-compose -f org2/org2CA.yml up -d
sleep 3s

docker-compose -f org1/org1CA.yml up -d
sleep 3s

docker-compose -f org0/ordererCA.yml up -d