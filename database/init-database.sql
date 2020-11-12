-- public.accounts definition

-- Drop table

-- DROP TABLE public.accounts;

CREATE TABLE public.accounts (
	id uuid NOT NULL,
	account_id varchar NOT NULL,
	"password" varchar NOT NULL,
	is_active bool NOT NULL,
	created_at timestamp(0) NOT NULL,
	updated_at timestamp(0) NOT NULL,
	CONSTRAINT accounts_pk_id PRIMARY KEY (id),
	CONSTRAINT accounts_un_account_id UNIQUE (account_id)
);

-- Permissions

ALTER TABLE public.accounts OWNER TO "spectra-admin-dev";
GRANT ALL ON TABLE public.accounts TO "spectra-admin-dev";


-- public.permissions definition

-- Drop table

-- DROP TABLE public.permissions;

CREATE TABLE public.permissions (
	id uuid NOT NULL,
	"permission" varchar NOT NULL,
	description varchar NULL,
	CONSTRAINT permissions_pk_id PRIMARY KEY (id),
	CONSTRAINT permissions_un_permission_name UNIQUE (permission)
);

-- Permissions

ALTER TABLE public.permissions OWNER TO "spectra-admin-dev";
GRANT ALL ON TABLE public.permissions TO "spectra-admin-dev";


-- public.roles definition

-- Drop table

-- DROP TABLE public.roles;

CREATE TABLE public.roles (
	id uuid NOT NULL,
	"role" varchar NOT NULL,
	description varchar NULL,
	CONSTRAINT roles_pk_id PRIMARY KEY (id),
	CONSTRAINT roles_un_role_name UNIQUE (role)
);

-- Permissions

ALTER TABLE public.roles OWNER TO "spectra-admin-dev";
GRANT ALL ON TABLE public.roles TO "spectra-admin-dev";


-- public.account_info definition

-- Drop table

-- DROP TABLE public.account_info;

CREATE TABLE public.account_info (
	id uuid NOT NULL,
	account_id varchar NOT NULL,
	firstname varchar NOT NULL,
	lastname varchar NULL,
	email varchar NOT NULL,
	phone_number varchar NOT NULL,
	CONSTRAINT account_info_pk PRIMARY KEY (id),
	CONSTRAINT account_info_fk FOREIGN KEY (account_id) REFERENCES accounts(account_id) ON UPDATE SET NULL ON DELETE SET NULL
);

-- Permissions

ALTER TABLE public.account_info OWNER TO "spectra-admin-dev";
GRANT ALL ON TABLE public.account_info TO "spectra-admin-dev";


-- public."privileges" definition

-- Drop table

-- DROP TABLE public."privileges";

CREATE TABLE public."privileges" (
	"role" varchar NOT NULL,
	"permission" varchar NOT NULL,
	id uuid NOT NULL,
	CONSTRAINT privileges_pk PRIMARY KEY (id),
	CONSTRAINT privileges_fk FOREIGN KEY (role) REFERENCES roles(role) ON UPDATE SET NULL ON DELETE SET NULL,
	CONSTRAINT privileges_fk_1 FOREIGN KEY (permission) REFERENCES permissions(permission) ON UPDATE SET NULL ON DELETE SET NULL
);

-- Permissions

ALTER TABLE public."privileges" OWNER TO "spectra-admin-dev";
GRANT ALL ON TABLE public."privileges" TO "spectra-admin-dev";