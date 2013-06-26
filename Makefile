generate:	table.go
	tools/generate.rb > table.go

demo:
	go run tools/demo.go

