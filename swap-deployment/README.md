## Setup
```
kubectl apply -f manifests/
```

動作確認
```
kubectl get svc
```
handsonのEXTERNAL IPにアクセスする．（nginxの初期画面が表示されればOK）

## use Telepresence
Deploymentを入れ替える
```
docker build . -t swap-image
telepresence --swap-deployment handson --docker-run --rm -p 80:80 swap-image
```

## Cleanup
```
kubectl delete -f manifests/
```