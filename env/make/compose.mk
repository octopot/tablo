PROJECT = tablo

_base    = docker-compose -p $(PROJECT) -f env/docker/compose/docker-compose.base.yml
_compose = $(_base) -f env/docker/compose/docker-compose.dev.yml


.PHONY: demo
demo:
	$(eval _compose = $(_base) -f env/docker/compose/docker-compose.demo.yml)
	@echo $(_compose)

.PHONY: e2e
e2e:
	$(eval _compose = $(_base) -f env/docker/compose/docker-compose.e2e.yml)
	@echo $(_compose)

.PHONY: up
up:
	@$(_compose) up -d

.PHONY: fresh-up
fresh-up:
	@$(_compose) up -d --build --force-recreate

.PHONY: down
down:
	@$(_compose) down

.PHONY: destroy
destroy:
	@$(_compose) down --volumes --rmi local

.PHONY: status
status:
	@$(_compose) ps


.PHONY: compose-clean
compose-clean:
	@$(_compose) rm -f

.PHONY: compose-config
compose-config:
	@$(_compose) config
