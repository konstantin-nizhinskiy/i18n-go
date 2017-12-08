CREATE OR REPLACE FUNCTION public.partitioning_lang()
  RETURNS trigger AS
$BODY$
BEGIN
  BEGIN
    EXECUTE 'INSERT INTO '||TG_TABLE_NAME||'_'||NEW.lang||' SELECT $1.*' USING NEW;
    EXCEPTION
    WHEN undefined_table THEN
      EXECUTE 'CREATE TABLE '||TG_TABLE_NAME||'_'||NEW.lang||' (CONSTRAINT '||TG_TABLE_NAME||'_'||NEW.lang||'_pkey PRIMARY KEY (id),CHECK ( lang='''||NEW.lang||''' ))INHERITS ('||TG_TABLE_NAME||')';
      EXECUTE 'INSERT INTO '||TG_TABLE_NAME||'_'||NEW.lang||' SELECT $1.*' USING NEW;
  END;
  RETURN NULL;
END;
$BODY$
LANGUAGE plpgsql VOLATILE
COST 100;
ALTER FUNCTION public.partitioning_lang()
OWNER TO postgres;


CREATE TABLE public.error_translate
(
  id character varying(500) NOT NULL,
  lang character varying(2) NOT NULL,
  txt_translate character varying(5000) NOT NULL,
  date_cr timestamp without time zone NOT NULL DEFAULT now(),
  CONSTRAINT pk_error_translate PRIMARY KEY (id)
)
WITH (
  OIDS=FALSE
);
ALTER TABLE public.error_translate
  OWNER TO postgres;

-- Trigger: error_translate_trigger_p on public.error_translate

-- DROP TRIGGER error_translate_trigger_p ON public.error_translate;

CREATE TRIGGER error_translate_trigger_p
  BEFORE INSERT
  ON public.error_translate
  FOR EACH ROW
  EXECUTE PROCEDURE public.partitioning_lang();
