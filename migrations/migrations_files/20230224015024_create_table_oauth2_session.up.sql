CREATE TABLE public.oauth2_session (
	id uuid NOT NULL,
	user_id uuid NULL,
	is_active bool NULL DEFAULT true,
	created_at timestamptz NULL,
	client_id text NOT NULL,
	client_secret text NOT NULL,
	CONSTRAINT oauth2_session_pkey PRIMARY KEY (id),
	CONSTRAINT fk_oauth2_session_user FOREIGN KEY (user_id) REFERENCES public."user"(id) ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX idx_oauth2_session_id ON public.oauth2_session USING btree (id);
CREATE INDEX idx_oauth2_session_user_id ON public.oauth2_session USING btree (user_id);
