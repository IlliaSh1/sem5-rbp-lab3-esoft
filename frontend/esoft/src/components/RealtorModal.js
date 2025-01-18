import React, { useState } from 'react';
import '../styles/Modal.css';

const RealtorModal = ({ realtor, isEditMode, onClose, onSave }) => {
    const [formData, setFormData] = useState(realtor || {
        lastName: '',
        firstName: '',
        middleName: '',
        commissionShare: '',
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
        if (
            formData.commissionShare !== '' &&
            (isNaN(formData.commissionShare) || formData.commissionShare < 0 || formData.commissionShare > 100)
        ) {
            newErrors.commissionShare = 'Доля от комиссии должна быть числом от 0 до 100.';
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
                <h2>{isEditMode ? 'Редактировать риэлтора' : 'Добавить риэлтора'}</h2>
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
                    type="number"
                    name="commissionShare"
                    placeholder="Доля от комиссии"
                    value={formData.commissionShare}
                    onChange={handleChange}
                />
                {errors.commissionShare && <p className="error">{errors.commissionShare}</p>}
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

export default RealtorModal;
