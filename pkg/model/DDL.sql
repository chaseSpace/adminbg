/*
DDL: DATA DEFINE LANGUAGE

###
- Basic principle of database design
    -   Recommend
        -   Snake-case naming is recommended(besides index).

    -   Not-Recommend
        -   Foreign keys are not recommended.
*/

CREATE DATABASE IF NOT EXISTS adminbg;
USE adminbg;

-- User
DROP TABLE IF EXISTS adminbg_user;
CREATE TABLE adminbg_user
(
    uid           INT PRIMARY KEY AUTO_INCREMENT COMMENT 'uid is more used inside the system, unique',
    account_id    VARCHAR(50)                    NOT NULL COMMENT 'Use for sign-in, unique, cant modify in general',
    encrypted_pwd VARCHAR(40)                    NOT NULL,
    salt          VARCHAR(20)                    NOT NULL,
    nick_name     VARCHAR(50)                    NOT NULL DEFAULT '',
    phone         VARCHAR(50)                    NOT NULL DEFAULT '',
    email         VARCHAR(50)                    NOT NULL DEFAULT '',
    sex           ENUM ('MAN','WOMAN','UNKNOWN') NOT NULL DEFAULT 'UNKNOWN',
    remark        VARCHAR(100)                   NOT NULL DEFAULT '',
    status        ENUM ('NORMAL','SUSPEND')      NOT NULL DEFAULT 'NORMAL',
    created_at    DATETIME(3)                    NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    updated_at    DATETIME(3)                    NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    deleted_at    DATETIME(3)                    NULL,
    UNIQUE KEY `idx_account` (account_id),
    KEY `idx_status_deletedAt` (status, deleted_at)
)
    AUTO_INCREMENT = 1000
    ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4;


-- test data
-- >> Plain password is 123
INSERT INTO adminbg_user (encrypted_pwd, salt, account_id, uid)
VALUES ('85e25c1e193df1df5ada40fa52d3de6c713a242f', 'salt', 'admin', 1);
-- select sha1(concat('123','salt')) ==> '85e25c1e193df1df5ada40fa52d3de6c713a242f' ;


-- User_group_ref
DROP TABLE IF EXISTS adminbg_user_group_ref;
CREATE TABLE adminbg_user_group_ref
(
    id         INT PRIMARY KEY AUTO_INCREMENT,
    uid        INT         NOT NULL,
    group_id   INT         NOT NULL,
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    UNIQUE KEY `idx_uidGroupId` (uid, group_id)
)
    ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4;


-- Group
DROP TABLE IF EXISTS adminbg_usergroup;
CREATE TABLE adminbg_usergroup
(
    group_id   INT PRIMARY KEY AUTO_INCREMENT,
    group_name VARCHAR(50) NOT NULL,
    role_id    INT         NOT NULL DEFAULT 0,
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    deleted_at DATETIME(3) NULL,
    UNIQUE KEY `idx_deleteAt_groupName` (deleted_at, group_name),
    KEY `idx_roleId` (role_id)
)
    ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4;

-- You have to execute this two SQLs to insert a id=1 AUTO_INCREMENT column.
INSERT INTO adminbg_usergroup(group_id, group_name, role_id)
VALUES (1, 'DefaultGroup', 1);

-- Role
DROP TABLE IF EXISTS adminbg_role;
CREATE TABLE adminbg_role
(
    role_id    INT PRIMARY KEY AUTO_INCREMENT,
    role_name  VARCHAR(50)               NOT NULL,
    status     ENUM ('NORMAL','SUSPEND') NOT NULL DEFAULT 'NORMAL',
    created_at DATETIME(3)               NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    updated_at DATETIME(3)               NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    deleted_at DATETIME(3)               NULL,
    KEY `idx_deleteAt` (deleted_at),
    KEY `idx_status` (status)
)
    ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4;
-- You have to execute this two SQLs to insert a id=1 AUTO_INCREMENT column.
INSERT INTO adminbg_role(role_id, role_name, status)
VALUES (1, 'DefaultRole', 'NORMAL');

-- Role_mf_ref
DROP TABLE IF EXISTS adminbg_role_mf_ref;
CREATE TABLE adminbg_role_mf_ref
(
    id         INT PRIMARY KEY AUTO_INCREMENT,
    role_id    INT         NOT NULL,
    mf_id      INT         NOT NULL,
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    UNIQUE KEY `idx_roleIdMfId` (role_id, mf_id)
)
    ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4;


-- Menu_and_function
DROP TABLE IF EXISTS adminbg_menu_and_function;
CREATE TABLE adminbg_menu_and_function
(
    mf_id        INT PRIMARY KEY AUTO_INCREMENT,
    mf_name      VARCHAR(50)               NOT NULL,
    path         VARCHAR(50)               NOT NULL,
    parent_id    INT                       NOT NULL DEFAULT 100,
    level        TINYINT                   NOT NULL DEFAULT 0,
    type         ENUM ('MENU', 'FUNCTION') NOT NULL,
    menu_route   VARCHAR(100)              NOT NULL DEFAULT '' COMMENT 'Be used in the frontend',
    menu_display ENUM ('Y','N')            NOT NULL DEFAULT 'Y' COMMENT 'Be used in the frontend',
    sort_num     SMALLINT UNSIGNED         NOT NULL,
    created_at   DATETIME(3)               NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    updated_at   DATETIME(3)               NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    deleted_at   DATETIME(3)               NULL,
    KEY `idx_deleteAt` (deleted_at),
    KEY `idx_level` (level),
    KEY `idx_type` (type)
)
    ENGINE = InnoDB
    AUTO_INCREMENT = 2000
    DEFAULT CHARSET = utf8mb4;
-- DefaultMenu, also called root class menu cannot be deleted
INSERT INTO adminbg_menu_and_function (mf_id, mf_name, path, parent_id, level, type, menu_route, menu_display, sort_num)
VALUES (100, 'DefaultMenu', '100/', 0, 0, 'MENU', '', 'N', 0);

-- Mf_api_ref
DROP TABLE IF EXISTS adminbg_mf_api_ref;
CREATE TABLE adminbg_mf_api_ref
(
    id         INT PRIMARY KEY AUTO_INCREMENT,
    mf_id      INT         NOT NULL,
    api_id     INT         NOT NULL,
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    UNIQUE KEY `idx_mfIdApiId` (mf_id, api_id)
)
    ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4;


-- API
DROP TABLE IF EXISTS adminbg_api;
CREATE TABLE adminbg_api
(

    api_id     INT PRIMARY KEY AUTO_INCREMENT,
    identity   VARCHAR(50)  NOT NULL,
    remark     VARCHAR(100) NOT NULL DEFAULT '',
    created_at DATETIME(3)  NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    updated_at DATETIME(3)  NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    deleted_at DATETIME(3)  NULL,
    KEY `idx_deleteAt` (deleted_at),
    UNIQUE KEY `idx_identity` (identity)
)
    ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4;


-- ---------------------------- split line -----------------------------------

-- Operation_log
DROP TABLE IF EXISTS adminbg_operation_log;
CREATE TABLE adminbg_operation_log
(
    op_id       INT PRIMARY KEY AUTO_INCREMENT,
    type        ENUM ('[sign-in]','[sign-out]','[other]') NOT NULL,
    op_uid      INT                                       NOT NULL,
    op_username VARCHAR(50)                               NOT NULL DEFAULT '' COMMENT 'snapshot for username',
    remark      VARCHAR(50)                               NOT NULL,
    created_at  DATETIME(3)                               NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    updated_at  DATETIME(3)                               NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3)
)
