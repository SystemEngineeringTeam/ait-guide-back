#!/bin/bash
set -euC

# EXTENSION_SCRIPTはコンテナが起動するたびにコールされるため、
# import処理が実施済かフラグファイルの有無をチェック
if [ -f /import/done ]; then
    echo "Skip import process"
    return
fi

# データを全削除
echo "delete database started."
rm -rf /data/databases
rm -rf /data/transactions
echo "delete database finished."

# CSVデータのインポート
echo "Start the data import process"
neo4j-admin import \
  --nodes=/import/point/point.csv \
  --relationships=/import/route/route.csv
echo "Complete the data import process"

# import処理の完了フラグファイルの作成
echo "Start creating flag file"
touch /import/done
echo "Complete creating flag file"
