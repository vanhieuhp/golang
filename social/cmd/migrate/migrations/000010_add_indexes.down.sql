drop index if exists idx_comments_content;
drop index if exists idx_posts_title;
drop index if exists idx_posts_tag;
drop index if exists idx_users_username;
drop index if exists idx_posts_user_id;
drop index if exists idx_comments_post_id;
drop extension if exists pg_trgm;