-- drop the table if it exists
DROP TABLE IF EXISTS depositrequest;

-- and create the new one
CREATE TABLE depositrequest(
   id          INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
   requester   VARCHAR( 32 ) NOT NULL DEFAULT '',
   user        VARCHAR( 32 ) NOT NULL DEFAULT '',
   school      VARCHAR( 32 ) NOT NULL DEFAULT '',
   degree      VARCHAR( 32 ) NOT NULL DEFAULT '',
   status      ENUM('pending', 'submitted') NOT NULL DEFAULT 'pending',
   request_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
   deposit_date TIMESTAMP NULL
) CHARACTER SET utf8 COLLATE utf8_bin;

INSERT INTO depositrequest( id, requester, user, school, degree ) VALUES( 1, "dpg3k", "dpg3k", "Engineering", "Ph.D" );