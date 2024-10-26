.SILENT:
.PHONY: init compile-dev compile-prod test

init:
	pip install maturin

compile-dev:
	maturin develop

compile-prod:
	maturin develop --release

test:
	python -m main 1000
