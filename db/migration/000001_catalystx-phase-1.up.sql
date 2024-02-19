CREATE TABLE "serviceareatypes" (
  "id" bigserial PRIMARY KEY,
  "serviceareatype_name" varchar UNIQUE NOT NULL,
  "serviceareatype_desc" varchar NOT NULL
);

CREATE TABLE "areas" (
  "id" bigserial PRIMARY KEY,
  "area_name" varchar UNIQUE NOT NULL,
  "area_desc" varchar NOT NULL
);

CREATE TABLE "clutters" (
  "id" bigserial PRIMARY KEY,
  "clutter_name" varchar UNIQUE NOT NULL,
  "clutter_desc" varchar NOT NULL
);

CREATE TABLE "sitetypes" (
  "id" bigserial PRIMARY KEY,
  "type_name" varchar UNIQUE NOT NULL,
  "type_desc" varchar NOT NULL
);

CREATE TABLE "vendors" (
  "id" bigserial PRIMARY KEY,
  "vendor_name" varchar UNIQUE NOT NULL,
  "vendor_desc" varchar NOT NULL
);

CREATE TABLE "techs" (
  "id" bigserial PRIMARY KEY,
  "tech_name" varchar UNIQUE NOT NULL,
  "tech_desc" varchar NOT NULL
);

CREATE TABLE "bands" (
  "id" bigserial PRIMARY KEY,
  "band_name" varchar UNIQUE NOT NULL,
  "band_desc" varchar NOT NULL,
  "size" bigint NOT NULL DEFAULT 0,
  "start_freq" bigint NOT NULL DEFAULT 0,
  "end_freq" bigint NOT NULL DEFAULT 0,
  "tech_id" bigint NOT NULL
);

CREATE TABLE "carriers" (
  "id" bigserial PRIMARY KEY,
  "carrier_name" varchar UNIQUE NOT NULL,
  "carrier_desc" varchar NOT NULL,
  "size" bigint NOT NULL DEFAULT 0,
  "start_freq" bigint NOT NULL DEFAULT 0,
  "end_freq" bigint NOT NULL DEFAULT 0,
  "band_id" bigint NOT NULL
);

CREATE TABLE "continents" (
  "id" bigserial PRIMARY KEY,
  "continent_name" varchar UNIQUE NOT NULL,
  "continent_desc" varchar NOT NULL
);

CREATE TABLE "countries" (
  "id" bigserial PRIMARY KEY,
  "country_name" varchar UNIQUE NOT NULL,
  "country_desc" varchar NOT NULL,
  "continent_id" bigint NOT NULL
);

CREATE TABLE "states" (
  "id" bigserial PRIMARY KEY,
  "state_name" varchar UNIQUE NOT NULL,
  "state_desc" varchar NOT NULL,
  "country_id" bigint NOT NULL,
  "area_id" bigint NOT NULL
);

CREATE TABLE "cities" (
  "id" bigserial PRIMARY KEY,
  "city_name" varchar UNIQUE NOT NULL,
  "city_desc" varchar NOT NULL,
  "state_id" bigint NOT NULL
);

CREATE TABLE "districts" (
  "id" bigserial PRIMARY KEY,
  "district_name" varchar UNIQUE NOT NULL,
  "district_desc" varchar NOT NULL,
  "city_id" bigint NOT NULL
);

CREATE TABLE "towns" (
  "id" bigserial PRIMARY KEY,
  "town_name" varchar UNIQUE NOT NULL,
  "town_desc" varchar NOT NULL,
  "district_id" bigint NOT NULL
);

CREATE TABLE "blocks" (
  "id" bigserial PRIMARY KEY,
  "block_name" varchar UNIQUE NOT NULL,
  "block_desc" varchar NOT NULL,
  "total_population" bigint NOT NULL,
  "town_id" bigint NOT NULL,
  "clutter_id" bigint NOT NULL
);

CREATE TABLE "properties" (
  "id" bigserial PRIMARY KEY,
  "property_name" varchar UNIQUE NOT NULL,
  "lat" real NOT NULL DEFAULT 0,
  "long" real NOT NULL DEFAULT 0,
  "block_id" bigint NOT NULL
);

