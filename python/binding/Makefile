.SILENT:
.PHONY: init compile-dev compile-prod test

init:
	pip install maturin

compile-dev:
	maturin develop

test: compile-dev
	python -m main 100

compile:
	maturin develop --release

run: 
	python -m main 300
