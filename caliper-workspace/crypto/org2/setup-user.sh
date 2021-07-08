################################################################
##############       org1-user1 setup                  ##############
################################################################
rm -r ./user1
mkdir  -p user1/ca
sleep 2s
cp ./ca/server/crypto/ca-cert.pem ./user1/ca/org1-ca-cert.pem
sleep 1s
export FABRIC_CA_CLIENT_HOME=./user1
export FABRIC_CA_CLIENT_TLS_CERTFILES=ca/org1-ca-cert.pem
export FABRIC_CA_CLIENT_MSPDIR=msp
fabric-ca-client enroll -d -u https://user-org2:org2UserPW@0.0.0.0:7055

sleep 2s

# here we are creating admincerts dierctory in  every peer msp so that all the peers including the orderer have there  org's admin cert there
# mkdir peers/peer1/msp/admincerts
# mkdir peers/peer2/msp/admincerts
# mkdir admin/msp/admincerts
# sleep 2s
# cp ./admin/msp/signcerts/cert.pem ./peers/peer1/msp/admincerts/org1-admin-cert.pem
# cp ./admin/msp/signcerts/cert.pem ./peers/peer2/msp/admincerts/org1-admin-cert.pem
# cp ./admin/msp/signcerts/cert.pem  ./admin/msp/admincerts/org1-admin-cert.pem
# sleep 2s

