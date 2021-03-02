# CHANGELOG

## v0.1.0 (2021-3-2)

### Added

- feat: test integration for user validate

- feat: return 403 forbidden when user reach task limit

- feat: add sqlite3 to configs

- feat: re-add sqlite

- feat: unittest for list tasks

- feat: moce sql.NullString to storage layer

- feat: implement task limit

- feat: split postgres from generic database/sql

- feat: ping db before do anything

- feat: return valid with error when validate user

- feat: implement sql string for user usecase

- feat: move jwt key to configs

- feat: use viper to init database and port

- feat: use viper for reading configs

- feat: init 3 layers in main

- feat: use pointer for usecase

- feat: init jwt generator

- feat: init transport layer

- feat: init usecase

- feat: no more sqlite because we use abstrace database/sql

- feat: complete add task in transport layer

- feat: use nowFn and uuidGenerateFn

- feat: implement add task in transport layer

- feat: implement task path in transport layer

- feat: implement user transports layer

- feat: split Storage to TaskStorage and UserStorage

- feat: skeleton transport and use case layer

- feat(storages): add storage interface

### Fixed

- fix: correct reaching limit task

- fix: return error in string for response

### Others

- docs: add structure in README

- test: integration for add task, list tasks

- chore: move integration configs inside tests

- chore: set up skeleton for integration tests

- chore(readme): add sqlite

- chore(readme): update wording

- chore: add make run and add guide to README

- chore: add Makefile for local test

- test: unittest for task use case reach limit

- test: unittest for task use case when database failed

- test: unittest for task usecase to list tasks

- test: full unit test for user usecase validate

- test(usecase): unit test for user use case

- chore: add gomock for storage layer

- chore: check error to fix lint

- chore: add github action

- chore: use postgres params for config

- chore: run local postgres without ssl

- build: update go.mod

- chore: ignore unwanted node modules

- refactor: move const to split file

- refactor: change module name to prevent import from original repo

- chore: add docker-compose with postgres

- change FE requirements

- add more explaination about the test in overview

- add requirement for using Redux for FE

- update requirement for readme

- add README with uml docs

- add README with uml docs

- add postman collection for test project

- Initial commit
