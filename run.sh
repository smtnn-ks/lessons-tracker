POSTGRES_URL="postgres://localuser:localpassword@localhost:5432/postgres"
MIGRATIONS_DIR="./migrations"

migrate() {
    goose postgres $POSTGRES_URL --dir $MIGRATIONS_DIR ${@:2}
}

docker_compose() {
    docker compose -f infra/docker-compose.yml -f infra/docker-compose.local.yml ${@:2}
}

install_tailwind() {
    curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-macos-arm64
    chmod +x tailwindcss-macos-arm64
    mv tailwindcss-macos-arm64 tailwind/tailwindcss
}

########################################

command="$1"
if [ -z "$command" ]
then
 echo "run.sh [command] [args]"
 exit 0;
else
 $command $@
fi
