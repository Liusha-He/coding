from invoke import task


@task(name="compile-dev")
def compile_def(c):
    c.run("""maturin develop""")


@task(name="test")
def test(c):
    c.run("""PYTHONPATH=./ python examples/test.py 100""")


@task(name="compile")
def compile(c):
    c.run("""maturin develop --release""")


@task(name="lab")
def lab(c):
    c.run("""jupyter lab""")
