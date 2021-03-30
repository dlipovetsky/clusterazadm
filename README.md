# clusterazadm

## Installation

```shell
go get github.com/dlipovetsky/clusterazadm
```

## Usage

### Create Azure Service Principal (if it does not exist)

```shell
NAME="$(whoami)-test"
az ad sp create-for-rbac --role contributor --name "$NAME" --sdk-auth > "$NAME.json"
```

### Patch CAPZ Bootstrap Credentials Secret

```shell
AZURE_AUTH_LOCATION="$NAME.json" ./clusterazadm > patch.json
kubectl -n capz-system patch secret capz-manager-bootstrap-credentials --patch-file patch.json
```
