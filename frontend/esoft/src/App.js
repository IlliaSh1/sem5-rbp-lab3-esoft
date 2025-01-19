import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import HomePage from './pages/HomePage';
import ClientsPage from './pages/ClientsPage';
import './styles/App.css';
import RealtorsPage from "./pages/RealtorsPage";
import RealEstatePage from "./pages/RealEstatePage";

const App = () => {
    return (
        <Router>
            <Routes>
                <Route path="/" element={<HomePage />} />
                <Route path="/clients" element={<ClientsPage />} />
                <Route path="/realtors" element={<RealtorsPage />} />
                <Route path="/realestates" element={<RealEstatePage />} />
            </Routes>
        </Router>
    );
};

export default App;
