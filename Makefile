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
		repos/codecrafters-io/build-your-own-interpreter/course-definition.yml \
		| jq -r .content \
		| base64 -d \
		> internal/test_helpers/course_definition.yml

update_tester_utils:
	go get -u github.com/codecrafters-io/tester-utils

test_dev: build
	cd /Users/ryang/Developer/byox/craftinginterpreters && \
	CODECRAFTERS_REPOSITORY_DIR=/Users/ryang/Developer/byox/craftinginterpreters \
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
	CODECRAFTERS_REPOSITORY_DIR=./craftinginterpreters/build/gen/chap04_scanning \
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

test_parsing_w_jlox: build
	CODECRAFTERS_REPOSITORY_DIR=./craftinginterpreters/build/gen/chap06_parsing \
	CODECRAFTERS_TEST_CASES_JSON="[ \
		{\"slug\":\"sc2\",\"tester_log_prefix\":\"stage_201\",\"title\":\"Stage #201: Parsing: Booleans\"}, \
		{\"slug\":\"ra8\",\"tester_log_prefix\":\"stage_202\",\"title\":\"Stage #202: Parsing: Number literals\"}, \
		{\"slug\":\"th5\",\"tester_log_prefix\":\"stage_203\",\"title\":\"Stage #203: Parsing: String literals\"}, \
		{\"slug\":\"xe6\",\"tester_log_prefix\":\"stage_204\",\"title\":\"Stage #204: Parsing: Parentheses\"}, \
		{\"slug\":\"mq1\",\"tester_log_prefix\":\"stage_205\",\"title\":\"Stage #205: Parsing: Unary operators\"}, \
		{\"slug\":\"wa9\",\"tester_log_prefix\":\"stage_206\",\"title\":\"Stage #206: Parsing: Factors\"}, \
		{\"slug\":\"yf2\",\"tester_log_prefix\":\"stage_207\",\"title\":\"Stage #207: Parsing: Terms\"}, \
		{\"slug\":\"uh4\",\"tester_log_prefix\":\"stage_208\",\"title\":\"Stage #208: Parsing: Comparison\"}, \
		{\"slug\":\"ht8\",\"tester_log_prefix\":\"stage_209\",\"title\":\"Stage #209: Parsing: Equality\"}, \
		{\"slug\":\"wz8\",\"tester_log_prefix\":\"stage_210\",\"title\":\"Stage #210: Parsing: Errors\"} \
	]" \
	$(shell pwd)/dist/main.out


test_evaluation_w_jlox: build
	CODECRAFTERS_REPOSITORY_DIR=./craftinginterpreters/build/gen/chap07_evaluating \
	CODECRAFTERS_TEST_CASES_JSON="[ \
		{\"slug\":\"iz6\",\"tester_log_prefix\":\"stage_301\",\"title\":\"Stage #301: Evaluation: Literals: Booleans & Nil\"}, \
		{\"slug\":\"lv1\",\"tester_log_prefix\":\"stage_302\",\"title\":\"Stage #302: Evaluation: Literals: Strings & Numbers\"}, \
		{\"slug\":\"oq9\",\"tester_log_prefix\":\"stage_303\",\"title\":\"Stage #303: Evaluation: Parentheses\"}, \
		{\"slug\":\"dc1\",\"tester_log_prefix\":\"stage_304\",\"title\":\"Stage #304: Evaluation: Unary operators\"}, \
		{\"slug\":\"bp3\",\"tester_log_prefix\":\"stage_305\",\"title\":\"Stage #305: Evaluation: Multiplicative operators\"}, \
		{\"slug\":\"jy2\",\"tester_log_prefix\":\"stage_306\",\"title\":\"Stage #306: Evaluation: Additive operators\"}, \
		{\"slug\":\"jx8\",\"tester_log_prefix\":\"stage_307\",\"title\":\"Stage #307: Evaluation: Concatenation operator\"}, \
		{\"slug\":\"et4\",\"tester_log_prefix\":\"stage_308\",\"title\":\"Stage #308: Evaluation: Relational operators\"}, \
		{\"slug\":\"hw7\",\"tester_log_prefix\":\"stage_309\",\"title\":\"Stage #309: Evaluation: Equality operators\"}, \
		{\"slug\":\"gj9\",\"tester_log_prefix\":\"stage_310\",\"title\":\"Stage #310: Evaluation: Runtime errors: Unary\"}, \
		{\"slug\":\"yu6\",\"tester_log_prefix\":\"stage_311\",\"title\":\"Stage #311: Evaluation: Runtime errors: Multiplication\"}, \
		{\"slug\":\"cq1\",\"tester_log_prefix\":\"stage_312\",\"title\":\"Stage #312: Evaluation: Runtime errors: Addition\"}, \
		{\"slug\":\"ib5\",\"tester_log_prefix\":\"stage_313\",\"title\":\"Stage #313: Evaluation: Runtime errors: Comparisons\"} \
	]" \
	$(shell pwd)/dist/main.out

