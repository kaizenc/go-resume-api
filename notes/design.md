# API Design Notes

## API routes and subroutes
* v1/health_check
* v1/job
    * v1/job/<company>
    * v1/job/range [POST]
* v1/education
    * v1/education/<college>
    * v1/education/range [POST]
* v1/interests
* v1/languages
    * v1/languages/<language>
* v1/frameworks
    * v1/frameworks/<framework>
* v1/projects
    * v1/project/<name>

## Job Object
* company
* position
* startdate
* enddate
* location
* responsibilities
* description
* languages
* frameworks

## Education Object
* College
* Degree
* GPA
* Honors Programs

## Language/Framework Objects
* name
* jobs used (return jobs themselves in array)
* projects used (return projects themselves in array)

## Project Object
* name
* description
* last updated
* languages
* framework

