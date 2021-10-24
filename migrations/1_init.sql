package migration

-- +migrate Up
CREATE TABLE "public"."users_booking" (
  "user_id" text COLLATE "pg_catalog"."default" NOT NULL,
  "username" text COLLATE "pg_catalog"."default",
  "password" text COLLATE "pg_catalog"."default",
  "date_booking" text COLLATE "pg_catalog"."default",
  "time_booking" text COLLATE "pg_catalog"."default",
  "time_booking_backup" text COLLATE "pg_catalog"."default",
  "amenities" text COLLATE "pg_catalog"."default",
  "amenity" text COLLATE "pg_catalog"."default",
  "homes" text COLLATE "pg_catalog"."default",
  "home_id" text COLLATE "pg_catalog"."default",
  "created_booking" timestamptz(6),
  "completed_booking" timestamptz(6)
)
;


INSERT INTO "public"."users_booking" VALUES ('1f327e07-f1b6-11ea-b0dd-0ae0afb34104', 'tonthatngoc.dr@gmail.com', 'Dell0101', '15', '06:00 AM - 07:00 AM', '07:00 AM - 08:00 AM', 'TENNIS NEW', 'Tennis 2 New', 'no', '0', '2021-01-01 02:05:00.001045+00', '0001-01-01 00:00:00+00');
INSERT INTO "public"."users_booking" VALUES ('1f327e07-f1b6-11ea-b0dd-0ae0afb34103', 'tonthatngoc.dr@gmail.com', 'Dell0101', '15', '06:00 AM - 07:00 AM', '07:00 AM - 08:00 AM', 'TENNIS NEW', 'Tennis 1 New', 'no', '0', '2021-01-01 02:05:00.002043+00', '0001-01-01 00:00:00+00');
INSERT INTO "public"."users_booking" VALUES ('1f327e07-f1b6-11ea-b0dd-0ae0afb34109', 'nhantran2002@gmail.com', 'Sunrise123@', '19', '06:00 AM - 07:00 AM', '07:00 AM - 08:00 AM', 'TENNIS NEW', 'Tennis 1 New', 'no', '0', '2021-01-04 17:12:00.001021+00', '0001-01-01 00:00:00+00');
INSERT INTO "public"."users_booking" VALUES ('1f327e07-f1b6-11ea-b0dd-0ae0afb34101', 'sangnt76@gmail.com', 'Quynh@nh05', '16', '07:00 AM - 08:00 AM', '07:00 AM - 08:00 AM', 'TENNIS NEW', 'Tennis 1 New', 'no', '0', '2021-01-01 16:59:30.021489+00', '2021-01-01 17:00:01.25171+00');
INSERT INTO "public"."users_booking" VALUES ('1f327e07-f1b6-11ea-b0dd-0ae0afb34102', 'sangnt76@gmail.com', 'Quynh@nh05', '16', '07:00 AM - 08:00 AM', '07:00 AM - 08:00 AM', 'TENNIS NEW', 'Tennis 2 New', 'no', '0', '2021-01-01 16:59:30.02325+00', '2021-01-01 17:00:01.254707+00');
INSERT INTO "public"."users_booking" VALUES ('1f327e07-f1b6-11ea-b0dd-0ae0afb34106', 'landquan2@gmail.com', 'Nguyenthuy1', '13', '06:00 AM - 07:00 AM', '07:00 AM - 08:00 AM', 'TENNIS NEW', 'Tennis 2 New', 'yes', 'T2-A06-03', '2020-10-28 17:00:00.008038+00', '2020-10-28 17:00:22.62863+00');
INSERT INTO "public"."users_booking" VALUES ('1f327e07-f1b6-11ea-b0dd-0ae0afb34107', 'landquan2@gmail.com', 'Nguyenthuy1', '13', '06:00 AM - 07:00 AM', '07:00 AM - 08:00 AM', 'TENNIS NEW', 'Tennis 1 New', 'yes', 'T5-B08-12', '2020-10-28 17:00:15.003297+00', '2020-10-28 17:01:37.70653+00');

-- +migrate Down
DROP TABLE users_booking;