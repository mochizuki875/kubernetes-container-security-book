# ケース5: コンテナからコンテナホストを操作されてしまった

`Podのデプロイ`
```bash
kubectl apply -f sample-web.yaml
```

`Podに含まれるコンテナへの侵入`
```bash
kubectl exec -it sample-web -- /bin/bash

root@sample-web:/# hostname
```

`Nodeへの侵入`
```bash
root@sample-web:/# nsenter -t 1 -a /bin/bash
```

`Nodeに対する操作`
```bash
bash-5.0# hostname

bash-5.0# pwd

bash-5.0# touch create-from-container.txt

bash-5.0# ls

bash-5.0# exit

root@sample-web:/# exit
```

`minikube sshコマンドによるNodeへの接続`
```bash
kubectl get pod sample-web -o wide

minikube ssh -n minikube
```

`Nodeに侵入できたことの確認`
```bash
hostname

ls /

exit
```

`Podのデプロイ`
```bash
kubectl apply -f hostpath.yaml
```

`コンテナにマウントしたディレクトリに対するファイル作成`
```bash
kubectl exec -it hostpath -- /bin/bash

root@hostpath:/# touch /host-etc/create-from-container-hostpath.txt

root@hostpath:/# ls -l /host-etc

root@hostpath:/# exit
```

`コンテナにマウントしたNodeのディレクトリの確認`
```bash
minikube ssh -n minikube

ls -l /etc

exit
```

`Podのデプロイ`
```bash
kubectl run ubuntu -it --image=ubuntu:22.04 --rm --restart=Never -- /bin/bash
```

`システムクロックの変更`
```bash
root@ubuntu:/# date -s "2000/01/01 00:00:00"
```

`libcap2-binのインストール`
```bash
root@ubuntu:/# apt update

root@ubuntu:/# apt install -y libcap2-bin
```

`コンテナプロセスに付与されたCapabilityの確認`
```bash
root@ubuntu:/# getpcaps $$

root@ubuntu:/# exit
```

`コンテナプロセスに付与されたCapabilityの確認`
```bash
kubectl apply -f capability-drop-and-add.yaml

kubectl exec -it capability-drop-and-add -- /bin/bash

root@capability-drop-and-add:/# apt update

root@capability-drop-and-add:/# apt install -y libcap2-bin

root@capability-drop-and-add:/# getpcaps 1

root@capability-drop-and-add:/# exit
```

`コンテナ内でのunshareコマンドの実行`
```bash
kubectl apply -f seccomp-runtime-default.yaml

kubectl exec -it seccomp-runtime-default -- /bin/bash

root@seccomp-runtime-default:/# unshare -r /bin/bash

root@seccomp-runtime-default:/# exit
```

`コンテナプロセスに適用されたSeccompプロファイルの確認`
```bash
minikube ssh -n minikube

sudo crictl inspect $(sudo crictl ps --name=seccomp-runtime-default -q)

exit
```

`Node上での/etcの権限確認`
```bash
minikube ssh -n minikube

ls -la /

exit
```

`コンテナイメージのビルド（ubuntu-user-1000:22.04）`
```bash
docker build -t <Docker ID>/ubuntu-user-1000:22.04 .
```

`コンテナイメージのアップロード（ubuntu-user-1000:22.04）`
```bash
docker push <Docker ID>/ubuntu-user-1000:22.04
```

`コンテナ内でのコンテナプロセス実行ユーザーの確認`
```bash
kubectl apply -f run-as-user-default-1000.yaml

kubectl exec -it run-as-user-default-1000 -- /bin/bash

user01@run-as-user-default-1000:~$ id

user01@run-as-user-default-1000:~$ pwd

user01@run-as-user-default-1000:~$ ps axo pid,uid,gid,comm,args

user01@run-as-user-default-1000:~$ exit
```

`Nodeでのコンテナプロセス実行ユーザーの確認`
```bash
minikube ssh -n minikube

ps axfo pid,uid,gid,comm,args

exit
```

