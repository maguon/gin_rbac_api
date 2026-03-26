CREATE USER yepin_dev WITH PASSWORD 'yepin_dev_2023';
CREATE DATABASE gin_rbac owner yepin_dev;
GRANT ALL ON DATABASE gin_rbac TO yepin_dev; 

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;
-- CREATE FUNCTION UPDATE_TIMESTAMP_FUNC
create or replace function update_timestamp_func() returns trigger as $$ begin new.updated_at = current_timestamp;
return new;
end $$ language plpgsql;
-- ----------------------------
-- Table structure for admin_info
-- ----------------------------
DROP TABLE IF EXISTS "public"."admin_info";
CREATE TABLE "public"."admin_info" (
    "id" bigserial NOT NULL,
    "created_at" timestamp with time zone NOT NULL DEFAULT NOW(),
    "updated_at" timestamp with time zone NOT NULL DEFAULT NOW(),
    "last_login_at" timestamp with time zone NOT NULL DEFAULT NOW(),
    "status" smallint NOT NULL,
    "username" character varying(50),
    "password" character varying(200) NOT NULL,
    "avatar" character varying(200) ,
    "phone" character varying(50) NOT NULL,
    "email" character varying(100),
    "gender" smallint,
    PRIMARY KEY (id)
);
create trigger admin_info_upt before
update on admin_info for each row execute procedure update_timestamp_func();
select setval('admin_info_id_seq', 1000, false);

CREATE UNIQUE INDEX admin_phohe_unique ON admin_info (phone); 


-- ----------------------------
-- Table structure for admin_role
-- ----------------------------
CREATE TABLE "public"."admin_role" (
    "id" bigserial NOT NULL,
    "created_at" timestamp with time zone NOT NULL DEFAULT NOW(),
    "updated_at" timestamp with time zone NOT NULL DEFAULT NOW(),
    "status" smallint NOT NULL DEFAULT 1,
    "role_name" character varying(50) NOT NULL,
    "default_url" character varying(200) NOT NULL,
    "remark" character varying(400),
    PRIMARY KEY (id)
);
create trigger admin_role_upt before
update on admin_role for each row execute procedure update_timestamp_func();
select setval('admin_role_id_seq', 1000, false);

-- ----------------------------
-- Table structure for admin_role_rel
-- ----------------------------
CREATE TABLE "public"."admin_role_rel" (
    "id" bigserial NOT NULL,
    "created_at" timestamp with time zone NOT NULL DEFAULT NOW(),
    "updated_at" timestamp with time zone NOT NULL DEFAULT NOW(),
    "admin_id" int8 NOT NULL,
    "admin_role_id" int8 NOT NULL,
    PRIMARY KEY (id)
);
create trigger admin_role_rel_upt before
update on admin_role_rel for each row execute procedure update_timestamp_func();
CREATE UNIQUE INDEX admin_role_rel_unique ON admin_role_rel (admin_id,admin_role_id); 


-- ----------------------------
-- Table structure for admin_menu
-- ----------------------------
CREATE TABLE "public"."admin_menu" (
    "id" bigserial NOT NULL,
    "created_at" timestamp with time zone NOT NULL DEFAULT NOW(),
    "updated_at" timestamp with time zone NOT NULL DEFAULT NOW(),
    "status" smallint NOT NULL,
    "level" smallint NOT NULL DEFAULT 0,
    "parent_id" int8 DEFAULT NULL,
    "url" character varying(200) DEFAULT NULL,
    "menu_name" character varying(50) NOT NULL,
    "sort" int8 DEFAULT NULL,
    "icon" character varying(50) NOT NULL,
    "component" character varying(50) DEFAULT NULL,
    PRIMARY KEY (id)
);
create trigger admin_menu_upt before
update on admin_menu for each row execute procedure update_timestamp_func();
select setval('admin_menu_id_seq', 1000, false);


-- -------------------------------------
-- Table structure for admin_menu_btn
-- -------------------------------------
CREATE TABLE "public"."admin_menu_btn" (
    "id" bigserial NOT NULL,
    "created_at" timestamp with time zone NOT NULL DEFAULT NOW(),
    "updated_at" timestamp with time zone NOT NULL DEFAULT NOW(),
    "status" smallint NOT NULL,
    "btn_name" character varying(50) NOT NULL,
    "admin_menu_id" int8 NOT NULL,
    "remark" character varying(50) DEFAULT NULL,
    PRIMARY KEY (id)
);
create trigger admin_menu_btn_upt before
update on admin_menu_btn for each row execute procedure update_timestamp_func();
select setval('admin_menu_btn_id_seq', 1000, false);


