import React from 'react';
import '../styles/RealtorsPage.css';

const RealtorList = ({ realtors, onUpdate, onDelete }) => {
    return (
        <table className="realtor-table">
            <thead>
            <tr>
                <th>Фамилия</th>
                <th>Имя</th>
                <th>Отчество</th>
                <th>Доля от комиссии</th>
                <th>Действия</th>
            </tr>
            </thead>
            <tbody>
            {realtors.map((realtor, index) => (
                <tr key={index}>
                    <td>{realtor.lastName}</td>
                    <td>{realtor.firstName}</td>
                    <td>{realtor.middleName}</td>
                    <td>{realtor.commissionShare || '—'}</td>
                    <td className="actions">
                        <button onClick={() => onUpdate(realtor)}>Изменить</button>
                        <button onClick={() => onDelete(realtor)}>Удалить</button>
                    </td>
                </tr>
            ))}
            </tbody>
        </table>
    );
};

export default RealtorList;
