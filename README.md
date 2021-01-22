# adminbg

## Overview
The Minimum, the simplest administration background system including authority management in Go. 

## Using third-party package
- gin
- gorm

No third party authority management package is be used.

## Install

Coming...

## Feature

-   User management
-   User group management
-   Role management
-   Menu&Function management(Contains feature management)
-   API management
-   Logging management
    -   Operation logs(contains sign-in related logs)

<!-- 
## 前端功能
-  系统管理
	-	用户管理
	    -   增删改查
	    -   此页面包含对用户绑定组的操作（一个用户可绑定多个组）
    -	用户组管理
        -   有不可删的默认组
        -   增删改查
	-	角色管理
	    -   有不可删的默认角色
	    -   增删改查
	-	菜单管理
	    -   菜单、以及叶子菜单下的功能管理（增删查改）
	    -   此页面包含对功能绑定API的操作（一个功能可绑定多个API）
	-   API管理（单独开放给技术管理员角色）
	    -   增删查改（普通账户不应被授予API的任何管理权限）
	-	日志管理
            -	操作日志(包含登录相关)
-->

## Develop progress
It's been a few weeks.

Detailed:
-   SignIn related
    -   ✔️ /SignIn 
    -   ✔️ /SignOut 
-   System management
    -   User management related
        -   ✔️ /NewUser 
        -   ✔️ /UpdateUser 
        -   /GetUser(user-info)
        -   /GetUserList(user-info)
        -   /DeleteUsers
    -   UserGroup management
        -   /NewUserGroup
        -   /UpdateUserGroup
        -   /GetUserGroup
        -   /GetUserGroupList
        -   /DeleteUserGroup(Zero group-id group would be reserved)
    -   Role management
        -   /NewRole
        -   /UpdateRole
        -   /GetRole
        -   /GetRoleList
        -   /DeleteRole(Zero role-id role would be reserved)
    -   Menu&Function management(Contains Binding management of functions and APIs)
        -   /NewMenu
        -   /UpdateMenu
        -   /GetMenuList(Contain child-menus&functions of leaf-menu)
        -   /DeleteMenus
        -   /NewFunction
        -   /DeleteFunctions
        -   /GetAPIList(Filter by params)
        -   /UpdateFuncAndAPIBindingInfo(Bind/Unbind depends on params)
    -   API management
        -   /NewAPI
        -   /UpdateAPI
        -   /GetAPIList(As mentioned above)
        -   /DeleteAPI
    -   Logging management
        -   /GetLoggingList
        -   ~~/DeleteLogging~~
        
        
Note: 
-   Design by default, system would not provide an API that can delete operation logging.
  
<!-- 
详细
-   登录相关
    -   ✔️ /SignIn 
    -   ✔️ /SignOut 
-   用户管理相关
    -   /NewUser √
    -   /ModifyUser √
-->

<!-- 

## 二次开发说明

**尽可能不在根目录下新增目录，业务代码只需写在pkg/目录中，可在pkg/目录下新建子目录**

作者保证本项目尽可能使用足够优秀的设计和简洁的代码实现，不会添加任何多余的功能。

-->
