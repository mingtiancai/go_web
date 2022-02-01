drop database t;
create database t;
drop user gaoming;
create user gaoming with password 'tiancai';
grant all privileges on database t to gaoming;