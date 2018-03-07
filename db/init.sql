
CREATE TABLE IF NOT EXISTS todos (
  id SERIAL,
  title VARCHAR(255) NOT NULL ,
  complete BOOLEAN NOT NULL DEFAULT FALSE,
  created TIMESTAMP NOT NULL DEFAULT NOW()
);

INSERT INTO todos (title) VALUES ('Create lots of bugs');
INSERT INTO todos (title) VALUES ('Fix lots of bugs');
