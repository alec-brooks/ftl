-- migrate:up
CREATE
    EXTENSION IF NOT EXISTS pgcrypto;

-- Function for deployment notifications.
CREATE OR REPLACE FUNCTION notify_deployment_event() RETURNS TRIGGER AS
$$
DECLARE
    payload JSONB;
BEGIN
    IF TG_OP = 'DELETE'
    THEN
        payload = jsonb_build_object(
                'table', TG_TABLE_NAME,
                'action', TG_OP,
                'old', old.name
            );
    ELSE
        payload = jsonb_build_object(
                'table', TG_TABLE_NAME,
                'action', TG_OP,
                'new', new.name
            );
    END IF;
    PERFORM pg_notify('notify_events', payload::text);
    RETURN NULL;
END;
$$ LANGUAGE plpgsql;

-- Function for runner notifications because we need the deleted row
-- in order to find the module name, for deleting routing table events.
CREATE OR REPLACE FUNCTION notify_routing_event() RETURNS TRIGGER AS
$$
DECLARE
    payload JSONB;
BEGIN
    IF TG_OP = 'DELETE'
    THEN
        payload = jsonb_build_object(
                'table', TG_TABLE_NAME,
                'action', TG_OP,
                'old', old.key
            );
    ELSE
        payload = jsonb_build_object(
                'table', TG_TABLE_NAME,
                'action', TG_OP,
                'new', new.key
            );
    END IF;
    PERFORM pg_notify('notify_events', payload::text);
    RETURN NULL;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE modules
(
    id       BIGINT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    language VARCHAR        NOT NULL,
    name     VARCHAR UNIQUE NOT NULL
);

CREATE TABLE deployments
(
    id           BIGINT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    created_at   TIMESTAMPTZ    NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc'),
    module_id    BIGINT         NOT NULL REFERENCES modules (id) ON DELETE CASCADE,
    -- Unique name for this deployment in the form <module-name>-<random>.
    "name"       VARCHAR UNIQUE NOT NULL,
    -- Proto-encoded module schema.
    "schema"     BYTEA          NOT NULL,
    -- Labels are used to match deployments to runners.
    "labels"     JSONB          NOT NULL DEFAULT '{}',
    min_replicas INT            NOT NULL DEFAULT 0
);

CREATE UNIQUE INDEX deployments_name_idx ON deployments (name);
CREATE INDEX deployments_module_id_idx ON deployments (module_id);
-- Only allow one deployment per module.
CREATE UNIQUE INDEX deployments_unique_idx ON deployments (module_id)
    WHERE min_replicas > 0;

CREATE TRIGGER deployments_notify_event
    AFTER INSERT OR UPDATE OR DELETE
    ON deployments
    FOR EACH ROW
EXECUTE PROCEDURE notify_deployment_event();

CREATE TABLE artefacts
(
    id         BIGINT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc'),
    -- SHA256 digest of the content.
    digest     BYTEA UNIQUE NOT NULL,
    content    BYTEA        NOT NULL
);

CREATE UNIQUE INDEX artefacts_digest_idx ON artefacts (digest);

CREATE TABLE deployment_artefacts
(
    artefact_id   BIGINT      NOT NULL REFERENCES artefacts (id) ON DELETE CASCADE,
    deployment_id BIGINT      NOT NULL REFERENCES deployments (id) ON DELETE CASCADE,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc'),
    executable    BOOLEAN     NOT NULL,
    -- Path relative to the module root.
    path          VARCHAR     NOT NULL
);

CREATE INDEX deployment_artefacts_deployment_id_idx ON deployment_artefacts (deployment_id);

CREATE TYPE runner_state AS ENUM (
    -- The Runner is available to run deployments.
    'idle',
    -- The Runner is reserved but has not yet deployed.
    'reserved',
    -- The Runner has been assigned a deployment.
    'assigned',
    -- The Runner is dead.
    'dead'
    );

