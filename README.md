# Seacloud
Seacloud is a private cloud disk project that is written in Golang, and you can use Seacloud to store files, folders, and so on. Within the plan, Seacloud will support the following features:

Todo:
##1. user management
###1) User's classification
There are two types of users: ordinary users and administrators, after installing seacloud, a user (username: root, password: 123456) was already created.You can log on by this user.
###2) Add a user
Only administrators can add users
###3) Import user
Adding a user can only be added one by one, which is very inconvenient in some scenarios (such as a teacher wanting to add a lot of students at one time), so here is a feature for importing users. Support for importing users in a. txt,. xls file format, followed by LDAP support.
###4) Delete user(s)
Only administrators can delete users and support bulk deletion.
###5) User Login/Logout
Users can log in to the system after entering the correct user name and password. After you log in, you can exit the login, it will delete the session information, the next time you log on will enter the username and password again.
###6) Edit personal information
Users can edit personal information, mainly the following fields: User name, personal description, telephone, avatar and so on. Personal descriptions and telephone calls are optional.
###7) Modify password
User can modify password
###8) Log off
User can log off account
##2 File management
###1) Upload, download
This is the basic function.We also support large file uploads.
###2) Copy, move files
Supports copying, moving files/folders to specified directories
###3) Delete file
Support for deleting files
###4) File Recycle Bin
Deleted files are temporarily stored here, the file Recycle Bin has a default retention period of 30 days, more than 30 days of files will automatically clean up, do not occupy disk space. Files that have not exceeded the retention period can be restored
###5) Preview
Supports previews of common formats, such as PDFs, text files, pictures, audio, video, etc.
###6) File share
Files/folders support the generation of shared links, shared links can be encrypted, and other people (who may (not) be Seacloud users) can access files/folders through this shared link. Seacloud users can share files and folders with each other
##3 other
Support for advanced functions such as distributed, load balancing, etc.

Powered by Beego and Golang.
