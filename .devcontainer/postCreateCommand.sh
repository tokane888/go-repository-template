#!/bin/bash

# sudoで実行される前提で各種パッケージインストール

apt-get update -y
apt-get install -y dnsutils pre-commit
