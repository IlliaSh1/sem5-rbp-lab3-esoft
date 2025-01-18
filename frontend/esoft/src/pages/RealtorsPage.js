import React, { useState } from 'react';
import Sidebar from '../components/Sidebar';
import RealtorList from '../components/RealtorList';
import RealtorModal from '../components/RealtorModal';
import '../styles/RealtorsPage.css';
import {Link} from "react-router-dom";

const RealtorsPage = () => {
    const [realtors, setRealtors] = useState([
        {
            lastName: 'Кочкина',
            firstName: 'Валентина',
            middleName: 'Петровна',
            commissionShare: '50',
            linkedToNeed: false,
            linkedToOffer: false,
        },
        {
            lastName: 'Фомина',
            firstName: 'Александра',
            middleName: 'Ивановна',
            commissionShare: '10',
            linkedToNeed: false,
            linkedToOffer: false,
        },
    ]);
    const [filteredRealtors, setFilteredRealtors] = useState(realtors);
    const [modalOpen, setModalOpen] = useState(false);
    const [currentRealtor, setCurrentRealtor] = useState(null);
    const [isEditMode, setIsEditMode] = useState(false);

    const handleCreate = () => {
        setIsEditMode(false);
        setCurrentRealtor({
            lastName: '',
            firstName: '',
            middleName: '',
            commissionShare: '',
        });
        setModalOpen(true);
    };

    const handleUpdate = (realtor) => {
        setIsEditMode(true);
        setCurrentRealtor(realtor);
        setModalOpen(true);
    };

    const handleDelete = (realtorToDelete) => {
        const updatedRealtors = realtors.filter(realtor => realtor !== realtorToDelete);
        setRealtors(updatedRealtors);
        setFilteredRealtors(updatedRealtors);
    };

    const handleSave = (newRealtor) => {
        if (isEditMode) {
            const updatedRealtors = realtors.map(realtor =>
                realtor === currentRealtor ? newRealtor : realtor
            );
            setRealtors(updatedRealtors);
            setFilteredRealtors(updatedRealtors);
        } else {
            setRealtors([...realtors, newRealtor]);
            setFilteredRealtors([...realtors, newRealtor]);
        }
        setModalOpen(false);
    };

    const handleSearch = (query) => {
        setFilteredRealtors(
            realtors.filter((realtor) =>
                `${realtor.lastName} ${realtor.firstName} ${realtor.middleName}`
                    .toLowerCase()
                    .includes(query.toLowerCase())
            )
        );
    };

    return (
        <div className="realtors-page">
            <header className="app-header">
                <div className="header-content">
                    <h1 style={{textAlign: 'center', flex: 1}}>Риэлторы</h1>
                    <Link to="/" className="logo" title="На главную">
                        🏠
                    </Link>
                </div>
                <div className="stats">
                    <div className="stat-card">
                        <h3>Риэлторы</h3>
                        <p>{realtors.length}</p>
                    </div>
                    <div className="stat-card">
                        <h3>Связаны с потребностью</h3>
                        <p>0</p>
                    </div>
                    <div className="stat-card">
                        <h3>Связаны с предложением</h3>
                        <p>0</p>
                    </div>
                </div>
            </header>
            <div className="content">
                <Sidebar onCreate={handleCreate} onSearch={handleSearch}/>
                <RealtorList
                    realtors={filteredRealtors}
                    onUpdate={handleUpdate}
                    onDelete={handleDelete}
                />
            </div>
            {modalOpen && (
                <RealtorModal
                    realtor={currentRealtor}
                    isEditMode={isEditMode}
                    onClose={() => setModalOpen(false)}
                    onSave={handleSave}
                />
            )}
        </div>
    );
};

export default RealtorsPage;
