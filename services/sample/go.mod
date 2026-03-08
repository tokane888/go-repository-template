// TODO: モジュール名調整
module github.com/tokane888/go-repository-template/services/sample

go 1.25.7

require (
	github.com/joho/godotenv v1.5.1
	github.com/tokane888/go-repository-template/pkg/logger v0.0.0
)

replace github.com/tokane888/go-repository-template/pkg/logger => ../../pkg/logger
