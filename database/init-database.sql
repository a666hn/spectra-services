-- Generate Account Table

-- public.accounts definition

-- Drop table

-- DROP TABLE public.accounts;

CREATE TABLE public.accounts (
	id uuid NOT NULL,
	firstname varchar NOT NULL,
	lastname varchar NOT NULL,
	fullname varchar NOT NULL,
	nickname varchar NULL,
	username varchar NOT NULL,
	email varchar NOT NULL,
	phone_number varchar NOT NULL,
	address_id varchar NULL,
	created_at timestamp(0) NOT NULL,
	updated_at timestamp(0) NOT NULL,
	created_by varchar NULL,
	updated_by varchar NULL,
	family_org_id uuid NOT NULL,
	CONSTRAINT account_id_pk PRIMARY KEY (id),
	CONSTRAINT accounts_unique_key UNIQUE (username, email, phone_number)
);
CREATE INDEX accounts_fullname_idx ON public.accounts USING btree (fullname, username, email, phone_number, address_id);


-- public.accounts foreign keys

ALTER TABLE public.accounts ADD CONSTRAINT accounts_fk FOREIGN KEY (id) REFERENCES family_org(id) ON UPDATE SET NULL ON DELETE SET NULL;

-- Generate Account Address Tables

-- public.account_address definition

-- Drop table

-- DROP TABLE public.account_address;

CREATE TABLE public.account_address (
	account_id uuid NOT NULL,
	province varchar NULL,
	city varchar NULL,
	district varchar NULL,
	village varchar NULL,
	street varchar NULL,
	postal_code varchar NULL,
	created_at timestamp(0) NOT NULL,
	updated_at timestamp(0) NOT NULL,
	created_by varchar NULL,
	updated_by varchar NULL
);


-- public.account_address foreign keys

ALTER TABLE public.account_address ADD CONSTRAINT account_address_fk FOREIGN KEY (account_id) REFERENCES accounts(id) ON UPDATE SET NULL ON DELETE SET NULL;

-- Generate Family ORG Table

-- public.family_org definition

-- Drop table

-- DROP TABLE public.family_org;

CREATE TABLE public.family_org (
	id uuid NOT NULL,
	"name" varchar NOT NULL,
	identitynumber varchar NOT NULL,
	status bool NOT NULL DEFAULT false,
	created_at timestamp(0) NOT NULL,
	updated_at timestamp(0) NOT NULL,
	created_by varchar NULL,
	updated_by varchar NULL,
	CONSTRAINT family_org_pk PRIMARY KEY (id)
);
CREATE INDEX family_org_name_idx ON public.family_org USING btree (name, identitynumber, status);

-- Generate Role Table

-- public.roles definition

-- Drop table

-- DROP TABLE public.roles;

CREATE TABLE public.roles (
	id uuid NOT NULL,
	"name" varchar NOT NULL,
	description varchar NULL,
	CONSTRAINT roles_pk PRIMARY KEY (id),
	CONSTRAINT roles_un UNIQUE (name)
);

-- Generate Permissions Table

-- public.permissions definition

-- Drop table

-- DROP TABLE public.permissions;

CREATE TABLE public.permissions (
	id uuid NOT NULL,
	"name" varchar NOT NULL,
	description varchar NULL,
	CONSTRAINT permissions_pk PRIMARY KEY (id),
	CONSTRAINT permissions_un UNIQUE (name)
);

-- Generate Privileges Table

-- public."privileges" definition

-- Drop table

-- DROP TABLE public."privileges";

CREATE TABLE public."privileges" (
	role_id uuid NOT NULL,
	role_name varchar NOT NULL,
	permission_id uuid NOT NULL,
	permission_name varchar NOT NULL
);


-- public."privileges" foreign keys

ALTER TABLE public."privileges" ADD CONSTRAINT privileges_fk FOREIGN KEY (role_id) REFERENCES roles(id) ON UPDATE SET NULL ON DELETE SET NULL;
ALTER TABLE public."privileges" ADD CONSTRAINT privileges_fk_1 FOREIGN KEY (permission_id) REFERENCES permissions(id) ON UPDATE SET NULL ON DELETE SET NULL;