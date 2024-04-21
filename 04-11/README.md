# Kubernetesクラスタに対するポリシー制御

`Namespaceの作成`
```bash
kubectl apply -f ns-psa.yaml
```

`Podのデプロイ（失敗）`
```bash
kubectl apply -f exploit-pod.yaml
```

`Podのデプロイ`
```bash
kubectl apply -f baseline-pod.yaml

kubectl get pod baseline-pod -n ns-psa
```

`Namespaceの作成`
```bash
kubectl create namespace ns-vap
```

`ValidatingAdmissionPolicyとValidatingAdmissionPolicyBindingの作成`
```bash
kubectl apply -f deny-privileged-container-policy.yaml

kubectl apply -f deny-privileged-container-policy-binding.yaml
```

`Podのデプロイ（失敗）`
```bash
kubectl apply -f privileged-pod.yaml
```

`リソースの削除`
```bash
kubectl delete pod baseline-pod -n ns-psa

kubectl delete ValidatingAdmissionPolicyBinding deny-privileged-container-policy-binding.example.com

kubectl delete ValidatingAdmissionPolicy deny-privileged-container-policy.example.com

kubectl delete namespace ns-psa ns-vap
```
