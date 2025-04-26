-- Create the database if it doesn't exist
IF NOT EXISTS (SELECT * FROM sys.databases WHERE name = 'testdb')
BEGIN
    CREATE DATABASE testdb;
END
GO

USE testdb;
GO

-- Create Users table if it doesn't exist
IF NOT EXISTS (SELECT * FROM sys.tables WHERE name = 'users')
BEGIN
    CREATE TABLE users (
        id INT IDENTITY(1,1) PRIMARY KEY,
        name NVARCHAR(100) NOT NULL,
        email NVARCHAR(100) NOT NULL,
        created_at DATETIME2 NOT NULL,
        updated_at DATETIME2 NOT NULL
    );
END
GO

-- Insert some sample data
INSERT INTO users (name, email, created_at, updated_at)
VALUES 
    ('Sample User', 'sample@example.com', GETDATE(), GETDATE()),
    ('Another User', 'another@example.com', GETDATE(), GETDATE());
GO 