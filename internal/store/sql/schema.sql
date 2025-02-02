CREATE TABLE IF NOT EXISTS receipt (
  id INTEGER PRIMARY KEY,
  item_qty INTEGER,
  item_id INTEGER UNIQUE,
  user_id INTEGER,
  CONSTRAINT fk_item FOREIGN KEY (item_id)
    REFERENCES items(id)
);

CREATE TABLE IF NOT EXISTS receipts_users (
  receipt_id INTEGER NOT NULL,
  user_id INTEGER NOT NULL,
  UNIQUE(receipt_id, user_id)
);

CREATE TABLE IF NOT EXISTS items (
  id INTEGER PRIMARY KEY,
  item TEXT NOT NULL,
  price REAL NOT NULL
);

CREATE TABLE IF NOT EXISTS users (
  id INTEGER PRIMARY KEY,
  name TEXT NOT NULL
);
