-- Write your migrate up statements here
CREATE TABLE IF NOT EXISTS transferencia (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	valor DECIMAL NOT NULL,
	tipo INT REFERENCES tipo_transferencia(id)
)
---- create above / drop below ----
DROP TABLE IF EXISTS transferencia;
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
