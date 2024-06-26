.PHONY: release build test test_with_bash copy_course_file

current_version_number := $(shell git tag --list "v*" | sort -V | tail -n 1 | cut -c 2-)
next_version_number := $(shell echo $$(($(current_version_number)+1)))

release:
	git tag v$(next_version_number)
	git push origin main v$(next_version_number)

build:
	go build -o dist/main.out ./cmd/tester

test:
	TESTER_DIR=$(shell pwd) go test -v ./internal/

test_and_watch:
	onchange '**/*' -- go test -v ./internal/

copy_course_file:
	hub api \
		repos/codecrafters-io/build-your-own-grep/contents/course-definition.yml \
		| jq -r .content \
		| base64 -d \
		> internal/test_helpers/course_definition.yml

update_tester_utils:
	go get -u github.com/codecrafters-io/tester-utils

test_dev: build
	cd /Users/ryang/Developer/byox/craftinginterpreters && \
	CODECRAFTERS_SUBMISSION_DIR=/Users/ryang/Developer/byox/craftinginterpreters \
	CODECRAFTERS_TEST_CASES_JSON="[ \
		{\"slug\":\"ry8\",\"tester_log_prefix\":\"stage_101\",\"title\":\"Stage #1: Scanning: Empty File\"}, \
		{\"slug\":\"ol4\",\"tester_log_prefix\":\"stage_102\",\"title\":\"Stage #2: Scanning: Parenthese\"}, \
		{\"slug\":\"oe8\",\"tester_log_prefix\":\"stage_103\",\"title\":\"Stage #3: Scanning: Braces\"}, \
		{\"slug\":\"xc5\",\"tester_log_prefix\":\"stage_104\",\"title\":\"Stage #4: Scanning: single-character tokens\"} \
	]" \
	$(shell pwd)/dist/main.out
