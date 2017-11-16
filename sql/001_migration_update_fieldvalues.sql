
-- change some existing field names

UPDATE fieldvalues set field_value = "Department of Civil & Environmental Engineering" WHERE field_name = "department" AND field_value = "Department of Civil Engineering" LIMIT 1;
UPDATE fieldvalues set field_value = "Department of Electrical & Computing Engineering" WHERE field_name = "department" AND field_value = "Department of Electrical Engineering" LIMIT 1;
UPDATE fieldvalues set field_value = "Department of Materials Science & Engineering" WHERE field_name = "department" AND field_value = "Department of Materials Science and Engineering" LIMIT 1;
UPDATE fieldvalues set field_value = "Department of Systems & Information Engineering" WHERE field_name = "department" AND field_value = "Department of Systems Engineering" LIMIT 1;

-- add new department values
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "department","Batten School of Leadership and Public Policy");
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "department","Department of Engineering Science");

-- add new degree values
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "degree","MLA (Master of Landscape Architecture)");
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "degree","BARH (Bachelor of Architectural History)");
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "degree","MPP (Master of Public Policy)");
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "degree","BUEP (Bachelor of Urban and Environmental Planning)");

-- end of file