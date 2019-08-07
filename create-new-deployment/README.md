## Setup
```
kubectl apply -f manifests/
```

## use Telepresence

Deoloymentを追加する
```
cd new_deploy
docker build . -t new_deploy
telepresence --new-deployment new-deploy --docker-run --rm -p 80:80 -e APP_B_ENDPOINT="http://secret/secret" new_deploy
```
手元にSecretリソースがなくても変更後のコンテナがSecretリソースにアクセスできていることを確認する．

dockerコマンドで環境変数の書き換える．
```
telepresence --swap-deployment handson --docker-run --rm -p 80:80 -e MY_SECRET=new_secret secret
```
localで，変数を変更して試すことができる．

ファイル読み込みでも同じようにアクセス可能なのできになる人は試してみるといいかも？

## Cleanup
```
kubectl delete -f manifests/
```