test_statements_w_jlox: build
	CODECRAFTERS_REPOSITORY_DIR=./craftinginterpreters/build/gen/chap08_statements \
	CODECRAFTERS_TEST_CASES_JSON="[ \
		{\"slug\":\"xy1\",\"tester_log_prefix\":\"stage_401\",\"title\":\"Stage #401: Statements: Single print statements\"}, \
		{\"slug\":\"oe4\",\"tester_log_prefix\":\"stage_402\",\"title\":\"Stage #402: Statements: Multiple print statements\"}, \
		{\"slug\":\"fi3\",\"tester_log_prefix\":\"stage_403\",\"title\":\"Stage #403: Statements: Expression statements\"}, \
		{\"slug\":\"yg2\",\"tester_log_prefix\":\"stage_404\",\"title\":\"Stage #404: Statements: Variable: declaration\"}, \
		{\"slug\":\"sv7\",\"tester_log_prefix\":\"stage_405\",\"title\":\"Stage #405: Statements: Variable: runtime errors\"}, \
		{\"slug\":\"bc1\",\"tester_log_prefix\":\"stage_406\",\"title\":\"Stage #406: Statements: Variable: initialization\"}, \
		{\"slug\":\"dw9\",\"tester_log_prefix\":\"stage_407\",\"title\":\"Stage #407: Statements: Variable: redeclaration\"}, \
		{\"slug\":\"pl3\",\"tester_log_prefix\":\"stage_408\",\"title\":\"Stage #408: Statements: Assignment operation\"}, \
		{\"slug\":\"vr5\",\"tester_log_prefix\":\"stage_409\",\"title\":\"Stage #409: Statements: Block syntax\"}, \
		{\"slug\":\"fb4\",\"tester_log_prefix\":\"stage_410\",\"title\":\"Stage #410: Statements: Scope: nested & shadowing\"} \
	]" \
	$(shell pwd)/dist/main.out

test_control_flow_w_jlox: build
	CODECRAFTERS_REPOSITORY_DIR=./craftinginterpreters/build/gen/chap09_control \
	CODECRAFTERS_TEST_CASES_JSON="[ \
		{\"slug\":\"ne3\",\"tester_log_prefix\":\"stage_501\",\"title\":\"Stage #501: Control flow: If statements\"}, \
		{\"slug\":\"st5\",\"tester_log_prefix\":\"stage_502\",\"title\":\"Stage #502: Control flow: Else statements\"}, \
		{\"slug\":\"fh8\",\"tester_log_prefix\":\"stage_503\",\"title\":\"Stage #503: Control flow: If-else if-else statements\"}, \
		{\"slug\":\"xj4\",\"tester_log_prefix\":\"stage_504\",\"title\":\"Stage #504: Control flow: Nested if statements\"}, \
		{\"slug\":\"wk8\",\"tester_log_prefix\":\"stage_505\",\"title\":\"Stage #505: Control flow: Logical OR operator\"}, \
		{\"slug\":\"jx4\",\"tester_log_prefix\":\"stage_506\",\"title\":\"Stage #506: Control flow: Logical AND operator\"}, \
		{\"slug\":\"qy3\",\"tester_log_prefix\":\"stage_507\",\"title\":\"Stage #507: Control flow: While statements\"}, \
		{\"slug\":\"bw6\",\"tester_log_prefix\":\"stage_508\",\"title\":\"Stage #508: Control flow: For statements\"}, \
		{\"slug\":\"vt1\",\"tester_log_prefix\":\"stage_509\",\"title\":\"Stage #509: Control flow: Errors\"} \
	]" \
	$(shell pwd)/dist/main.out