CREATE TABLE "sites" (
  "id" bigserial PRIMARY KEY,
  "site_name" varchar UNIQUE NOT NULL,
  "site_name_old" varchar NOT NULL,
  "site_id_givin" varchar NOT NULL,
  "site_id_givin_old" varchar NOT NULL,
  "lac" varchar NOT NULL,
  "rac" varchar NOT NULL,
  "rnc" varchar NOT NULL,
  "site_on_air_date" timestamptz NOT NULL,
  "property_id" bigint NOT NULL,
  "sitetype_id" bigint NOT NULL,
  "vendor_id" bigint NOT NULL
);

CREATE TABLE "cells" (
  "id" bigserial PRIMARY KEY,
  "cell_name" varchar UNIQUE NOT NULL,
  "cell_name_old" varchar NOT NULL,
  "cell_id_givin" varchar NOT NULL,
  "cell_id_givin_old" varchar NOT NULL,
  "sector_name" varchar NOT NULL,
  "uplinkuarfcn" varchar NOT NULL,
  "downlinkuarfcn" varchar NOT NULL,
  "dlprscramblecode" varchar NOT NULL,
  "azimuth" varchar NOT NULL,
  "height" varchar NOT NULL,
  "etilt" varchar NOT NULL,
  "mtilt" varchar NOT NULL,
  "antennatype" varchar NOT NULL,
  "antennamodel" varchar NOT NULL,
  "ecgi" varchar NOT NULL,
  "site_id" bigint NOT NULL,
  "carrier_id" bigint NOT NULL,
  "serviceareatype_id" bigint NOT NULL
);

CREATE TABLE "traffic" (
  "id" bigserial PRIMARY KEY,
  "traffic_date" timestamptz NOT NULL,
  "avgdailydldatamb" real NOT NULL DEFAULT 0,
  "avgdailyuldatamb" real NOT NULL DEFAULT 0,
  "avgdailytotdatamb" real NOT NULL DEFAULT 0,
  "avgdailytotvoicemin" real NOT NULL DEFAULT 0,
  "avgdailytotvideomin" real NOT NULL DEFAULT 0,
  "qci1_data" real NOT NULL DEFAULT 0,
  "qci6_data" real NOT NULL DEFAULT 0,
  "qci8_data" real NOT NULL DEFAULT 0,
  "qci_other_data" real NOT NULL DEFAULT 0,
  "avgdailytotvoicemin4g" real NOT NULL DEFAULT 0,
  "avgdailytotvoicemintotal" real NOT NULL DEFAULT 0,
  "userdlthroughput" real NOT NULL DEFAULT 0,
  "dlpacketlossrate" real NOT NULL DEFAULT 0,
  "overallpsdropcallrate" real NOT NULL DEFAULT 0,
  "bhdldatamb" real NOT NULL DEFAULT 0,
  "bhupdatamb" real NOT NULL DEFAULT 0,
  "bhtotdatamb" real NOT NULL DEFAULT 0,
  "bhtotvoicemin" real NOT NULL DEFAULT 0,
  "bhtotvideomin" real NOT NULL DEFAULT 0,
  "bhcsusers" real NOT NULL DEFAULT 0,
  "bhhsupausers" real NOT NULL DEFAULT 0,
  "bhhsdpausers" real NOT NULL DEFAULT 0,
  "bhr99uldl" real NOT NULL DEFAULT 0,
  "powercapacity" real NOT NULL DEFAULT 0,
  "powerutilization" real NOT NULL DEFAULT 0,
  "codecapacity" real NOT NULL DEFAULT 0,
  "codeutilization" real NOT NULL DEFAULT 0,
  "ceulcapacity" real NOT NULL DEFAULT 0,
  "ceulutilization" real NOT NULL DEFAULT 0,
  "cedlcapacity" real NOT NULL DEFAULT 0,
  "cedlutilization" real NOT NULL DEFAULT 0,
  "iubcapacity" real NOT NULL DEFAULT 0,
  "iubutlization" real NOT NULL DEFAULT 0,
  "bhrrcusers" real NOT NULL DEFAULT 0,
  "cell_id" bigint NOT NULL
);

