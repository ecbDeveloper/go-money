-- Write your migrate up statements here
CREATE TABLE IF NOT EXISTS transferencia (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	id_conta UUID NOT NULL REFERENCES conta(id),
	valor DECIMAL NOT NULL,
	tipo INTEGER NOT NULL REFERENCES tipo_transferencia(id),
	data_transferencia TIMESTAMPTZ DEFAULT now()
)
---- create above / drop below ----
DROP TABLE IF EXISTS transferencia;
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
