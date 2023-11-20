import Header from '../components/Header';
import {BrowserRouter as Router, Route, Routes} from "react-router-dom";
import Gobuster from "../components/Gobuster";
import React from "react";

const Main = ({ toggleMode }) => {
    return (
        <Router>
            <Header />
            <Routes>
                <Route path="/gobuster" element={<Gobuster/>} />
            </Routes>
        </Router>
    );
};

export default Main;
