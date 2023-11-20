import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';

import Header from './components/Header';
import Gobuster from './components/Gobuster';
import Main from "./pages/Main";

function App() {
    return (
        <Main />
    );
}

export default App;