-- Runners are processes that are available to run modules.
CREATE TABLE runners
(
    id                  BIGINT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    -- Unique identifier for this runner, generated at startup.
    key                 UUID UNIQUE  NOT NULL,
    created             TIMESTAMPTZ  NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc'),
    last_seen           TIMESTAMPTZ  NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc'),
    -- If the runner is reserved, this is the time at which the reservation expires.
    reservation_timeout TIMESTAMPTZ,
    state               runner_state NOT NULL DEFAULT 'idle',
    endpoint            VARCHAR      NOT NULL,
    -- Some denormalisation for performance. Without this we need to do a two table join.
    module_name         VARCHAR,
    deployment_id       BIGINT       REFERENCES deployments (id) ON DELETE SET NULL,
    labels              JSONB        NOT NULL DEFAULT '{}'
);

CREATE UNIQUE INDEX runners_key ON runners (key);
CREATE UNIQUE INDEX runners_endpoint_not_dead_idx ON runners (endpoint) WHERE state <> 'dead';
CREATE INDEX runners_module_name_idx ON runners (module_name);
CREATE INDEX runners_state_idx ON runners (state);
CREATE INDEX runners_deployment_id_idx ON runners (deployment_id);
CREATE INDEX runners_labels_idx ON runners USING GIN (labels);

CREATE TRIGGER runners_update_notify_event
    AFTER UPDATE
    ON runners
    FOR EACH ROW
    WHEN (OLD.state <> NEW.state OR
          OLD.deployment_id <> NEW.deployment_id OR
          OLD.labels <> NEW.labels OR
          OLD.module_name <> NEW.module_name)
EXECUTE PROCEDURE notify_routing_event();

CREATE TRIGGER runners_notify_event
    AFTER INSERT OR DELETE
    ON runners
    FOR EACH ROW
EXECUTE PROCEDURE notify_routing_event();

CREATE TABLE ingress_routes
(
    method        VARCHAR NOT NULL,
    path          VARCHAR NOT NULL,
    -- The deployment that should handle this route.
    deployment_id BIGINT  NOT NULL REFERENCES deployments (id) ON DELETE CASCADE,
    -- Duplicated here to avoid having to join from this to deployments then modules.
    module        VARCHAR NOT NULL,
    verb          VARCHAR NOT NULL
);

CREATE INDEX ingress_routes_method_path_idx ON ingress_routes (method, path);

-- Inbound requests.
CREATE TABLE ingress_requests
(
    id          BIGINT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    key         UUID UNIQUE NOT NULL,
    source_addr VARCHAR     NOT NULL
);

CREATE UNIQUE INDEX ingress_requests_key_idx ON ingress_requests (key);

CREATE TYPE controller_state AS ENUM (
    'live',
    'dead'
    );

CREATE TABLE controller
(
    id        BIGINT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    key       UUID UNIQUE      NOT NULL,
    created   TIMESTAMPTZ      NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc'),
    last_seen TIMESTAMPTZ      NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc'),
    state     controller_state NOT NULL DEFAULT 'live',
    endpoint  VARCHAR          NOT NULL
);

CREATE UNIQUE INDEX controller_endpoint_not_dead_idx ON controller (endpoint) WHERE state <> 'dead';

CREATE TYPE event_type AS ENUM (
    'call',
    'log',
    'deployment'
    );

CREATE TABLE events
(
    time_stamp    TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc'),

    deployment_id BIGINT      NOT NULL REFERENCES deployments (id) ON DELETE CASCADE,
    request_id    BIGINT      NULL REFERENCES ingress_requests (id) ON DELETE CASCADE,

    type          event_type  NOT NULL,

    -- Type-specific keys used to index events for searching.
    custom_key_1  VARCHAR     NULL,
    custom_key_2  VARCHAR     NULL,
    custom_key_3  VARCHAR     NULL,
    custom_key_4  VARCHAR     NULL,

    payload       JSONB       NOT NULL
);

CREATE INDEX events_timestamp_idx ON events (time_stamp);
CREATE INDEX events_deployment_id_idx ON events (deployment_id);
CREATE INDEX events_request_id_idx ON events (request_id);
CREATE INDEX events_type_idx ON events (type);
CREATE INDEX events_custom_key_1_idx ON events (custom_key_1);
CREATE INDEX events_custom_key_2_idx ON events (custom_key_2);
CREATE INDEX events_custom_key_3_idx ON events (custom_key_3);
CREATE INDEX events_custom_key_4_idx ON events (custom_key_4);

-- migrate:down