CREATE INDEX ON "serviceareatypes" ("serviceareatype_name");

CREATE INDEX ON "areas" ("area_name");

CREATE INDEX ON "clutters" ("clutter_name");

CREATE INDEX ON "sitetypes" ("type_name");

CREATE INDEX ON "vendors" ("vendor_name");

CREATE INDEX ON "techs" ("tech_name");

CREATE INDEX ON "bands" ("band_name");

CREATE INDEX ON "carriers" ("carrier_name");

CREATE INDEX ON "continents" ("continent_name");

CREATE INDEX ON "countries" ("country_name");

CREATE INDEX ON "states" ("state_name");

CREATE INDEX ON "cities" ("city_name");

CREATE INDEX ON "districts" ("district_name");

CREATE INDEX ON "towns" ("town_name");

CREATE INDEX ON "blocks" ("block_name");

CREATE INDEX ON "properties" ("property_name");

CREATE INDEX ON "properties" ("lat");

CREATE INDEX ON "properties" ("long");

CREATE INDEX ON "properties" ("lat", "long");

CREATE INDEX ON "sites" ("site_name");

CREATE INDEX ON "sites" ("site_name_old");

CREATE INDEX ON "sites" ("site_id_givin");

CREATE INDEX ON "sites" ("site_id_givin_old");

CREATE INDEX ON "sites" ("site_on_air_date");

CREATE INDEX ON "cells" ("cell_name");

CREATE INDEX ON "cells" ("cell_name_old");

CREATE INDEX ON "cells" ("cell_id_givin");

CREATE INDEX ON "cells" ("cell_id_givin_old");

CREATE INDEX ON "cells" ("ecgi");

CREATE INDEX ON "traffic" ("cell_id");

ALTER TABLE "bands" ADD FOREIGN KEY ("tech_id") REFERENCES "techs" ("id");

ALTER TABLE "carriers" ADD FOREIGN KEY ("band_id") REFERENCES "bands" ("id");

ALTER TABLE "countries" ADD FOREIGN KEY ("continent_id") REFERENCES "continents" ("id");

ALTER TABLE "states" ADD FOREIGN KEY ("country_id") REFERENCES "countries" ("id");

ALTER TABLE "states" ADD FOREIGN KEY ("area_id") REFERENCES "areas" ("id");

ALTER TABLE "cities" ADD FOREIGN KEY ("state_id") REFERENCES "states" ("id");

ALTER TABLE "districts" ADD FOREIGN KEY ("city_id") REFERENCES "cities" ("id");

ALTER TABLE "towns" ADD FOREIGN KEY ("district_id") REFERENCES "districts" ("id");

ALTER TABLE "blocks" ADD FOREIGN KEY ("town_id") REFERENCES "towns" ("id");

ALTER TABLE "blocks" ADD FOREIGN KEY ("clutter_id") REFERENCES "clutters" ("id");

ALTER TABLE "properties" ADD FOREIGN KEY ("block_id") REFERENCES "blocks" ("id");

ALTER TABLE "sites" ADD FOREIGN KEY ("property_id") REFERENCES "properties" ("id");

ALTER TABLE "sites" ADD FOREIGN KEY ("sitetype_id") REFERENCES "sitetypes" ("id");

ALTER TABLE "sites" ADD FOREIGN KEY ("vendor_id") REFERENCES "vendors" ("id");

ALTER TABLE "cells" ADD FOREIGN KEY ("site_id") REFERENCES "sites" ("id");

ALTER TABLE "cells" ADD FOREIGN KEY ("carrier_id") REFERENCES "carriers" ("id");

ALTER TABLE "cells" ADD FOREIGN KEY ("serviceareatype_id") REFERENCES "serviceareatypes" ("id");

ALTER TABLE "traffic" ADD FOREIGN KEY ("cell_id") REFERENCES "cells" ("id");
