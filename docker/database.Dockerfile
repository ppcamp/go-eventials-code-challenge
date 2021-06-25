# Import da imagem
FROM postgres:12-alpine

ARG POSTGRES_USER
ARG POSTGRES_PASSWORD

LABEL maintainer="ppcamp"
LABEL description="Inicia uma instância do banco com os dados"

# Configurando a linguagem
# RUN localedef -i pt_BR -c -f UTF-8 -A /usr/share/locale/locale.alias pt_BR.UTF-8
# ENV LANG pt_BR.utf8

# Configurando o diretório de trabalho
WORKDIR /docker-entrypoint-initdb.d/

# Copiando os dados para a imagem
COPY q1_catalog.csv .

# Extraindo os dados
# RUN tar -xf ./dataIntegrationChallenge.tgz

# Copiando a query sql que será executada
COPY default.sql .

# Fix default database (When start, will fix)
# COPY setup.sh /docker-entrypoint-initdb.d/