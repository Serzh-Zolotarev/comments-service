DROP TABLE IF EXISTS comments;

CREATE TABLE comments (
    id SERIAL PRIMARY KEY,
    news_id INT  NOT NULL,
    comment_id INT,
    content TEXT NOT NULL,
    author TEXT NOT NULL,
    pub_time BIGINT NOT NULL
);

INSERT INTO comments (id, news_id, comment_id, content, author ,pub_time)
VALUES (1, 1, 0, 'Some nice comment', 'Buddy',0);