#!/bin/bash
set -euxo pipefail

add_account() {
  local address="$1"
  local amount="$2"
  echo $(jq ". += [{ \"address\": \"$address\", \"amount\":  \"$amount\"}]")
}