# go-sqlite3-api

Simple api that serves up data from a sqlite3 database

Image built to run on arm7 using alpine base

## Build:

clone this repository  

`git clone https://github.com/Christian-Bull/go-sqlite3-api.git`

build docker image  

`docker build --tag <whateveryouwant> .`

run it  

`docker run -p 2000:2000 -e sqldatabase=$db -v <mount location on host>:/go-api/database <name of image>`
