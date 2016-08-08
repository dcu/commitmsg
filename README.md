# Commit message validator

commitmsg checks that a commit message is well formatted.

## Install

```shell
go get -u -v github.com/dcu/commitmsg
```

## Check title

There are 2 options to check the title, one is `-c` which checks that the title is capitalized. The other one is `-l` which checks the max length of the title, example:

```shell
commitmsg hook .git/COMMIT-MSG -c -l 52
```

## Check body

The check body option `-b` checks that the body contains the given word, example:

```shell
commitmsg hook .git/COMMIT-MSG -b "JIRA-"
```

## Check author email

The check author email `-e` checks the domain of the email, example:

```shell
commitmsg hook .git/COMMIT-MSG -e "mycompany.com"
```

