# ケース8: PodからKubernetesクラスタを不正に操作されてしまった

`Podのデプロイ`
```bash
kubectl apply -f sample-sa.yaml

kubectl apply -f sample-crb.yaml

kubectl apply -f sample-pod.yaml
```


`Podのに含まれるコンテナへの侵入`
```bash
kubectl exec -it sample-pod -- /bin/bash

root@sample-pod:/# hostname
```


`Podに含まれるコンテナ内でkubectlコマンドのインストール`
```bash
root@sample-pod:/# apt update

root@sample-pod:/# apt install -y curl

root@sample-pod:/# curl -LO https://dl.k8s.io/release/v1.30.0/bin/linux/amd64/kubectl

root@sample-pod:/# chmod +x kubectl

root@sample-pod:/# mv ./kubectl /usr/local/bin/kubectl
```

`Podに含まれるコンテナからKubernetesクラスタの情報を参照`
```bash
root@sample-pod:/# kubectl get nodes

root@sample-pod:/# kubectl get namespaces

root@sample-pod:/# kubectl get pods
```

`Podに含まれるコンテナからPodのデプロイ`
```bash
root@sample-pod:/# cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Pod
metadata:
  name: malicious-pod
spec:
  hostPID: true
  containers:
    - name: ubuntu
      image: ubuntu:22.04
      command: ["/bin/sh", "-c", "while :; do sleep 10; done"]
      securityContext:
        privileged: true
EOF
```

`Podの状態確認`
```bash
root@sample-pod:/# kubectl get pod malicious-pod

root@sample-pod:/# exit
```

`ClusterRoleの確認`
```bash
kubectl get clusterrole cluster-admin -o yaml
```

`Pod内のコンテナにマウントされたServiceAccountの情報`
```bash
kubectl get pod sample-pod -o yaml

kubectl exec -it sample-pod -- ls /var/run/secrets/kubernetes.io/serviceaccount
```

`curlコマンドによるKubernetesクラスタへのアクセス`
```bash
kubectl exec -it sample-pod -- /bin/bash

root@sample-pod:/# SERVICE_ACCOUNT_TOKEN=`cat /var/run/secrets/kubernetes.io/serviceaccount/token`

root@sample-pod:/# curl -H "Authorization: Bearer ${SERVICE_ACCOUNT_TOKEN}" --cacert /var/run/secrets/kubernetes.io/serviceaccount/ca.crt https://${KUBERNETES_SERVICE_HOST}/api/v1/namespaces

root@sample-pod:/# exit
```

`default ServiceAccountの確認`
```bash
kubectl get serviceaccount default
```

`default ServiceAccountを紐付けたPodのデプロイ`
```bash
kubectl apply -f default-sa-pod.yaml

kubectl get pod default-sa-pod -o yaml
```

`PodからKubernetesクラスタを操作できないことの確認`
```bash
kubectl exec -it default-sa-pod -- /bin/bash

root@default-sa-pod:/# kubectl get nodes

root@default-sa-pod:/# exit
```

`ServiceAccount情報がマウントされていないことの確認`
```bash
kubectl apply -f no-sa-token-pod.yaml

kubectl get pod no-sa-token-pod -o yaml

kubectl exec -it no-sa-token-pod -- ls /var/run/secrets/kubernetes.io/serviceaccount
```

`PodからKubernetesクラスタにアクセスできないことの確認`
```bash
kubectl exec -it no-sa-token-pod -- /bin/bash

root@no-sa-token-pod:/# kubectl get nodes

root@no-sa-token-pod:/# exit
```

`検証に使用したリソースの削除`
```bash
kubectl delete pod sample-pod \
    malicious-pod \
    default-sa-pod \
    no-sa-token-pod

kubectl delete clusterrolebinding sample-crb

kubectl delete serviceaccount sample-sa
```