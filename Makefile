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
		repos/codecrafters-io/build-your-own-interpreter/contents/course-definition.yml \
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
		{\"slug\":\"xc5\",\"tester_log_prefix\":\"stage_104\",\"title\":\"Stage #4: Scanning: Single-character tokens\"}, \
		{\"slug\":\"ea6\",\"tester_log_prefix\":\"stage_105\",\"title\":\"Stage #5: Scanning: Lexical errors\"}, \
		{\"slug\":\"mp7\",\"tester_log_prefix\":\"stage_106\",\"title\":\"Stage #6: Scanning: Equality operators\"}, \
		{\"slug\":\"bu3\",\"tester_log_prefix\":\"stage_107\",\"title\":\"Stage #7: Scanning: Negation operators\"}, \
		{\"slug\":\"et2\",\"tester_log_prefix\":\"stage_108\",\"title\":\"Stage #8: Scanning: Relational operators\"}, \
		{\"slug\":\"ml2\",\"tester_log_prefix\":\"stage_109\",\"title\":\"Stage #9: Scanning: Comments\"}, \
		{\"slug\":\"er2\",\"tester_log_prefix\":\"stage_110\",\"title\":\"Stage #10: Scanning: Whitespaces\"}, \
		{\"slug\":\"tz7\",\"tester_log_prefix\":\"stage_111\",\"title\":\"Stage #11: Scanning: Multi-line errors\"}, \
		{\"slug\":\"ue7\",\"tester_log_prefix\":\"stage_112\",\"title\":\"Stage #12: Scanning: String literals\"}, \
		{\"slug\":\"kj0\",\"tester_log_prefix\":\"stage_113\",\"title\":\"Stage #13: Scanning: Number literals\"}, \
		{\"slug\":\"ey7\",\"tester_log_prefix\":\"stage_114\",\"title\":\"Stage #14: Scanning: Identifiers\"}, \
		{\"slug\":\"pq5\",\"tester_log_prefix\":\"stage_115\",\"title\":\"Stage #15: Scanning: Reserved words\"} \
	]" \
	$(shell pwd)/dist/main.out


test_scanning_w_jlox: build
	CODECRAFTERS_SUBMISSION_DIR=./craftinginterpreters/build/gen/chap04_scanning \
	CODECRAFTERS_TEST_CASES_JSON="[ \
		{\"slug\":\"ry8\",\"tester_log_prefix\":\"stage_101\",\"title\":\"Stage #1: Scanning: Empty File\"}, \
		{\"slug\":\"ol4\",\"tester_log_prefix\":\"stage_102\",\"title\":\"Stage #2: Scanning: Parenthese\"}, \
		{\"slug\":\"oe8\",\"tester_log_prefix\":\"stage_103\",\"title\":\"Stage #3: Scanning: Braces\"}, \
		{\"slug\":\"xc5\",\"tester_log_prefix\":\"stage_104\",\"title\":\"Stage #4: Scanning: Single-character tokens\"}, \
		{\"slug\":\"ea6\",\"tester_log_prefix\":\"stage_105\",\"title\":\"Stage #5: Scanning: Lexical errors\"}, \
		{\"slug\":\"mp7\",\"tester_log_prefix\":\"stage_106\",\"title\":\"Stage #6: Scanning: Equality operators\"}, \
		{\"slug\":\"bu3\",\"tester_log_prefix\":\"stage_107\",\"title\":\"Stage #7: Scanning: Negation operators\"}, \
		{\"slug\":\"et2\",\"tester_log_prefix\":\"stage_108\",\"title\":\"Stage #8: Scanning: Relational operators\"}, \
		{\"slug\":\"ml2\",\"tester_log_prefix\":\"stage_109\",\"title\":\"Stage #9: Scanning: Comments\"}, \
		{\"slug\":\"er2\",\"tester_log_prefix\":\"stage_110\",\"title\":\"Stage #10: Scanning: Whitespaces\"}, \
		{\"slug\":\"tz7\",\"tester_log_prefix\":\"stage_111\",\"title\":\"Stage #11: Scanning: Multi-line errors\"}, \
		{\"slug\":\"ue7\",\"tester_log_prefix\":\"stage_112\",\"title\":\"Stage #12: Scanning: String literals\"}, \
		{\"slug\":\"kj0\",\"tester_log_prefix\":\"stage_113\",\"title\":\"Stage #13: Scanning: Number literals\"}, \
		{\"slug\":\"ey7\",\"tester_log_prefix\":\"stage_114\",\"title\":\"Stage #14: Scanning: Identifiers\"}, \
		{\"slug\":\"pq5\",\"tester_log_prefix\":\"stage_115\",\"title\":\"Stage #15: Scanning: Reserved words\"} \
	]" \
	$(shell pwd)/dist/main.out

test_all: test_dev test_scanning_w_jlox
