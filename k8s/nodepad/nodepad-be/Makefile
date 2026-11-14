.PHONY: sync run api openapi ent ent-new

include .env
export

ent:
	go generate ./ent

# make ent-new name=User
ent-new:
	go run -mod=mod entgo.io/ent/cmd/ent new ${name}

# make migrate-diff name=init
migrate-diff: ent
	atlas migrate diff $(name) \
		--dir "file://ent/migrate/migrations" \
		--to "ent://ent/schema" \
		--dev-url "docker://postgres/16/dev?search_path=public"

# make migrate-apply
migrate-apply: ent
	atlas migrate apply \
		--dir "file://ent/migrate/migrations" \
		--url "$(DATABASE_DSN)" \
		--allow-dirty

openapi:
	goctl api swagger \
	--api api/nodepad.api \
	--dir . \
	--yaml

api:
	goctl api go \
	-api api/nodepad.api \
	-dir .
	$(MAKE) openapi

run:
	go run nodepad.go
