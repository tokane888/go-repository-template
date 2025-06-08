FROM mcr.microsoft.com/devcontainers/go:1-1.23-bookworm

# install air
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b /usr/local/bin

# install apt packages
RUN apt-get update -y
RUN apt-get install -y dnsutils pre-commit
