import React, { useState } from 'react';
import './App.css';
import {v4 as getId} from "uuid";

import ShoppingList from './components/ShoppingList';
import ShoppingListForm from './components/ShoppingListForm';
import Item from './models/item';


function App() {
  const [items, setItems] = useState<Item[]>([]);

  const AddItem = (item: string, quantity: number) => {
    console.log("this is AddItem");
    setItems([...items, {id: getId(), product: item, quantity: quantity}]);
  };

  return (
    <div className="App">
      <ShoppingList items={items} />
      <ShoppingListForm onAddItem={AddItem} />
    </div>
  );
}

export default App;
