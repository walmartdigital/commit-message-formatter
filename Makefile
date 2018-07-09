build:
	cp package.json package/package.json
	packr -v -z

publish:
	rm -R ./dist
	goreleaser
	npm publish
