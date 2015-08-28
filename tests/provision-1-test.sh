#!/usr/bin/env roundup
#
#/ usage:  rerun stubbs:test -m deis -p provision [--answers <>]
#

# Helpers
# -------
[[ -f ./functions.sh ]] && . ./functions.sh

describe "provision"

it_stops_if_seeing_error() {
  rigger configure <<EOF
1.9.1
1
EOF
  PROVIDER=noop rigger provision
}
