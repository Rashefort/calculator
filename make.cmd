@ECHO ON
go build -ldflags "-s" calculator.go syntax.go stack.go notations.go
RENAME "calculator.exe" "=.exe"
