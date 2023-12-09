import React, { useState } from 'react';
import Modal from 'react-modal';

function Nmap() {
    const [nmapParams, setNmapParams] = useState({
        targetSpecification: [],
        hostDiscovery: {
            listScan: false,
            pingScan: false,
            skipHostDiscovery: false,
        },
        portSpecification: {
            ports: '',
            excludePorts: '',
            fastMode: false,
            sequentialScan: false,
        },
        miscOptions: {
            enableIPv6Scan: false,
            enableOSDetection: false,
            enableVersionDetection: false,
            enableScriptScanning: false,
            enableTraceroute: false,
            printVersionNumber: false,
        },
    });
    const handleInputChange = (category, field, value) => {
        setNmapParams((prevParams) => ({
            ...prevParams,
            [category]: {
                ...prevParams[category],
                [field]: value,
            },
        }));
    };

    const downloadSvg = async () => {
        try {
            const svgResponse = await fetch('http://localhost:12345/api/nmap/report', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(nmapParams)
            });
            const svgBlob = await svgResponse.blob();
            const svgObjectURL = URL.createObjectURL(svgBlob);

            const link = document.createElement('a');
            link.href = svgObjectURL;
            link.download = 'downloaded_content.xml';
            document.body.appendChild(link);
            link.click();
            document.body.removeChild(link);
        } catch (error) {
            console.error('Error downloading xml:', error);
        }
    };

    const [modalIsOpen, setModalIsOpen] = useState(false);
    const [htmlContent, setHtmlContent] = useState('');
    const [loading, setLoading] = useState(false);

    const openModal = () => setModalIsOpen(true);
    const closeModal = () => setModalIsOpen(false);

    const scanDomain = async () => {
        try {
            setLoading(true);
            const response = await fetch('http://localhost:12345/api/nmap/scan', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(nmapParams),
            });

            const html = await response.text();
            setHtmlContent(html);
            openModal();
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
                    {/* Target Specification */}
                    <div className="mb-3">
                        <label className="form-label">Target Specification:</label>
                        <input
                            type="text"
                            className="form-control"
                            value={nmapParams.targetSpecification.join(', ')}
                            onChange={(e) =>
                                setNmapParams((prevParams) => ({
                                    ...prevParams,
                                    targetSpecification: e.target.value.split(',').map((s) => s.trim()),
                                }))
                            }
                        />
                    </div>

                    {/* Host Discovery Options */}
                    <div className="mb-3">
                        <label className="form-label">Host Discovery Options:</label>
                        <div className="form-check">
                            <input
                                type="checkbox"
                                className="form-check-input"
                                id="listScan"
                                checked={nmapParams.hostDiscovery.listScan}
                                onChange={(e) => handleInputChange('hostDiscovery', 'listScan', e.target.checked)}
                            />
                            <label className="form-check-label" htmlFor="listScan">
                                List Scan
                            </label>
                        </div>
                        <div className="form-check">
                            <input
                                type="checkbox"
                                className="form-check-input"
                                id="pingScan"
                                checked={nmapParams.hostDiscovery.pingScan}
                                onChange={(e) => handleInputChange('hostDiscovery', 'pingScan', e.target.checked)}
                            />
                            <label className="form-check-label" htmlFor="pingScan">
                                Ping Scan
                            </label>
                        </div>
                        <div className="form-check">
                            <input
                                type="checkbox"
                                className="form-check-input"
                                id="skipHostDiscovery"
                                checked={nmapParams.hostDiscovery.skipHostDiscovery}
                                onChange={(e) => handleInputChange('hostDiscovery', 'skipHostDiscovery', e.target.checked)}
                            />
                            <label className="form-check-label" htmlFor="skipHostDiscovery">
                                Skip Host Discovery
                            </label>
                        </div>
                    </div>

                    {/* Port Specification Options */}
                    <div className="mb-3">
                        <label className="form-label">Port Specification Options:</label>
                        <input
                            type="text"
                            className="form-control"
                            placeholder="Ports (e.g., 22, 80-100)"
                            value={nmapParams.portSpecification.ports}
                            onChange={(e) => handleInputChange('portSpecification', 'ports', e.target.value)}
                        />
                        <input
                            type="text"
                            className="form-control mt-2"
                            placeholder="Exclude Ports (e.g., 25, 8080)"
                            value={nmapParams.portSpecification.excludePorts}
                            onChange={(e) => handleInputChange('portSpecification', 'excludePorts', e.target.value)}
                        />
                        <div className="form-check mt-2">
                            <input
                                type="checkbox"
                                className="form-check-input"
                                id="fastMode"
                                checked={nmapParams.portSpecification.fastMode}
                                onChange={(e) => handleInputChange('portSpecification', 'fastMode', e.target.checked)}
                            />
                            <label className="form-check-label" htmlFor="fastMode">
                                Fast Mode
                            </label>
                        </div>
                        <div className="form-check">
                            <input
                                type="checkbox"
                                className="form-check-input"
                                id="sequentialScan"
                                checked={nmapParams.portSpecification.sequentialScan}
                                onChange={(e) => handleInputChange('portSpecification', 'sequentialScan', e.target.checked)}
                            />
                            <label className="form-check-label" htmlFor="sequentialScan">
                                Sequential Scan
                            </label>
                        </div>
                    </div>

                    {/* Misc Options */}
                    <div className="mb-3">
                        <label className="form-label">Misc Options:</label>
                        <div className="form-check">
                            <input
                                type="checkbox"
                                className="form-check-input"
                                id="enableIPv6Scan"
                                checked={nmapParams.miscOptions.enableIPv6Scan}
                                onChange={(e) => handleInputChange('miscOptions', 'enableIPv6Scan', e.target.checked)}
                            />
                            <label className="form-check-label" htmlFor="enableIPv6Scan">
                                Enable IPv6 Scan
                            </label>
                        </div>
                        <div className="form-check">
                            <input
                                type="checkbox"
                                className="form-check-input"
                                id="enableOSDetection"
                                checked={nmapParams.miscOptions.enableOSDetection}
                                onChange={(e) => handleInputChange('miscOptions', 'enableOSDetection', e.target.checked)}
                            />
                            <label className="form-check-label" htmlFor="enableOSDetection">
                                Enable OS Detection
                            </label>
                        </div>
                        {/* Add similar checkboxes for other Misc Options */}
                    </div>

                    {/* Button to trigger scan */}
                    <button className="btn btn-primary" onClick={scanDomain}>
                        Scan Domain
                    </button>
                    {loading ? (
                        <p>Loading...</p>
                    ) : (
                        <Modal
                            isOpen={modalIsOpen}
                            onRequestClose={closeModal}
                            contentLabel="HTML Content Modal"
                            style={{
                                overlay: {
                                    zIndex: 1000, // Adjust as needed
                                },
                                content: {
                                    zIndex: 1001, // Adjust as needed
                                    position: 'fixed', // or 'absolute'
                                },
                            }}
                        >
                            <div className="d-flex justify-content-end align-items-center mb-2">
                                <a
                                    href="#"
                                    onClick={downloadSvg}
                                    className="btn btn-primary ml-2"
                                    style={{
                                        padding: '10px',
                                    }}
                                >
                                    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor"
                                         className="bi bi-download" viewBox="0 0 16 16">
                                        <path
                                            d="M.5 9.9a.5.5 0 0 1 .5.5v2.5a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-2.5a.5.5 0 0 1 1 0v2.5a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2v-2.5a.5.5 0 0 1 .5-.5"/>
                                        <path
                                            d="M7.646 11.854a.5.5 0 0 0 .708 0l3-3a.5.5 0 0 0-.708-.708L8.5 10.293V1.5a.5.5 0 0 0-1 0v8.793L5.354 8.146a.5.5 0 1 0-.708.708l3 3z"/>
                                    </svg>
                                </a>
                                <button
                                    type="button"
                                    className="btn btn-danger close"
                                    onClick={closeModal}
                                    style={{
                                        padding: '10px',
                                    }}
                                >
                                    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor"
                                         className="bi bi-x" viewBox="0 0 16 16">
                                        <path
                                            d="M4.646 4.646a.5.5 0 0 1 .708 0L8 7.293l2.646-2.647a.5.5 0 0 1 .708.708L8.707 8l2.647 2.646a.5.5 0 0 1-.708.708L8 8.707l-2.646 2.647a.5.5 0 0 1-.708-.708L7.293 8 4.646 5.354a.5.5 0 0 1 0-.708"></path>
                                    </svg>
                                </button>
                            </div>
                            <div dangerouslySetInnerHTML={{__html: htmlContent}}/>
                        </Modal>
                    )}
                </div>
            </div>
        </div>
    )
        ;
}

export default Nmap;