-- -------------------------------------
-- Table structure for admin_role_menu
-- -------------------------------------
CREATE TABLE "public"."admin_role_menu" (
    "id" bigserial NOT NULL,
    "created_at" timestamp with time zone NOT NULL DEFAULT NOW(),
    "updated_at" timestamp with time zone NOT NULL DEFAULT NOW(),
    "admin_menu_id" int8 NOT NULL,
    "admin_role_id" int8 NOT NULL,
    "menu_btn" jsonb  NULL,
    PRIMARY KEY (id)
);
create trigger admin_role_menu_upt before
update on admin_role_menu for each row execute procedure update_timestamp_func();
select setval('admin_role_menu_id_seq', 1000, false);
CREATE UNIQUE INDEX admin_role_menu_unique ON admin_role_menu (admin_menu_id,admin_role_id); 

-- -------------------------------------
-- Table structure for admin_sys_api
-- -------------------------------------
CREATE TABLE "public"."admin_sys_api" (
    "id" bigserial NOT NULL,
    "created_at" timestamp with time zone NOT NULL DEFAULT NOW(),
    "updated_at" timestamp with time zone NOT NULL DEFAULT NOW(),
    "status" smallint NOT NULL,
    "method" character varying(20) NOT NULL,
    "url" character varying(200) NOT NULL,
    "tag" character varying(50) NOT NULL,
    "remark" character varying(200) DEFAULT NULL,
    PRIMARY KEY (id)
);
create trigger admin_sys_api_upt before
update on admin_sys_api for each row execute procedure update_timestamp_func();
select setval('admin_sys_api_id_seq', 1000, false);


-- -------------------------------------
-- Table structure for admin_casbin
-- -------------------------------------
CREATE TABLE "public"."admin_casbin" (
    "id" bigserial NOT NULL,
    "ptype" character varying(100) DEFAULT NULL,
    "v0" character varying(100) DEFAULT NULL,
    "v1" character varying(100) DEFAULT NULL,
    "v2" character varying(100) DEFAULT NULL,
    "v3" character varying(100) DEFAULT NULL,
    "v4" character varying(100) DEFAULT NULL,
    "v5" character varying(100) DEFAULT NULL,
    PRIMARY KEY (id)
);
select setval('admin_casbin_id_seq', 1000, false);
CREATE UNIQUE INDEX admin_casbin_unique ON admin_casbin (ptype,v0,v1,v2,v3,v4,v5); 

-- ----------------------------
-- Table structure for admin_token
-- ----------------------------
CREATE TABLE "public"."admin_token" (
    "id" bigserial NOT NULL,
    "token" character varying(400),
    PRIMARY KEY (id)
);
create trigger admin_token_upt before
update on admin_token for each row execute procedure update_timestamp_func();
select setval('admin_token_id_seq', 1000, false);

--CREATE TABLE date_base
CREATE TABLE IF NOT EXISTS public.date_base
(
    id integer NOT NULL,
    day integer NOT NULL,
    week integer NOT NULL,
    month integer NOT NULL,
    year integer NOT NULL,
    y_month integer NOT NULL,
    y_week integer NOT NULL,
    PRIMARY KEY (id)
);
SELECT cron.schedule('add_date_sdl', '0 16 * * *', $$insert into date_base (id,day,week,month,year,y_month,y_week) values
( CAST (to_char(current_timestamp, 'YYYYMMDD') AS NUMERIC)  ,
CAST(ltrim(to_char(current_timestamp, 'DD'),'0') AS NUMERIC) ,
CAST(ltrim(to_char(current_timestamp, 'WW'),'0')  AS NUMERIC) ,
CAST(ltrim(to_char(current_timestamp, 'MM'),'0') AS NUMERIC) ,
CAST(to_char(current_timestamp, 'YYYY') AS NUMERIC),
CAST(to_char(current_timestamp, 'YYYYMM') AS NUMERIC) ,
CAST(to_char(current_timestamp, 'YYYYWW') AS NUMERIC));$$);
