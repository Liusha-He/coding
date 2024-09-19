interface GameOverProps {
    winner: (string | null);
    onReStart: () => void;
}

export default function GameOver({winner, onReStart}: GameOverProps): JSX.Element {
    return (
        <div id="game-over">
            <h2>Game Over!</h2>
            {winner && <p>{winner} won!</p>}
            {!winner && <p>It's a draw!</p>}
            <p>
                <button onClick={onReStart}>
                    Restart
                </button>
            </p>
        </div>
    )
}
