# vagrant 설치

```bash
$ vagrant init
$ vim Vagrantfile

# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
# kube-control1 VM
 config.vm.define "kube-control1" do |config|
  config.vm.box = "ubuntu/focal64"
  config.vm.provider "virtualbox" do |vb|
    vb.name = "kube-control1"
    vb.cpus = 2
    vb.memory = 4096
  end
  config.vm.hostname = "kube-control1"
  config.vm.network "public_network", ip: "172.30.1.11"
 end

 config.vm.define "kube-control2" do |config|
  config.vm.box = "ubuntu/focal64"
  config.vm.provider "virtualbox" do |vb|
    vb.name = "kube-control2"
    vb.cpus = 2
    vb.memory = 4096
  end
  config.vm.hostname = "kube-control2"
  config.vm.network "public_network", ip: "172.30.1.12"
 end

 config.vm.define "kube-control3" do |config|
  config.vm.box = "ubuntu/focal64"
  config.vm.provider "virtualbox" do |vb|
    vb.name = "kube-control3"
    vb.cpus = 2
    vb.memory = 4096
  end
  config.vm.hostname = "kube-control3"
  config.vm.network "public_network", ip: "172.30.1.13"
 end

# kube-node1 VM
 config.vm.define "kube-worker1" do |config|
  config.vm.box = "ubuntu/focal64"
  config.vm.provider "virtualbox" do |vb|
    vb.name = "kube-worker1"
    vb.cpus = 4
    vb.memory = 8192
  end
  config.vm.hostname = "kube-worker1"
  config.vm.network "public_network", ip: "172.30.1.21"
 end

 config.vm.define "kube-worker2" do |config|
  config.vm.box = "ubuntu/focal64"
  config.vm.provider "virtualbox" do |vb|
    vb.name = "kube-worker2"
    vb.cpus = 4
    vb.memory = 8192
  end
  config.vm.hostname = "kube-worker2"
  config.vm.network "public_network", ip: "172.30.1.22"
 end

# Hostmanager Plugin
config.hostmanager.enabled = true
config.hostmanager.manage_guest = true

# Provision
 config.vm.provision "shell", inline: <<-SHELL
  sed -i 's/ChallengeResponseAuthentication no/ChallengeResponseAuthentication yes/g' /etc/ssh/sshd_config
  sed -i 's/archive.ubuntu.com/ftp.daum.net/g' /etc/apt/sources.list
  sed -i 's/security.ubuntu.com/ftp.daum.net/g' /etc/apt/sources.list
  systemctl restart ssh
  apt install -y net-tools
  route add default gw 172.30.1.254
  route delete default gw 10.0.2.2
  apt update
  apt install -y chrony ntp
  service ntp restart
  ntpq -p
  swapoff -a
  sudo apt-get update && sudo apt-get install -y apt-transport-https ca-certificates curl software-properties-common gnupg2
  curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key --keyring /etc/apt/trusted.gpg.d/docker.gpg add -
  sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
  sudo apt-get update && sudo apt-get install -y containerd.io=1.2.13-2 docker-ce=5:19.03.11~3-0~ubuntu-$(lsb_release -cs) docker-ce-cli=5:19.03.11~3-0~ubuntu-$(lsb_release -cs)
  setenforce 0
  mkdir -p /etc/docker
  cat <<EOF | sudo tee /etc/docker/daemon.json
{
  "exec-opts": ["native.cgroupdriver=systemd"],
  "log-driver": "json-file",
  "log-opts": {
    "max-size": "100m"
  },
  "storage-driver": "overlay2"
}
EOF
  sudo curl -fsSLo /usr/share/keyrings/kubernetes-archive-keyring.gpg https://packages.cloud.google.com/apt/doc/apt-key.gpg
  echo "deb [signed-by=/usr/share/keyrings/kubernetes-archive-keyring.gpg] https://apt.kubernetes.io/ kubernetes-xenial main" | sudo tee /etc/apt/sources.list.d/kubernetes.list
  sudo apt-get update
  sudo apt-get install -y kubelet=1.22.9-00 kubeadm=1.22.9-00 kubectl=1.22.9-00
  sudo systemctl daemon-reload
  sudo systemctl restart docker
  sudo systemctl restart kubelet
 SHELL
end

$ vagrant up

$ vagrant ssh control1
```

