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
rm -r peers
sleep 1s



echo "                                  "
echo "##################################"
echo "###  running steupOrg2CA.sh    ###"
echo "##################################"
echo "                                  "

./setupOrg2CA.sh
sleep 20s

echo "                                  "
echo "##################################"
echo "###  running steup-peer1.sh    ###"
echo "##################################"
echo "                                  "

./setup-peer1.sh
sleep 10s
echo "                                                 "
echo "#################################################"
echo "###          renaming                        ####"
echo "### ./peers/peer1/tls-msp/keystore           ####"
echo "### ./peers/peer1/tls-msp/keystore/key.pem   ####"
echo "#################################################"
echo "                                                 "

for peer1key in "./peers/peer1/tls-msp/keystore"/*
do
  echo "$peer1key" 
  mv   $peer1key  ./peers/peer1/tls-msp/keystore/key.pem
done

sleep 2s 

echo "                                  "
echo "##################################"
echo "###  running steup-peer2.sh    ###"
echo "##################################"
echo "                                  "

./setup-peer2.sh
sleep 10s
echo "                                                 "
echo "#################################################"
echo "###          renaming                        ####"
echo "### ./peers/peer2/tls-msp/keystore           ####"
echo "### ./peers/peer2/tls-msp/keystore/key.pem   ####"
echo "#################################################"
echo "                                                 "

for peer2key in "./peers/peer2/tls-msp/keystore"/*
do
  echo "$peer2key" 
  mv   $peer2key  ./peers/peer2/tls-msp/keystore/key.pem
done

sleep 2s 


echo "                                  "
echo "##################################"
echo "###  running setup-admin.sh    ###"
echo "##################################"
echo "                                  "
./setup-admin.sh
sleep 10s

echo "                                  "
echo "##################################"
echo "###  running setup-org1-msp.sh ###"
echo "##################################"
echo "                                  "
./setup-org2-msp.sh
sleep 5s


echo "                                  "
echo "##################################"
echo "###  copying config.yaml to    ###"
echo "##################################"
echo "                                  "

cp config.yaml msp/config.yaml
sleep2s

echo "                                  "
echo "##################################"
echo "###  running setup-user.sh     ###"
echo "##################################"
echo "                                  "
./setup-user.sh