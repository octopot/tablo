DOCKER_IMAGE_NAME    = octopot/tablo
DOCKER_IMAGE_BACKUP  = quay.io/octopot/tablo
DOCKER_IMAGE_VERSION = 1.x

_active = $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_VERSION)
_latest = $(DOCKER_IMAGE_NAME):latest
_image  = $(_active)
_name   = $(subst /,-,$(DOCKER_IMAGE_NAME))-dev

.PHONY: docker-backup
docker-backup: _active_backup = $(DOCKER_IMAGE_BACKUP):$(DOCKER_IMAGE_VERSION)
docker-backup: _latest_backup = $(DOCKER_IMAGE_BACKUP):latest
docker-backup:
	@docker tag $(_active) $(_active_backup)
	@docker tag $(_latest) $(_latest_backup)
	@docker push $(_active_backup)
	@docker push $(_latest_backup)

.PHONY: docker-build
docker-build:
	docker build -f env/docker/service/Dockerfile \
	             -t $(_active) \
	             -t $(_latest) \
	             --force-rm --no-cache --pull --rm \
	             .

.PHONY: docker-publish
docker-publish: docker-build docker-push

.PHONY: docker-push
docker-push:
	@docker push $(_active)
	@docker push $(_latest)

.PHONY: docker-run
docker-run:
	@docker run --rm -it \
	            --name $(_name) \
	            -p 8080:8080 \
	            -p 8090:8090 \
	            -p 8091:8091 \
	            -p 8092:8092 \
	            $(_image) run --with-monitoring --with-profiling
