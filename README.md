# adminbg

#### Overview
The Minimum, the simplest administration background system including authority management in Go. 

#### Using third-party package
- gin
- gorm

No third party authority management package is be used.

#### Install

Coming...

#### Feature

-   User management
-   User group management
-   Role management
-   Menu management(Contains feature management)
-   Logging management

<!-- 
#### 前端功能
-  系统管理
	-	用户管理
    -	用户组管理（组管理、组绑定角色管理，有不可删的默认组）
	-	角色管理（有不可删的默认角色）
	-	菜单管理
	    -   菜单、以及叶子菜单下的功能管理（增删查改）
	    -   功能与API的关联管理
	-   API管理（单独开放给技术管理员角色）
	    -   增删查改
	-	日志管理
            -	登录日志
            -	操作日志
-->

#### Develop progress
This is just started.

Detailed:
-   SignIn related
    -   /SignIn √
    -   /SignOut √
-   User management related
    -   /NewUser √
    -   /ModifyUser √
    
<!-- 
详细
-   登录相关
    -   /SignIn √
    -   /SignOut √
-   用户管理相关
    -   /NewUser √
    -   /ModifyUser √
-->

<!-- 

#### 二次开发说明

**尽可能不在根目录下新增目录，业务代码只需写在pkg/目录中，可在pkg/目录下新建子目录**

作者保证本项目尽可能使用足够优秀的设计和简洁的代码实现，不会添加任何多余的功能。

-->