package syntax_sugar

interface Downloader {
    fun download()
}

interface Player {
    fun play()
}

class FileDownloader(private val file: String): Downloader {
    override fun download() {
        println("$file Downloaded")
    }
}

class FilePlayer(private val file: String): Player {
    override fun play() {
        println("$file Playing")
    }
}

class NonDelegatedMedia(
    private val downloader: Downloader,
    private val player: Player
): Downloader, Player {
    override fun download() {
        downloader.download()
    }

    override fun play() {
        player.play()
    }
}

class DelegatedMedia(
    private val downloader: Downloader,
    private val player: Player
): Downloader by downloader, Player by player {
    fun anotherMethod() {
        println("Another method...")
    }
}

fun main() {
    val file = "test.mp4"

    val media1 = NonDelegatedMedia(
        FileDownloader(file), FilePlayer(file)
    )

    media1.download()
    media1.play()

    val media2 = DelegatedMedia(
        FileDownloader(file), FilePlayer(file)
    )

    media2.download()
    media2.play()
    media2.anotherMethod()
}
