-- Write your migrate up statements here
CREATE TABLE IF NOT EXISTS tipo_transferencia (
	id SERIAL PRIMARY KEY,
	tipo VARCHAR(50)
);

INSERT INTO tipo_transferencia (tipo) 
VALUES ('Depósito'), ('Saque');
---- create above / drop below ----
DROP TABLE IF EXISTS tipo_transferencia;
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
