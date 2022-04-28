#!/bin/sh

ARGUMENTS=$(echo "$ARGOCD_APP_PARAMETERS" | jq -r '.[] | select(.name == "values-files").array | .[] | "--values=" + .')
PARAMETERS=$(echo "$ARGOCD_APP_PARAMETERS" | jq -r '.[] | select(.name == "helm-parameters").map | to_entries | map("\(.key)=\(.value)") | .[] | "--set=" + .')
ARGUMENTS=$(printf "%s\n%s" "$ARGUMENTS" "$PARAMETERS")

echo "$ARGUMENTS" | xargs helm template .
