CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- TODO:debt implement by specification
CREATE DOMAIN "emoji" AS TEXT;

CREATE TABLE "board"
(
    "id"          UUID      NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    "title"       TEXT      NOT NULL,
    "emoji"       EMOJI     NULL                 DEFAULT NULL,
    "description" TEXT      NULL                 DEFAULT NULL,
    "created_at"  TIMESTAMP NOT NULL             DEFAULT now(),
    "updated_at"  TIMESTAMP NULL                 DEFAULT NULL,
    "deleted_at"  TIMESTAMP NULL                 DEFAULT NULL
);

CREATE TABLE "column"
(
    "id"          UUID      NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    "board_id"    UUID      NOT NULL,
    "title"       TEXT      NOT NULL,
    "emoji"       EMOJI     NULL                 DEFAULT NULL,
    "description" TEXT      NULL                 DEFAULT NULL,
    "created_at"  TIMESTAMP NOT NULL             DEFAULT now(),
    "updated_at"  TIMESTAMP NULL                 DEFAULT NULL,
    "deleted_at"  TIMESTAMP NULL                 DEFAULT NULL
);

CREATE TABLE "card"
(
    "id"          UUID      NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    "column_id"   UUID      NOT NULL,
    "title"       TEXT      NOT NULL,
    "emoji"       EMOJI     NULL                 DEFAULT NULL,
    "description" TEXT      NULL                 DEFAULT NULL,
    "created_at"  TIMESTAMP NOT NULL             DEFAULT now(),
    "updated_at"  TIMESTAMP NULL                 DEFAULT NULL,
    "deleted_at"  TIMESTAMP NULL                 DEFAULT NULL
);
