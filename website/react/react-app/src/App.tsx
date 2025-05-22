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
                <Route path="/a" element={<TestHTML />} />
                <Route path="/b" element={<TestHTML />} />
                <Route path="/c" element={<TestHTML />} />
                <Route path="/d" element={<TestHTML />} />
                <Route path="/e" element={<TestHTML />} />
                <Route path="/f" element={<TestHTML />} />
                <Route path="/g" element={<TestHTML />} />
                <Route path="/h" element={<TestHTML />} />
                <Route path="/i" element={<TestHTML />} />
                <Route path="/j" element={<TestHTML />} />
                <Route path="/k" element={<TestHTML />} />
                <Route path="/l" element={<TestHTML />} />
                <Route path="/m" element={<TestHTML />} />
                <Route path="/n" element={<TestHTML />} />
                <Route path="/o" element={<TestHTML />} />
                <Route path="/p" element={<TestHTML />} />
                <Route path="/q" element={<TestHTML />} />
                <Route path="/r" element={<TestHTML />} />
                <Route path="/s" element={<TestHTML />} />
                <Route path="/t" element={<TestHTML />} />
                <Route path="/u" element={<TestHTML />} />
                <Route path="/v" element={<TestHTML />} />
                <Route path="/w" element={<TestHTML />} />
                <Route path="/x" element={<TestHTML />} />
                <Route path="/y" element={<TestHTML />} />
                <Route path="/z" element={<TestHTML />} />
            </Routes>
        </div>
    )
}

export default App;