import React, {useRef} from "react";

interface ShoppingListFormProps {
    onAddItem: (item: string, quantity: number) => void;
}

export default function ShoppingListForm({ onAddItem }: ShoppingListFormProps): JSX.Element {
    const textInputRef = useRef<HTMLInputElement>(null);
    const numberInputRef = useRef<HTMLInputElement>(null);

    function handleSubmit(event: React.FormEvent) {
        event.preventDefault();

        const newProduct = textInputRef.current!.value;
        const quantity = parseInt(numberInputRef.current!.value);

        onAddItem(newProduct, quantity);

        textInputRef.current!.value = "";
        numberInputRef.current!.value = "";
    }

    return (
        <form onSubmit={handleSubmit}>
            <input type="text" placeholder="item name" ref={textInputRef}/>
            <input type="number" min={0} ref={numberInputRef} />
            <button type="submit">Add</button>
        </form>
    )
}
