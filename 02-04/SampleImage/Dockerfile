# [ビルドステップ0] ベースイメージを指定
FROM ubuntu:22.04

# [ビルドステップ1] ファイルAを作成
RUN touch file_a

# [ビルドステップ2] ファイルAに文字列を書き込み
RUN echo "test" > file_a

# [ビルドステップ3] ファイルBを作成
RUN touch file_b

# [ビルドステップ4] ファイルAを削除
RUN rm file_a

# コンテナ起動時の実行コマンドを設定
CMD ["/bin/sh", "-c", "while :; do sleep 10; done"]