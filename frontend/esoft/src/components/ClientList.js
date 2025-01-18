import React from 'react';
import '../styles/ClientsPage.css';

const ClientList = ({ clients, onUpdate, onDelete }) => {
    return (
        <table className="client-table">
            <thead>
            <tr>
                <th>Фамилия</th>
                <th>Имя</th>
                <th>Отчество</th>
                <th>Номер телефона</th>
                <th>Электронная почта</th>
                <th>Действия</th>
            </tr>
            </thead>
            <tbody>
            {clients.map((client, index) => (
                <tr key={index}>
                    <td>{client.lastName}</td>
                    <td>{client.firstName}</td>
                    <td>{client.middleName}</td>
                    <td>{client.phone}</td>
                    <td>{client.email}</td>
                    <td className="actions">
                        <button onClick={() => onUpdate(client)}>Изменить</button>
                        <button onClick={() => onDelete(client)}>Удалить</button>
                    </td>
                </tr>
            ))}
            </tbody>
        </table>
    );
};

export default ClientList;
