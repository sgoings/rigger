function eval-provider-file {
  local filename="rigger/${1}"

  local file="${PROVIDER_DIR}/${PROVIDER:-}/${filename}"
  rerun_log debug "running ${file}"

  if [ -f "${file}" ]; then
    pushd ${PROVIDER_DIR}/${PROVIDER} &> /dev/null
    rerun_log debug "pwd for provider is $(pwd)"
    eval "$(cat ${file})"
    popd &> /dev/null
  else
    not-implemented ${FUNCNAME[1]}
  fi
}

function not-implemented {
  local function_name="${1:-${FUNCNAME[1]}}"

  rerun_log debug "No implementation of ${function_name} in ${PROVIDER:-}"
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

function _load-provider-config {
  eval-provider-file "config.sh"
}
