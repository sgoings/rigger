function choose-deis-version {

  :

  # local options=(
  #                 "Released version"
  #                 "Official GitHub Repository"
  #               )

  # choice-prompt "What Deis would you like to use?" options[@] 1 answer

  # case ${answer} in
  #   1)
  #     prompt "Enter Deis version:" VERSION 1.9.0
  #     ;;
  #   2)
  #     prompt "Enter Deis branch/tag/sha1:" VERSION master
  #     ;;
  # esac
}

function configure-deisctl-tunnel {
  prompt "Enter Deisctl tunnel IP address:" DEISCTL_TUNNEL 127.0.0.1:2222
}

function configure-deis-version {
  prompt "Enter Deis version:" DEIS_VERSION 1.10.0
}

function configure-go {
  ORIGINAL_PATH="${PATH}"
  export ORIGINAL_PATH

  prompt "What's your GOPATH?" GOPATH "${HOME}/go"

  export PATH="${GOPATH}/bin:${PATH}"
}

function configure-deis-root {
  # Needed to run provisioning (provisioning scripts located in repo)
  prompt "Where is the Deis repository located?" DEIS_ROOT "${GOPATH:-${HOME}}/src/github.com/deis/deis"
}

function configure-ipaddr {
  prompt "What's the ip address of your Docker environment?" HOST_IPADDR "$(guess-ipaddr)"
}

function configure-registry {
  case ${PROVIDER} in
    vagrant)
      prompt "Where can I find your Docker registry?" DEV_REGISTRY "$(guess-registry)"
      ;;
    *)
      prompt "What's a publicly available Docker registry I can use?" DEV_REGISTRY "${SUGGEST_DEV_REGISTRY:-}"
      prompt "And an organization/user I can push to (include trailing /)?" IMAGE_PREFIX "${SUGGEST_IMAGE_PREFIX:-}"
      ;;
  esac
}

function configure-app-deployment {
  prompt "What ssh key should I use for application deployment?" DEIS_TEST_AUTH_KEY "${SUGGEST_DEIS_SSH_KEY:-}"
}

function configure-ssh {
  prompt "What ssh key should I use (for deisctl/ssh)?" DEIS_TEST_SSH_KEY "${SUGGEST_DEIS_SSH_KEY:-}"
}

function configure-dns {
  prompt "What wildcard domain name is available for me to use?" DEIS_TEST_DOMAIN "${SUGGEST_DEIS_TEST_DOMAIN:-}"
}

function configure-provider {
  if [ -z "${PROVIDER:-}" ]; then
    local options=(
                    "Vagrant"
                    "Amazon Web Services (AWS)"
                    "Digital Ocean"
                  )

    choice-prompt "What cloud provider would you like to use?" options[@] 1 answer

    case ${answer} in
      1)
        PROVIDER=vagrant
        ;;
      2)
        PROVIDER=aws
        ;;
      3)
        PROVIDER=digitalocean
        ;;
    esac
  fi
}
