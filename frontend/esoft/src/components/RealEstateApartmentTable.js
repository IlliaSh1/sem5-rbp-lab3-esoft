import React from "react";
import "../styles/RealEstateTable.css";

const RealEstateApartmentTable = ({ properties, onEdit, onDelete }) => {
    return (
        <table className="real-estate-table">
            <thead>
            <tr>
                <th>Город</th>
                <th>Улица</th>
                <th>№ Дома</th>
                <th>№ Кв.</th>
                <th>Широта</th>
                <th>Долгота</th>
                <th>Этаж</th>
                <th>Кол-во комнат</th>
                <th>Площадь</th>
                <th>Действия</th>
            </tr>
            </thead>
            <tbody>
            {properties.map((property, index) => (
                <tr key={index}>
                    <td>{property.city}</td>
                    <td>{property.street}</td>
                    <td>{property.houseNumber}</td>
                    <td>{property.apartmentNumber}</td>
                    <td>{property.latitude}</td>
                    <td>{property.longitude}</td>
                    <td>{property.floorCount}</td>
                    <td>{property.roomCount}</td>
                    <td>{property.area} м²</td>
                    <td>
                        <button onClick={() => onEdit(property)}>Изменить</button>
                        <button onClick={() => onDelete(property)}>Удалить</button>
                    </td>
                </tr>
            ))}
            </tbody>
        </table>
    );
};

export default RealEstateApartmentTable;
