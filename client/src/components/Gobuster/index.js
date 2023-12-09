import React, { useState, useEffect } from 'react';

function Gobuster() {
    console.log('Component rendering...');
    const [domain, setDomain] = useState('');
    const [showOption, setShowOption] = useState('none'); // 'none', 'CNAME', 'IP'

    const [timeout, setTimeout] = useState('1s');
    const [threads, setThreads] = useState(10);
    const [results, setResults] = useState([]);
    const [loading, setLoading] = useState(false);


    const scanDomain = async () => {
        try {
            setLoading(true);
            console.log("fetching")
            const response = await fetch('http://localhost:12345/api/gobuster/scan', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    domain,
                    showOption,
                    timeout: `${timeout}`, // Make sure it's a string
                    threads,
                }),
            });
            console.log("fetched")
            setResults(await response.json());
        } catch (error) {
            console.error('Error fetching data:', error);
        }finally {
            setLoading(false);
        }
    };

    return (
        <div className="container mt-5 pt-5">
            <div className="row justify-content-center">
                <div className="col-md-6 col">
                    <div className="input-group">
                        <input
                            type="text"
                            className="form-control"
                            placeholder="Something to search for oleg"
                            aria-label="Something to search for oleg"
                            onChange={e => setDomain(e.target.value)}
                        />
                        <button className="btn btn-outline-secondary" type="button" onClick={scanDomain}>
                            Let's go!
                        </button>
                    </div>
                    <div className="mt-3">
                        <label>Show Option:</label>
                        <div className="form-check">
                            <input
                                type="radio"
                                className="form-check-input"
                                id="showNone"
                                checked={showOption === 'none'}
                                onChange={() => setShowOption('none')}
                            />
                            <label className="form-check-label" htmlFor="showNone">
                                None
                            </label>
                        </div>
                        <div className="form-check">
                            <input
                                type="radio"
                                className="form-check-input"
                                id="showCNAME"
                                checked={showOption === 'CNAME'}
                                onChange={() => setShowOption('CNAME')}
                            />
                            <label className="form-check-label" htmlFor="showCNAME">
                                Show CNAME
                            </label>
                        </div>
                        <div className="form-check">
                            <input
                                type="radio"
                                className="form-check-input"
                                id="showIPs"
                                checked={showOption === 'IP'}
                                onChange={() => setShowOption('IP')}
                            />
                            <label className="form-check-label" htmlFor="showIPs">
                                Show IPs
                            </label>
                        </div>
                    </div>
                    <div className="mt-3">
                        <label>Timeout:</label>
                        <input
                            type="text"
                            className="form-control"
                            value={timeout}
                            onChange={e => setTimeout(e.target.value)}
                        />
                    </div>
                    <div className="mt-3">
                        <label>Threads:</label>
                        <input
                            type="number"
                            className="form-control"
                            value={threads}
                            onChange={e => setThreads(parseInt(e.target.value, 10))}
                        />
                    </div>
                    {loading ? (
                        <p>Loading...</p>
                    ) : (
                        <ul className="list-group mt-5">
                            {results.map((item, index) => (
                                <li key={index}>{item}</li>
                            ))}
                        </ul>
                    )}
                </div>
            </div>
        </div>
    );
}

export default Gobuster;
