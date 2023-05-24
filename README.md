# Objective

Implement a virtual file system with user and file management capabilities using GoLang 1.20+.

# Assignment Details

Your task is to create a virtual file system as an interactive command line program. This system should
have the following capabilities:

1. User Management:

   - Allow users to register a unique, case insensitive username.
   - Users can have an arbitrary number of folders and files.

2. Folder Management:

   - Users can create, delete, and rename folders.
   - Folder names must be unique within the user's scope and are case insensitive.
   - Folders have an optional description field.

3. File Management:
   - Users can create, delete, and list all files within a specified folder.
   - File names must be unique within the same folder and are case insensitive.
   - Files have an optional description field.

# Command Specification

Your program should support the following commands:

Note:

- All the messages of successful or warning command executions should output to STDOUT.
- All the Error messages should output to STDERR.
- The token surrounded by [...] is a user input/variable.
- The question mark(?) within the [...]? indicates that token/user input is optional.

## 1. User Registration

- `register [username]`

  Response:

  - `Add [username] successfully.`
  - `Error: The [username] has already existed.`
  - `Error: The [username] contain invalid chars.`

## 2. Folder Management

- `create-folder [username] [foldername] [description]?`

  Response:

  - `Create [foldername] successfully.`
  - `Error: The [username] doesn't exist.`
  - `Error: The [foldername] contain invalid chars.`

- `delete-folder [username] [foldername]`

  Response:

  - `Delete [foldername] successfully.`
  - `Error: The [username] doesn't exist.`
  - `Error: The [foldername] doesn't exist.`

- `list-folders [username] [--sort-name|--sort-created] [asc|desc]`

  Response:

  - List all the folders within the `[username]` scope in following formats:  
    `[foldername] [description] [created at] [username]`

    Each field should be separated by whitespace or tab characters. The `[created at]` is a human-readable date/time format.

    The order of printed folder information is determined by the `--sort-name` or `--sort-created` combined with `asc` or `desc` flags.

    The `--sort-name` flag means sorting by `[foldername]` .

    If neither `--sort-name` nor `--sort-created` is provided, sort the list by `[foldername]` in ascending order.

  - Warning: The `[username]` doesn't have any folders.
  - Error: The `[username]` doesn't exist.
  - Prompt the user the usage of the command if there is an invalid flag.(should output to STDERR)

- `rename-folder [username] [foldername] [new-folder-name]`

  Response:

  - Rename `[foldername]` to `[new-folder-name]` successfully.
  - Error: The `[username]` doesn't exist.
  - Error: The `[foldername]` doesn't exist.

## 3. File Management

- `create-file [username] [foldername] [filename] [description]?`

  Response:

  - Create [filename] in [username] / [foldername] successfully.
  - Error: The [username] doesn't exist.
  - Error: The [foldername] doesn't exist.
  - Error: The [filename] contains invalid chars.

- `delete-file [username] [foldername] [filename]`

  Response:

  - Delete [filename] in [username] / [foldername] successfully.
  - Error: The [username] doesn't exist.
  - Error: The [foldername] doesn't exist.
  - Error: The [filename] doesn't exist.

- `list-files [username] [foldername] [--sort-name|--sort-created] [asc|desc]`

  Response:

  - List files with the following fields:  
    `[filename] [description] [created at] [foldername] [username]`

    Each field should be separated by whitespace or tab characters. The `[created at]` is a human-readable date/time format.

    The order of printed file information is determined by the `--sort-name` or `--sort-created` combined with `asc` or `desc` flags.

    The `--sort-name` means sorting by `[filename]` .

    If neither `--sort-name` nor `--sort-created` is provided, sort the list by `[filename]` in ascending order.

  - Warning: The folder is empty.
  - Error: The [username] doesn't exist.
  - Error: The [foldername] doesn't exist.
  - Prompt the user the usage of the command if there is an invalid flag.(should output to STDERR)

# Input Validation and Restrictions

The program is expected to perform input validation. If the command is not recognized by the
program,  
it should notify the user that it is an invalid command or display the usage of the program.
To simplify the implementation, it is acceptable for the `[username]` , `[foldername]`, [filename] , and `[new-folder-name]` not to contain whitespace characters.

As part of the input validation process, you should also define and enforce
user input restrictions for usernames, folder names, and file names,
such as maximum input length and valid character sets.
Please ensure that these restrictions are clearly documented in your README.md file.
Consider real-world constraints and potential issues when defining these restrictions,
keeping in mind the usability of the virtual file system and the need to prevent issues
such as excessively long inputs or invalid characters.

**[BONUS]** If you wish to challenge yourself, you can choose to implement a more advanced version that allows these tokens to have whitespace characters.  
In that case, all tokens with whitespace characters should be enclosed in double quotes,  
e.g., "New Folder" . Ensure that the program can handle such inputs correctly.

# Example

The "#" below is a prompt to inform the user that they can type commands. The following examples
demonstrate the usage of various commands in the virtual file system:
Register two users, user1 and user2

```zsh
# register user1
Add user1 successfully.
```

Register user2

```zsh
# register user2
Add user2 successfully.
```

Create a folder for user1 and user2 with the same folder name

```zsh
# create-folder user1 folder1
Create folder1 successfully.
# create-folder user2 folder1
Create folder1 successfully.
```

Attempt to create a folder with an existing name for user1

```zsh
# create-folder user1 folder1
Error: folder1 has already existed.
```

Create a folder with a description for user1

```zsh
# create-folder user1 folder2 this-is-folder-2
Create folder2 successfully.
```

List folders for user1 sorted by name in ascending order

```zsh
# list-folders user1 --sort-name asc
folder1 2023-01-01 15:00:00 user1
folder2 this-is-folder-2 2023-01-01 15:00:10 user1
```

List folders for user2

```zsh
# list-folders user2
folder1 2023-01-01 15:05:00 user2
```

Create a file with a description for user1 in folder1

```zsh
# create-file user1 folder1 file1 this-is-file1
Create file1 in user1/folder1 successfully.
```

Create a file named config with a description for user1 in folder1

```zsh
# create-file user1 folder1 config a-config-file
Create config in user1/folder1 successfully.
```

Attempt to create an existing file.

```zsh
# create-file user1 folder1 config a-config-file
Error: the config has already existed.
```

Attempt to create an file for an unregistered user.

```zsh
# create-file user-abc folder-abc config a-config-file
Error: The user-abc doesn't exist.
```

Attempt to type a unsupported command

```zsh
# list data
Error: Unrecognized command
```

Attempt to list files with incorrect flags

```zsh
# list-files user1 folder1 --sort a
Usage: list files [username] [foldername] [--sort-name|--sort-created] [asc|
desc]
```

```zsh
# list-files user1 folder1 --sort-name desc
file1 this-is-file1 2023-01-01 15:00:20 folder1 user1
config a-config-file 2023-01-01 15:00:30 folder1 user1
```
