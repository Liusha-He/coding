package syntax_sugar

data class Person(
    var name: String,
    var age: Int,
    var employed: Boolean,
)

var number: Int? = 100
val str = "hello world"

fun main() {
    /*
    Use with/run to execute logic
     */
    val arr = with(str) {
        val arr = mutableListOf<String>()
        for (s in this) {
            if (s != ' ') arr.add(s.toString()) else continue
        }
        arr
    }
    println(arr)

    /*
    use let to execute logic with null check
     */
    val x = number?.let {
        val number2 = it + 1
        number2
    } ?: 0

    println(x)
    println(number)

    /*
    use apply to modify object
     */
    val liusha = Person("liusha", 34, true)
    liusha.apply {
        this.name = "Louis"
    }
    println(liusha)
}
