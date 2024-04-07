CREATE TABLE videos (
    id uuid PRIMARY KEY,
    title TEXT,
    description TEXT,
    publish_time TIMESTAMP,
    thumbnails TEXT[]
);
