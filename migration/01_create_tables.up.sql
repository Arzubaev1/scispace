

CREATE TABLE "users" (
    "id" UUID PRIMARY KEY,
    "fullname" VARCHAR(50) NOT NULL,
    "institution" VARCHAR(50) NOT NULL,
    "department" VARCHAR(50) NOT NULL,
    "degree" VARCHAR(50) NOT NULL,
    "email" VARCHAR(50) NOT NULL,
    "password" VARCHAR(50) NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);


CREATE TABLE "questions" (
    "id" UUID PRIMARY KEY,
    "title" VARCHAR NOT NULL,
    "context" VARCHAR NOT NULL,
    "tags" VARCHAR,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);
