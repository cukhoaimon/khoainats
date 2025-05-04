CREATE TABLE IF NOT EXISTS public.users(
    id              uuid PRIMARY KEY        NOT NULL DEFAULT uuid_generate_v4(),
    email           text                    NOT NULL,
    password        text,
    created_at      timestamp               DEFAULT now() NOT NULL,
    updated_at      timestamp,
    deleted_at      timestamp
);

CREATE UNIQUE INDEX idx_users_email ON users(email) WHERE deleted_at IS NULL;
CREATE UNIQUE INDEX idx_users_id ON users(id) WHERE deleted_at IS NULL;

CREATE TABLE IF NOT EXISTS public.organizations (
    id              uuid PRIMARY KEY        NOT NULL DEFAULT uuid_generate_v4(),
    owner_email     text NOT NULL,
    domain          text NOT NULL,
    description     text,
    created_at      timestamp               DEFAULT now() NOT NULL,
    created_by      uuid NOT NULL,
    updated_at      timestamp,
    updated_by      uuid,
    deleted_at      timestamp,
    deleted_by      uuid
);

CREATE UNIQUE INDEX idx_organizations ON organizations (id) WHERE deleted_at IS NULL;

CREATE TABLE IF NOT EXISTS public.login_events (
    id              uuid PRIMARY KEY        NOT NULL DEFAULT uuid_generate_v4(),
    ip_address      text                    NOT NULL,
    code            text                    NOT NULL,
    attempts        int                     DEFAULT 0 NOT NULL,
    created_at      timestamp               DEFAULT now() NOT NULL,
    succeed_at      timestamp,
    last_tried_at   timestamp
);

CREATE UNIQUE INDEX idx_login_events ON login_events(id);


CREATE TYPE principal_type AS ENUM (
    'user',
    'service'
);

CREATE TABLE IF NOT EXISTS public.principals(
    id              uuid PRIMARY KEY        NOT NULL DEFAULT uuid_generate_v4(),
    type            principal_type,
    organization_id uuid                    NOT NULL,
    created_at      timestamp               DEFAULT now() NOT NULL,
    deleted_at      timestamp,
    deleted_by      uuid,

    constraint fk_principals_organization foreign key (organization_id) references organizations (id)
);

CREATE UNIQUE INDEX idx_principals ON principals(id) WHERE deleted_at IS NULL;

CREATE TYPE principal_attribute AS ENUM (
    'email',       -- for user
    'mac_address'  -- for device
);

CREATE TABLE IF NOT EXISTS public.principal_attributes (
    id                  uuid PRIMARY KEY    NOT NULL DEFAULT uuid_generate_v4(),
    principal_id        uuid                NOT NULL,
    attribute           principal_attribute NOT NULL,
    attribute_value     text                NOT NULL,

    created_at          timestamp           DEFAULT now() NOT NULL,
    created_by          uuid                NOT NULL,
    deleted_at          timestamp,
    deleted_by          uuid,

    constraint fk_principal_attributes_principals foreign key (principal_id) references principals (id)
);

CREATE UNIQUE INDEX idx_principal_attributes ON principal_attributes(principal_id) WHERE deleted_at IS NULL;

CREATE TABLE IF NOT EXISTS public.access_tokens (
    id                  uuid PRIMARY KEY    NOT NULL DEFAULT uuid_generate_v4(),
    principal_id        uuid                NOT NULL,
    organization_id     uuid                NOT NULL,
    roles               text[]              NOT NULL,

    created_at          timestamp           DEFAULT now() NOT NULL,
    updated_at          timestamp,
    revoked_at          timestamp,
    revoked_by          uuid,

    constraint fk_access_token_principals foreign key (principal_id) references principals(id),
    constraint fk_access_token_organization foreign key (organization_id) references organizations(id)

);

CREATE UNIQUE INDEX idx_access_tokenss ON access_tokens(id);


