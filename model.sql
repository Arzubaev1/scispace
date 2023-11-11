CREATE TABLE "users" (
    "id" UUID PRIMARY KEY,
    "fullname" VARCHAR(55) NOT NULL,
    "institution" VARCHAR(55) NOT NULL,
    "department" VARCHAR(55) NOT NULL,
    "degree" VARCHAR(55) NOT NULL,
    "email" VARCHAR(55) NOT NULL,
    "password" VARCHAR(55) NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);
