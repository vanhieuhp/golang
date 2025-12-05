alter table posts
add constraint fk_user foreign key (user_id) references users(id) on delete cascade;