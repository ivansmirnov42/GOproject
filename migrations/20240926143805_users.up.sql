CREATE TABLE users (
                       ID SERIAL PRIMARY KEY,
                       Email VARCHAR(255) NOT NULL,
                       Password VARCHAR(255) NOT NULL,
                       Deleted_At TIMESTAMP WITH TIME ZONE,
                       Created_At TIMESTAMP WITH TIME ZONE NULL,
                       Updated_At TIMESTAMP WITH TIME ZONE NULL
);