# adminbg

## Overview
The Minimum, the Simplest administration background system including authority management in Go. 

## Using third-party package
- gin
- gorm

No third party authority management package used.

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
    -   通用API
        -   获取当前用户可访问的菜单信息（包含子菜单，不包含也不需要功能）
-->

## Develop progress
It's started at 2020/12/10.

Detailed APIs:
-   SignIn related
    -   ✔️ /SignIn 
    -   ✔️ /SignOut 
-   Common APIs
    -   /GetAvailableMenuList(Contain child-menus, not contain child-functions.)
-   System management
    -   User management related
        -   ✔️ /NewUser 
        -   ✔️ /UpdateUser 
        -   ✔️ /QueryUser
        -   ✔️ /GetUserList(Only super administrator can call it in general.)
        -   /DeleteUser(But the data will remain.)
    -   UserGroup management
        -   ✔️ /NewUserGroup
        -   ✔️ /UpdateUserGroup(Support editing and deleting. Be careful Default-Group can't be deleted)
        -   ✔️ /QueryUserGroup
        -   ✔️ /GetUserGroupList(Only super administrator can call it in general.)
    -   Role management
        -   /NewRole
        -   /UpdateRole
        -   /GetRole
        -   /GetRoleList
        -   /DeleteRole(Default-Role can't be deleted)
    -   Menu&Function management(Contains Binding management of functions and APIs)
        -   ✔️ /NewMenu
        -   ✔️ /GetMenuList(Contain child-menus && functions of leaf-menus)
        -   ✔️ /UpdateMenu
        -   ✔️ /DeleteMenus(It's better to give a prompt to delete all child-menus and all child -functions at front end)
        -   ✔️ /NewFunction
        -   ✔️ /UpdateFunction
        -   ✔️ /GetAPIList(Filter by params)
        -   ✔️ /UpdateFuncAndAPIBindInfo(Bind/Unbind depends on params)
    -   API management
        -   ✔️ /NewAPI
        -   ✔️ /UpdateAPI(Update/Delete depends on params)
        -   ✔️ /GetAPIList(**Same as mentioned above**)
    -   Logging management
        -   /GetLog
        -   /DeleteLog(_Not Expose_)


Note: 
-   By default, system would not expose an API that can delete operation log(But it would be implemented).
-   In addition to the specific comments, almost all the Delete-APIs mentioned above are logical deletions.
