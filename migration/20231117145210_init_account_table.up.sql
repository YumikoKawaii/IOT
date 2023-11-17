create table if not exist `accounts` (
    `id` int primary key auto_increment,
    `username` varchar(100) character set utf8 not null,
    `password` varchar(100) not null,
    `created_at` datetime,
    `updated_at` datetime,
    `status` varchar(20)
) engine=InnoDB default charset=utf8mb4 collate=utf8mb4_unicode_ci;