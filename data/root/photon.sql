/******************************
**	File:	photon.sql
**	Name:	Photon
**	Desc:	Photon Microservice's Persistence script
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
** 2    12/29/2019      Alonso R      Add role enum
*******************************/

/******************************
    Role enum
*******************************/
CREATE TYPE ROLE_ENUM AS ENUM (
    'ROLE_STUDENT',
    'ROLE_TEACHER',
    'ROLE_OPERATOR',
    'ROLE_MANAGER',
    'ROLE_ADMIN',
    'ROLE_ROOT',
    'ROLE_SUPPORT'
);

/******************************
    Users table
    ROLES: 
        -   Student:                ROLE_STUDENT
        -   Teacher:                ROLE_TEACHER
        -   Warehouse Operator:     ROLE_OPERATOR
        -   Warehouse Manager:      ROLE_MANAGER
        -   Admin:                  ROLE_ADMIN
        -   Root:                   ROLE_ROOT
        -   Support:                ROLE_SUPPORT
*******************************/
CREATE TABLE IF NOT EXISTS USERS (
    ID          BIGSERIAL NOT NULL PRIMARY KEY,
    NAME        VARCHAR(255) NOT NULL,
    SURNAME     VARCHAR(255) NOT NULL,
    BIRTH       DATE NOT NULL DEFAULT CURRENT_DATE,
    USERNAME    VARCHAR(100) NOT NULL UNIQUE,
    PASSWORD    TEXT NOT NULL,
    IMAGE       TEXT DEFAULT NULL,
    ROLE        ROLE_ENUM DEFAULT 'ROLE_STUDENT',
    ACTIVE      BOOLEAN DEFAULT TRUE,
    CREATED_AT  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UPDATED_AT  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

/******************************
    Courses table
*******************************/
CREATE TABLE IF NOT EXISTS COURSES (
    ID          BIGSERIAL NOT NULL PRIMARY KEY,
    NAME        VARCHAR(255) NOT NULL,
    GRADE       INT DEFAULT 1,
    FK_TEACHER  BIGSERIAL NOT NULL REFERENCES USERS(ID) ON DELETE CASCADE,
    CREATED_AT  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UPDATED_AT  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

/******************************
    Classrooms table
    [(Course <-> Student relation table)]
*******************************/
CREATE TABLE IF NOT EXISTS CLASSROOMS (
    FK_STUDENT  BIGSERIAL NOT NULL REFERENCES USERS(ID) ON DELETE CASCADE,
    FK_COURSE   BIGSERIAL NOT NULL REFERENCES COURSES(ID) ON DELETE CASCADE,
    JOINED_AT   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

/******************************
    Materials table
*******************************/
CREATE TABLE IF NOT EXISTS MATERIALS (
    ID          BIGSERIAL NOT NULL PRIMARY KEY,
    NAME        VARCHAR(255),
    DESCRIPTION TEXT DEFAULT NULL,
    CAPACITY    VARCHAR(100) DEFAULT NULL,
    IMAGE       TEXT DEFAULT NULL,
    CREATED_AT  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UPDATED_AT  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

/******************************
    Material types table
*******************************/
CREATE TABLE IF NOT EXISTS MATERIAL_TYPES (
    ID          SERIAL NOT NULL PRIMARY KEY,
    NAME        VARCHAR(255)
);

/******************************
    Material groups table
    Material grouped by types
    [(Material <-> Type relation table)]
*******************************/
CREATE TABLE IF NOT EXISTS MATERIAL_GROUPS (
    FK_MATERIAL BIGSERIAL NOT NULL REFERENCES MATERIALS(ID) ON DELETE CASCADE,
    FK_TYPE     SERIAL NOT NULL REFERENCES MATERIAL_TYPES(ID) ON DELETE CASCADE
);

/******************************
    Tasks table
*******************************/
CREATE TABLE IF NOT EXISTS TASKS (
    ID          BIGSERIAL NOT NULL PRIMARY KEY,
    TITLE       VARCHAR(255) NOT NULL,
    DESCRIPTION TEXT,
    IS_TEAM     BOOLEAN DEFAULT FALSE,
    FK_COURSE   BIGSERIAL NOT NULL REFERENCES COURSES(ID) ON DELETE CASCADE,
    ISSUED_AT   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    EXPIRES_AT  TIMESTAMP NOT NULL,
    UPDATED_AT  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

/******************************
    Task's materials table
    [(Task -> Material relation table)]
*******************************/
CREATE TABLE IF NOT EXISTS TASK_MATERIALS (
    FK_TASK     BIGSERIAL NOT NULL REFERENCES TASKS(ID) ON DELETE CASCADE,
    FK_MATERIAL SERIAL NOT NULL REFERENCES MATERIALS(ID) ON DELETE CASCADE
);

/******************************
    Warehouses table
        TYPES ENUM: 
        -   0:      MAIN WAREHOUSE
        -   1:      LAB WAREHOUSE
        -   2:      OTHER WAREHOUSE
*******************************/
CREATE TABLE IF NOT EXISTS WAREHOUSES (
    ID          SERIAL NOT NULL PRIMARY KEY,
    NAME        VARCHAR(255) NOT NULL,
    DESCRIPTION TEXT DEFAULT NULL,
    TYPE        INT DEFAULT 0,
    CREATED_AT  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UPDATED_AT  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

/******************************
    Material Stocks table
*******************************/
CREATE TABLE IF NOT EXISTS STOCKS (
    FK_WAREHOUSE    SERIAL NOT NULL REFERENCES WAREHOUSES(ID) ON DELETE CASCADE,
    FK_MATERIAL     SERIAL NOT NULL REFERENCES MATERIALS(ID) ON DELETE CASCADE,
    QUANTITY        INT DEFAULT 0,
    UPDATED_AT      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

/******************************
    Orders table
    STATUS ENUM: 
        -   0:      REQUESTED BY STUDENT
        -   1:      ACCEPTED BY OPERATOR
        -   2:      READY
        -   3:      DELIVERED
        -   4:      COMPLETED

        -   5:      CANCELED BY STUDENT
        -   6:      REJECTED BY OPERATOR
        -   7:      INCOMPLETED
*******************************/
CREATE TABLE IF NOT EXISTS ORDERS (
    ID          BIGSERIAL NOT NULL PRIMARY KEY,
    FK_STUDENT  BIGSERIAL NOT NULL REFERENCES USERS(ID),
    FK_TASK     BIGSERIAL NOT NULL REFERENCES TASKS(ID),
    STATUS      INT DEFAULT 0,
    ISSUED_AT   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    EXPIRES_AT  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

/******************************
    Bans table
    REASON ENUM: 
        -   0:      MISSING / BROKEN MATERIAL
        -   1:      LATE MATERIAL DELIVERY
        -   2:      OTHER
*******************************/
CREATE TABLE IF NOT EXISTS BANS (
    ID          BIGSERIAL NOT NULL PRIMARY KEY,
    TITLE       VARCHAR(255) NOT NULL,
    COMMENT     TEXT NOT NULL,
    REASON      INT DEFAULT 0,
    FK_STUDENT  BIGSERIAL NOT NULL REFERENCES USERS(ID),
    FK_ORDER    BIGSERIAL NOT NULL REFERENCES ORDERS(ID),
    FK_ISSUER   BIGSERIAL NOT NULL REFERENCES USERS(ID),
    ISSUED_AT   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    EXPIRES_AT  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create support schema
CREATE SCHEMA IF NOT EXISTS SUPPORT;

/******************************
    Support Tickets table
*******************************/
CREATE TABLE IF NOT EXISTS SUPPORT.TICKETS (
    ID          BIGSERIAL NOT NULL PRIMARY KEY,
    TITLE       VARCHAR(255) NOT NULL,
    COMMENT     TEXT NOT NULL,
    ISSUED_AT   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FK_ISSUER   BIGSERIAL NOT NULL REFERENCES USERS(ID),
    FK_SUPPORT  BIGSERIAL NOT NULL REFERENCES USERS(ID)
);

/******************************
    Support Logs table
*******************************/
CREATE TABLE IF NOT EXISTS SUPPORT.LOGS (
	ID			BIGSERIAL NOT NULL PRIMARY KEY,
	DESCRIPTION	TEXT NOT NULL,
	EXECUTED_AT	TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	EXECUTED_BY	VARCHAR(255) NOT NULL
);