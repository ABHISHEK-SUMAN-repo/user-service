CREATE TABLE "user" (
    "ID" SERIAL PRIMARY KEY,
    "First_name" VARCHAR(100) NOT NULL,
    "Last_name" VARCHAR(100) NOT NULL,
    "Password" VARCHAR(255) NOT NULL,
    "Email" VARCHAR(255) NOT NULL,
    "Phone" VARCHAR(20) NOT NULL,
    "Token" VARCHAR(255),
    "User_type" VARCHAR(10) NOT NULL,
    "Refresh_token" VARCHAR(255),
    "Created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "Updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "Created_by" VARCHAR(100),
    "Updated_by" VARCHAR(100)
);