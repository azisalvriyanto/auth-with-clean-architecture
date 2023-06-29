FROM mysql:8.0

COPY ./docs/*.sql /docker-entrypoint-initdb.d/
