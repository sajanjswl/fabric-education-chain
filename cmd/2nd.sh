




docker exec -e "CORE_PEER_MSPCONFIGPATH=/tmp/hyperledger/org1/admin/msp" -e "CORE_PEER_ADDRESS=peer1-org1:7051" cli-org1 peer lifecycle chaincode approveformyorg -o orderer1-org0:7050  --channelID mychannel --name basic --version 1.0 --package-id basic_1.0:ff54eebe63611ce3f035834941f67354256275992b5b4c1184307efe6a46c2d0 --sequence 1 --tls --cafile /tmp/hyperledger/org1/peer1/tls-msp/tlscacerts/tls-0-0-0-0-7052.pem
sleep 2s
echo "                                                                "
echo "################################################################"
echo "############### approving the chaincode in peer1-org1     ######"
echo "################################################################"
echo "                                                                "



docker exec -e "CORE_PEER_MSPCONFIGPATH=/tmp/hyperledger/org2/admin/msp" -e "CORE_PEER_ADDRESS=peer1-org2:7051" cli-org2 peer lifecycle chaincode approveformyorg -o orderer1-org0:7050  --channelID mychannel --name basic --version 1.0 --package-id basic_1.0:ff54eebe63611ce3f035834941f67354256275992b5b4c1184307efe6a46c2d0 --sequence 1 --tls --cafile /tmp/hyperledger/org2/peer1/tls-msp/tlscacerts/tls-0-0-0-0-7052.pem
sleep 2s
echo "                                                                "
echo "################################################################"
echo "############### approving the chaincode in peer1-org2     ######"
echo "################################################################"
echo "                                                                "



docker exec -e "CORE_PEER_MSPCONFIGPATH=/tmp/hyperledger/org1/admin/msp" -e "CORE_PEER_ADDRESS=peer1-org1:7051" cli-org1 peer lifecycle chaincode checkcommitreadiness  --channelID mychannel --name basic --version 1.0  --sequence 1 --tls --cafile /tmp/hyperledger/org1/peer1/tls-msp/tlscacerts/tls-0-0-0-0-7052.pem --output json
sleep 2s
echo "                                                                "
echo "################################################################"
echo "############### check commitreadniess of chaincode        ######"
echo "################################################################"
echo "                                                                "


docker exec -e "CORE_PEER_MSPCONFIGPATH=/tmp/hyperledger/org1/admin/msp" -e "CORE_PEER_ADDRESS=peer1-org1:7051" cli-org1 peer lifecycle chaincode commit -o orderer1-org0:7050  --channelID mychannel --name basic --version 1.0 --sequence 1 --tls --cafile /tmp/hyperledger/org1/peer1/tls-msp/tlscacerts/tls-0-0-0-0-7052.pem  --peerAddresses peer1-org1:7051 --tlsRootCertFiles /tmp/hyperledger/org1/peer1/assets/tls-ca/tls-ca-cert.pem --peerAddresses peer1-org2:7051 --tlsRootCertFiles /tmp/hyperledger/org1/peer1/assets/tls-ca/tls-ca-cert.pem
sleep 2s
echo "                                                                "
echo "################################################################"
echo "############### commiting chaincode to the channel        ######"
echo "################################################################"
echo "                                                                "




docker exec -e "CORE_PEER_MSPCONFIGPATH=/tmp/hyperledger/org1/admin/msp" -e "CORE_PEER_ADDRESS=peer1-org1:7051" cli-org1 peer lifecycle chaincode querycommitted --channelID mychannel --name basic --cafile /tmp/hyperledger/org1/peer1/tls-msp/tlscacerts/tls-0-0-0-0-7052.pem --output json
sleep 2s
echo "                                                                "
echo "################################################################"
echo "############### check if chaincode is committed to the block ###"
echo "################################################################"
echo "                                                                "



docker exec -e "CORE_PEER_MSPCONFIGPATH=/tmp/hyperledger/org1/admin/msp" -e "CORE_PEER_ADDRESS=peer1-org1:7051" cli-org1 peer chaincode invoke -o orderer1-org0:7050 \
 --tls --cafile /tmp/hyperledger/org1/peer1/tls-msp/tlscacerts/tls-0-0-0-0-7052.pem  --peerAddresses peer1-org1:7051 --tlsRootCertFiles /tmp/hyperledger/org1/peer1/assets/tls-ca/tls-ca-cert.pem --peerAddresses peer1-org2:7051 --tlsRootCertFiles /tmp/hyperledger/org1/peer1/assets/tls-ca/tls-ca-cert.pem  -C mychannel -n basic  -c '{"function":"InitLedger","Args":[]}'
sleep 2s
echo "                                                                "
echo "################################################################"
echo "############### invoking the chaincode peer1-org1            ###"
echo "################################################################"
echo "                                                                "


docker exec -e "CORE_PEER_MSPCONFIGPATH=/tmp/hyperledger/org1/admin/msp" -e "CORE_PEER_ADDRESS=peer1-org1:7051" cli-org1 peer chaincode query -C mychannel -n basic -c '{"Args":["QueryStudent","1816129"]}'

sleep 2s
echo "                                                                "
echo "################################################################"
echo "############### querry the   chaincode peer1-org1            ###"
echo "################################################################"
echo "                                                                "




#  docker exec -e "CORE_PEER_MSPCONFIGPATH=/tmp/hyperledger/org1/admin/msp" -e "CORE_PEER_ADDRESS=peer1-org1:7051" cli-org1 peer chaincode query -C mychannel -n basic -c '{"Args":["GenrateReport","1816129"]}'

# sleep 2s
# echo "                                                                "
# echo "################################################################"
# echo "############### generate report the   chaincode peer1-org1     ##"
# echo "################################################################"
# echo "                                                                "

#  docker exec ipfs ipfs cat bafyreibkgtfrdsxwvaean5niv6klvgoflzm74k6ntcuikyz5jgduzuk574

# docker exec -e "CORE_PEER_MSPCONFIGPATH=/tmp/hyperledger/org1/admin/msp" -e "CORE_PEER_ADDRESS=peer1-org1:7051" cli-org1 peer channel getinfo -c mychannel
# sleep 2s
# echo "                                                                "
# echo "################################################################"
# echo "############### querying channel info                    #######"
# echo "################################################################"
# echo "                                                                "
















