/*
DDL: DATA DEFINE LANGUAGE

###
- Basic principle of database design
    -   Recommend
        -   Snake-case naming is recommended(besides index).

    -   Not-Recommend
        -   Foreign keys are not recommended.
*/

# User
DROP TABLE IF EXISTS adminbg_user;
CREATE TABLE adminbg_user
(
    uid           INT PRIMARY KEY AUTO_INCREMENT,
    encrypted_pwd varchar(40) NOT NULL,
    salt          varchar(10) NOT NULL,
    user_name     VARCHAR(50) NOT NULL,
    group_id      INT         NOT NULL DEFAULT 0,
    UNIQUE KEY `idx_userName` (user_name),
    KEY `idx_groupId` (group_id)
)
    ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4;


# Group
DROP TABLE IF EXISTS adminbg_user_group;
CREATE TABLE adminbg_user_group
(
    group_id   INT PRIMARY KEY AUTO_INCREMENT,
    group_name VARCHAR(50) NOT NULL,
    role_id    INT         NOT NULL DEFAULT 0,
    KEY `idx_roleId` (role_id)
)
    ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4;


# Role
DROP TABLE IF EXISTS adminbg_role;
CREATE TABLE adminbg_role
(
    role_id   INT PRIMARY KEY AUTO_INCREMENT,
    role_name VARCHAR(50) NOT NULL
)
    ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4;
