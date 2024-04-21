# コンテナの振る舞い監視

`Falcoのインストール`
```bash
helm repo add falcosecurity https://falcosecurity.github.io/charts

helm repo update

helm install falco falcosecurity/falco \
-n falco --create-namespace \
--version 4.3.0 \
--set tty=true \
--set driver.kind=modern_ebpf \
--set "falcoctl.config.artifact.install.refs={falco-rules:2,falco-incubating-rules:2,falco-sandbox-rules:2}" \
--set "falcoctl.config.artifact.follow.refs={falco-rules:2,falco-incubating-rules:2,falco-sandbox-rules:2}" \
--set "falco.rules_file={/etc/falco/k8s_audit_rules.yaml,/etc/falco/rules.d,/etc/falco/falco_rules.yaml,/etc/falco/falco-incubating_rules.yaml,/etc/falco/falco-sandbox_rules.yaml}"

kubectl get pods -n falco
```

`Falcoのログ表示`
```bash
kubectl logs ds/falco -c falco -n falco -f | grep exploit-pod
```

`Podのデプロイ`
```bash
kubectl apply -f exploit-pod.yaml
```

`Podに含まれるコンテナへの攻撃`
```bash
kubectl exec -it exploit-pod -- /bin/bash

root@exploit-pod:/# nsenter -t 1 -a /bin/bash
```

`リソースの削除`
```bash
helm uninstall falco -n falco

kubectl delete namespace falco

kubectl delete pod exploit-pod
```
