# Project Overview

A virtual file system with user, folder and file management using GoLang 1.20+.

## Table of Contents

- [Main Features](#main-features)
- [Commands](#commands)
- [Contact](#contact)

## Main Features

- Character Validation: Characters that cannot be included are [\\/:*?"<>|\s]. These characters are not allowed in usernames, folder names, or file names.
- Length Validation: Username must not exceed 50 characters, folder name must not exceed 100 characters, and file name must not exceed 255 characters.

## Commands

`register [username]`
</br>
</br>
<img src="./demo/register.gif" alt="register"/>
</br>
</br>

`create-folder [username] [foldername] [description]?`
</br>
</br>
<img src="./demo/create_folder.gif" alt="create_folder"/>
</br>
</br>

`delete-folder [username] [foldername]`
</br>
</br>
<img src="./demo/delete_folder.gif" alt="delete_folder"/>
</br>
</br>

`list-folders [username] [--sort-name|--sort-created] [asc|desc]`
</br>
</br>
<img src="./demo/list-folders.gif" alt="list-folders"/>
</br>
</br>

`rename-folder [username] [foldername] [new-folder-name]`
</br>
</br>
<img src="./demo/rename_folder.gif" alt="rename_folder"/>
</br>
</br>

`create-file [username] [foldername] [filename] [description]?`
</br>
</br>
<img src="./demo/create-file.gif" alt="create-file"/>
</br>
</br>

`delete-file [username] [foldername] [filename]`
</br>
</br>
<img src="./demo/delete-file.gif" alt="delete-file"/>
</br>
</br>

`list-files [username] [foldername] [--sort-name|--sort-created] [asc|desc]`
</br>
</br>
<img src="./demo/list-files.gif" alt="list-files"/>
</br>
</br>

## Contact

👨‍💻Wei-Han, Wang
<br/>

📧Email: s13602507586@gmail.com
