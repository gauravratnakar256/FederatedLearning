
mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config

kubectl apply -f  https://projectcalico.docs.tigera.io/manifests/calico.yaml

kubectl taint nodes --all node-role.kubernetes.io/control-plane-

