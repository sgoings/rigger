function eval-provider-file {
  local filename="${1}"

  {
    cd ${DEIS_ROOT}/contrib/${PROVIDER}
    if [ -f ${1} ]; then
      eval "$(cat ${1})"
    else
      not-implemented
    fi
  }
}

function _setup-provider-dependencies {
  eval-provider-file "install.sh"
}

function _create {
  eval-provider-file "create.sh"
}

function _destroy {
  eval-provider-file "destroy.sh"
}

function _check-cluster {
  eval-provider-file "check.sh"
}
