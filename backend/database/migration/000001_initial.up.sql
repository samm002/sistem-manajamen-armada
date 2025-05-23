-- Create "vehicle_locations" table
CREATE TABLE "public"."vehicle_locations" (
  "vehicle_id" character varying(10) NOT NULL,
  "latitude" numeric NOT NULL,
  "longitude" numeric NOT NULL,
  "timestamp" bigint NOT NULL
);
-- Create index "idx_vehicleid_timestamp" to table: "vehicle_locations"
CREATE UNIQUE INDEX "idx_vehicleid_timestamp" ON "public"."vehicle_locations" ("vehicle_id", "timestamp");
