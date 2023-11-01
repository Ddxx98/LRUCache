import React, { useState } from 'react';
import './App.css';
import axios from 'axios'; // Import Axios

function App() {
    const [key, setKey] = useState('');
    const [value, setValue] = useState('');
    const [cacheValue, setCacheValue] = useState('');

    const fetchCacheValue = async () => {
        try {
            const response = await axios.get(`http://localhost:8080/get?key=${key}`);
            setCacheValue(response.data.value);
            console.log(response);
        } catch (error) {
            console.error(error);
            setCacheValue('Key not found');
        }
    };

    const setCacheKey = async () => {
        try {
            const response = await axios.post('http://localhost:8080/set', {
                key: key,
                value: value,
                expiration_seconds: 5, // Set your desired expiration time
            });
            console.log(key, value, response);
        } catch (error) {
            console.error(error);
        }
    };

    return (
        <div className="App">
            <h1>LRU Cache Demo</h1>
            <div className='setting'>
                <input
                    type="text"
                    placeholder="Key"
                    value={key}
                    onChange={(e) => setKey(e.target.value)}
                />
                <input
                    type="text"
                    placeholder="Value"
                    value={value}
                    onChange={(e) => setValue(e.target.value)}
                />
                <div>
                    <button onClick={setCacheKey}>Set Key</button>
                </div>
            </div>
            <div className='getting'>
                <div>
                    <input
                        type="text"
                        placeholder="Get Key"
                        value={key}
                        onChange={(e) => setKey(e.target.value)}
                    />
                    <div>
                        <button onClick={fetchCacheValue}>Get Value</button>
                    </div>
                </div>
                <div>
                    <p>Cache Value: {cacheValue}</p>
                </div>
            </div>
        </div>
    );
}

export default App;