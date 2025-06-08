FROM mcr.microsoft.com/devcontainers/go:1-1.23-bookworm

# Go 1.24 を手動でインストール（lefthookのインストールに必要なため）
# テンプレ改修TODO: 公式で1.24のイメージが出たら、これを削除
ENV GO_VERSION=1.24.0
RUN curl -LO https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz && \
    rm -rf /usr/local/go && \
    tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz && \
    rm go${GO_VERSION}.linux-amd64.tar.gz
# go toolを1つインストールして強制的に /go/pkg を作らせ、/go/pkgの所有者を調整
# テンプレ改修TODO: go 1.24公式イメージが出たらここも削除
RUN go install github.com/evilmartians/lefthook@latest || true
RUN chown -R vscode:golang /usr/local/go /go/pkg /go/bin /go/src
ENV PATH="/usr/local/go/bin:${PATH}"

# install air
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b /usr/local/bin

# install apt packages
RUN apt-get update -y
RUN apt-get install -y dnsutils

RUN go install github.com/evilmartians/lefthook@latest
