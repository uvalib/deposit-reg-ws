-- drop the table if it exists
DROP TABLE IF EXISTS fieldmaps;

-- and create the new one
CREATE TABLE fieldmaps(
   id          INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
   source_id   INT NOT NULL,
   map_id      INT NOT NULL,
   create_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
) CHARACTER SET utf8 COLLATE utf8_bin;