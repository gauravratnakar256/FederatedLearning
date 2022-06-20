apt-get install ebtables ethtool
apt-get -y update
apt-get install -y docker.io
apt-get -y update 
apt-get install -y apt-transport-https
apt-get install -y curl 
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add -
apt-add-repository "deb http://apt.kubernetes.io/ kubernetes-xenial main"
apt-get -y update
apt-get install -y kubelet kubeadm kubectl
sed -i '/^ExecStart/ s/$/ --exec-opt native.cgroupdriver=systemd/' /lib/systemd/system/docker.service
systemctl daemon-reload
systemctl restart docker

sleep 10s
#kubeadm init --pod-network-cidr=192.168.0.0/16
