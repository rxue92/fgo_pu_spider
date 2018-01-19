-- Table: public.fgo_pu

-- DROP TABLE public.fgo_pu;

CREATE TABLE public.fgo_pu
(
  id integer NOT NULL DEFAULT nextval('fgo_pu_id_seq'::regclass),
  rid bigint NOT NULL,
  sname character varying(255) NOT NULL DEFAULT ''::character varying,
  rname character varying(64) NOT NULL DEFAULT ''::character varying,
  servant character varying(64) NOT NULL DEFAULT ''::character varying,
  rarity integer NOT NULL,
  create_at integer NOT NULL,
  main_up_5 character(64) NOT NULL,
  main_up_4 character(64) NOT NULL
)
WITH (
  OIDS=FALSE
);
ALTER TABLE public.fgo_pu
  OWNER TO postgres;
