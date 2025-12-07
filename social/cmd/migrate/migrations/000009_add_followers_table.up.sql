create table if not exists followers
(
    user_id     bigserial                              not null,
    follower_id bigserial                              not null,
    created_at  timestamp with time zone default now() not null,
    primary key (user_id, follower_id),
    constraint fk_followers_user_id foreign key (user_id) references users (id) on delete cascade,
    constraint fk_followers_follower_id foreign key (follower_id) references users (id) on delete cascade
)