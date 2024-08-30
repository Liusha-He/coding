import React from "react";

interface GreeterProps {
    name: string;
}

function Greeter({ name }: GreeterProps): JSX.Element {
    return <h1>Hello {name}!</h1>
}

// Also can do FunctionComponent 
// const Greeter: React.FC = () => {
//     return <h1>Hello</h1>
// };

export default Greeter;
