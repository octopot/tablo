FROM postgres:alpine AS storage

LABEL maintainer="Kamil Samigullin <kamil@samigullin.info>"

COPY initdb.d /docker-entrypoint-initdb.d
