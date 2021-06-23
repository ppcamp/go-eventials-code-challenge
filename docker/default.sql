CREATE DATABASE yawoen;
\c yawoen;

-- Cria a tabela principal
CREATE TABLE companies (
	id SERIAL PRIMARY KEY,
	name VARCHAR(200),
	addressZip VARCHAR(5)
);
-- Faz uma copia desta tabela
CREATE TEMP TABLE temp_tbl AS SELECT * FROM companies LIMIT 0;

-- Insere o CSV na tabela temporária
COPY temp_tbl(name,addressZip) FROM '/docker-entrypoint-initdb.d/q1_catalog.csv' DELIMITER ';' CSV HEADER;

-- Move os dados para a tabela original realizando o tratamento destes
INSERT INTO companies (name,addressZip) SELECT UPPER(name),addressZip FROM temp_tbl;

-- Deleta a tabela temporária
DROP TABLE temp_tbl;