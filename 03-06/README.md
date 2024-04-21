# ケース6: コンテナを改竄されてしまった

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

`Podに含まれるコンテナへの侵入`
```bash
kubectl exec -it sample-web -- /bin/bash

root@sample-web:/# hostname
```

`Webサイトを提供するファイルの確認`
```bash
root@sample-web:/# cat /var/www/html/index.html
```

`Webサイトのファイルの改竄`
```bash
root@sample-web:/# sed -i -e 's/https\:\/\/minikube.sigs.k8s.io\/docs\//danger.html/' /var/www/html/index.html
```

`コンテナ内でのmountコマンド実行結果`
```bash
root@sample-web:/# mount

root@sample-web:/# exit
```

`リストX.6.10. upperdirの確認`
```bash
minikube ssh -n minikube

sudo cat <upperdir>/var/www/html/index.html

exit
```

`Podのデプロイ`
```bash
kubectl apply -f read-only-rootfs.yaml
```

`ルートファイルシステムが読み取り専用であることの確認`
```bash
kubectl exec -it read-only-rootfs -- /bin/bash

root@read-only-rootfs:/# mount

root@read-only-rootfs:/# touch test

root@read-only-rootfs:/# exit
```

`Podのデプロイ（失敗）`
```bash
kubectl apply -f read-only-rootfs-httpd.yaml

kubectl get pod read-only-rootfs-httpd
```

`コンテナログの確認`
```bash
kubectl logs read-only-rootfs-httpd
```

`Podのデプロイ`
```bash
kubectl apply -f read-only-rootfs-httpd-emptydir.yaml

kubectl get pod read-only-rootfs-httpd-emptydir
```

`emptyDirをマウントしたコンテナ内でのmountコマンド実行結果`
```bash
kubectl exec -it read-only-rootfs-httpd-emptydir -- /bin/bash

root@read-only-rootfs-httpd-emptydir:/usr/local/apache2# mount

root@read-only-rootfs-httpd-emptydir:/usr/local/apache2# exit
```

`検証に使用したリソースの削除`
```bash
kubectl delete pod sample-web \
    read-only-rootfs \
    read-only-rootfs-httpd \
    read-only-rootfs-httpd-emptydir

kubectl delete service sample-web
```