--
-- School of Architecture
INSERT INTO fieldmaps( source_id, map_id ) VALUES(
  ( SELECT id from fieldvalues where field_name = "department" AND field_value = "School of Architecture" ),
  ( SELECT id from fieldvalues where field_name = "degree" AND field_value = "MAR (Master of Architecture)" )
);
INSERT INTO fieldmaps( source_id, map_id ) VALUES(
  ( SELECT id from fieldvalues where field_name = "department" AND field_value = "School of Architecture" ),
  ( SELECT id from fieldvalues where field_name = "degree" AND field_value = "MLA (Master of Landscape Architecture)" )
);

--
-- Department of Architectural History
INSERT INTO fieldmaps( source_id, map_id ) VALUES(
  ( SELECT id from fieldvalues where field_name = "department" AND field_value = "Department of Architectural History" ),
  ( SELECT id from fieldvalues where field_name = "degree" AND field_value = "MARH (Master of Architectural History)" )
);
INSERT INTO fieldmaps( source_id, map_id ) VALUES(
  ( SELECT id from fieldvalues where field_name = "department" AND field_value = "Department of Architectural History" ),
  ( SELECT id from fieldvalues where field_name = "degree" AND field_value = "BARH (Bachelor of Architectural History)" )
);

--
-- Batten School of Leadership and Public Policy
INSERT INTO fieldmaps( source_id, map_id ) VALUES(
  ( SELECT id from fieldvalues where field_name = "department" AND field_value = "Batten School of Leadership and Public Policy" ),
  ( SELECT id from fieldvalues where field_name = "degree" AND field_value = "MPP (Master of Public Policy)" )
);

--
-- Department of Biomedical Engineering
INSERT INTO fieldmaps( source_id, map_id ) VALUES(
  ( SELECT id from fieldvalues where field_name = "department" AND field_value = "Department of Biomedical Engineering" ),
  ( SELECT id from fieldvalues where field_name = "degree" AND field_value = "BS (Bachelor of Science)" )
);

--
-- Department of Chemical Engineering
INSERT INTO fieldmaps( source_id, map_id ) VALUES(
  ( SELECT id from fieldvalues where field_name = "department" AND field_value = "Department of Chemical Engineering" ),
  ( SELECT id from fieldvalues where field_name = "degree" AND field_value = "BS (Bachelor of Science)" )
);

--
-- Department of Civil & Environmental Engineering
INSERT INTO fieldmaps( source_id, map_id ) VALUES(
  ( SELECT id from fieldvalues where field_name = "department" AND field_value = "Department of Civil & Environmental Engineering" ),
  ( SELECT id from fieldvalues where field_name = "degree" AND field_value = "BS (Bachelor of Science)" )
);

--
-- Department of Computer Science
INSERT INTO fieldmaps( source_id, map_id ) VALUES(
  ( SELECT id from fieldvalues where field_name = "department" AND field_value = "Department of Computer Science" ),
  ( SELECT id from fieldvalues where field_name = "degree" AND field_value = "MCS (Master of Computer Science)" )
);
INSERT INTO fieldmaps( source_id, map_id ) VALUES(
  ( SELECT id from fieldvalues where field_name = "department" AND field_value = "Department of Computer Science" ),
  ( SELECT id from fieldvalues where field_name = "degree" AND field_value = "BS (Bachelor of Science)" )
);

--
-- Department of Electrical & Computing Engineering
INSERT INTO fieldmaps( source_id, map_id ) VALUES(
  ( SELECT id from fieldvalues where field_name = "department" AND field_value = "Department of Electrical & Computing Engineering" ),
  ( SELECT id from fieldvalues where field_name = "degree" AND field_value = "BS (Bachelor of Science)" )
);

--
-- Department of Engineering Physics
INSERT INTO fieldmaps( source_id, map_id ) VALUES(
  ( SELECT id from fieldvalues where field_name = "department" AND field_value = "Department of Engineering Physics" ),
  ( SELECT id from fieldvalues where field_name = "degree" AND field_value = "ME (Master of Engineering)" )
);
INSERT INTO fieldmaps( source_id, map_id ) VALUES(
  ( SELECT id from fieldvalues where field_name = "department" AND field_value = "Department of Engineering Physics" ),
  ( SELECT id from fieldvalues where field_name = "degree" AND field_value = "MEP (Master of Engineering Physics)" )
);

--
-- Department of Engineering Science
INSERT INTO fieldmaps( source_id, map_id ) VALUES(
  ( SELECT id from fieldvalues where field_name = "department" AND field_value = "Department of Engineering Science" ),
  ( SELECT id from fieldvalues where field_name = "degree" AND field_value = "ME (Master of Engineering)" )
);
INSERT INTO fieldmaps( source_id, map_id ) VALUES(
  ( SELECT id from fieldvalues where field_name = "department" AND field_value = "Department of Engineering Science" ),
  ( SELECT id from fieldvalues where field_name = "degree" AND field_value = "BS (Bachelor of Science)" )
);

--
-- Department of Materials Science & Engineering
INSERT INTO fieldmaps( source_id, map_id ) VALUES(
  ( SELECT id from fieldvalues where field_name = "department" AND field_value = "Department of Materials Science & Engineering" ),
  ( SELECT id from fieldvalues where field_name = "degree" AND field_value = "MMSE (Master of Materials Science and Engineering)" )
);
INSERT INTO fieldmaps( source_id, map_id ) VALUES(
  ( SELECT id from fieldvalues where field_name = "department" AND field_value = "Department of Materials Science & Engineering" ),
  ( SELECT id from fieldvalues where field_name = "degree" AND field_value = "ME (Master of Engineering)" )
);

--
-- Department of Mechanical and Aerospace Engineering
INSERT INTO fieldmaps( source_id, map_id ) VALUES(
  ( SELECT id from fieldvalues where field_name = "department" AND field_value = "Department of Mechanical and Aerospace Engineering" ),
  ( SELECT id from fieldvalues where field_name = "degree" AND field_value = "BS (Bachelor of Science)" )
);

--
-- Department of Systems & Information Engineering
INSERT INTO fieldmaps( source_id, map_id ) VALUES(
  ( SELECT id from fieldvalues where field_name = "department" AND field_value = "Department of Systems & Information Engineering" ),
  ( SELECT id from fieldvalues where field_name = "degree" AND field_value = "BS (Bachelor of Science)" )
);

--
-- Department of Urban and Environmental Planning
INSERT INTO fieldmaps( source_id, map_id ) VALUES(
  ( SELECT id from fieldvalues where field_name = "department" AND field_value = "Department of Urban and Environmental Planning" ),
  ( SELECT id from fieldvalues where field_name = "degree" AND field_value = "MUEP (Master of Urban and Environmental Planning)" )
);
INSERT INTO fieldmaps( source_id, map_id ) VALUES(
  ( SELECT id from fieldvalues where field_name = "department" AND field_value = "Department of Urban and Environmental Planning" ),
  ( SELECT id from fieldvalues where field_name = "degree" AND field_value = "BUEP (Bachelor of Urban and Environmental Planning)" )
);

-- end of file