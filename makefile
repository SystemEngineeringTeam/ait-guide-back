.DEFAULT_GOAL := help
-include .env

.PHONY: build
build:## コンテナビルド
	# コンテナビルド中
	@docker-compose build

.PHONY: up
up: ## コンテナ起動
	# コンテナ起動中
	@docker-compose up -d

.PHONY: build-up
build-up: build up ## コンテナビルド＆起動

.PHONY: down
down: ## コンテナ停止
	@docker-compose down

.PHONY: restart
restart: down up ## コンテナ再起動

.PHONY: logs
logs:## コンテナログ表示
	@docker-compose logs -f

.PHONY: db
db: ## DBコンテナに入る
	docker exec -it $(MYSQL_CONTAINER_NAME) mysql -u $(MYSQL_USER) -p $(MYSQL_PASS)

.PHONY: go
go: ## goコンテナに入る
	docker exec -it $(API_CONTAINER_NAME) /bin/bash

.PHONY: help
help: ## ヘルプ
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'