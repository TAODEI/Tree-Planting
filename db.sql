Create database TP;
Use TP

create table users(
  id          int not null auto_increment ,
  student_id  varchar(100)	not null ,
  content     text ,
  primary key (id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;

