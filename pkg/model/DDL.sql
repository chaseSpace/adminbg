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
    account_id    VARCHAR(50)                    NOT NULL COMMENT 'Use for sign in, unique, cant modify in general',
    encrypted_pwd VARCHAR(40)                    NOT NULL,
    salt          VARCHAR(20)                    NOT NULL,
    nick_name     VARCHAR(50)                    NOT NULL DEFAULT '',
    phone         VARCHAR(50)                    NOT NULL DEFAULT '',
    email         VARCHAR(50)                    NOT NULL DEFAULT '',
    sex           ENUM ('MAN','WOMAN','UNKNOWN') NOT NULL DEFAULT 'UNKNOWN',
    remark        VARCHAR(100)                   NOT NULL DEFAULT '',
    group_id      INT                            NOT NULL DEFAULT 0,
    status        ENUM ('NORMAL','SUSPEND')      NOT NULL DEFAULT 'normal',
    created_at    DATETIME(3)                    NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    updated_at    DATETIME(3)                    NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    deleted_at    DATETIME(3)                    NULL,
    UNIQUE KEY `idx_account` (account_id),
    KEY `idx_groupId` (group_id),
    KEY `idx_deletedAt` (deleted_at)
)
    ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4;
-- test data
-- >> Plain password is 123
INSERT INTO adminbg_user (encrypted_pwd, salt, account_id, uid)
VALUES ('85e25c1e193df1df5ada40fa52d3de6c713a242f', 'salt', 'admin', 1);
# select sha1(concat('123','salt')) = '85e25c1e193df1df5ada40fa52d3de6c713a242f' ;

# Group
DROP TABLE IF EXISTS adminbg_user_group;
CREATE TABLE adminbg_user_group
(
    group_id   INT PRIMARY KEY AUTO_INCREMENT,
    group_name VARCHAR(50) NOT NULL,
    role_id    INT         NOT NULL DEFAULT 0,
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    deleted_at DATETIME(3) NULL,
    KEY `idx_roleId` (role_id)
)
    ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4;


# Role
DROP TABLE IF EXISTS adminbg_role;
CREATE TABLE adminbg_role
(
    role_id    INT PRIMARY KEY AUTO_INCREMENT,
    role_name  VARCHAR(50) NOT NULL,
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    deleted_at DATETIME(3) NULL
)
    ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4;
