INSERT INTO items (
  id, item, price
) VALUES ( 0, "Banana", 0.49 );

INSERT INTO items (
  id, item, price
) VALUES ( 1, "Popcorn with herbs", 3.49 );

INSERT INTO items (
  id, item, price
  ) VALUES ( 2, "Peeled tomatoes", 2.49 );

INSERT INTO items (
  id, item, price
) VALUES ( 3, "Tomato paste", 1.99 );

INSERT INTO users (
  id, name
) VALUES ( 0, "Jon" );

INSERT INTO users (
  id, name
) VALUES ( 1, "Paul" );

INSERT INTO receipt (
  item_id, item_qty, user_id
) VALUES ( 1, 2, 2 );

INSERT INTO receipt (
  item_id, item_qty, user_id
) VALUES ( 3, 1, 1 );

INSERT INTO receipt (
  item_id, item_qty, user_id
) VALUES ( 0, 1, 0 );
