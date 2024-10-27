import { GameTurn } from "../models/GameTurn";

interface LogProps {
    turns: GameTurn[];
}

export default function Log({ turns }: LogProps): JSX.Element {
    
    return (
        <ol id="log">
            {turns.map(turn => (
                <li key={`${turn.square.rowIndex}${turn.square.colIndex}`}>
                    {turn.player} selected {turn.square.rowIndex},{turn.square.colIndex}
                </li>
            ))}
        </ol>
    )
}
