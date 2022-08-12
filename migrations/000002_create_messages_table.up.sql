CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS messages (
    id text UNIQUE NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
    recipient_id text NOT NULL,
    sender_id text NOT NULL,
    message_type text default 'message',
    message_text text NOT NULL,
    created_at timestamp with time zone default CURRENT_TIMESTAMP
);

-- Seed data
INSERT INTO messages (
    id, sender_id, recipient_id, message_text
) VALUES (
    uuid_generate_v4(), '2', '1', 'Hey Caro! How are you?'
);

INSERT INTO messages (
    id, sender_id, recipient_id, message_text
) VALUES (
    uuid_generate_v4(), '2', '1', 'I just got back from vacation.'
);

INSERT INTO messages (
    id, sender_id, recipient_id, message_text
) VALUES (
    uuid_generate_v4(), '1', '2', 'Barb! Doing well! How are you?'
);

INSERT INTO messages (
    id, sender_id, recipient_id, message_text
) VALUES (
    uuid_generate_v4(), '1', '3', 'Mark, I have got a question for you.'
);

INSERT INTO messages (
    id, sender_id, recipient_id, message_text
) VALUES (
    uuid_generate_v4(), '3', '1', 'Hi Caroline. Go for it.'
);

INSERT INTO messages (
    id, sender_id, recipient_id, message_text, message_type
) VALUES (
    uuid_generate_v4(), '3', '3', 'Remind me to send Barb an email next week.', 'reminder'
);
