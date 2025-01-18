import React, { useState } from 'react';
import '../styles/Sidebar.css';

const Sidebar = ({ onCreate, onUpdate, onDelete, onSearch }) => {
    const [search, setSearch] = useState('');

    const handleSearch = (e) => {
        setSearch(e.target.value);
        onSearch(e.target.value);
    };

    return (
        <aside className="sidebar">
            <input
                type="text"
                placeholder="Поиск по ФИО"
                className="search-input"
                value={search}
                onChange={handleSearch}
            />
            <div className="actions">
                <button onClick={onCreate}>Создать</button>
            </div>
        </aside>
    );
};

export default Sidebar;
