FROM mysql:8.0.23

COPY ./docs/*.sql /docker-entrypoint-initdb.d/
