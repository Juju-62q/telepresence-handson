## Setup
```
kubectl apply -f manifests/
```

動作確認
```
kubectl get svc
```
handsonのEXTERNAL IPにアクセスする．
下記構造体が表示されればOK
```
http://${EXTERNAL IP}/secret
{"secret": "secret is secret"}
http://${EXTERNAL IP}/config
{"config": "configmap is configmap"}
```

## use Telepresence
変更を確認する．
さらに変更するのもOK!!
```
cd secret_test
${EDITOR} main.go
```

Deoloymentを入れ替える
```
docker build . -t secret
telepresence --swap-deployment handson --docker-run --rm -p 80:80 secret
```
手元にSecretリソースがなくても変更後のコンテナがSecretリソースにアクセスできていることを確認する．

## Cleanup
```
kubectl delete -f manifests/
```
