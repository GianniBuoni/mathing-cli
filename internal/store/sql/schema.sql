CREATE TABLE IF NOT EXISTS receipt (
  id INTEGER PRIMARY KEY,
  item_qty INTEGER NOT NULL,
  item_id INTEGER UNIQUE NOT NULL,
  CONSTRAINT fk_item FOREIGN KEY (item_id)
    REFERENCES items(id)
    ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS receipts_users (
  receipt_id INTEGER NOT NULL,
  user_id INTEGER NOT NULL,
  CONSTRAINT fk_receipt FOREIGN KEY (receipt_id)
    REFERENCES receipt(id)
    ON DELETE CASCADE,
  CONSTRAINT fk_user FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE CASCADE,
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

PRAGMA foreign_keys = ON;
