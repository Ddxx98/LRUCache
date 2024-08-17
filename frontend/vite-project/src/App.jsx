import React, { useState, useEffect } from "react";
import "./App.css";
import axios from "axios"; // Import Axios

function App() {
    const [key, setKey] = useState("");
    const [formData, setFormData] = useState({
        key: '',
        value: '',
        time: ''
      });
    const [cacheValue, setCacheValue] = useState("");

    const handleChange = (e) => {
        setFormData({ ...formData, [e.target.name]: e.target.value });
      };

    // useEffect(() => {
    //     fetchCacheValue()
    // },[keydata])

    const fetchCacheValue = async () => {
        try {
            const response = await (await axios.get(`http://localhost:8080/get?key=${key}`))
            setCacheValue(response.data.value);
            console.log(response);
        } catch (error) {
            console.error(error);
            setCacheValue("Key not found");
        }
    };

    const setCacheKey = async () => {
        try {
            const response = await axios.post("http://localhost:8080/set", {
                key: formData.key,
                value: formData.value,
                expirationSeconds: Number(formData.time), // Set your desired expiration time (in seconds)formData.time, // Set your desired expiration time
            });
            console.log(formData.key, formData.value, formData.time, response);
            setFormData({
                key : '',
                name : '',
                time : '',
              });
        } catch (error) {
            console.error(error);
        }
        // console.log(formData);
    };

    return (
        <div className="App">
            <h1>LRU Cache Demo</h1>
            <div className="setting">
                <label>
                    <p>Key:</p>
                    <input
                        type="text"
                        name="key"
                        placeholder="Key"
                        value={formData.key}
                        onChange={handleChange}
                    />
                </label> 
                <label>
                    <p>Value:</p>
                    <input
                        type="text"
                        name="value"
                        placeholder="Value"
                        value={formData.value}
                        onChange={handleChange}
                    />
                </label>
                <label>
                    <p>Expiration time in seconds :</p>
                    <input
                        type="text"
                        name="time"
                        placeholder="Time in seconds"
                        value={formData.time}
                        onChange={handleChange}
                    />
                </label>
                <div>
                    <button onClick={setCacheKey} >Set Key</button>
                </div>
            </div>
            <div className="getting">
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
