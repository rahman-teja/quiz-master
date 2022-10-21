# Quipper Quiz Master


This is a technical exercise for software engineering candidates.

## Task

- Quiz Master

## Tech

This assigment using

- [Go](https://go.dev/) - awesome language
- [Make](https://www.gnu.org/software/make/manual/make.html)

## Installation

Install the dependencies and start the server.
```sh
./bin/setup
```

For run apps

```sh
./bin/quiz_master
```

## Verify Cli Apps

Verify the apps by typing
```sh
help
```

This apps have 9 commands:

- create_question <no> <question> <answer> | Creates a question
- update_question <no> <question> <answer> | Updates a question
- delete_question <no> | Deletes a question
- question <no> | Shows a question
- questions | Shows question list
- answer_question <no_question> <answer> | Answer a question
- answers | Shows answer list
- help | show command list
- exit | exit quiz master

You can testing those command using:
```sh
- create_question 1 "How many letters are there in the English alphabet?" 26

- update_question 1 "How many vowels are there in the English alphabet?" 5

- delete_question 1

- question 1

- questions

- answer_question 1 5

- answers

- help

- exit

```

### Authors

**Rahman Teja Wicaksono** - *Candidate*  - *rahman.9h.23@gmail.com*
- [Gitlab](https://gitlab.com/rteja-library3)
- [Linkedin](https://www.linkedin.com/in/rahman-teja-wicaksono-644518191/)