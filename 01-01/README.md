# コンテナセキュリティを学ぶ前に

`コンテナイメージのビルド`
```bash
ls

docker build -t <Docker ID>/sample-web:intro .
```

`コンテナイメージの確認`
```bash
docker images <Docker ID>/sample-web
```

`DockerHubの認証`
```bash
docker login -u <Docker ID>
```

`コンテナイメージのアップロード`
```bash
docker push <Docker ID>/sample-web:intro
```

`コンテナイメージの削除`
```bash
docker rmi <Docker ID>/sample-web:intro
```

`コンテナイメージの確認`
```bash
docker images <Docker ID>/sample-web
```

`コンテナイメージの取得`
```bash
docker pull <Docker ID>/sample-web:intro
```

`コンテナイメージの確認`
```bash
docker images <Docker ID>/sample-web
```

`コンテナの起動`
```bash
docker run -d --name sample-container -p 8080:80 <Docker ID>/sample-web:intro
```

`コンテナの確認`
```bash
docker ps -f name=sample-container
```

`コンテナへのアクセス`
```bash
curl http://127.0.0.1:8080
```

`コンテナの停止と削除`
```bash
docker stop sample-container

docker rm sample-container
```

`Podのデプロイ`
```bash
kubectl run sample-pod --image=<Docker ID>/sample-web:intro
```

`Podの確認`
```bash
kubectl get pod sample-pod
```

`Nodeの確認`
```bash
kubectl get nodes

kubectl get pod sample-pod -o wide
```

`コンテナに含まれるコマンド実行`
```bash
kubectl exec -it sample-pod -- /bin/bash

root@sample-pod:/# hostname

root@sample-pod:/# exit
```

`Podの詳細情報確認①`
```bash
kubectl get pod sample-pod -o yaml
```

`Podの詳細情報確認②`
```bash
kubectl describe pod sample-pod
```

`Podの削除`
```bash
kubectl delete pod sample-pod
```

`Podのデプロイ`
```bash
kubectl apply -f sample-pod.yaml
```

`Podに割り当てられたIPアドレスの確認`
```bash
kubectl get pod sample-pod -o wide
```

`Namespaceの確認`
```bash
kubectl get namespaces
```

`Namespaceの作成`
```bash
kubectl create namespace sample-namespace

kubectl get namespaces
```

`作成したNamespaceに存在するPodの確認`
```bash
kubectl get pods -n sample-namespace
```

`作成したNamespaceに存在するPodの確認`
```bash
kubectl apply -f sample-pod-2.yaml

kubectl get pods -n sample-namespace
```

`default Namespaceに存在するPodの確認`
```bash
kubectl get pods

kubectl get pods -n default
```

`kube-system Namespaceに存在するPodの確認`
```bash
kubectl get pods -n kube-system
```

`検証に使用したリソースの削除`
```bash
kubectl delete pod sample-pod-2 -n sample-namespace

kubectl delete namespace sample-namespace
```

`Service（ClusterIP）の作成`
```bash
kubectl apply -f sample-svc-clusterip.yaml
```

`Serviceの確認`
```bash
kubectl get service sample-svc
```

`Kubernetesクラスタ内でのServiceを介したアクセス確認`
```bash
kubectl run -it --rm curl --image=curlimages/curl -- /bin/sh

~ $ curl http://sample-svc:80

~ $ exit
```

`Service（LoadBalancer）の作成`
```bash
kubectl apply -f sample-svc-loadbalancer.yaml
```

`Serviceの確認`
```bash
kubectl get service sample-svc
```

`minikube tunnelコマンドの実行`
```bash
minikube tunnel
```

`Serviceの確認`
```bash
kubectl get service sample-svc
```

`Kubernetesクラスタ外からのServiceを介したアクセス確認`
```bash
curl http://<EXTERNAL-IP>:80
```

`検証に使用したリソースの削除`
```bash
kubectl delete pod sample-pod

kubectl delete service sample-svc
```

`Podのデプロイ`
```bash
kubectl run hello --image=busybox:1.28 -- /bin/sh -c "while :; do echo hello; sleep 5; done"

kubectl get pod hello
```

`コンテナログの確認`
```bash
kubectl logs hello
```

`コンテナ内で実行されるプロセスの確認`
```bash
kubectl exec -it hello -- ps aux
```

`コンテナホストで実行されているプロセスの確認`
```bash
kubectl get pod hello -o wide

minikube ssh -n minikube

ps auxf

exit
```

`コンテナホストのinitプロセスのNamespaceの確認`
```bash
minikube ssh -n minikube

sudo ls -la /proc/1/ns

exit
```

`Podに含まれるコンテナプロセスのNamespaceの確認`
```bash
kubectl exec -it hello -- ls -la /proc/1/ns
```

`検証に使用したリソースの削除`
```bash
kubectl delete pod hello
```