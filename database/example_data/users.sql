-- Add example users
INSERT INTO
    users (username, salt, password_hash, email)
VALUES
    (
        'admin',
        'salt',
        'password',
        'something@something.com'
    );

INSERT INTO
    users (username, salt, password_hash, email)
VALUES
    (
        'user',
        'salt',
        'password',
        'another@something.com'
    );
