/******************************
**	File:	photon_triggers.sql
**	Name:	Photon Triggers
**	Desc:	Photon Database's triggers script
**	Auth:	Damascus Mexico Group
**	Date:	2019
**	Lisc:	Confidential - Closed
**	Copy:	Damascus Mexico Group, Inc. 2019 All rights reserved.
**************************
** Change History
**************************
** PR   Date        Author  		Description 
** --   --------   -------   		------------------------------------
** 1    12/27/2019      Alonso R      Initial setup
*******************************/

CREATE PROCEDURAL LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION UPDATE_USER_LOG()
    RETURNS TRIGGER AS $BODY$
BEGIN
    INSERT INTO SUPPORT.LOGS(DESCRIPTION, EXECUTED_BY) VALUES (
		FORMAT('Operation executed at %s by %s', TG_TABLE_NAME, CURRENT_USER),
		CURRENT_USER
	);
    RETURN NULL;
END;
$BODY$
LANGUAGE plpgsql VOLATILE COST 100;


CREATE TRIGGER USER_UPDATED
AFTER UPDATE ON USERS 
FOR EACH ROW
EXECUTE PROCEDURE UPDATE_USER_LOG();

-- DROP TRIGGER USER_UPDATED ON USERS;

UPDATE USERS SET SURNAME = 'Ruiz Esparza Acuna' WHERE ID = 1;
SELECT * FROM SUPPORT.LOGS;