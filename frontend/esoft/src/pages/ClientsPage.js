import React, { useState } from 'react';
import Sidebar from '../components/Sidebar';
import ClientList from '../components/ClientList';
import Modal from '../components/Modal';
import '../styles/ClientsPage.css';
import { Link } from 'react-router-dom';

const ClientsPage = () => {
    const [clients, setClients] = useState([
        {
            lastName: 'Кочкина',
            firstName: 'Валентина',
            middleName: 'Петровна',
            phone: '+79025673827',
            email: 'kochkin-86@yandex.ru',
            linkedToNeed: false,
            linkedToOffer: false,
        },
        {
            lastName: 'Фомина',
            firstName: 'Александра',
            middleName: 'Ивановна',
            phone: '+79376980876',
            email: 'aleksand_@gmail.com',
            linkedToNeed: false,
            linkedToOffer: false,
        },
    ]);

    const [filteredClients, setFilteredClients] = useState(clients);
    const [modalOpen, setModalOpen] = useState(false);
    const [currentClient, setCurrentClient] = useState(null);
    const [isEditMode, setIsEditMode] = useState(false);

    const handleCreate = () => {
        setIsEditMode(false);
        setCurrentClient({
            lastName: '',
            firstName: '',
            middleName: '',
            phone: '',
            email: '',
            linkedToNeed: false,
            linkedToOffer: false,
        });
        setModalOpen(true);
    };

    const handleUpdate = (client) => {
        setIsEditMode(true);
        setCurrentClient(client);
        setModalOpen(true);
    };

    const handleDelete = (clientToDelete) => {
        if (clientToDelete.linkedToNeed || clientToDelete.linkedToOffer) {
            alert('Нельзя удалить клиента, связанного с потребностью или предложением.');
            return;
        }
        const updatedClients = clients.filter((client) => client !== clientToDelete);
        setClients(updatedClients);
        setFilteredClients(updatedClients);
    };

    const handleSave = (newClient) => {
        if (isEditMode) {
            const updatedClients = clients.map((client) =>
                client === currentClient ? newClient : client
            );
            setClients(updatedClients);
            setFilteredClients(updatedClients);
        } else {
            setClients([...clients, newClient]);
            setFilteredClients([...clients, newClient]);
        }
        setModalOpen(false);
    };

    const handleSearch = (query) => {
        setFilteredClients(
            clients.filter((client) =>
                `${client.lastName} ${client.firstName} ${client.middleName}`
                    .toLowerCase()
                    .includes(query.toLowerCase())
            )
        );
    };

    return (
        <div className="clients-page">
            <header className="app-header">
                <div className="header-content">
                    <h1 style={{ textAlign: 'center', flex: 1 }}>Клиенты</h1>
                    <Link to="/" className="logo" title="На главную">
                        🏠
                    </Link>
                </div>
                <div className="stats">
                    <div className="stat-card">
                        <h3>Клиенты</h3>
                        <p>{clients.length}</p>
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
                <Sidebar onCreate={handleCreate} onSearch={handleSearch} />
                <ClientList
                    clients={filteredClients}
                    onUpdate={handleUpdate}
                    onDelete={handleDelete}
                />
            </div>
            {modalOpen && (
                <Modal
                    client={currentClient}
                    isEditMode={isEditMode}
                    onClose={() => setModalOpen(false)}
                    onSave={handleSave}
                />
            )}
        </div>
    );
};

export default ClientsPage;
