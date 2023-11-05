# go resume api

* an excuse to write/refresh my skills with go and an interesting way of showing off my resume

### TODO:

- [ ] write the basic application
    - [ ] finish all top level api routes
- [ ] attach a database
    - [ ] decide which db and write a doc explaining why
    - [ ] have the api routes call the database
- [ ] finish rest of the api routes
- [ ] write tests
    - [ ] unit tests for basic api routes, integration tests for cross application functions
    - [ ] performance (load) tests
- [ ] create basic ci/cd pipeline
    - [ ] runs all tests
    - [ ] checks commit messages using gitlint
    - [ ] (stretch) tests using toxiproxy?
- [ ] host it via k8s
    - [ ] host database first
    - [ ] host application after
    - [ ] load test using t3a instances, then improve
- [ ] route it to kaizenc.dev/resume/api

### Stretch Goals

* full dynamic staging environments using seeded databases
* implement a self hosted logging solution (elk stack?)
* implement a self hosted monitoring solution (prometheus + grafana)
* [!!] create a frontend to match