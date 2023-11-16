drop table if exists `user`;
create table `user` (
                        `id` bigint(20) not null auto_increment,
                        `user_id` bigint(20) not null,
                        `username` varchar(64) collate utf8mb4_general_ci not null,
                        `password` varchar(64) collate utf8mb4_general_ci not null,
                        `email` varchar(64) collate utf8mb4_general_ci,
                        `gender` tinyint(4) not null default '0',
                        `create_time` timestamp null default current_timestamp,
                        `update_time` timestamp null default current_timestamp on update current_timestamp,
                        primary key (`id`),
                        unique key `idx_username` (`username`) using btree,
                        unique key `idx_user_id` (`user_id`) using btree
)engine=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

drop table if exists `community`;
create table `community` (
                        `id` bigint(20) not null auto_increment,
                        `community_id` bigint(20) not null,
                        `community_name` varchar(128) collate utf8mb4_general_ci not null,
                        `introduction` varchar(256) collate utf8mb4_general_ci not null,
                        `create_time` timestamp null default current_timestamp,
                        `update_time` timestamp null default current_timestamp on update current_timestamp,
                        primary key (`id`),
                        unique key `idx_community` (`community_id`) using btree,
                        unique key `idx_community_name` (`community_name`) using btree
)engine=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

insert into `community` (id,community_id, community_name, introduction) VALUES ('0','0','cpp','hello cpp');

select community_id,community_name,introduction,create_time from community where community_id=1

drop table if exists `post`;
create table `post` (
                        `id` bigint(20) not null auto_increment,
                        `post_id` bigint(20) not null,
                        `title` varchar(128) collate utf8mb4_general_ci not null,
                        `content` varchar(8192) collate utf8mb4_general_ci not null,
                        `author_id` bigint(20) collate utf8mb4_general_ci,
                        `community_id` bigint(20) not null default '0',
                        `status` int not null default '0',
                        `create_time` timestamp null default current_timestamp,
                        `update_time` timestamp null default current_timestamp on update current_timestamp,
                        primary key (`id`),
                        unique key `idx_post_id` (`post_id`) using btree
)engine=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


select post_id,title
from post
where id in (46164855144579073)
order by FIND_IN_SET(post_id,46164855144579073)

select post_id,title,content,author_id,community_id from post where post_id=46031311390900225;