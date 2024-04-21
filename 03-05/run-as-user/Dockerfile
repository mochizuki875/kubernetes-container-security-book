FROM ubuntu:22.04

# user01というユーザー（uid=1000）とグループ（gid=1000）およびホームディレクトリの作成
RUN useradd -m -U -u 1000 user01

# コンテナの実行ユーザーとしてuser01を指定
USER user01

# ワークディレクトリ（/home/user01）の設定
WORKDIR /home/user01

CMD ["/bin/sh", "-c", "while :; do sleep 10; done"]