CREATE TABLE IF NOT EXISTS history (
  admin_id serial PRIMARY KEY ,
  form_id serial,
  comments text NOT NULL,
  edit_made TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
);