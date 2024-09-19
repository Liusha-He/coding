interface BoardProps {
    onSelectSquare: (rowIndex: number, colIndex: number) => void;
    // activePlayerSymbol: string;
    // turns: GameTurn[];
    board: (string | null)[][];
}

export default function GameBoard({ onSelectSquare, board }: BoardProps): JSX.Element {
    return (
        <ol id="game-board">
            {board.map( (row: (string | null)[], rowIndex: number) => (
                <li key={rowIndex}>
                    <ol>
                        {row.map( (playerSymbol: (string | null), colIndex: number) => (
                            <li key={colIndex}>
                                <button 
                                onClick={() => onSelectSquare(rowIndex, colIndex)}
                                disabled={playerSymbol !== null}>
                                    {playerSymbol}
                                </button>
                            </li>
                        ) )}
                    </ol>
                </li>
            ))}
        </ol>
    )
}
