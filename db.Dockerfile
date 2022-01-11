FROM mysql:8.0.27

COPY ./database/*.sql /docker-entrypoint-initdb.d/

