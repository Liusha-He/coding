.SILENT:
.PHONY: init compile-dev compile-prod test

init:
	pip install maturin

compile-dev:
	maturin develop

test: 
	pytest tests/ -svv --cov --cov-fail-under=90

compile:
	maturin develop --release

lab:
	jupyter lab

# run: 
# 	python -m main 300
