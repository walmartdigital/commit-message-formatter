build:
	packr -v -z

publish:
	rm -R ./dist
	goreleaser
	npm publish
