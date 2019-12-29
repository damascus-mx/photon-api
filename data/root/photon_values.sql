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

INSERT INTO USERS(NAME, SURNAME, BIRTH, USERNAME, PASSWORD) VALUES (
    'Luis Alonso',
    'Ruiz Esparza',
    '09-20-1998',
    '1117250020',
    'password'
); 

UPDATE USERS SET NAME = 'Alonso' WHERE ID = 1;

INSERT INTO USERS(NAME, SURNAME, BIRTH, USERNAME, PASSWORD, ROLE) VALUES (
    'Miguel',
    'Ruiz Esparza',
    '09-29-1959',
    '1117250021',
    'password',
    'ROLE_TEACHER'
);

INSERT INTO COURSES(NAME, GRADE, FK_TEACHER) VALUES (
    'Quimica Basica I',
    1,
    2
);