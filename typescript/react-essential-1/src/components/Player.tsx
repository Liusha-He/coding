import React, { useState } from "react";

interface PlayerProps {
    initialName: string;
    symbol: string;
    isActive: boolean;
    onChangeName: (symbol: string, newName: string) => void;
}

export default function Player({ initialName, symbol, isActive, onChangeName }: PlayerProps): JSX.Element {
    const [ playerName, setPlayerName ] = useState(initialName);
    const [ isEditing, setIsEditing ] = useState(false);

    function handleEdit() {
        setIsEditing((editing) => !editing );
        if (isEditing) {
            onChangeName(symbol, playerName);
        };
    }

    function handleChange(event: React.FormEvent<HTMLInputElement>) {
        console.log(event);
        setPlayerName(event.currentTarget.value);
    }

    let editablePlayerName = <span className="player-name">{ playerName }</span>;
    
    if (isEditing) {
        editablePlayerName = <input type="text" required value={playerName} onChange={handleChange} />
    }

    return (
        <li className={isActive ? "active" : undefined}>
            <span className='player'>
            { editablePlayerName }
            <span className='player-symbol'>{ symbol }</span>
            </span>
            <button onClick={handleEdit}>{ isEditing ? "Save" : "Edit" }</button>
          </li>
    )
}