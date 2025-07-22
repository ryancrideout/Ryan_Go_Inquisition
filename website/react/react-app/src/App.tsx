import React from "react";
import { Route, Routes } from "react-router-dom";
import TestHTML from "./pages/testHTML";

const JsonEndpoint = ({ endpoint }: {endpoint: string}) => {
    const data = {
        id: Math.floor(Math.random() * 1000),
        endpoint: endpoint,
        timestamp: new Date().toISOString(),
        users: Array.from({length: Math.floor(Math.random() * 10) + 1}, (_, i) => ({
            id: i + 1,
            name: `User${i + 1}`,
            email: `user${i + 1}@example.com`
        })),
        products: Array.from({length: Math.floor(Math.random() * 5) + 1}, (_, i) => ({
            id: i + 1,
            name: `Product ${i + 1}`,
            price: Math.floor(Math.random() * 100) + 10
        })),
        randomValue: Math.random(),
        status: 'success'
    };

    return <pre style={{ textAlign: 'left', padding: '20px' }}>{JSON.stringify(data, null, 2)}</pre>;
};

function App() {
    // Mock values to help with the endpoint setup.
    const letters = 'abcdefghijklmnopqrstuvwxyz'.split('');
    const numbers = Array.from({length: 1000}, (_, i) => i + 1);
    const categories = ['users', 'products', 'orders', 'analytics', 'reports', 'settings'];

    return (
        <div style={{ textAlign: "center" }}>
            <Routes>
                <Route path="/" element={<h1>Memories are broken, the truth goes unspoken.</h1>} />
                <Route path="/test" element={<TestHTML />} />
                {letters.map(letter => (
                    <Route key={letter} path={`/${letter}`} element={<JsonEndpoint endpoint={letter} />} />
                ))}
                {numbers.map(num => (
                    <Route key={num} path={`/endpoint${num}`} element={<JsonEndpoint endpoint={`endpoint${num}`} />} />
                ))}
                {categories.map(category => (
                    <Route key={category} path={`/${category}`} element={<JsonEndpoint endpoint={category} />} />
                ))}
                {/* <Route path="/a" element={<TestHTML />} />
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
                <Route path="/z" element={<TestHTML />} /> */}
            </Routes>
        </div>
    )
}

export default App;