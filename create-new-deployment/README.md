## Setup
```
kubectl apply -f manifests/
```

確認
```
kubectl get all
```
`service/secret`, `deployment.apps/secret` があればOK

## use Telepresence

Deoloymentを追加する
```
cd new_app
docker build . -t new_app
telepresence --new-deployment new-app --expose 80:80 --docker-run --rm -p 80:80 -e APP_B_ENDPOINT="http://secret" new_app 
```

別タブでリソースの変化を確認する
```
kubectl get all
```

deployment, replicaset, pod, serviceにnew-appが追加されていることが確認できる。

K8sのCoreDNSで名前解決をしながら疎通してみる
```
telepresence --run curl http://new-app
```

下記のようなレスポンスが確認できる
```
{"message":"this is added!","secret":"secret is secret"}
```

これでtelepresenceでクラスタに自在にアクセスできることや、新しいdeploymentを作ることができるということがわかった。

## Cleanup
```
cd ..
kubectl delete -f manifests/
```