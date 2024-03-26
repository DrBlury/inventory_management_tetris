-- Add example items
INSERT INTO
  items (
    variant,
    name,
    description,
    buy_value,
    sell_value,
    max_stack,
    height,
    width,
    rawshape,
    type
  )
VALUES
  (
    'sword',
    'Sword',
    'A sword',
    10.0,
    5.0,
    1,
    3,
    3,
    '..#.#.#..',
    'weapon'
  );

INSERT INTO
  items (
    variant,
    name,
    description,
    buy_value,
    sell_value,
    max_stack,
    height,
    width,
    rawshape,
    type
  )
VALUES
  (
    'shield',
    'Shield',
    'A shield',
    10.0,
    5.0,
    1,
    3,
    3,
    '..#####..',
    'armor'
  );
