export interface Square {
    rowIndex: number;
    colIndex: number;
}

export interface GameTurn {
    square: Square;
    player: string;
}
