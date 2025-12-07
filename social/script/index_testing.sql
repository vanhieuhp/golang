do
$$
    begin
        for r in 1..100000
            loop
                insert into comments (post_id, user_id, content)
                values (129, 197,
                        'Super OMG! this is amazing post!');
            end loop;
    end;
$$;

select count(*)
from comments
where post_id = 129
  and user_id = 197;

explain analyse
select *
from comments
where content ILIKE '%post%';

CREATE EXTENSION IF NOT EXISTS pg_trgm;
create index idx_comments_content on comments using gin (content gin_trgm_ops);