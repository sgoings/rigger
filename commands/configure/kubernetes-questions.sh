function configure-deis-repo {
  configure-kubernetes-repo
}

function configure-deis-root {
  configure-kubernetes-root
}

function configure-user-type {
  local answer
  local options=(
                  "Release"
                  "Path"
                  "Git"
                )

  choice-prompt "Where can I find the version of Kubernetes you want?" options[@] 1 DEIS_SOURCE

  case ${DEIS_SOURCE} in
    1) # released version
      configure-deis-version
      export DEIS_GIT_REPO="${SUGGEST_KUBERNETES_GIT_REPO}"
      export DEIS_GIT_VERSION="${VERSION}"
      configure-deis-repo
      export GOPATH="${DEIS_ID_DIR}/go"
      export DEIS_ROOT="${GOPATH}/src/github.com/kubernetes/kubernetes"
      save-vars GOPATH DEIS_ROOT
      ;;
    2) # path based version
      ;;
    3) # Git based version
      configure-deis-repo
      export GOPATH="${DEIS_ID_DIR}/go"
      export DEIS_ROOT="${GOPATH}/src/github.com/kubernetes/kubernetes"
      save-vars GOPATH DEIS_ROOT
      ;;
  esac
  
  export DEIS_SOURCE
  save-vars DEIS_SOURCE
}

function configure-kubernetes-root {
  # Needed to run provisioning (provisioning scripts located in repo)
  prompt "Where is the Kubernetes repository located?" DEIS_ROOT "${GOPATH:-${HOME}}/src/github.com/kubernetes/kubernetes"

  save-vars DEIS_ROOT
}

function configure-kubernetes-repo {
  prompt "Enter Kubernetes git repo url:" DEIS_GIT_REPO "${SUGGEST_KUBERNETES_GIT_REPO}"
  prompt "Enter Kubernetes git branch/tag/sha1:" DEIS_GIT_VERSION "${SUGGEST_KUBERNETES_GIT_VERSION}"

  export VERSION="${DEIS_GIT_VERSION}"
  save-vars DEIS_GIT_REPO DEIS_GIT_VERSION VERSION
}

function configure-registry {
  :
}

function configure-ssh {
  :
}

function configure-app-deployment {
  :
}

function configure-dns {
  :
}

  # configure-user-type

  # configure-go
  # configure-deis-root

  # source-config

  # if [ ${DEIS_SOURCE} -ne 2 ]; then
  #   checkout-deis "${DEIS_ROOT}" "${DEIS_GIT_VERSION:-${VERSION}}"
  # else
  #   (
  #     cd ${DEIS_ROOT}
  #     export DEIS_GIT_VERSION="$(git describe --long --dirty --abbrev=10 --tags --always)"
  #     export VERSION="${DEIS_GIT_VERSION}"
  #     save-vars DEIS_GIT_VERSION VERSION
  #   )
  # fi

  # choose-provider

  # setup-provider "${PROVIDER:-}"

  # configure-provider

  # load-env "${RIGGER_VARS_FILE}"

  # create-docker-env
  # configure-registry
  # activate-docker-machine-env

  # configure-ssh

  # configure-app-deployment

  # configure-dns

  # update-link "${RIGGER_VARS_FILE}"
