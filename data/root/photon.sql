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
    User Role enum
*******************************/
CREATE TYPE USER_ROLE AS ENUM (
    'ROLE_STUDENT',
    'ROLE_TEACHER',
    'ROLE_OPERATOR',
    'ROLE_MANAGER',
    'ROLE_ADMIN',
    'ROLE_ROOT',
    'ROLE_SUPPORT',
    'ROLE_EMPLOYEE',
    'ROLE_PRINCIPAL'
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
        -   Faculty's Employee:     ROLE_EMPLOYEE
        -   Faculty's Principal:    ROLE_PRINCIPAL
*******************************/
CREATE TABLE IF NOT EXISTS USERS (
    ID          BIGSERIAL NOT NULL PRIMARY KEY,
    NAME        VARCHAR(255) NOT NULL,
    SURNAME     VARCHAR(255) NOT NULL,
    BIRTH       DATE NOT NULL DEFAULT CURRENT_DATE,
    USERNAME    VARCHAR(100) NOT NULL UNIQUE,
    PASSWORD    TEXT NOT NULL,
    IMAGE       TEXT DEFAULT NULL,
    ROLE        USER_ROLE DEFAULT 'ROLE_STUDENT',
    ACTIVE      BOOLEAN DEFAULT TRUE,
    CREATED_AT  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UPDATED_AT  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

/******************************
    Colleges -Tenants- table
    [(College <- User)]
*******************************/
CREATE TABLE IF NOT EXISTS COLLEGES (
    ID          BIGSERIAL NOT NULL PRIMARY KEY,
    NAME        TEXT NOT NULL,
    SHORTNAME   VARCHAR(100),
    FK_OWNER    BIGSERIAL NOT NULL REFERENCES USERS(ID),
    STATE       VARCHAR(10) NOT NULL, -- Using ISO 3166-2:COUNTRY 3-Digit state acronym
    COUNTRY     VARCHAR(4) NOT NULL, -- Using ISO 3166-1 alpha-2 country acronym
    FOUNDED_AT  DATE NOT NULL DEFAULT CURRENT_DATE,
    CREATED_AT  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UPDATED_AT  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

/******************************
    Faculties table
    [(Faculty <- College)]
*******************************/
CREATE TABLE IF NOT EXISTS FACULTIES (
    ID          BIGSERIAL NOT NULL PRIMARY KEY,
    NAME        TEXT NOT NULL,
    FK_COLLEGE  BIGSERIAL NOT NULL REFERENCES COLLEGES(ID) ON DELETE CASCADE,
    CREATED_AT  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UPDATED_AT  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

/******************************
    Faculty - Users table
    [(Faculty <-> User)]
*******************************/
CREATE TABLE IF NOT EXISTS FACULTY_USERS (
    FK_FACULTY  BIGSERIAL NOT NULL REFERENCES FACULTIES(ID) ON DELETE CASCADE,
    FK_USER     BIGSERIAL NOT NULL REFERENCES USERS(ID) ON DELETE CASCADE
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
        ex. funnel, spatula
*******************************/
CREATE TABLE IF NOT EXISTS MATERIAL_TYPES (
    ID          SERIAL NOT NULL PRIMARY KEY,
    NAME        VARCHAR(255)
);

/******************************
    Material presentations table
        ex. glass, porcelain
*******************************/
CREATE TABLE IF NOT EXISTS MATERIAL_PRESENTATIONS (
    ID          SERIAL NOT NULL PRIMARY KEY,
    NAME        VARCHAR(255)
);

/******************************
    Material presentation groups table
    Material grouped by presentations
    [(Material <-> Presentation relation table)]
*******************************/
CREATE TABLE IF NOT EXISTS MATERIAL_PRESENTATION_GROUPS (
    FK_MATERIAL         BIGSERIAL NOT NULL REFERENCES MATERIALS(ID) ON DELETE CASCADE,
    FK_PRESENTATION     SERIAL NOT NULL REFERENCES MATERIAL_PRESENTATIONS(ID) ON DELETE CASCADE
);

/******************************
    Material groups table
    Material grouped by types/categories
    [(Material <-> Type relation table)]
*******************************/
CREATE TABLE IF NOT EXISTS MATERIAL_GROUPS (
    FK_MATERIAL BIGSERIAL NOT NULL REFERENCES MATERIALS(ID) ON DELETE CASCADE,
    FK_TYPE     SERIAL NOT NULL REFERENCES MATERIAL_TYPES(ID) ON DELETE CASCADE
);

/******************************
    Warehouse type enum
*******************************/
CREATE TYPE WAREHOUSE_TYPE AS ENUM (
    'MAIN_WAREHOUSE',
    'LAB_WAREHOUSE',
    'OTHER_WAREHOUSE'
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
    TYPE        WAREHOUSE_TYPE DEFAULT 'LAB_WAREHOUSE',
    CREATED_AT  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UPDATED_AT  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

/******************************
    Warehouse's operators table
    [(Warehouse <- Operator relation table)]
*******************************/
CREATE TABLE IF NOT EXISTS WAREHOUSE_OPERATORS (
    ID              SERIAL NOT NULL PRIMARY KEY,
    FK_OPERATOR     BIGSERIAL NOT NULL UNIQUE REFERENCES USERS(ID) ON DELETE CASCADE,
    FK_WAREHOUSE    SERIAL NOT NULL REFERENCES WAREHOUSES(ID) ON DELETE CASCADE,
    CREATED_AT      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    EXPIRES_AT      DATE DEFAULT NULL
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
    Tasks table
*******************************/
CREATE TABLE IF NOT EXISTS TASKS (
    ID          BIGSERIAL NOT NULL PRIMARY KEY,
    TITLE       VARCHAR(255) NOT NULL,
    DESCRIPTION TEXT DEFAULT NULL,
    IS_TEAM     BOOLEAN DEFAULT FALSE,
    FK_COURSE   BIGSERIAL NOT NULL REFERENCES COURSES(ID) ON DELETE CASCADE,
    ISSUED_AT   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    EXPIRES_AT  TIMESTAMP NOT NULL,
    UPDATED_AT  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

/******************************
    Files for tasks table
*******************************/
CREATE TABLE IF NOT EXISTS TASKS_FILES (
    ID          BIGSERIAL NOT NULL PRIMARY KEY,
    FK_TASK     BIGSERIAL NOT NULL REFERENCES TASKS(ID) ON DELETE CASCADE,
    FILE_URL    TEXT NOT NULL,
    UPLOADED_AT TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
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
    Order status enum
*******************************/
CREATE TYPE ORDER_STATUS AS ENUM (
    'REQUESTED_BY_STUDENT',
    'ACCEPTED_BY_OPERATOR',
    'READY',
    'DELIVERED',
    'COMPLETED',
    'CANCELED_BY_STUDENT',
    'REJECTED_BY_OPERATOR',
    'INCOMPLETED'
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
    ID              BIGSERIAL NOT NULL PRIMARY KEY,
    FK_STUDENT      BIGSERIAL NOT NULL REFERENCES USERS(ID),
    FK_TASK         BIGSERIAL NOT NULL REFERENCES TASKS(ID),
    FK_RECEIVED_BY  BIGSERIAL REFERENCES USERS(ID),
    STATUS          ORDER_STATUS DEFAULT 'REQUESTED_BY_STUDENT',
    ISSUED_AT       TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    EXPIRES_AT      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

/******************************
    Ban reasons enum
*******************************/
CREATE TYPE BAN_REASON AS ENUM (
    'DAMAGED_MATERIAL',
    'MISSING_MATERIAL',
    'LATE_DELIVERY',
    'OTHER'
);

/******************************
    Bans table
    REASON ENUM: 
        -   0:      BROKEN MATERIAL
        -   1:      MISSING MATERIAL
        -   2:      LATE MATERIAL DELIVERY
        -   3:      OTHER
*******************************/
CREATE TABLE IF NOT EXISTS BANS (
    ID          BIGSERIAL NOT NULL PRIMARY KEY,
    TITLE       VARCHAR(255) NOT NULL,
    COMMENT     TEXT NOT NULL,
    REASON      BAN_REASON DEFAULT 'OTHER',
    FK_STUDENT  BIGSERIAL NOT NULL REFERENCES USERS(ID),
    FK_ORDER    BIGSERIAL REFERENCES ORDERS(ID),
    FK_ISSUER   BIGSERIAL NOT NULL REFERENCES USERS(ID),
    ISSUED_AT   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    EXPIRES_AT  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create support schema
CREATE SCHEMA IF NOT EXISTS SUPPORT;

/******************************
    Ticket reason enum
*******************************/
CREATE TYPE TICKET_REASON AS ENUM (
    'BAN_ISSUE',
    'ORDER_ISSUE',
    'WAREHOUSE_ISSUE',
    'MATERIAL_ISSUE',
    'COURSE_ISSUE',
    'USER_ISSUE',
    'OTHER'
);

/******************************
    Support Tickets table
*******************************/
CREATE TABLE IF NOT EXISTS SUPPORT.TICKETS (
    ID          BIGSERIAL NOT NULL PRIMARY KEY,
    TITLE       VARCHAR(255) NOT NULL,
    COMMENT     TEXT NOT NULL,
    REASON      TICKET_REASON NOT NULL DEFAULT 'OTHER',
    ISSUED_AT   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FK_ISSUER   BIGSERIAL NOT NULL REFERENCES USERS(ID),
    FK_SUPPORT  BIGSERIAL REFERENCES USERS(ID)
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

/******************************
    Support User devices table
    Pub/Sub microservice-related
*******************************/
CREATE TABLE IF NOT EXISTS SUPPORT.USER_DEVICES(
    ID              BIGSERIAL NOT NULL PRIMARY KEY,
    FK_USER         BIGSERIAL NOT NULL REFERENCES USERS(ID),
    DEVICE_ADDRESS  TEXT NOT NULL,
    DEVICE_NAME     TEXT DEFAULT NULL,
    DEVICE_OS       TEXT DEFAULT NULL,
    CREATED_AT      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);