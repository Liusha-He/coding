import { useState } from "react";

import Player from "./components/Player";
import GameBoard from "./components/GameBoard";
import Log from "./components/Log";
import { GameTurn } from './models/GameTurn';

function deriveActivePlayer(gameTurns: GameTurn[]): string {
  let currentPlayer = "X";

  if (gameTurns.length > 0 && gameTurns[0].player === "X") {
    currentPlayer = "O";
  }
  
  return currentPlayer;
}

function App(): JSX.Element {
  const [gameTurns, setGameTurns] = useState<GameTurn[]>([]);
  const activePlayer = deriveActivePlayer(gameTurns);

  function handleSelectSquare(rowIndex: number, colIndex: number) {
    setGameTurns((prevTurns: GameTurn[]) => {
      const currentPlayer = deriveActivePlayer(prevTurns);

      const updatedTurns: GameTurn[] = [
        {
          square: {
            rowIndex: rowIndex,
            colIndex: colIndex,
          },
          player: currentPlayer
        },
        ...prevTurns];
        return updatedTurns;
    });
  }

  return (
    <main>
      <div id="game-container">
        <ol id="players" className="highlight-player">
          <Player 
          initialName='Player 1' 
          symbol='X' 
          isActive={activePlayer === "X"} />
          <Player 
          initialName='Player 2' 
          symbol='O' 
          isActive={activePlayer === "O"} />
        </ol>

        <GameBoard 
        onSelectSquare={handleSelectSquare}
        turns={gameTurns} />
      </div>
      <Log turns={gameTurns} />
    </main>
  );
}

export default App;
