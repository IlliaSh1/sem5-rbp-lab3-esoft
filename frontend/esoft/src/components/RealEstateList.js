import React from 'react';
import '../styles/RealEstatePage.css';

const RealEstateList = ({ properties, onUpdate, onDelete }) => {
    return (
        <table className="real-estate-table">
            <thead>
            <tr>
                <th>Тип</th>
                <th>Город</th>
                <th>Улица</th>
                <th>№ Дома</th>
                <th>№ Квартиры</th>
                <th>Широта</th>
                <th>Долгота</th>
                <th>Дополнительно</th>
                <th>Действия</th>
            </tr>
            </thead>
            <tbody>
            {properties.map((property, index) => (
                <tr key={index}>
                    <td>{property.type}</td>
                    <td>{property.city}</td>
                    <td>{property.street}</td>
                    <td>{property.houseNumber}</td>
                    <td>{property.apartmentNumber}</td>
                    <td>{property.latitude}</td>
                    <td>{property.longitude}</td>
                    <td>{property.additional}</td>
                    <td className="actions">
                        <button onClick={() => onUpdate(property)}>Изменить</button>
                        <button
                            onClick={() => {
                                if (property.linkedToOffer) {
                                    alert('Нельзя удалить объект, связанный с предложением.');
                                    return;
                                }
                                onDelete(property);
                            }}
                        >
                            Удалить
                        </button>
                    </td>
                </tr>
            ))}
            </tbody>
        </table>
    );
};

export default RealEstateList;
