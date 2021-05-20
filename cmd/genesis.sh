
sudo rm ../org1/peers/peer1/assets/mychannel.block
export COMPOSE_PROJECT_NAME=cmd
docker-compose -f compose.yml down --volumes
sleep 10s
rm -r ../org0/orderers/orderer1/genesis.block
rm -r ../org0/orderers/orderer2/genesis.block
rm -r ../org0/orderers/orderer3/genesis.block
sleep 2s
rm -r ../org1/peers/peer1/assets/channel.tx
sleep 2s
rm -r ../org1/peers/peer1/assets/mychannel.block
rm -r ../org2/peers/peer1/assets/mychannel.block
sleep 2s

echo "                                                                "
echo "################################################################"
echo "############### cleaning the environment  ######################"
echo "################################################################"
echo "                                                                "
./configtxgen -profile SampleMultiNodeEtcdRaft -channelID byfn-sys-channel -outputBlock ../org0/orderers/orderer1/genesis.block
sleep 2s
cp ../org0/orderers/orderer1/genesis.block ../org0/orderers/orderer2/genesis.block
sleep 2s 
cp ../org0/orderers/orderer1/genesis.block ../org0/orderers/orderer3/genesis.block
echo "                                                          "
echo "################################################################"
echo "############### Generating the genesis block  ##################"
echo "################################################################"
echo "                                                                "

export CHANNEL_NAME=mychannel

./configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ../org1/peers/peer1/assets/channel.tx -channelID mychannel
sleep 2s
echo "                                                                "
echo "################################################################"
echo "############### Generating the transaction channel  ############"
echo "################################################################"
echo "                                                                "

export COMPOSE_PROJECT_NAME=cmd
docker-compose -f compose.yml up -d
sleep 10s

echo "                                                                "
echo "################################################################"
echo "############### starting the network                ############"
echo "################################################################"
echo "                                                                "


docker exec -e "CORE_PEER_MSPCONFIGPATH=/tmp/hyperledger/org1/admin/msp"  cli-org1 peer channel create -c mychannel -f /tmp/hyperledger/org1/peer1/assets/channel.tx -o orderer1-org0:7050 --outputBlock /tmp/hyperledger/org1/peer1/assets/mychannel.block --tls --cafile /tmp/hyperledger/org1/peer1/tls-msp/tlscacerts/tls-0-0-0-0-7052.pem
sleep 2s
cp ../org1/peers/peer1/assets/mychannel.block ../org2/peers/peer1/assets/mychannel.block


echo "                                                                "
echo "################################################################"
echo "############### creating  channel                   ############"
echo "################################################################"
echo "                                                                "


docker exec -e "CORE_PEER_MSPCONFIGPATH=/tmp/hyperledger/org1/admin/msp" -e "CORE_PEER_ADDRESS=peer1-org1:7051" cli-org1 peer channel join -b /tmp/hyperledger/org1/peer1/assets/mychannel.block
sleep 2s
echo "                                                                "
echo "################################################################"
echo "############### joining peer1-org1 to my channel    ############"
echo "################################################################"
echo "                                                                "







docker exec -e "CORE_PEER_MSPCONFIGPATH=/tmp/hyperledger/org2/admin/msp" -e "CORE_PEER_ADDRESS=peer1-org2:7051" cli-org2 peer channel join -b /tmp/hyperledger/org2/peer1/assets/mychannel.block
sleep 2s
echo "                                                                "
echo "################################################################"
echo "############### joining peer1-org2 to my channel    ############"
echo "################################################################"
echo "                                                                "




docker exec -e "CORE_PEER_MSPCONFIGPATH=/tmp/hyperledger/org1/admin/msp" -e "CORE_PEER_ADDRESS=peer1-org1:7051" cli-org1 peer lifecycle chaincode package basic.tar.gz --path /opt/gopath/src/github.com/hyperledger/fabric-samples/chaincode/sacc/ --lang golang --label basic_1.0
sleep 5s
echo "                                                                "
echo "################################################################"
echo "############### packgaging chaincode                     #######"
echo "################################################################"
echo "                                                                "

docker exec -e "CORE_PEER_MSPCONFIGPATH=/tmp/hyperledger/org1/admin/msp" -e "CORE_PEER_ADDRESS=peer1-org1:7051" cli-org1 peer lifecycle chaincode install basic.tar.gz
sleep 2s
echo "                                                                "
echo "################################################################"
echo "############### installing chaincode on peer1-org1       #######"
echo "################################################################"
echo "                                                                "


docker exec -e "CORE_PEER_MSPCONFIGPATH=/tmp/hyperledger/org2/admin/msp" -e "CORE_PEER_ADDRESS=peer1-org2:7051" cli-org2 peer lifecycle chaincode install basic.tar.gz
sleep 2s
echo "                                                                "
echo "################################################################"
echo "############### installing chaincode on peer1-org2      #######"
echo "################################################################"
echo "                                                                "




# docker exec -e "CORE_PEER_MSPCONFIGPATH=/tmp/hyperledger/org2/admin/msp" -e "CORE_PEER_ADDRESS=peer1-org2:7051" cli-org2 peer lifecycle chaincode queryinstalled
# sleep 2s
# echo "                                                                "
# echo "################################################################"
# echo "############### quering the package id                   #######"
# echo "################################################################"
# echo "                                                                "

