## gcloud setup

```
gcloud config configurations create telepresence
gcloud config set account ${email}
gcloud config set compute/region asia-northeast1
gcloud config set compute/zone asia-northeast1-a
gcloud auth login
gcloud projects create telepresence-handson-$(perl -e 'printf ("%04d\n",rand(10000))') --set-as-default
```
この時作成されるプロジェクトIDはこの後利用する．

課金アカウントのID確認
```
gcloud beta billing accounts list
```

GKEのAPI有効化
```
gcloud beta billing projects link ${project_id} --billing-account ${billing_account_id}
gcloud services enable compute.googleapis.com
gcloud services enable container.googleapis.com
```

## create cluster
```
gcloud container clusters create telepresence-handson
gcloud container clusters get-credentials telepresence-handson
```

## install telepresence
```
brew cask install osxfuse
brew install datawire/blackbird/telepresence
```