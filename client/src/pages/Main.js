import Header from '../components/Header';
import Gobuster from "../components/Gobuster";
import Nmap from "../components/Nmap";
import {BrowserRouter as Router, Route, Routes} from "react-router-dom";
import React from "react";

const Main = ({ toggleMode }) => {
    return (
        <Router>
            <Header />
            <Routes>
                <Route path="/gobuster" element={<Gobuster/>} />
                <Route path="/nmap" element={<Nmap/>} />
            </Routes>
        </Router>
    );
};

export default Main;
