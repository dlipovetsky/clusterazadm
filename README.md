# Usage

```shell
NAME="$(whoami)-test"
az ad sp create-for-rbac --role contributor --name "$NAME" --sdk-auth > "$NAME.json"
AZURE_AUTH_LOCATION="$NAME.json" ./clusterazadm
```
