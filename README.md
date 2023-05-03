# IF2211 Strategi Algoritma - Tugas Besar 3 - Backend

## 🤖 ImaGenieKelar Bot
```
📢 "Simple ChatGPT using String Matching Algorithms and Regular Expression"
```

## **📜 Table of Contents**
* [Program Description](#-program-description)
* [Required Program](#%EF%B8%8F-required-program)
* [How to Run The Program](#-how-to-run-the-program-local)
* [Implementation Screenshots](#-implementation-screenshots)
* [Progress Report](#-progress-report)
* [Folders and Files Description](#-folders-and-files-description)
* [Author](#-authors)

## **📄 Program Description**
In this assignment, the program is required to develop a simple ChatGPT application by applying the simplest QA approach. The search for the most similar question to the question given by the user is done using the Knuth-Morris-Pratt (KMP) and Boyer-Moore (BM) string matching algorithms. Regex is used to determine the format of the question. If there is no exact match between the user's question and the questions in the database through the KMP or BM algorithms, then the chatbot will use the most similar question with at least 90% similarity. If there is no question with a similarity above 90%, then the chatbot will provide a maximum of 3 options for the most similar questions to be selected by the user. The similarity calculation is calculated using Levenshtein Distance. This application should also have several features such as text questions, calculator, date feature, add question and answer to the database, and delete question from the database. The classification is done using regex and is classified like everyday language.

## **🛠️ Required Program**
| Required Program | Reference Link                            |
|------------------|-------------------------------------------|
| Go (Golang)      | [Go (Golang)](https://go.dev/doc/install) |
| Go Fiber         | [Go Fiber](https://gofiber.io/)           |
| Air              | [Air](https://github.com/cosmtrek/air)    |
| PostgreSQL       | [PostgreSQL](https://www.postgresql.org/) |
| ReactJS          | [ReactJS](https://react.dev/)             |
| Chakra UI        | [ChakraUI](https://chakra-ui.com/)        |
| Docker           | [Docker](https://www.docker.com/)         |

## **💻 How to Run The Program (Local)**

1. Make .env file with this criteria
```sh
DB_HOST = db
DB_USERNAME = username
DB_NAME = your_db_name
DB_PASSWORD = password
```
2. Compose your docker containers
Your local server should serve at port 5000 for the backend services and port 5432 for the Postgres database <br>
```sh
docker compose up
```
3. Interract with Endpoints (Look up the documentation)
```sh
https://docs.google.com/document/d/1_3Z_u122nycifeYYGF1Ud6ceMrI4KzqmiWVB5Tq0-vA/edit#heading=h.rmo71eiaumdw
```

## **📷 Implementation Screenshots**

[//]: # (<img src="docs/assets/.png">)

## **📃 Progress Report**

| Features                                               | Yes      | No |
|--------------------------------------------------------|----------|----|
| Adding questions and answers to database               | &check;  |    |
| Removing questions and answers to database             | &check;  |    |
| Mathematical calculation queries                       | &check;  |    |
| Date-related queries                                   | &check;  |    |
| Queries with multiple questions and multiple answers	 | &check;  |    |
| Play games with the bot                                | &check;  |    |


## **📂 Folders and Files Description**
```bash
Backend
├───algorithms
│   ├───calculator
|       ├───calculator.go
|       └───deque.go
│   ├───date
|       ├───date.go
│   └───utils
|       ├───boyer_moore.go
|       ├───knuth_morris_pratt.go
|       └───levenshtein_distance.go
├───cmd
|   ├───main.go
|   └───routes.go
├───database
|   └───database.go
├───extra
|   ├───random_pick.go
|   └───rock_paper_scissor.go
├───handlers
│   ├───chat_session
|       ├───create_chat_session.go
|       ├───delete_chat_session.go
|       └───get_chat_session.go
│   ├───messages
|       ├───get_chat_messages.go
|       └───message_handler.go
│   ├───query_utils
|       ├───query_utils.go
│   ├───question_answer
|       └───get_qa.go
│   └───user_query
|       ├───qa_regex_handler.go
|       └───qa_string_matching_handler.go
├───models
|   ├───message.go
|   ├───question_answer.go
|   └───session.go
├───test
|   ├───boyermoore_test.go
|   ├───calculator_test.go
|   ├───levenshtein_test.go
|   ├───random_pick_test.go
|   └───rps_test.go
├───.air.toml
├───.env
├───docker-compose.yml
├───Dockerfile
├───go.mod
└───go.sum
```

## **👨‍💻 Authors**
| Name                      | Student ID | Role / Job Desk     |
|---------------------------|------------|---------------------|
| Juan Christopher Santoso  | 13521116   | Fullstack Developer |
| Nicholas Liem             | 13521135   | Backend Developer   |
| Nathania Calista Djunaedi | 13521139   | Frontend Developer  |
