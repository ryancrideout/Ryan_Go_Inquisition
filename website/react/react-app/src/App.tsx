// import ListGroup from "./components/ListGroup"

// function App() {
//     const items = [
//         'New York',
//         'San Francisco',
//         'Tokyo',
//         'London',
//         'Paris'
//     ]

//     const handleSelectItem = (item: string) => {
//         console.log(item);
//     }

//     return <div><ListGroup items={items} heading="Cities" onSelectItem={handleSelectItem}></ListGroup></div>
// }

import { useState } from "react";
import Alert from "./components/Alert";
import Button from "./components/Button";

function App() {
    const [alertVisible, setAlertVisibility] = useState(false)

    return (
        <div>
            { alertVisible && <Alert onClose={() => setAlertVisibility(false)}><h1>Adunarino you chickens</h1></Alert> }
            <Button color="primary" onClick={() => setAlertVisibility(true)}>Adun, Button</Button>
        </div>
    )
}

export default App