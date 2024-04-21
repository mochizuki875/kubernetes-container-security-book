FROM ubuntu:22.04

# 秘密情報を使用してコマンドを実行
# --secretオプションで渡した秘密情報を使用
RUN --mount=type=secret,id=password echo Use credential: $(cat /run/secrets/password)

CMD ["/bin/sh", "-c", "while :; do sleep 10; done"]