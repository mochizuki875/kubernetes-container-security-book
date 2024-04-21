# ケース10: Podに対して不正な通信が行われてしまった

`Namespaceの作成`
```bash
kubectl create namespace sample-ns
```

`PodとServiceのデプロイ`
```bash
kubectl apply -f wordpress-mysql.yaml

kubectl apply -f wordpress.yaml
```

`Podのデプロイ`
```bash
kubectl apply -f vulnerable-pod.yaml
```

`Podの状態確認`
```bash
kubectl get pods -n sample-ns
```

`minikube tunnelコマンドによるServiceの公開`
```bash
minikube tunnel
```

`Serviceの状態確認`
```bash
kubectl get service -n sample-ns
```

`Podに含まれるコンテナへの侵入`
```bash
kubectl exec -it vulnerable-pod -n sample-ns -- /bin/bash

root@vulnerable-pod:/# hostname
```

`mysql-clientのインストール`
```bash
root@vulnerable-pod:/# apt update

root@vulnerable-pod:/# apt install -y mysql-client
```

`Vulnerable PodからMySQL Podへのアクセス`
```bash
root@vulnerable-pod:/# mysql -h wordpress-mysql -u wordpress -ppassword
```

`MySQL PodからWordPressの情報を奪取`
```bash
mysql> show databases;

mysql> use wordpress;

mysql> show tables;

mysql> SELECT user_nicename,user_email FROM wp_users;

mysql> exit

root@vulnerable-pod:/# exit
```

`全ての通信を禁止するNetworkPolicyの適用`
```bash
kubectl apply -f netpol-deny-all-ingress-and-egress.yaml
```

`各Podに対するNetworkPolicyの適用`
```bash
kubectl apply -f wordpress-netpol.yaml

kubectl apply -f wordpress-mysql-netpol.yaml

kubectl apply -f vulnerable-netpol.yaml
```

`NetworkPolicyの確認`
```bash
kubectl get networkpolicy -n sample-ns
```

`deny-all-ingress-and-egressの確認`
```bash
kubectl describe networkpolicy deny-all-ingress-and-egress -n sample-ns
```

`wordpress-netpolの確認`
```bash
kubectl describe networkpolicy wordpress-netpol -n sample-ns
```

`NetworkPolicyによりVulnerable PodからMySQL Podへの通信が禁止されていることの確認`
```bash
kubectl exec -it vulnerable-pod -n sample-ns -- /bin/bash

root@vulnerable-pod:/# mysql -h wordpress-mysql -u wordpress -ppassword

root@vulnerable-pod:/# exit
```

`NetworkPolicyにより新規作成したPodから外部への通信が禁止されていることの確認`
```bash
kubectl run -it --rm sample-pod --image=curlimages/curl -n sample-ns -- /bin/sh

~ $ curl -m 5 http://wordpress:80

~ $ curl -m 5 https://kubernetes.io

~ $ exit
```


`検証に使用したリソースの削除`
```bash
kubectl delete pod wordpress \
    wordpress-mysql \
    vulnerable-pod \
    -n sample-ns

kubectl delete service wordpress \
    wordpress-mysql \
    -n sample-ns

kubectl delete networkpolicy deny-all-ingress-and-egress \
    wordpress-netpol \
    wordpress-mysql-netpol \
    vulnerable-netpol \
    -n sample-ns

kubectl delete namespace sample-ns
```