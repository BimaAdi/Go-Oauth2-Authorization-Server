CREATE TABLE public.oauth2_token (
	id uuid NOT NULL,
	user_id uuid NULL,
	code varchar NOT NULL,
	created_at timestamptz NULL,
	CONSTRAINT oauth2_token_pkey PRIMARY KEY (id),
	CONSTRAINT fk_oauth2_token_user FOREIGN KEY (user_id) REFERENCES public."user"(id) ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX idx_oauth2_token_code ON public.oauth2_token USING btree (code);
CREATE INDEX idx_oauth2_token_id ON public.oauth2_token USING btree (id);
CREATE INDEX idx_oauth2_token_user_id ON public.oauth2_token USING btree (user_id);
