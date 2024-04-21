# ケース7: コンテナホストのリソースを過剰に使用されてしまった

`コンテナイメージのビルド（sample-web:case7）`
```bash
cd sample-web

ls

docker build -t <Docker ID>/sample-web:case7 .
```

`コンテナイメージのアップロード（sample-web:case7）`
```bash
docker push <Docker ID>/sample-web:case7
```

`PodとServiceのデプロイ`
```bash
kubectl apply -f sample-web-a.yaml

kubectl apply -f sample-web-b.yaml
```

`Podの確認`
```bash
kubectl get pods
```

`minikube tunnelコマンドによるServiceの公開`
```bash
minikube tunnel
```

`Serviceの確認`
```bash
kubectl get services
```

`sample-web-aに対するリクエスト送信`
```bash
curl <EXTERNAL-IP-A> -w "time_total: %{time_total}\n"
```

`sample-web-bに対するリクエスト送信`
```bash
curl <EXTERNAL-IP-B> -w "time_total: %{time_total}\n"
```

`Podの状態確認`
```bash
kubectl get pods -o wide
```

`sample-web-aに対する大量リクエストの送信`
```bash
for i in {1..10};\
do \
  while true; do curl -s <EXTERNAL-IP-A> > /dev/null; done & \
done
```

`sample-web-aに対するリクエスト送信`
```bash
curl <EXTERNAL-IP-A> -w "time_total: %{time_total}\n"
```

`sample-web-bに対するリクエスト送信`
```bash
curl <EXTERNAL-IP-B> -w "time_total: %{time_total}\n"
```

`sample-web-aに対する大量リクエスト送信の停止`
```bash
for i in {1..10}; do kill %$i; done
```

`Nodeのリソース使用状況`
```bash
minikube ssh -n minikube

top

exit
```

`Podの再デプロイ`
```bash
kubectl delete pod sample-web-a sample-web-b

kubectl apply -f sample-web-a-cap.yaml

kubectl apply -f sample-web-b-cap.yaml
```

`sample-web-aに対するリクエスト送信`
```bash
curl <EXTERNAL-IP-A> -w "time_total: %{time_total}\n"
```

`sample-web-bに対するリクエスト送信`
```bash
curl <EXTERNAL-IP-B> -w "time_total: %{time_total}\n"
```

`sample-web-aに対する大量リクエストの送信`
```bash
for i in {1..10};\
do \
  while true; do curl -s <EXTERNAL-IP-A> > /dev/null; done & \
done
```

`sample-web-aに対するリクエスト送信`
```bash
curl <EXTERNAL-IP-A> -w "time_total: %{time_total}\n"
```

`Nodeのリソース使用状況`
```bash
minikube ssh -n minikube

top

exit
```

`sample-web-bに対するリクエスト送信`
```bash
curl <EXTERNAL-IP-B> -w "time_total: %{time_total}\n"
```

`sample-web-aに対する大量リクエスト送信の停止`
```bash
for i in {1..10}; do kill %$i; done
```

`PodとServiceの削除`
```bash
kubectl delete pod sample-web-a sample-web-b

kubectl delete service sample-web-a sample-web-b
```

`Namespaceの作成`
```bash
kubectl create namespace limit-range-test
```

`LimitRangeの作成`
```bash
kubectl apply -f sample-limit-range.yaml
```

`Podのデプロイ`
```bash
kubectl apply -f sample-pod.yaml
```

`Podの詳細情報確認`
```bash
kubectl describe pod sample-pod -n limit-range-test
```

`Kubernetesクラスタを構成するNodeの確認`
```bash
kubectl get nodes
```

`taintの設定`
```bash
kubectl taint nodes k8s-cluster-node02 dedicated=true:NoSchedule
```

`Podのデプロイとデプロイ先Nodeの確認`
```bash
kubectl run pod-1 --image=nginx:1.25.5 --restart=Never

kubectl run pod-2 --image=nginx:1.25.5 --restart=Never

kubectl run pod-3 --image=nginx:1.25.5 --restart=Never

kubectl get pods -o wide
```

`Nodeに対するラベルの付与`
```bash
kubectl label node k8s-cluster-node02 node-restriction.kubernetes.io/dedicated=true
```

`占有NodeへのPodのデプロイ`
```bash
kubectl apply -f pod-dedicated.yaml

kubectl get pods -o wide
```

`検証に使用したリソースの削除`
```bash
kubectl delete pod sample-pod -n limit-range-test

kubectl delete limitrange sample-limit-range -n limit-range-test

kubectl delete namespace limit-range-test
```