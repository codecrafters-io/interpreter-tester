set -e

# Install dart
brew tap dart-lang/dart
brew install dart@2.19
echo 'export PATH="/opt/homebrew/opt/dart@2.19/bin:$PATH"' >> ~/.zshrc
source ~/.zshrc
dart --version
dart --disable-analytics

# Clone & compile craftinginterpreters
git clone https://github.com/munificent/craftinginterpreters.git
cd craftinginterpreters && make get
sed -i '' 's/Lox\.error(line, "Unexpected character\.");/Lox\.error(line, "Unexpected character: " + c);/g' java/com/craftinginterpreters/lox/Scanner.java
make java_chapters
cd ..
cp -r ./internal/test_helpers/jlox04/* ./craftinginterpreters/build/gen/chap04_scanning
cp -r ./internal/test_helpers/jlox06/* ./craftinginterpreters/build/gen/chap06_parsing
cp -r ./internal/test_helpers/jlox07/* ./craftinginterpreters/build/gen/chap07_evaluating
cp -r ./internal/test_helpers/jlox08/* ./craftinginterpreters/build/gen/chap08_statements
cp -r ./internal/test_helpers/jlox08/* ./craftinginterpreters/build/gen/chap09_control
cp -r ./internal/test_helpers/jlox08/* ./craftinginterpreters/build/gen/chap10_functions
cp -r ./internal/test_helpers/jlox08/* ./craftinginterpreters/build/gen/chap11_resolving
cp -r ./internal/test_helpers/jlox08/* ./craftinginterpreters/build/gen/chap12_classes
cp -r ./internal/test_helpers/jlox08/* ./craftinginterpreters/build/gen/chap13_inheritance
make test
