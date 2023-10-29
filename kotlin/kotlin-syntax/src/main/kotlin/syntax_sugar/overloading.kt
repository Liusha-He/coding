package syntax_sugar

fun printMessage() {
    println("Hello, World!")
}

fun printMessage(message: String) {
    println(message)
}

fun printMessage(message: String, times: Int) {
    for (i in 1..times) {
        println(message)
    }
}

fun main() {
    printMessage()               // Calls the first function
    printMessage("Hello Kotlin!") // Calls the second function
    printMessage("Hello Kotlin!", 3) // Calls the third function, printing the message 3 times
}
