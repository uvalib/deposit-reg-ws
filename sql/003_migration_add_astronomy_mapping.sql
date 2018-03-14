
--
-- Department of Astronomy
INSERT INTO fieldmaps( source_id, map_id ) VALUES(
  ( SELECT id from fieldvalues where field_name = "department" AND field_value = "Department of Astronomy" ),
  ( SELECT id from fieldvalues where field_name = "degree" AND field_value = "BS (Bachelor of Science)" )
);
INSERT INTO fieldmaps( source_id, map_id ) VALUES(
  ( SELECT id from fieldvalues where field_name = "department" AND field_value = "Department of Astronomy" ),
  ( SELECT id from fieldvalues where field_name = "degree" AND field_value = "MS (Master of Science)" )
);