test_functions_w_jlox: build
	CODECRAFTERS_REPOSITORY_DIR=./craftinginterpreters/build/gen/chap10_functions \
	CODECRAFTERS_TEST_CASES_JSON="[ \
		{\"slug\":\"av4\",\"tester_log_prefix\":\"stage_601\",\"title\":\"Stage #601: Functions: Native Functions\"}, \
		{\"slug\":\"pg8\",\"tester_log_prefix\":\"stage_602\",\"title\":\"Stage #602: Functions: No arguments\"}, \
		{\"slug\":\"lb6\",\"tester_log_prefix\":\"stage_603\",\"title\":\"Stage #603: Functions: With arguments\"}, \
		{\"slug\":\"px4\",\"tester_log_prefix\":\"stage_604\",\"title\":\"Stage #604: Functions: Syntax errors\"}, \
		{\"slug\":\"rd2\",\"tester_log_prefix\":\"stage_605\",\"title\":\"Stage #605: Functions: Return statements\"}, \
		{\"slug\":\"ey3\",\"tester_log_prefix\":\"stage_606\",\"title\":\"Stage #606: Functions: Higher order Functions\"}, \
		{\"slug\":\"fj7\",\"tester_log_prefix\":\"stage_607\",\"title\":\"Stage #607: Functions: Runtime Errors\"}, \
		{\"slug\":\"bz4\",\"tester_log_prefix\":\"stage_608\",\"title\":\"Stage #608: Functions: Scope\"}, \
		{\"slug\":\"gg6\",\"tester_log_prefix\":\"stage_609\",\"title\":\"Stage #609: Functions: Closures\"} \
	]" \
	$(shell pwd)/dist/main.out

test_resolving_w_jlox: build
	CODECRAFTERS_REPOSITORY_DIR=./craftinginterpreters/build/gen/chap11_resolving \
	CODECRAFTERS_TEST_CASES_JSON="[ \
		{\"slug\":\"de8\",\"tester_log_prefix\":\"stage_701\",\"title\":\"Stage #701: Resolving: Identifier Resolution\"}, \
		{\"slug\":\"pt7\",\"tester_log_prefix\":\"stage_702\",\"title\":\"Stage #704: Resolving: Self Initialization\"}, \
		{\"slug\":\"pz7\",\"tester_log_prefix\":\"stage_703\",\"title\":\"Stage #707: Resolving: Variable Re-declaration\"}, \
		{\"slug\":\"eh3\",\"tester_log_prefix\":\"stage_704\",\"title\":\"Stage #705: Resolving: Invalid Return\"} \
	]" \
	$(shell pwd)/dist/main.out

test_classes_w_jlox: build
	CODECRAFTERS_REPOSITORY_DIR=./craftinginterpreters/build/gen/chap12_classes \
	CODECRAFTERS_TEST_CASES_JSON="[ \
		{\"slug\":\"vf4\",\"tester_log_prefix\":\"stage_801\",\"title\":\"Stage #801: Class Declarations\"}, \
		{\"slug\":\"yk8\",\"tester_log_prefix\":\"stage_802\",\"title\":\"Stage #802: Class Instances\"}, \
		{\"slug\":\"yf3\",\"tester_log_prefix\":\"stage_803\",\"title\":\"Stage #803: Getters & Setters\"}, \
		{\"slug\":\"qr2\",\"tester_log_prefix\":\"stage_804\",\"title\":\"Stage #804: Instance Methods\"}, \
		{\"slug\":\"yd7\",\"tester_log_prefix\":\"stage_805\",\"title\":\"Stage #805: The `this` keyword\"}, \
		{\"slug\":\"dg2\",\"tester_log_prefix\":\"stage_806\",\"title\":\"Stage #806: Invalid usages of `this`\"}, \
		{\"slug\":\"ou5\",\"tester_log_prefix\":\"stage_807\",\"title\":\"Stage #807: Constructor calls\"}, \
		{\"slug\":\"eb9\",\"tester_log_prefix\":\"stage_808\",\"title\":\"Stage #808: Return within constructors\"} \
	]" \
	$(shell pwd)/dist/main.out

test_all: test_scanning_w_jlox test_parsing_w_jlox test_evaluation_w_jlox test_statements_w_jlox test_control_flow_w_jlox test_functions_w_jlox test_resolving_w_jlox test_classes_w_jlox