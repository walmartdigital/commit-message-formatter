build:
	packr2 build

publish:
	if [ -d ./dist ]; then rm -Rf ./dist; fi
	goreleaser
	npm publish
