
CREATE TABLE "ish_joyi"(
    "id" UUID PRIMARY KEY,
    "name" VARCHAR(50) NOT NULL
)
CREATE TABLE "mutahassislik"(
    "id" UUID PRIMARY KEY,
    "name" VARCHAR(50) NOT NULL
)
CREATE TABLE "fan_tarmogi"(
    "id" UUID PRIMARY KEY,
    "name" VARCHAR(50) NOT NULL
)
CREATE TABLE "mavzu"(
    "id" UUID PRIMARY KEY,
    "name" VARCHAR(50) NOT NULL
)
CREATE TABLE "oqituvchi"(
    "id" UUID PRIMARY KEY,
    "first_name"  VARCHAR(50) NOT NULL,
    "last_name"  VARCHAR(50) NOT NULL,
    "middle_name"  VARCHAR(50) NOT NULL,
    "date_of_birth"  VARCHAR(50) NOT NULL,
    "ish_joyi"  UUID NOT NULL,
    "mutahassislik"  UUID  NOT NULL,
    "fan_tarmogi"  UUID  NOT NULL,
    "mavzular"  UUID  NOT NULL,
    "email"  VARCHAR(50) NOT NULL,
    "password"  VARCHAR(50) NOT NULL,
    "phone_number"  VARCHAR(50) NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP 
)

CREATE TABLE "tadqiqotchi"(
    "id" UUID PRIMARY KEY,
    "first_name"  VARCHAR(50) NOT NULL,
    "last_name"  VARCHAR(50) NOT NULL,
    "middle_name"  VARCHAR(50) NOT NULL,
    "date_of_birth"  VARCHAR(50) NOT NULL,
    "oqish_joyi"  UUID NOT NULL,
    "fan_tarmogi"  UUID  NOT NULL,
    "mavzu"  UUID  NOT NULL,
    "email"  VARCHAR(50) NOT NULL,
    "password"  VARCHAR(50) NOT NULL,
    "phone_number"  VARCHAR(50) NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP 
)

CREATE TABLE "other"(
    "id" UUID PRIMARY KEY,
    "first_name"  VARCHAR(50) NOT NULL,
    "last_name"  VARCHAR(50) NOT NULL,
    "middle_name"  VARCHAR(50) NOT NULL,
    "date_of_birth"  VARCHAR(50) NOT NULL,
    "oqish_joyi"  UUID NOT NULL,
    "yonalish"  UUID  NOT NULL,
    "mavzular"  UUID  NOT NULL,
    "email"  VARCHAR(50) NOT NULL,
    "password"  VARCHAR(50) NOT NULL,
    "phone_number"  VARCHAR(50) NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP 
)

CREATE TABLE "questions" (
    "id" UUID PRIMARY KEY,
    "title" VARCHAR NOT NULL,
    "context" VARCHAR NOT NULL,
    "tags" VARCHAR,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);


CREATE TABLE "posts" (
    "id" UUID PRIMARY KEY,
    "title" VARCHAR NOT NULL,
    "context" VARCHAR NOT NULL,
    "link" VARCHAR,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);


CREATE TABLE "reports" (
    "id" UUID PRIMARY KEY,
    "report_status" VARCHAR NOT NULL,
    "description" VARCHAR,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);


CREATE TABLE "tools" (
    "id" UUID PRIMARY KEY,
    "tool_name" VARCHAR NOT NULL,
    "link" VARCHAR,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);


CREATE TABLE "databases" (
    "id" UUID PRIMARY KEY,
    "database_name" VARCHAR NOT NULL,
    "link" VARCHAR,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);
CREATE TABLE "maqola"(
    "id" UUID PRIMARY KEY,
    "name" VARCHAR NOT NULL,
    "tavsifi" TEXT NOT NULL,
    "qoshimcha_linklar" VARCHAR NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
)