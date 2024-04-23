CREATE TABLE "user" (
  "username" varchar(255) PRIMARY KEY,
  "email" varchar(255) NOT NULL,
  "fullname" varchar(255) NOT NULL,
  "password" varchar(255) NOT NULL,
  "role" varchar(50) NOT NULL,
  "enable" INT NOT NULL,
  "created_datetime" timestamptz NOT NULL DEFAULT (now()),
  "updated_datetime" DATE DEFAULT (now())
);

CREATE TABLE "user_connection" (
  "id" INT PRIMARY KEY,
  "username" varchar(255) NOT NULL,
  "last_connection_date_time" DATE NOT NULL,
  "last_action_date_time" DATE NOT NULL,
  "logged_on_date_time" DATE NOT NULL,
  "expired_date_time" DATE NOT NULL,
  "session_key" varchar(255) NOT NULL,
  "client_device" varchar(255) NOT NULL,
  "created_datetime" DATE NOT NULL,
  "updated_datetime" DATE NOT NULL
);

CREATE TABLE "user_pool" (
  "username" varchar(255) NOT NULL PRIMARY KEY,
  "login_retry_count" INT NOT NULL,
  "duplicated_login_count" INT NOT NULL
);

ALTER TABLE "user_connection" ADD FOREIGN KEY ("username") REFERENCES "user" ("username");

CREATE INDEX ON "user_connection" ("username");

CREATE INDEX ON "user_pool" ("username");
