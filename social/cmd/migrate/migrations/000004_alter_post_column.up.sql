alter table posts add column created_at timestamp(0) with time zone not null default now();
alter table posts add column updated_at timestamp(0) with time zone not null default now();
