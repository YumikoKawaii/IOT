create table if not exist devices (
    `id` int primary key auto_increment,
    `type` varchar(30) not null,
    `code` varchar(30) not null,
    `owner` varchar(100) not null,
    `topic` varchar(100) not null
) engine=InnoDB default charset=utf8mb4 collate=utf8mb4_unicode_ci;