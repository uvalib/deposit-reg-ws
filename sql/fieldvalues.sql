-- drop the table if it exists
DROP TABLE IF EXISTS fieldvalues;

-- and create the new one
CREATE TABLE fieldvalues(
   id          INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
   field_name  VARCHAR( 32 ) NOT NULL DEFAULT '',
   field_value VARCHAR( 80 ) NOT NULL DEFAULT '',
   create_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
) CHARACTER SET utf8 COLLATE utf8_bin;

-- set of default department values
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "department","Curry School of Education");
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "department","Department of Anthropology");
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "department","Department of Architectural History");
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "department","Department of Art");
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "department","Department of Astronomy");
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "department","Department of Biochemistry and Molecular Genetics");
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "department","Department of Biology");
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "department","Department of Biomedical Engineering");
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "department","Department of Biophysics");
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "department","Department of Chemical Engineering");
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "department","Department of Chemistry");
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "department","Department of Civil Engineering");
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "department","Department of Classics");
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "department","Department of Computer Engineering");
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "department","Department of Computer Science");
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "department","Department of Drama");
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "department","Department of East Asian Studies");
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "department","Department of Economics");
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "department","Department of Electrical Engineering");
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "department","Department of Engineering Physics");
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "department","Department of English");
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "department","Department of Environmental Sciences");
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "department","Department of Foreign Affairs");
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "department","Department of French");
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "department","Department of Germanic Languages and Literatures");
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "department","Department of History");
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "department","Department of Materials Science and Engineering");
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "department","Department of Mathematics");
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "department","Department of Mechanical and Aerospace Engineering");
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "department","Department of Microbiology, Immunology, and Cancer Biology");
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "department","Department of Molecular Physiology and Biological Physics");
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "department","Department of Molecular, Cell and Developmental Biology");
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "department","Department of Music");
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "department","Department of Neuroscience");
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "department","Department of Pathology");
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "department","Department of Pharmacology");
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "department","Department of Philosophy");
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "department","Department of Physics");
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "department","Department of Politics");
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "department","Department of Psychology");
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "department","Department of Religious Studies");
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "department","Department of Slavic Languages and Literatures");
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "department","Department of Sociology");
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "department","Department of Spanish, Italian, and Portuguese");
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "department","Department of Statistics");
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "department","Department of Systems Engineering");
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "department","Department of Urban and Environmental Planning");
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "department","School of Nursing");

-- set of default deposit_type values
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "degree", "AM" );
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "degree", "CGE" );
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "degree", "DNP" );
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "degree", "EDD" );
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "degree", "MA" );
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "degree", "MAR" );
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "degree", "MARH" );
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "degree", "MCS" );
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "degree", "ME" );
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "degree", "MED" );
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "degree", "MEP" );
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "degree", "MFA" );
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "degree", "MMSE" );
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "degree", "MS" );
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "degree", "MUEP" );
INSERT INTO fieldvalues( field_name, field_value ) VALUES( "degree", "PHD" );
