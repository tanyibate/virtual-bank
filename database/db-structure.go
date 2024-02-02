package database

const CreateQuery string = `
  CREATE TABLE IF NOT EXISTS accounts (
  account_number INTEGER NOT NULL PRIMARY KEY,
  name VARCHAR(255),
  pin INTEGER,
  balance

  description TEXT
  );`
