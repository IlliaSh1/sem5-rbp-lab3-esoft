import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import HomePage from './pages/HomePage';
import ClientsPage from './pages/ClientsPage';
import './styles/App.css';
import RealtorsPage from "./pages/RealtorsPage";

const App = () => {
    return (
        <Router>
            <Routes>
                <Route path="/" element={<HomePage />} />
                <Route path="/clients" element={<ClientsPage />} />
                <Route path="/realtors" element={<RealtorsPage />} />
            </Routes>
        </Router>
    );
};

export default App;
