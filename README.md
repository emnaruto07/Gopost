# Gopost

## Objective
Create a command line application that allows interacting with a sqlite3 DB to create/retrieve/update/delete job posts

## Specification
The command will have five sub-commands:
1. create [Create]
2. get [Retrieve]
3. find [Searching]
4. update [Update]
5. delete [Delete]

These commands will be called from the terminal like:
```
gopost <command-name> <command-arguments>
```

for e.g.:
```
gopost create
```
```
gopost find "job name"
```

## Sub Commands
### create
This command will create a new job post in the DB and a new job id will be generated. The data will be prompted from the user like so:
```
$ gopost create
Creating a new job post, please the enter the following data:
name: <user enters name and presses enter>
salary: <user enters salary and presses enter>
Job was successfully posted to the DB with the ID = 237
```

### get
This command will return the job post data for the given id
```
$ gopost get 237
Job post found:
name: <from the DB>
salary: <from the DB>
```

### find
This command will find the jobs based on a search name
```
$ gopost find "developer"
[237] "senior developer required"
[ 40] "go developer needed"
```

### update
This command will update the job data of the given id, it will act just like create. It wont update a field if user enters an empty string
```
$ gopst update 237
Updating job by id 237, enter new values (just press enter keep a field unchanged):
name: <user just hits enters, so name wont be updated>
salary: <user enters a new salary, so it must be updated>
The following fields were update [salary]
```

### delete
This command deletes a job post by the given id
```
$ gopost delete 237
Job Post was deleted
```
