swapoff -a
sed -i '/swap/ s/^/#/' /etc/fstab
echo 'net/bridge/bridge-nf-call-ip6tables = 1' >> /etc/ufw/sysctl.conf
echo 'net/bridge/bridge-nf-call-iptables = 1' >> /etc/ufw/sysctl.conf
echo 'net/bridge/bridge-nf-call-arptables = 1' >> /etc/ufw/sysctl.conf

#reboot
#sudo su
#apt-get install ebtables ethtool