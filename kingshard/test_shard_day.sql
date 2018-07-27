create table test_shard_day_20160306(
  mtime DATETIME NOT NULL ,
  game int DEFAULT 0,
  val1 int DEFAULT 0,
  val2 int DEFAULT 0
) ENGINE=innodb;
create table test_shard_day_20160307(
  mtime DATETIME NOT NULL ,
  game int DEFAULT 0,
  val1 int DEFAULT 0,
  val2 int DEFAULT 0
) ENGINE=innodb;
create table test_shard_day_20160308(
  mtime DATETIME NOT NULL ,
  game int DEFAULT 0,
  val1 int DEFAULT 0,
  val2 int DEFAULT 0
) ENGINE=innodb;
create table test_shard_day_20160309(
  mtime DATETIME NOT NULL ,
  game int DEFAULT 0,
  val1 int DEFAULT 0,
  val2 int DEFAULT 0
) ENGINE=innodb;

insert INTO test_shard_day_20160306 (mtime, game, val1, val2) VALUES ('2016-03-06 06:00:00', '1', 26, 6);
insert INTO test_shard_day_20160306 (mtime, game, val1, val2) VALUES ('2016-03-06 06:00:00', '2', 26, 16);
insert INTO test_shard_day_20160307 (mtime, game, val1, val2) VALUES ('2016-03-07 07:00:00', '1', 27, 7);
insert INTO test_shard_day_20160307 (mtime, game, val1, val2) VALUES ('2016-03-07 07:00:00', '2', 27, 17);
insert INTO test_shard_day_20160308 (mtime, game, val1, val2) VALUES ('2016-03-08 08:00:00', '1', 28, 8);
insert INTO test_shard_day_20160308 (mtime, game, val1, val2) VALUES ('2016-03-08 08:00:00', '2', 28, 18);
insert INTO test_shard_day_20160309 (mtime, game, val1, val2) VALUES ('2016-03-09 09:00:00', '1', 29, 19);
insert INTO test_shard_day_20160309 (mtime, game, val1, val2) VALUES ('2016-03-09 09:00:00', '2', 29, 9);
