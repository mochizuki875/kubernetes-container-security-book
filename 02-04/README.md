# ケース4: コンテナイメージから秘密情報を奪取されてしまった
`コンテナイメージのビルド（image-env-pass:latest）`
```bash
docker build --no-cache --progress=plain --build-arg PASSWORD=pass12345 -t image-env-pass:latest .
```

`コンテナ内で秘密情報が参照できないことの確認`
```bash
docker images image-env-pass:latest

docker run -it --rm image-env-pass:latest env | grep PASSWORD
```

`コンテナイメージのビルド履歴から秘密情報を奪取`
```bash
docker history image-env-pass:latest
```

`コンテナイメージのビルド（image-file-pass:latest）`
```bash
docker build --no-cache --progress=plain -t image-file-pass:latest .
```

`コンテナ内で秘密情報が参照できないことの確認`
```bash
docker run -it --rm image-file-pass:latest cat ./PASSWORD
```

`コンテナイメージのエクスポートと展開`
```bash
docker save image-file-pass:latest -o image-file-pass.tar

mkdir image-file-pass

tar xvf image-file-pass.tar -C ./image-file-pass

for layer in $(cat image-file-pass/manifest.json | jq -c '.[0].Layers[]' | sed s/\"//g); do \
  layer_dir=image-file-pass/blobs/sha256/layer-$(eval echo ${layer} | awk '{sub("blobs/sha256/", "");print $0;}') ; \
  mkdir ${layer_dir} ; \
  tar xvf image-file-pass/${layer} -C ${layer_dir} ; \
done

tree -L 4 -a image-file-pass
```

`秘密情報の奪取`
```bash
cat image-file-pass/blobs/sha256/layer-xxxxx/PASSWORD
```

`コンテナイメージのビルド（sample-image:latest）`
```bash
docker build -t sample-image:latest .
```

`コンテナイメージのエクスポートと展開`
```bash
docker save sample-image:latest -o sample-image.tar

mkdir sample-image

tar xvf sample-image.tar -C ./sample-image

for layer in $(cat sample-image/manifest.json | jq -c '.[0].Layers[]' | sed s/\"//g); do \
  layer_dir=sample-image/blobs/sha256/layer-$(eval echo ${layer} | awk '{sub("blobs/sha256/", "");print $0;}') ; \
  mkdir ${layer_dir} ; \
  tar xvf sample-image/${layer} -C ${layer_dir} ; \
done
```

`コンテナイメージの構成確認`
```bash
tree -L 4 -a sample-image
```

`コンテナイメージのメタデータ確認（index.json）`
```bash
cat sample-image/index.json | jq .
```

`コンテナイメージのメタデータ確認（manifest.json）`
```bash
cat sample-image/manifest.json | jq .
```

`コンテナイメージのメタデータ確認（Config）`
```bash
cat sample-image/blobs/sha256/xxxxx | jq .
```

`コンテナのルートファイルシステムに含まれるファイル・ディレクトリの確認`
```bash
docker run -it sample-image:latest ls /
```

`コンテナイメージに残存しているファイルの確認`
```bash
cat sample-image/blobs/sha256/layer-xxxxx/file_a
```

`Build secretを使用したコンテナイメージのビルド（image-secret:latest）`
```bash
docker build --no-cache --progress=plain --secret id=password,src=PASSWORD -t image-secret:latest .
```

`コンテナに秘密情報が含まれていないことの確認`
```bash
docker run -it --rm image-secret:latest cat /run/secrets/password
```
