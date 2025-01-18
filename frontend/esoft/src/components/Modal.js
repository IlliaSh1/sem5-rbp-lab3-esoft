import React, { useState } from 'react';
import '../styles/Modal.css';

const Modal = ({ client, isEditMode, onClose, onSave }) => {
    const [formData, setFormData] = useState(client || {
        lastName: '',
        firstName: '',
        middleName: '',
        phone: '',
        email: '',
    });
    const [errors, setErrors] = useState({});

    const validate = () => {
        const newErrors = {};
        const nameRegex = /^[А-ЯЁ][а-яё]+$/;
        if (!nameRegex.test(formData.lastName)) {
            newErrors.lastName = 'Фамилия должна начинаться с заглавной буквы и содержать только русские буквы.';
        }
        if (!nameRegex.test(formData.firstName)) {
            newErrors.firstName = 'Имя должно начинаться с заглавной буквы и содержать только русские буквы.';
        }
        if (!nameRegex.test(formData.middleName)) {
            newErrors.middleName = 'Отчество должно начинаться с заглавной буквы и содержать только русские буквы.';
        }
        const phoneRegex = /^\+7\d{10}$/;
        if (!phoneRegex.test(formData.phone)) {
            newErrors.phone = 'Телефон должен начинаться с +7 и содержать 11 цифр.';
        }
        const emailRegex = /^[a-zA-Z0-9._%+-]+@(yandex\.ru|mail\.ru|gmail\.com)$/;
        if (!emailRegex.test(formData.email)) {
            newErrors.email = 'Почта должна заканчиваться на @yandex.ru, @mail.ru или @gmail.com.';
        }
        setErrors(newErrors);
        return Object.keys(newErrors).length === 0;
    };

    const handleChange = (e) => {
        setFormData({ ...formData, [e.target.name]: e.target.value });
    };

    const handleSubmit = () => {
        if (validate()) {
            onSave(formData);
        }
    };

    return (
        <div className="modal-overlay">
            <div className="modal-content">
                <h2>{isEditMode ? 'Редактировать клиента' : 'Добавить клиента'}</h2>
                <input
                    type="text"
                    name="lastName"
                    placeholder="Фамилия"
                    value={formData.lastName}
                    onChange={handleChange}
                />
                {errors.lastName && <p className="error">{errors.lastName}</p>}
                <input
                    type="text"
                    name="firstName"
                    placeholder="Имя"
                    value={formData.firstName}
                    onChange={handleChange}
                />
                {errors.firstName && <p className="error">{errors.firstName}</p>}
                <input
                    type="text"
                    name="middleName"
                    placeholder="Отчество"
                    value={formData.middleName}
                    onChange={handleChange}
                />
                {errors.middleName && <p className="error">{errors.middleName}</p>}
                <input
                    type="text"
                    name="phone"
                    placeholder="Телефон"
                    value={formData.phone}
                    onChange={handleChange}
                />
                {errors.phone && <p className="error">{errors.phone}</p>}
                <input
                    type="email"
                    name="email"
                    placeholder="Электронная почта"
                    value={formData.email}
                    onChange={handleChange}
                />
                {errors.email && <p className="error">{errors.email}</p>}
                <div className="modal-actions">
                    <button className="action-button" onClick={handleSubmit}>
                        Сохранить
                    </button>
                    <button className="cancel-button" onClick={onClose}>
                        Отмена
                    </button>
                </div>
            </div>
        </div>
    );
};

export default Modal;
