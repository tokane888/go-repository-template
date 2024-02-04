FROM mcr.microsoft.com/devcontainers/go:1-1.21-bullseye

# debian公式repoのpostgresql-client(psql)が古く、postgres本体のmajor verと不一致であるためpostgres公式repoからpsqlインストール
RUN sudo sh -c 'echo "deb https://apt.postgresql.org/pub/repos/apt $(lsb_release -cs)-pgdg main" > /etc/apt/sources.list.d/pgdg.list'
RUN wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | sudo apt-key add -
RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends \
     postgresql-client
RUN go install -v \
    golang.org/x/tools/cmd/goimports@latest \
    golang.org/x/tools/gopls@latest \
    github.com/go-delve/delve/cmd/dlv@latest

# [Optional] Uncomment the next lines to use go get to install anything else you need
# USER vscode
# RUN go get -x <your-dependency-or-tool>
# USER root

# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1
