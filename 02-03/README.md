# ケース3: 改竄されたコンテナイメージを使用してしまった

`コンテナイメージのビルド（sample-web:case3）`
```bash
ls

docker build -t <Docker ID>/sample-web:case3 .
```

`コンテナイメージのアップロード（sample-web:case3）`
```bash
docker push <Docker ID>/sample-web:case3
```

`PodとServiceのデプロイ`
```bash
kubectl apply -f sample-web.yaml
```

`minikube tunnelコマンドの実行`
```bash
minikube tunnel
```

`起動状態の確認`
```bash
kubectl get all -l app=sample-web
```

`改竄されたコンテナイメージのビルド（sample-web:case3）`
```bash
ls

docker build -t <Docker ID>/sample-web:case3 .
```

`改竄されたコンテナイメージのアップロード（sample-web:case3）`
```bash
docker push <Docker ID>/sample-web:case3
```

`Podの再デプロイ`
```bash
kubectl delete pod sample-web

kubectl apply -f sample-web.yaml
```

`コンテナイメージのビルド（sample-web:case3）`
```bash
$ ls

$ docker build -t <Docker ID>/sample-web:case3 .
```

`コンテナイメージのアップロード（sample-web:case3）`
```bash
docker push <Docker ID>/sample-web:case3
```

`ダイジェスト値を指定したPodのデプロイ`
```bash
kubectl delete pod sample-web

kubectl apply -f sample-web-digest.yaml

kubectl get pod sample-web
```

`ダイジェスト値を指定したコンテナイメージの取得が行われたことの確認`
```bash
kubectl describe pod sample-web
```

`Cosignのインストール`
```bash
curl -O -L "https://github.com/sigstore/cosign/releases/download/v2.2.3/cosign-linux-amd64"

sudo mv cosign-linux-amd64 /usr/local/bin/cosign

sudo chmod +x /usr/local/bin/cosign
```

`キーペアの作成`
```bash
cosign generate-key-pair

ls
```

`コンテナイメージへの署名`
```bash
cosign sign --key cosign.key <Docker ID>/sample-web:@<ダイジェスト値>
```

`コンテナイメージの署名検証`
```bash
cosign verify --key cosign.pub <Docker ID>/sample-web:case3 | jq .
```

`Policy Controllerのインストール`
```bash
helm repo add sigstore https://sigstore.github.io/helm-charts

helm repo update

helm install policy-controller sigstore/policy-controller \
--create-namespace -n cosign-system \
--version 0.6.7 \
--devel

kubectl get pods -n cosign-system
```

`CosignのキーペアからSecretを作成`
```bash
ls

kubectl create secret generic mysecret --from-file=cosign.pub=./cosign.pub -n cosign-system
```

`ClusterImagePolicyの作成`
```bash
kubectl apply -f cip-key-secret.yaml
```

`Namespaceの作成とラベルの付与`
```bash
kubectl create namespace image-verify

kubectl label namespace image-verify policy.sigstore.dev/include=true
```

`署名を付与したコンテナイメージを使用してPodをデプロイしようとした場合`
```bash
kubectl run signed --image=<Docker ID>/sample-web:case3 --restart=Never -n image-verify

kubectl get pod signed -n image-verify
```

`署名が付与されていないコンテナイメージのアップロード`
```bash
docker pull nginx:latest

docker tag nginx:latest <Docker ID>/sample-web:case3-unsigned

docker push <Docker ID>/sample-web:case3-unsigned
```

`署名を付与していないコンテナイメージを使用してPodをデプロイしようとした場合`
```bash
kubectl run unsigned --image=<Docker ID>/sample-web:case3-unsigned --restart=Never -n image-verify

kubectl get pod unsigned -n image-verify
```

`検証に使用したリソースの削除`
```bash
kubectl delete pod sample-web

kubectl delete service sample-web

kubectl delete pod signed -n image-verify

kubectl delete namespace image-verify

kubectl delete ClusterImagePolicy cip-key-secret

helm uninstall policy-controller -n cosign-system

kubectl delete secret mysecret -n cosign-system

kubectl delete namespace cosign-system
```