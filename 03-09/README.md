# ケース9: コンテナの秘密情報が流出してしまった

`Secretのマニフェストの作成`
```bash
kubectl create secret generic sample-secret \
    --from-literal=username=user01 \
    --from-literal=password=password01 \
    --dry-run=client -o yaml > sample-secret.yaml
```

`Secretの作成`
```bash
kubectl apply -f sample-secret.yaml
```

`Secretの確認`
```bash
kubectl get secret sample-secret -o yaml
```

`Podのデプロイ`
```bash
kubectl apply -f sample-secret-pod.yaml
```

`コンテナの環境変数に秘密情報が設定されたことの確認`
```bash
kubectl exec -it sample-secret-pod -- env | grep USERNAME

kubectl exec -it sample-secret-pod -- env | grep PASSWORD
```

`秘密情報の復元`
```bash
echo "dXNlcjAx" | base64 --decode

echo "cGFzc3dvcmQwMQ==" | base64 --decode
```

`Base64による文字列変換と復元`
```bash
echo "abcdefg" | base64

echo "YWJjZGVmZwo=" | base64 --decode
```

`Sealed Secrets Controllerのインストール`
```bash
helm repo add sealed-secrets https://bitnami-labs.github.io/sealed-secrets

helm install sealed-secrets \
  sealed-secrets/sealed-secrets \
  -n kube-system \
  --set-string fullnameOverride=sealed-secrets-controller \
  --version=2.15.3

kubectl get pods -n kube-system -l app.kubernetes.io/name=sealed-secrets
```

`kubesealコマンドのインストール`
```bash
KUBESEAL_VERSION='0.26.2'

wget "https://github.com/bitnami-labs/sealed-secrets/releases/download/v${KUBESEAL_VERSION:?}/kubeseal-${KUBESEAL_VERSION:?}-linux-amd64.tar.gz"

tar -xvzf kubeseal-${KUBESEAL_VERSION:?}-linux-amd64.tar.gz kubeseal

sudo install -m 755 kubeseal /usr/local/bin/kubeseal
```

`SealedSecretマニフェストの作成`
```bash
kubeseal < sample-secret.yaml --format yaml --name sample-sealed-secret > sample-sealed-secret.yaml
```

`SealedSecretの作成`
```bash
kubectl apply -f sample-sealed-secret.yaml
```

`SealedSecretからSecretが自動作成されたことの確認`
```bash
kubectl get sealedsecret sample-sealed-secret

kubectl get secret sample-sealed-secret

kubectl get secret sample-sealed-secret -o yaml
```

`秘密情報の復元`
```bash
echo "dXNlcjAx" | base64 --decode

echo "cGFzc3dvcmQwMQ==" | base64 --decode
```

`Vaultのインストール`
```bash
helm repo add hashicorp https://helm.releases.hashicorp.com

helm install vault \
  hashicorp/vault \
  --set "server.dev.enabled=true" \
  -n vault-dev \
  --create-namespace \
  --version=0.27.0

kubectl get pods -n vault-dev
```

`Secrets Engineの確認`
```bash
kubectl exec -it vault-0 -n vault-dev -- /bin/sh

/ $ vault secrets list
```

`Vaultへの秘密情報の登録`
```bash
/ $ vault kv put secret/config username="user01" password="password01"
```

`Vaultに登録した秘密情報の確認`
```bash
/ $ vault kv get secret/config

/ $ exit
```

`ServiceAccountの作成`
```bash
$ kubectl create serviceaccount vault-kv-secret-sa
```

`Vaultの認証設定`
```bash
$ kubectl exec -it vault-0 -n vault-dev -- /bin/sh

/ $ vault auth enable kubernetes

/ $ vault write auth/kubernetes/config \
      kubernetes_host="https://$KUBERNETES_PORT_443_TCP_ADDR:443"
```

`Vaultの認可設定`
```bash
/ $ vault policy write app - <<EOF
path "secret/data/config" {
   capabilities = ["read"]
}
EOF

/ $ vault write auth/kubernetes/role/app \
      bound_service_account_names=vault-kv-secret-sa \
      bound_service_account_namespaces=default \
      policies=app \
      ttl=24h

/ $ exit
```

`External Secrets Operatorのインストール`
```bash
helm repo add external-secrets https://charts.external-secrets.io

helm repo update

helm install external-secrets \
  external-secrets/external-secrets \
  -n external-secrets \
  --create-namespace \
  --version=0.9.16

kubectl get pods -n external-secrets
```

`SecretStoreの作成`
```bash
kubectl apply -f secret-store-vault-k8s.yaml
```

`ExternalSecretの作成`
```bash
kubectl apply -f sample-external-secret.yaml
```

`ExternalSecretからSecretが自動作成されたことの確認`
```bash
kubectl get externalsecret sample-external-secret

kubectl get secret sample-secret-from-external-secret

kubectl get secret sample-secret-from-external-secret -o yaml
```

`秘密情報の復元`
```bash
echo "dXNlcjAx" | base64 --decode

echo "cGFzc3dvcmQwMQ==" | base64 --decode
```

`検証に使用したリソースの削除`
```bash
kubectl delete secret sample-secret

kubectl delete pod sample-secret-pod

kubectl delete sealedsecret sample-sealed-secret

helm uninstall sealed-secrets -n kube-system

kubectl delete secretstore vault-backend

kubectl delete externalsecret sample-external-secret

helm uninstall vault -n vault-dev

helm uninstall external-secrets -n external-secrets

kubectl delete serviceaccount vault-kv-secret-sa

kubectl delete namespace vault-dev

kubectl delete namespace external-secrets
```