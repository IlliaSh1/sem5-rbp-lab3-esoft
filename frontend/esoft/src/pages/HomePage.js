import React from 'react';
import { Link } from 'react-router-dom';
import '../styles/HomePage.css';

const HomePage = () => {
    return (
        <div className="app-container">
            {/* Header Section */}
            <header className="app-header">
                <div className="header-content">
                    <h1 style={{textAlign: 'center', flex: 1}}>Агентство недвижимости</h1>
                    <Link to="/" className="logo" title="На главную">
                        🏠
                    </Link>
                </div>
            </header>

            {/* Main Section */}
            <main className="app-main">
                <div className="grid-container">
                    <Link className="linkstyle" to="/clients">
                        <Card title="Клиенты" icon="👥"/>
                    </Link>
                    <Link className="linkstyle" to="/realtors">
                        <Card title="Риэлторы" icon="💼" />
                    </Link>
                    <Card title="Объекты недвижимости" icon="🏢" />
                    <Card title="Предложения" icon="🔄" />
                    <Card title="Потребности" icon="📋" />
                    <Card title="Сделки" icon="↔️" />
                </div>
            </main>
        </div>
    );
}

function Card({ title, icon }) {
    return (
        <div className="card">
            <div className="card-icon">{icon}</div>
            <div className="card-title">{title}</div>
        </div>
    );
}

export default HomePage;
