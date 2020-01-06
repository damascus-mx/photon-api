/******************************
**	File:	photon_values.sql
**	Name:	Photon Values
**	Desc:	Photon Database's initial values script
**	Auth:	Damascus Mexico Group
**	Date:	2019
**	Lisc:	Confidential - Closed
**	Copy:	Damascus Mexico Group, Inc. 2019 All rights reserved.
**************************
** Change History
**************************
** PR   Date        Author  		Description 
** --   --------   -------   		------------------------------------
** 1    12/28/2019      Alonso R      Initial setup
*******************************/

-- Create root user
INSERT INTO USERS(NAME, SURNAME, USERNAME, PASSWORD, ROLE) VALUES (
    'root',
    'root',
    '11111111',
    'root',
    'ROLE_ROOT'
);

-- Create students
INSERT INTO USERS(NAME, SURNAME, BIRTH, USERNAME, PASSWORD) VALUES (
    'Luis Alonso',
    'Ruiz Esparza',
    '09-20-1998',
    '1117250020',
    'password'
), (
    'Fernando',
    'Herrera Ahumada',
    '09-25-1998',
    '1117250021',
    'password'
), (
    'Jose Angel',
    'Rivera Lechuga',
    '07-29-1996',
    '1117250022',
    'password'
);

-- Create various user types
INSERT INTO USERS(NAME, SURNAME, BIRTH, USERNAME, PASSWORD, ROLE) VALUES (
    'Miguel',
    'Ruiz Esparza',
    '09-29-1959',
    '1117250023',
    'password',
    'ROLE_TEACHER'
), (
    'Fernando Ramon',
    'Ruelas Estrada',
    '09-20-1980',
    '1117250024',
    'password',
    'ROLE_ADMIN'
), (
    'Luz Imelda',
    'Acuna Arzaga',
    '10-06-1967',
    '1117250025',
    'password',
    'ROLE_MANAGER'
), (
    'Hiram Arturo',
    'Olivero Munoz',
    '06-26-1988',
    '1117250026',
    'password',
    'ROLE_OPERATOR'
), (
    'Edgar Alan',
    'Valerio Macias',
    '12-12-1975',
    '1117250027',
    'password',
    'ROLE_SUPPORT'
), (
    'Juan Carlos',
    'Arzaga Terrazas',
    '04-18-1985',
    '1117250028',
    'password',
    'ROLE_EMPLOYEE'
), (
    'Alfonso',
    'Arzaga Terrazas',
    '01-20-1990',
    '1117250029',
    'password',
    'ROLE_PRINCIPAL'
);

-- Create College -tenant-
INSERT INTO COLLEGES(NAME, SHORTNAME, FK_OWNER, COUNTRY, STATE, FOUNDED_AT) VALUES (
    'Universidad Autonoma de Chihuahua',
    'UACH',
    6,
    'MX',
    'MX-CHH',
    '12-08-1954'
);

-- Create faculty
INSERT INTO FACULTIES(NAME, FK_COLLEGE) VALUES (
    'Facultad de Ciencias Quimicas',
    1
);

-- Attach users to faculties
INSERT INTO FACULTY_USERS VALUES(
    1, 2
), (
    1, 3
), (
    1, 4
), (
    1, 5
), (
    1, 7
), (
    1, 8
), (
    1, 10
), (
    1, 11   
);

-- Create courses
INSERT INTO COURSES(NAME, GRADE, FK_TEACHER) VALUES (
    'Quimica Cuantica I',
    3,
    5
), (
    'Introduccion a quimica',
    1,
    5
);

-- Assign courses to students
-- Student - Course
INSERT INTO CLASSROOMS(FK_STUDENT, FK_COURSE) VALUES (
    2, 1
), (
    3, 1
), (
    3, 2
), (
    4, 2
);

-- Create materials
INSERT INTO MATERIALS(NAME, CAPACITY) VALUES (
    'Bureta',
    '50ml'
), (
    'Embudo Bunchner',
    '130mm'
);

-- Create material types
INSERT INTO MATERIAL_TYPES(NAME) VALUES (
    'Bureta'
), (
    'Embudo'
);

-- Create material presentations
INSERT INTO MATERIAL_PRESENTATIONS(NAME) VALUES (
    'Vidrio'
), (
    'Porcelana'
);

-- Assign presentations to materials
-- Material - Presentation
INSERT INTO MATERIAL_PRESENTATION_GROUPS VALUES (
    2, 2
), (
    1, 1
);

-- Assign types/categories to materials
-- Material - Type
INSERT INTO MATERIAL_GROUPS VALUES (
    1, 1
), (
    2, 2
);

-- Create main warehouse
INSERT INTO WAREHOUSES(NAME, DESCRIPTION, TYPE) VALUES (
    'Almacen principal',
    'Central de operaciones',
    'MAIN_WAREHOUSE'
);

-- Create lab warehouse
INSERT INTO WAREHOUSES(NAME, DESCRIPTION) VALUES (
    'Almacen laboratorio II',
    'Localizado en ala 1, planta baja'
);

-- Assign operator to warehouse
INSERT INTO WAREHOUSE_OPERATORS(FK_OPERATOR, FK_WAREHOUSE, EXPIRES_AT) VALUES (
    8, 1, '02-14-2020'
);

-- Add materials to warehouse's stocks
INSERT INTO STOCKS(FK_WAREHOUSE, FK_MATERIAL, QUANTITY) VALUES (
    1, 1, 90
), (
    1, 2, 48
), (
    2, 2, 15
);

-- Create tasks for courses
INSERT INTO TASKS(TITLE, DESCRIPTION, IS_TEAM, FK_COURSE, EXPIRES_AT) VALUES (
    'Practica gases I',
    'Temas relacionados con gases no nocivos',
    FALSE,
    2,
    '01-31-2020 12:00:00'
), (
    'Practica mercurio 3',
    'Realizacion de principios del mercurio',
    TRUE,
    1,
    '02-25-2020 15:00:00'
);

-- Assign task to file
INSERT INTO TASKS_FILES(FK_TASK, FILE_URL) VALUES (
    1, 'https://cdn.damascus-engineering.com/andromeda/users/drmanhattan.jpg'
);

-- Assign materials to task
INSERT INTO TASK_MATERIALS VALUES (
    1, 1
), (
    1, 2
);

-- Create order
INSERT INTO ORDERS(FK_STUDENT, FK_TASK, STATUS, EXPIRES_AT) VALUES (
    3, 1, 'REQUESTED_BY_STUDENT', '01-31-2020 11:30:00'
), (
	4, 1, 'REQUESTED_BY_STUDENT', '01-31-2020 11:30:00'
);

-- Ban student
INSERT INTO BANS(TITLE, COMMENT, REASON, FK_STUDENT, FK_ORDER, FK_ISSUER, EXPIRES_AT) VALUES (
    'Bloqueo de cuenta por material perdido',
    'El alumno Jose Angel Rivera Lechuga con la matricula 1117250022 ha perdido material perteneciente a la Facultad de Ciencias Quimicas.',
    'MISSING_MATERIAL',
    4,
    2,
    7,
    '12-31-2020 00:00:00'
);

-- User issues a support ticket
INSERT INTO SUPPORT.TICKETS(TITLE, COMMENT, REASON, FK_ISSUER) VALUES (
    'Bloqueo sin meritos',
    'Me he metido a mi cuenta hoy y aparece que estoy bloqueado, necesito ayuda.',
    'BAN_ISSUE',
    4
);