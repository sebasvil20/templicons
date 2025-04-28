build:
	go run ./cmd/generate
	templ generate

clean:
	rm -rf flags/*.templ
	rm -rf tabler/*.templ
