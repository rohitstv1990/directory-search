## US 123456 : As a developer I want to write a feature will search for the records based on some keywords and return the details about the user
  - Story Point: 1
  ### Tasks
    - Create a json file which holds data about a user `(1 hr)`
    - Develop a go program which will retrieve the user info `(2 hr)`
    - Containerize  the app `(1hr)`

  ### Acceptance Criteria:
    - When the docker image should be run it should return all the matching records

## Simple go utility to find records from address book present in a json file

### Prerequisite :
 - golang
 - docker

### Step to build the docker image

 - `docker build -t dir-search:latest .`
### Steps to run the docker image to get the records
  - docker run `<image_name>` `<argument>`

### command to run the go file :
 - go run main.go `<argument>`


