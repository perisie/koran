#!/usr/bin/env bash
which bash
bash --version

rm melchior.pem

echo "$PEM_MELCHIOR" > melchior.pem

chmod 600 melchior.pem

declare -A host_to_ip
host_to_ip["melchior"]=54.151.55.32

for HOST in melchior; do
    HOST_IP=${host_to_ip[$HOST]}
    echo "> („• ֊ •„) deploying $HOST"
    echo "> deleting old files..."
    ssh -o StrictHostKeyChecking=no -i $HOST.pem ec2-user@$HOST_IP "mkdir -p ~/build; mkdir -p ~/build/koran"
    ssh -o StrictHostKeyChecking=no -i $HOST.pem ec2-user@$HOST_IP "cd ~/build/koran; rm -rf *;"
    echo "> uploading new files..."
    scp -o StrictHostKeyChecking=no -i $HOST.pem -r ./* ec2-user@$HOST_IP:~/build/koran
    echo "> starting the service..."
    ssh -o StrictHostKeyChecking=no -i $HOST.pem ec2-user@$HOST_IP "cd ~/build/koran; go build; pkill koran"
    ssh -o StrictHostKeyChecking=no -i $HOST.pem ec2-user@$HOST_IP "mkdir -p ~/app; mkdir -p ~/app/koran"
    ssh -o StrictHostKeyChecking=no -i $HOST.pem ec2-user@$HOST_IP "mv -f ~/build/koran ~/app/koran"
    ssh -o StrictHostKeyChecking=no -i $HOST.pem ec2-user@$HOST_IP "cd ~/app/koran; ./koran > ./log.txt 2>&1 &"
    echo "> (^ヮ^)/ $HOST up!"
done

echo "> (^ヮ^)/ all done!"
