CREATE TABLE public.customers (
    customers_id bigserial NOT NULL,
    customers_email varchar NOT NULL,
    customers_createat int8 DEFAULT 0 NOT NULL,
    CONSTRAINT customers_id_pk PRIMARY KEY (customers_id),
    CONSTRAINT customers_email_unique UNIQUE (customers_email)
);

CREATE TABLE public.groups (
    groups_id bigserial NOT NULL,
    groups_token varchar NOT NULL,
    groups_createat int8 DEFAULT 0 NOT NULL,
    CONSTRAINT groups_id_pk PRIMARY KEY (groups_id),
    CONSTRAINT groups_token_unique UNIQUE (groups_token)
);

CREATE TABLE public.messages (
    messages_id bigserial NOT NULL,
    messages_customer_id bigint NOT NULL,
    messages_group_id bigint NOT NULL,
    messages_title varchar(50) NOT NULL,
    messages_content varchar NOT NULL,
    messages_createat int8 NULL,
    CONSTRAINT messages_id_pkey PRIMARY KEY (messages_id)
);
ALTER TABLE public.messages ADD CONSTRAINT messages_customer_id_fk FOREIGN KEY (messages_customer_id) REFERENCES public.customers(customers_id);
ALTER TABLE public.messages ADD CONSTRAINT messages_group_id_fk FOREIGN KEY (messages_group_id) REFERENCES public.groups(groups_id);
