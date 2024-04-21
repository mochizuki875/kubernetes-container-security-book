# 検証環境の準備

`CPUの仮想化支援機能が有効になっているかの確認`
```bash
egrep -c '(vmx|svm)' /proc/cpuinfo
```

`Google Cloudにおける仮想マシンの作成`
```bash
gcloud compute instances create ubuntu-vm \
  --enable-nested-virtualization \
  --zone=us-east1-b \
  --image-family="ubuntu-2204-lts" \
  --image-project="ubuntu-os-cloud" \
  --machine-type=n1-standard-8 \
  --min-cpu-platform="Intel Haswell" \
  --boot-disk-size=50GB
```

`Dockerのインストール`
```bash
sudo apt-get update

sudo apt-get install -y ca-certificates curl

sudo install -m 0755 -d /etc/apt/keyrings

sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc

sudo chmod a+r /etc/apt/keyrings/docker.asc

echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \
  $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

sudo apt-get update

VERSION_STRING=5:26.0.1-1~ubuntu.22.04~jammy

sudo apt-get install -y docker-ce=$VERSION_STRING docker-ce-cli=$VERSION_STRING containerd.io docker-buildx-plugin docker-compose-plugin

docker --version
```


`一般ユーザーに対するdockerコマンド実行権限の付与`
```bash
sudo usermod -aG docker $USER
```


`minikubeのインストール`
```bash
curl -LO https://storage.googleapis.com/minikube/releases/v1.33.0/minikube-linux-amd64

sudo install minikube-linux-amd64 /usr/local/bin/minikube

minikube version
```


`KVMのインストール`
```bash
sudo apt-get install -y qemu-kvm libvirt-daemon-system libvirt-clients bridge-utils

sudo adduser `id -un` libvirt

sudo adduser `id -un` kvm
```

`minikubeによるKubernetesクラスタの構築`
```bash
minikube start --driver=kvm2 --cpus=5 --memory=4g \
   --cni=calico --container-runtime=containerd \
   --kubernetes-version=v1.30.0
```


`Kubernetesクラスタの状態確認`
```bash
minikube status
```

`kubectlコマンドのインストール`
```bash
curl -LO https://dl.k8s.io/release/v1.30.0/bin/linux/amd64/kubectl

chmod +x ./kubectl

sudo mv ./kubectl /usr/local/bin/

kubectl version
```

`helmコマンドのインストール`
```bash
curl -LO https://get.helm.sh/helm-v3.14.4-linux-amd64.tar.gz

tar -zxvf helm-v3.14.4-linux-amd64.tar.gz

sudo mv linux-amd64/helm /usr/local/bin/helm

helm version
```
