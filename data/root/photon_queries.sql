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

-- GET ALL COURSES WITH TEACHERS
SELECT COURSES.NAME, COURSES.GRADE, USERS.NAME, USERS.SURNAME FROM COURSES INNER JOIN USERS ON COURSES.FK_TEACHER = USERS.ID;