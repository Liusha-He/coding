import { Square } from "./models/GameTurn";

const WINNING_COMBINATIONS: Square[][] = [
    [
        {rowIndex: 0, colIndex: 0},
        {rowIndex: 0, colIndex: 1},
        {rowIndex: 0, colIndex: 2},
    ],
    [
        {rowIndex: 1, colIndex: 0},
        {rowIndex: 1, colIndex: 1},
        {rowIndex: 1, colIndex: 2},
    ],
    [
        {rowIndex: 2, colIndex: 0},
        {rowIndex: 2, colIndex: 1},
        {rowIndex: 2, colIndex: 2},
    ],
    [
        {rowIndex: 0, colIndex: 0},
        {rowIndex: 1, colIndex: 0},
        {rowIndex: 2, colIndex: 0},
    ],
    [
        {rowIndex: 0, colIndex: 1},
        {rowIndex: 1, colIndex: 1},
        {rowIndex: 2, colIndex: 1},
    ],
    [
        {rowIndex: 0, colIndex: 2},
        {rowIndex: 1, colIndex: 2},
        {rowIndex: 2, colIndex: 2},
    ],
    [
        {rowIndex: 0, colIndex: 0},
        {rowIndex: 1, colIndex: 1},
        {rowIndex: 2, colIndex: 2},
    ],
    [
        {rowIndex: 0, colIndex: 2},
        {rowIndex: 1, colIndex: 1},
        {rowIndex: 2, colIndex: 0},
    ],
];

export default WINNING_COMBINATIONS;
