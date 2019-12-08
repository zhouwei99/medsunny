--  create DB
CREATE DATABASE medsunny
    WITH 
    OWNER = zhouwei
    ENCODING = 'UTF8'
    LC_COLLATE = 'Chinese (Simplified)_China.936'
    LC_CTYPE = 'Chinese (Simplified)_China.936'
    CONNECTION LIMIT = -1;


-- create table
-- sick ness
CREATE SEQUENCE public.sickness_pk_increase
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 99999999999999
    CACHE 1;

ALTER SEQUENCE public.sickness_pk_increase
    OWNER TO zhouwei;

-- Table: public.data_sickness

CREATE TABLE if not exists public.data_sickness
(
    id integer NOT NULL DEFAULT nextval('sickness_pk_increase'::regclass),
    uuid character(64)[] COLLATE pg_catalog."default" NOT NULL,
    name character varying(200)[] COLLATE pg_catalog."default" NOT NULL,
    symptom character varying(1000)[] COLLATE pg_catalog."default" NOT NULL,
    nursing character varying(500)[] COLLATE pg_catalog."default",
    medicals character varying(200)[] COLLATE pg_catalog."default",
    duration character varying(300)[] COLLATE pg_catalog."default",
    risks character varying(500)[] COLLATE pg_catalog."default",
    create_at timestamp without time zone NOT NULL,
    CONSTRAINT data_sickness_pkey PRIMARY KEY (id),
    CONSTRAINT data_sickness_name_key UNIQUE (name)
,
    CONSTRAINT data_sickness_uuid_key UNIQUE (uuid)

)

TABLESPACE pg_default;

ALTER TABLE public.data_sickness
    OWNER to zhouwei;