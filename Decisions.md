## General Structure

I have never structured a backend using Go before (read: never used Go) so I am using this article as a baseline for file structure:
https://medium.com/geekculture/how-to-structure-your-project-in-golang-the-backend-developers-guide-31be05c6fdd9

### /cmd folder

Where main.go lives. Should be a relatively light file.

### /internal folder

Business logic, db interaction. Transport -> Business -> Database.

Transport: network layer, where end user interacts with application.
Business: Business logic that supoorts app's core functions.
DB: Interact w/ db

#### internal directories

/app: All dependencies and logic are collected and run the app. Run method is called by main.go

/config: Init general app configs

/database: files contain methods for interacting with db

/models: structures of db tables

/services: business logic of application

/transport: http server settings, handlers, ports, etc.

### /build folder

Config files for project build (docker)



