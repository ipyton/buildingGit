# buildingGit
This is a version control system which is implemented by golang.
It supports some basic functions of git.
./database directory is used for storing the commit and files
./index used for accelerating and specifying what files to commit.
./diff for calculating the difference between the files.
./pack is used to compress text files
./remote is used to communicate with the remote repository.

different modes:
REGULAR_MODE = "100644"
EXECUTABLE_MODE = "100755"
DIRECTORY_MODE = "40000"


the order of implementation. init,add,commit. status, checkout
repository contains config database, almost everything

tree is commited 
index is staged
