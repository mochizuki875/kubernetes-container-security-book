# ケース2: コンテナイメージが流出してしまった

`コンテナイメージのビルド（sample-web:case2）`
```bash
ls

docker build -t <Docker ID>/sample-web:case2 .
```

`コンテナイメージのアップロード（sample-web:case2）`
```bash
docker push <Docker ID>/sample-web:case2
```

`Podのデプロイ`
```bash
kubectl apply -f sample-web.yaml

kubectl get pod sample-web
```

`DockerHubからのログアウト`
```bash
docker logout
```

`ローカル環境に存在するコンテナイメージの削除`
```bash
docker rmi <Docker ID>/sample-web:case2

docker images <Docker ID>/sample-web:case2
```

`一般ユーザーとしてコンテナイメージを取得`
```bash
docker pull <Docker ID>/sample-web:case2

docker images <Docker ID>/sample-web:case2
```

`コンテナイメージに含まれるファイルの取得`
```bash
docker save <Docker ID>/sample-web:case2 -o sample-web.tar

mkdir sample-web

tar xvf sample-web.tar -C ./sample-web

for layer in $(cat sample-web/manifest.json | jq -c '.[0].Layers[]' | sed s/\"//g); do \
  layer_dir=sample-web/blobs/sha256/layer-$(eval echo ${layer} | awk '{sub("blobs/sha256/", "");print $0;}') ; \
  mkdir ${layer_dir} ; \
  tar xvf sample-web/${layer} -C ${layer_dir} ; \
done

tree -L 4 -a sample-web

cat sample-web/blobs/sha256/layer-xxxxx/usr/local/apache2/htdocs/index.html
```

`Privateリポジトリへのコンテナイメージアップロード（sample-web-private:case2）`
```bash
docker tag <Docker ID>/sample-web:case2 <Docker ID>/sample-web-private:case2

docker images <Docker ID>/sample-web-private:case2

docker login -u <Docker ID>

docker push <Docker ID>/sample-web-private:case2
```


`DockerHubからのログアウト`
```bash
docker logout
```

`ローカル環境に存在するコンテナイメージの削除`
```bash
docker rmi <Docker ID>/sample-web-private:case2
```

`一般ユーザーとしてPrivateリポジトリからコンテナイメージを取得`
```bash
docker pull <Docker ID>/sample-web-private:case2
```

`DockerHubログイン後のPrivateリポジトリからのコンテナイメージ取得`
```bash
docker login -u <Docker ID>

docker pull <Docker ID>/sample-web-private:case2

docker images <Docker ID>/sample-web-private:case2
```

`Privateリポジトリのコンテナイメージを使用したPodのデプロイ（失敗）`
```bash
kubectl apply -f sample-web-private.yaml

kubectl get pod sample-web-private

kubectl describe pod sample-web-private

kubectl delete pod sample-web-private
```

`レジストリの認証情報を含むSecretの作成`
```bash
kubectl create secret docker-registry regcred \
        --docker-server=<レジストリURL> \
        --docker-username=<Docker ID> \
        --docker-password=<Password> \
        --docker-email=<アカウントメールアドレス>

kubectl get secret regcred
```

`Privateリポジトリのコンテナイメージを使用したPodのデプロイ`
```bash
kubectl apply -f sample-web-private-with-regcred.yaml

kubectl get pod sample-web-private
```

`検証に使用したリソースの削除`
```bash
kubectl delete pod sample-web sample-web-private

kubectl delete secret regcred
```