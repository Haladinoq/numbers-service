CREATE TABLE reservation (
                          id serial4 NOT NULL,
                          client text NOT NULL,
                          number int4 NOT NULL,
                          created_at timestamp NULL DEFAULT now(),
                          updated_at timestamp NULL DEFAULT now(),
                          CONSTRAINT reservation_pkey PRIMARY KEY (id),
                          UNIQUE (client,number)
);