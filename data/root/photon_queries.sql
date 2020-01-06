/******************************
**	File:	photon_queries.sql
**	Name:	Photon Queries
**	Desc:	Photon Database's queries script
**	Auth:	Damascus Mexico Group
**	Date:	2019
**	Lisc:	Confidential - Closed
**	Copy:	Damascus Mexico Group, Inc. 2019 All rights reserved.
**************************
** Change History
**************************
** PR   Date        Author  		Description 
** --   --------   -------   		------------------------------------
** 1    12/29/2019      Alonso R      Initial setup
*******************************/

-- Get all courses with teacher
SELECT COURSES.NAME, COURSES.GRADE, USERS.NAME, USERS.SURNAME FROM COURSES INNER JOIN USERS ON COURSES.FK_TEACHER = USERS.ID;

-- Get all users attached to 1st faculty
SELECT USERS.NAME, USERS.SURNAME FROM FACULTY_USERS INNER JOIN USERS ON FACULTY_USERS.FK_USER = USERS.ID 
WHERE FACULTY_USERS.FK_FACULTY = 1;