## HA 구성을 위한 로드밸런서 설치

```bash
$ sudo apt-get install -y haproxy
$ cat <<EOF >> /etc/haproxy/haproxy.cfg
frontend kubernetes-master-lb
bind 0.0.0.0:26443
option tcplog
mode tcp
default_backend kubernetes-master-nodes

backend kubernetes-master-nodes
mode tcp
balance roundrobin
option tcp-check
option tcplog
server master1 172.30.1.11:6443 check
server master2 172.30.1.12:6443 check
server master3 172.30.1.13:6443 check
EOF

systemctl restart haproxy && systemctl enable haproxy


```

## k8s Clustring

- CNI 의 경우 calico 사용

```bash
$ sudo kubeadm init --control-plane-endpoint "172.30.1.11:26443" --upload-certs --pod-network-cidr "192.168.0.0/16"

$ mkdir -p $HOME/.kube
$ sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
$ sudo chown $(id -u):$(id -g) $HOME/.kube/config
```

- master join 을 위한 kubeadm 명령어를 control2, control3 에도 등록

```bash
$ curl https://projectcalico.docs.tigera.io/manifests/calico.yaml -O

$ kubectl apply -f calico.yaml
```

- Calico CNI 적용

```bash
vagrant@kube-control1:~$ kubectl get pods -n kube-system
NAME                                       READY   STATUS    RESTARTS      AGE
calico-kube-controllers-5b5b79df9b-r2mqr   1/1     Running   0             21m
calico-node-29287                          1/1     Running   0             21m
calico-node-gg98w                          1/1     Running   0             21m
calico-node-pbkjs                          1/1     Running   0             12m
calico-node-qkb8p                          1/1     Running   0             21m
calico-node-xqgp6                          1/1     Running   0             11m
coredns-78fcd69978-75crp                   1/1     Running   0             31m
coredns-78fcd69978-hzm4x                   1/1     Running   0             31m
etcd-kube-control1                         1/1     Running   1             31m
etcd-kube-control2                         1/1     Running   0             27m
etcd-kube-control3                         1/1     Running   0             25m
kube-apiserver-kube-control1               1/1     Running   1             31m
kube-apiserver-kube-control2               1/1     Running   0             27m
kube-apiserver-kube-control3               1/1     Running   0             25m
kube-controller-manager-kube-control1      1/1     Running   2 (27m ago)   31m
kube-controller-manager-kube-control2      1/1     Running   0             27m
kube-controller-manager-kube-control3      1/1     Running   0             25m
kube-proxy-9zblp                           1/1     Running   0             25m
kube-proxy-klxh9                           1/1     Running   0             31m
kube-proxy-tfwvn                           1/1     Running   0             27m
kube-proxy-w9lp9                           1/1     Running   0             11m
kube-proxy-zmv44                           1/1     Running   0             12m
kube-scheduler-kube-control1               1/1     Running   2 (27m ago)   31m
kube-scheduler-kube-control2               1/1     Running   0             27m
kube-scheduler-kube-control3               1/1     Running   0             25m
```

- Calico CNI 적용 후 pod 상태 확인

```bash
vagrant@kube-control1:~$ kubectl get nodes
NAME            STATUS   ROLES                  AGE     VERSION
kube-control1   Ready    control-plane,master   29m     v1.22.9
kube-control2   Ready    control-plane,master   26m     v1.22.9
kube-control3   Ready    control-plane,master   24m     v1.22.9
kube-worker1    Ready    <none>                 11m     v1.22.9
kube-worker2    Ready    <none>                 9m44s   v1.22.9
```

- Calico CNI 적용 후 node 상태 확인

# Reference

[1] [Kubernetes] kubeadm Master 3개 설정 (Haproxy), https://velog.io/@pingping95/Kubernetes-kubeadm-Master-%EA%B3%A0%EA%B0%80%EC%9A%A9%EC%84%B1
[2] Install Calico networking and network policy for on-premises deployments, https://projectcalico.docs.tigera.io/getting-started/kubernetes/self-managed-onprem/onpremises
