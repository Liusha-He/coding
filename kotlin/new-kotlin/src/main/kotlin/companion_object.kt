package syntax_sugar

class AnotherPerson {
    companion object {
        val add: (Int, Int) -> Int = { x, y ->
            x + y
        }

        val divide: (Double, Double) -> Double = { x, y ->
            if (y != 0.0) x / y else 0.0
        }

        val multiply: (Double, Double) -> Double = { x, y ->
            x * y
        }

        val reduce: (Int, Int) -> Int = { x ,y ->
            x - y
        }
    }

    fun addOne(a: Int): Int = add(a, 1)
    fun divideThree(a: Int): Double = divide(a.toDouble(), 3.0)
    fun reduceTwo(x: Int): Int = reduce(x, 2)
}

fun main() {
    val p = AnotherPerson()
    println(p.addOne(53))
    println(p.divideThree(43))
    println(p.reduceTwo(100))
}
