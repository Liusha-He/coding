import { useState } from "react";

import Player from "./components/Player";
import GameBoard from "./components/GameBoard";
import Log from "./components/Log";
import GameOver from "./components/GameOver";
import { GameTurn } from './models/GameTurn';
import WINNING_COMBINATIONS from "./winning-combinations";

const initPlayers = {
  X: "Player 1",
  O: "Player 2",
}

const initialGameBoard: (string | null)[][] = [
  [null, null, null],
  [null, null, null],
  [null, null, null],
];

interface Players {
  [X: string]: string;
}

function deriveActivePlayer(gameTurns: GameTurn[]): string {
  let currentPlayer = "X";

  if (gameTurns.length > 0 && gameTurns[0].player === "X") {
    currentPlayer = "O";
  }

  return currentPlayer;
}

function App(): JSX.Element {
  const [players, setPlayers] = useState<Players>(initPlayers);
  const [gameTurns, setGameTurns] = useState<GameTurn[]>([]);
  const activePlayer = deriveActivePlayer(gameTurns);

  let gameBoard = [...initialGameBoard.map(array => [...array])];
  let winner: (string | null) = null;

  for (const turn of gameTurns) {
      const {square, player } = turn;
      const {rowIndex, colIndex} = square;

      gameBoard[rowIndex][colIndex] = player;
  };

  for (const combination of WINNING_COMBINATIONS) {
    const firstSymbol = gameBoard[combination[0].rowIndex][combination[0].colIndex];
    const secondSymbol = gameBoard[combination[1].rowIndex][combination[1].colIndex];
    const thirdSymbol = gameBoard[combination[2].rowIndex][combination[2].colIndex];
    
    if (firstSymbol && firstSymbol === secondSymbol && firstSymbol === thirdSymbol) {
      winner = players[firstSymbol];
    }
  };

  const hasDraw = (gameTurns.length === 9 && !winner);

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

  function handleRestart() {
    setGameTurns([]);
    // setPlayers(initPlayers);
  };

  function handlePlayers(symbol: string, newName: string) {
    setPlayers(prevPlayer => {
      return {
        ...prevPlayer,
        [symbol]: newName
      };
    });
  }

  return (
    <main>
      <div id="game-container">
        <ol id="players" className="highlight-player">
          <Player 
          initialName='Player 1' 
          symbol='X' 
          isActive={activePlayer === "X"} 
          onChangeName={handlePlayers} />
          <Player 
          initialName='Player 2' 
          symbol='O' 
          isActive={activePlayer === "O"} 
          onChangeName={handlePlayers} />
        </ol>
        {(winner || hasDraw) && <GameOver winner={winner} onReStart={handleRestart} />}
        <GameBoard 
        onSelectSquare={handleSelectSquare}
        board={gameBoard} />
      </div>
      <Log turns={gameTurns} />
    </main>
  );
}

export default App;
