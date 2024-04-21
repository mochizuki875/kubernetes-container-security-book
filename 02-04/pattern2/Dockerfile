FROM ubuntu:22.04

# 秘密情報の記載されたファイルを渡す
COPY PASSWORD ./PASSWORD

# ファイルに記載された秘密情報を使用してコマンドを実行
RUN echo Use credential: $(cat ./PASSWORD)

# 秘密情報の記載されたファイルを削除
RUN rm ./PASSWORD

CMD ["/bin/sh", "-c", "while :; do sleep 10; done"]