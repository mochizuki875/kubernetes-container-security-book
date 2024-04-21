FROM ubuntu:22.04

# --build-argで渡した変数をビルド時の環境変数に設定
ARG PASSWORD

# ビルド変数として渡した秘密情報を使用してコマンドを実行
RUN echo Use credential: $PASSWORD

CMD ["/bin/sh", "-c", "while :; do sleep 10; done"]