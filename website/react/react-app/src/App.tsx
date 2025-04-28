// import { useState } from "react";
// import Alert from "./components/Alert";
// import Button from "./components/Button";

// function App() {
//     const [alertVisible, setAlertVisibility] = useState(false)

//     return (
//         <div>
//             { alertVisible && <Alert onClose={() => setAlertVisibility(false)}><h1>Adunarino you chickens</h1></Alert> }
//             <Button color="primary" onClick={() => setAlertVisibility(true)}>Adun, Button</Button>
//         </div>
//     )
// }

// export default App

import React from "react";
import { Route, Routes } from "react-router-dom";
import TestHTML from "./pages/testHTML";

function App() {

    return (
        <div style={{ textAlign: "center" }}>
            <Routes>
                <Route path="/" element={<h1>Memories are broken, the truth goes unspoken.</h1>} />
                <Route path="/test" element={<TestHTML />} />
            </Routes>
        </div>
    )
}

export default App;