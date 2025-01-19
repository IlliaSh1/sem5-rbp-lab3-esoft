import React, { useState, useEffect } from 'react';
import '../styles/RealEstateModal.css';

const RealEstateModal = ({ property, isEditMode, onClose, onSave }) => {
    const [formData, setFormData] = useState({
        type: 'Дом',
        city: '',
        street: '',
        houseNumber: '',
        apartmentNumber: '',
        latitude: '',
        longitude: '',
        floorCount: '',
        roomCount: '',
        area: '',
        additional: '',
        linkedToOffer: false,
    });

    useEffect(() => {
        if (isEditMode && property) {
            setFormData(property);
        }
    }, [isEditMode, property]);

    const handleChange = (e) => {
        setFormData({ ...formData, [e.target.name]: e.target.value });
    };

    const handleSubmit = () => {
        const latitude = parseFloat(formData.latitude);
        const longitude = parseFloat(formData.longitude);

        if (latitude < -90 || latitude > 90) {
            alert('Широта должна быть от -90 до 90.');
            return;
        }

        if (longitude < -180 || longitude > 180) {
            alert('Долгота должна быть от -180 до 180.');
            return;
        }

        if (formData.area && isNaN(formData.area)) {
            alert('Площадь должна быть числом.');
            return;
        }

        onSave(formData);
    };

    return (
        <div className="modal-overlay">
            <div className="modal-content">
                <h2>{isEditMode ? 'Редактировать объект' : 'Добавить объект'}</h2>
                <select name="type" value={formData.type} onChange={handleChange}>
                    <option value="Дом">Дом</option>
                    <option value="Квартира">Квартира</option>
                    <option value="Земля">Земля</option>
                </select>
                <input
                    type="text"
                    name="city"
                    placeholder="Город"
                    value={formData.city}
                    onChange={handleChange}
                    pattern="^[А-ЯЁ][а-яё]*$"
                    title="Только русские буквы, первая заглавная"
                />
                <input
                    type="text"
                    name="street"
                    placeholder="Улица"
                    value={formData.street}
                    onChange={handleChange}
                    pattern="^[А-ЯЁ][а-яё]*$"
                    title="Только русские буквы, первая заглавная"
                />
                <input
                    type="number"
                    name="houseNumber"
                    placeholder="№ Дома"
                    value={formData.houseNumber}
                    onChange={handleChange}
                />
                <input
                    type="number"
                    name="apartmentNumber"
                    placeholder="№ Квартиры"
                    value={formData.apartmentNumber}
                    onChange={handleChange}
                />
                <input
                    type="number"
                    name="latitude"
                    placeholder="Широта"
                    value={formData.latitude}
                    onChange={handleChange}
                    min="-90"
                    max="90"
                />
                <input
                    type="number"
                    name="longitude"
                    placeholder="Долгота"
                    value={formData.longitude}
                    onChange={handleChange}
                    min="-180"
                    max="180"
                />
                {formData.type === "Дом" || formData.type === "Квартира" ? (
                    <>
                        <input
                            type="number"
                            name="floorCount"
                            placeholder="Этажность"
                            value={formData.floorCount}
                            onChange={handleChange}
                        />
                        <input
                            type="number"
                            name="roomCount"
                            placeholder="Кол-во комнат"
                            value={formData.roomCount}
                            onChange={handleChange}
                        />
                    </>
                ) : null}
                <input
                    type="number"
                    name="area"
                    placeholder="Площадь"
                    value={formData.area}
                    onChange={handleChange}
                />
                <textarea
                    name="additional"
                    placeholder="Дополнительная информация"
                    value={formData.additional}
                    onChange={handleChange}
                ></textarea>
                <button onClick={handleSubmit}>Сохранить</button>
                <button onClick={onClose}>Закрыть</button>
            </div>
        </div>
    );
};

export default RealEstateModal;
