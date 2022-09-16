#!/usr/bin/env bash

set -o errexit
set -o pipefail
set -o nounset

# Parameters
ACTION="$1"

deploy_all () { 
    docker compose up -d
}

restart () {
    docker compose build --no-cache
    deploy_all
}

stop () {
    docker compose stop
}

main () {
    case "$ACTION" in
        deploy_all)
            deploy_all
            ;;        
        restart)
            restart
            ;;
        stop)
            stop
            ;;
        *)
            echo $"Usage: $0 {deploy_all|restart|stop}"
            exit 1
    
    esac
}

# Start main function
main
