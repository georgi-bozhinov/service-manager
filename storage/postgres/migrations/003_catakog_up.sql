BEGIN;

ALTER TABLE brokers
    DELETE COLUMN catalog

ALTER TABLE brokers
    ADD CONSTRAINT UNIQUE broker_url

CREATE TABLE service_offerings (
    id varchar(100) PRIMARY KEY NOT NULL,
    name varchar(255) UNIQUE NOT NULL,
    description text NOT NULL DEFAULT '',
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,

    catalog_id varchar(255) NOT NULL,
    catalog_name varchar(255) NOT NULL,
    bindable smallint(1) NOT NULL DEFAULT '0',
    instances_retrievable smallint(1) NOT NULL DEFAULT '0',
    bindings_retrievable smallint(1) NOT NULL DEFAULT '0',
    plan_updatable NOT NULL smallint(1) DEFAULT '0',

    metadata json NOT NULL DEFAULT '{}',
    tags json NOT NULL DEFAULT '{}',
    requires json NOT NULL DEFAULT '{}',

    service_broker_id varchar(100) NOT NULL REFERENCES brokers (id) ON DELETE CASCADE
    UNIQUE (catalog_id, service_broker_id)
;

CREATE TABLE service_plans (
    id varchar(100) PRIMARY KEY NOT NULL,
    name varchar(255) NOT NULL,
    description text NOT NULL DEFAULT '',
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    catalog_name varchar(255) NOT NULL,
    catalog_id varchar(255) NOT NULL,

    free smallint(1) NOT NULL,

    bindable smallint(1) NOT NULL,
    plan_updatable smallint(1) NOT NULL,

    metadata json NOT NULL DEFAULT '{}',
    create_instance_schema json NOT NULL DEFAULT '{}',
    update_instance_schema json NOT NULL DEFAULT '{}',
    create_binding_schema json NOT NULL DEFAULT '{}',

    service_offering_id varchar(100) NOT NULL REFERENCES service_offerings (id) ON DELETE CASCADE,

    UNIQUE (service_offering_id, name),
);

COMMIT;