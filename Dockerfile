FROM golang:1.22.0-bookworm as build
LABEL MAINTAINER="shiny.gift6738@fastmail.com"

WORKDIR /app

RUN apt-get update && apt-get -y upgrade
# ADD file changes when new commits have been made - this forces a new checkout instead of using cache
ADD https://api.github.com/repos/skeletonkey/watch-my-ip/git/refs/heads/init_app version.json
RUN git clone -b init_app --single-branch --depth 1 https://github.com/skeletonkey/watch-my-ip.git watch-my-ip
RUN cd watch-my-ip && go build -o bin/watch-my-ip app/*.go

FROM golang:1.22.0-bookworm as prod
LABEL MAINTAINER="shiny.gift6738@fastmail.com"

ENV PROJECT_CONFIG_FILE="/app/config.json"

WORKDIR /app

COPY config.json .
COPY --from=build /app/watch-my-ip/bin/watch-my-ip watch-my-ip

RUN chmod 0755 /app/watch-my-ip

ENTRYPOINT [ "/app/watch-my-ip" ]