zip: clean
	CGO=0 GOOS=linux go build -o learn-lambda main.go
	zip handler.zip ./learn-lambda

clean:
	rm -f handler handler.zip learn-lambda