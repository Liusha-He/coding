import { useState } from "react";
import { GameTurn } from "../models/GameTurn";

interface BoardProps {
    onSelectSquare: (rowIndex: number, colIndex: number) => void;
    // activePlayerSymbol: string;
    turns: GameTurn[];
}

const initialGameBoard: (string | null)[][] = [
    [null, null, null],
    [null, null, null],
    [null, null, null],
];

export default function GameBoard({ onSelectSquare, turns }: BoardProps): JSX.Element {
    let gameBoard = initialGameBoard;

    for (const turn of turns) {
        const {square, player } = turn;
        const {rowIndex, colIndex} = square;

        gameBoard[rowIndex][colIndex] = player;
    }
    
    return (
        <ol id="game-board">
            {gameBoard.map( (row: (string | null)[], rowIndex: number) => (
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