`コンテナ内でのコンテナプロセス実行ユーザーの確認`
```bash
kubectl apply -f run-as-user-1000.yaml

kubectl exec -it run-as-user-1000 -- /bin/bash

I have no name!@run-as-user-1000:/$ id

I have no name!@run-as-user-1000:/$ ps axo pid,uid,gid,comm,args

I have no name!@run-as-user-1000:/$ exit
```

`Nodeでのコンテナプロセス実行ユーザーの確認`
```bash
minikube ssh -n minikube

ps axfo pid,uid,gid,comm,args

exit
```

`コンテナの動作確認（rootユーザーとしてコンテナを実行した場合）`
```bash
kubectl apply -f run-as-root-hello.yaml

kubectl exec -it run-as-root-hello -- tail hello.txt
```

`コンテナの動作確認（コンテナイメージに含まれない実行ユーザーを指定した場合）`
```bash
kubectl apply -f run-as-user-1000-hello-ng.yaml

kubectl exec -it run-as-user-1000-hello-ng -- tail hello.txt
```

`コンテナログの確認`
```bash
kubectl logs run-as-user-1000-hello-ng
```

`コンテナの動作確認（コンテナイメージに含まれる実行ユーザーを指定した場合）`
```bash
kubectl apply -f run-as-user-1000-hello-ok.yaml

kubectl exec -it run-as-user-1000-hello-ok -- tail hello.txt

kubectl exec -it run-as-user-1000-hello-ok -- ls -la /home/user01/hello.txt
```

`runAsNonRootによりPodの起動に失敗する例`
```bash
kubectl apply -f run-as-non-root.yaml

kubectl get pod run-as-non-root

kubectl describe pod run-as-non-root
```

`whoamiコマンドの実行`
```bash
whoami
```

`setuidの設定`
```bash
whoami

which whoami

ls -la /usr/bin/whoami

cp /usr/bin/whoami /home/ubuntu-user/whoami

chmod +s /home/ubuntu-user/whoami

ls -la /home/ubuntu-user/whoami
```

`setuidを行ったwhoamiコマンドの実行`
```bash
./whoami
```

`sudoコマンドの確認`
```bash
whoami

which sudo

ls -la /usr/bin/sudo

sudo su -

whoami
```

`コンテナイメージのビルド（ubuntu-user-1000:22.04-sudo）`
```bash
docker build -t <Docker ID>/ubuntu-user-1000:22.04-sudo .
```

`コンテナイメージのアップロード（ubuntu-user-1000:22.04-sudo）`
```bash
docker push <Docker ID>/ubuntu-user-1000:22.04-sudo
```

`コンテナ内での特権昇格の実行`
```bash
kubectl apply -f allow-privilege-escalation-true.yaml

kubectl exec -it allow-privilege-escalation-true -- /bin/bash

user01@allow-privilege-escalation-true:~$ id

user01@allow-privilege-escalation-true:~$ sudo su -

root@allow-privilege-escalation-true:~# id

root@allow-privilege-escalation-true:~# exit

user01@allow-privilege-escalation-true:~$ exit
```

`コンテナ内での特権昇格の実行（制限あり）`
```bash
kubectl apply -f allow-privilege-escalation-false.yaml

kubectl exec -it allow-privilege-escalation-false -- /bin/bash

user01@allow-privilege-escalation-false:~$ id

user01@allow-privilege-escalation-false:~$ sudo su -

user01@allow-privilege-escalation-false:~$ exit
```

`Trivyによるマニフェストスキャン（2024年4月時点）`
```bash
trivy config sample-web.yaml
```


`検証に使用したリソースの削除`
```bash
kubectl delete pod sample-web \
    hostpath \
    capability-drop-and-add \
    seccomp-runtime-default \
    run-as-user-default-1000 \
    run-as-user-1000 \
    run-as-root-hello \
    run-as-user-1000-hello-ng \
    run-as-user-1000-hello-ok \
    run-as-non-root \
    allow-privilege-escalation-true \
    allow-privilege-escalation-false
```
