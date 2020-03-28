--
-- init database sql
DROP DATABASE IF EXISTS dev;
create database dev default character set utf8 collate utf8_general_ci;

grant all on dev.* to dev_rw@'%' identified by 'dev';
