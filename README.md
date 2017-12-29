work with postgresql database

1.create data struct in pgbadmin4

-- create table userinfo.

CREATE TABLE userinfo
(
    id bigserial primary key,
    user_name character varying(20) COLLATE pg_catalog."default",
    user_password text COLLATE pg_catalog."default",
    email character varying(20) COLLATE pg_catalog."default",
    enabled boolean DEFAULT false,
    last_password_reset timestamp with time zone DEFAULT now()
)

-- insert some records

insert into userinfo(user_name,user_password,email) values('chris',md5('123456'),'chris@163.com'),
('peter',md5('123456'),'peter@163.com')

2.init prog

proj dir: /opt/go_prog

3.create dir:

cd /opt/go_prog
mkdir bin
mkdir src
mkdir pgk
cd src

4.get prog:
git clone https://github.com/longtian001/golang_prog_struct

5.install golang module:
cd app
export GOPATH=/opt/go_prog
go get -u github.com/golang/dep/cmd/dep
/opt/go_prog/bin/dep init
/opt/go_prog/bin/dep ensure

6.run prog:
go run main.go
