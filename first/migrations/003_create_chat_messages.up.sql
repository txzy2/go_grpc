CREATE TABLE IF NOT EXISTS chat_messages (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    ext_id VARCHAR(255) NOT NULL,
    chat_id INTEGER NOT NULL REFERENCES chats(id) ON DELETE CASCADE,
    msg TEXT NOT NULL,
    from_user_id INTEGER NOT NULL REFERENCES users(id),
    to_user_id INTEGER NOT NULL REFERENCES users(id),
    create_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    update_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_chat_messages_chat_id ON chat_messages(chat_id);
