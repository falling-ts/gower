# syntax=docker/dockerfile:1
FROM mysql/mysql-server:5.7.41
ENV MYSQL_ROOT_HOST="%"
ENV MYSQL_DATABASE="gower"
ENV MYSQL_DATABASE_CHAR="utf8mb4,utf8mb4_unicode_ci"
ENV MYSQL_ROOT_PASSWORD="root"
COPY entrypoint.sh .
RUN chmod +x entrypoint.sh
