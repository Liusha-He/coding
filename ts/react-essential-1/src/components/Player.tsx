import React, { useState } from "react";

interface PlayerProps {
    initialName: string;
    symbol: string;
    isActive: boolean;
}

export default function Player({ initialName, symbol, isActive }: PlayerProps): JSX.Element {
    const [ playerName, setPlayerName ] = useState(initialName);
    const [ isEditing, setIsEditing ] = useState(false);

    function handleEdit() {
        setIsEditing((editing) => !editing );
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