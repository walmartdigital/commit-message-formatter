build:
	cp package.json package/package.json
	packr -v -z

publish:
	if [ -d ./dist ]; then rm -Rf ./dist; fi
	goreleaser
	npm publish
