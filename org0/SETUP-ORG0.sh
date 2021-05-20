echo "                                  "
echo "##################################"
echo "###  cleaning old certificaes    ###"
echo "##################################"
echo "                                  "

rm -r admin
sleep 1s
rm -r ca
sleep 1s
rm -r msp
sleep 1s
rm -r orderers
sleep 1s

echo "                                  "
echo "##################################"
echo "###  running steupOrg0CA.sh    ###"
echo "##################################"
echo "                                  "

./setupOrg0CA.sh
sleep 20s




echo "                                  "
echo "##################################"
echo "###  running steup-orderer1.sh ###"
echo "##################################"
echo "                                  "

./setup-orderer1.sh
sleep 10s
echo "                                                 "
echo "#################################################"
echo "###          renaming                        ####"
echo "### ./orderers/orderer1/tls-msp/keystore           ####"
echo "### ./orderers/orderer1/tls-msp/keystore/key.pem   ####"
echo "#################################################"
echo "                                                 "

for orederer1key in "./orderers/orderer1/tls-msp/keystore"/*
do
  echo "$orederer1key" 
  mv   $orederer1key  ./orderers/orderer1/tls-msp/keystore/key.pem
done

sleep 2s 

echo "                                  "
echo "##################################"
echo "###  running steup-orderer2.sh  ##"
echo "##################################"
echo "                                  "

./setup-orderer2.sh
sleep 10s
echo "                                                 "
echo "#################################################"
echo "###          renaming                        ####"
echo "### ./orderers/orderer2/tls-msp/keystore           ####"
echo "### ./orderers/orderer2/tls-msp/keystore/key.pem   ####"
echo "#################################################"
echo "                                                 "

for orderer2key in "./orderers/orderer2/tls-msp/keystore"/*
do
  echo "$orderer2key" 
  mv   $orderer2key  ./orderers/orderer2/tls-msp/keystore/key.pem
done

sleep 2s 


echo "                                  "
echo "##################################"
echo "###  running steup-orderer3.sh  ##"
echo "##################################"
echo "                                  "

./setup-orderer3.sh
sleep 10s
echo "                                                       "
echo "#######################################################"
echo "###          renaming                              ####"
echo "### ./orderers/orderer3/tls-msp/keystore           ####"
echo "### ./orderers/orderer3/tls-msp/keystore/key.pem   ####"
echo "#######################################################"
echo "                                                       "

for orderer3key in "./orderers/orderer3/tls-msp/keystore"/*
do
  echo "$orderer3key" 
  mv   $orderer3key  ./orderers/orderer3/tls-msp/keystore/key.pem
done

sleep 2s 


echo "                                  "
echo "##################################"
echo "###  running setup-admin.sh    ###"
echo "##################################"
echo "                                  "
./setup-admin.sh
sleep 5s

echo "                                  "
echo "##################################"
echo "###  running setup-org0-msp.sh ###"
echo "##################################"
echo "                                  "
./setup-org0-msp.sh
sleep 5s


echo "                                  "
echo "##################################"
echo "###  copying config.yaml to    ###"
echo "##################################"
echo "                                  "

cp org0-config.yaml msp/config.yaml

cp orderer-config.yaml orderers/orderer1/msp/config.yaml
cp orderer-config.yaml orderers/orderer2/msp/config.yaml
cp orderer-config.yaml orderers/orderer3/msp/config.yaml
