# セキュリティが強化されたコンテナランタイムの使用

`minikubeにおけるgVisorの有効化`
```bash
minikube addons enable gvisor
```

`RuntimeClassの確認`
```bash
kubectl get runtimeclass
```

`Nodeの確認`
```bash
minikube ssh -n minikube

runsc --version

cat /etc/containerd/config.toml

exit
```

`Nodeのカーネル情報の確認`
```bash
minikube ssh -n minikube

uname -r

dmesg

exit
```

`Podのデプロイ（default-pod）`
```bash
kubectl apply -f default-pod.yaml
```

`コンテナのカーネル情報の確認（default-pod）`
```bash
kubectl exec -it default-pod -- uname -r

kubectl exec -it default-pod -- dmesg
```

`Podのデプロイ（gvisor-pod）`
```bash
kubectl apply -f gvisor-pod.yaml
```

`コンテナのカーネル情報の確認（gvisor-pod）`
```bash
kubectl exec -it gvisor-pod -- uname -r

kubectl exec -it gvisor-pod -- dmesg
```

`検証に使用したリソースの削除とアドオンの無効化`
```bash
kubectl delete pod default-pod gvisor-pod

minikube addons disable gvisor
```