import { useState } from "react";

interface BoardProps {
    onSelectSquare: () => void;
    activePlayerSymbol: string;
}

const initialGameBoard = [
    [null, null, null],
    [null, null, null],
    [null, null, null],
];

export default function GameBoard({ onSelectSquare, activePlayerSymbol}: BoardProps): JSX.Element {
    const [ gameBoard, setGameBoard ] = useState(initialGameBoard);

    function handleSelectSquare(rowIndex: number, colIndex: number) {
        setGameBoard( (previousGameBoard: any[][]) => {
            const updatedBoard = [...previousGameBoard.map(innerArray => [...innerArray])];
            updatedBoard[rowIndex][colIndex] = activePlayerSymbol;
            console.log(updatedBoard[rowIndex][colIndex]);
            return updatedBoard;
        });
        onSelectSquare();
    }

    return (
        <ol id="game-board">
            {gameBoard.map( (row: string[] | null[], rowIndex: number) => (
                <li key={rowIndex}>
                    <ol>
                        {row.map( (playerSymbol: string | null, colIndex: number) => (
                            <li key={colIndex}>
                                <button onClick={() => handleSelectSquare(rowIndex, colIndex)}>{playerSymbol}</button>
                            </li>
                        ) )}
                    </ol>
                </li>
            ))}
        </ol>
    )
}
