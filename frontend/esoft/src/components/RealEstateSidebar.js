import React, { useState } from "react";
import "../styles/RealEstateSidebar.css";
import RealEstateModal from "../components/RealEstateModal"; // импортируем модальное окно

const RealEstateSidebar = ({ setSelectedType }) => {
    const [isModalOpen, setIsModalOpen] = useState(false);
    const [selectedObjectType, setSelectedObjectType] = useState("house");

    // Функция для открытия модального окна для создания нового объекта
    const handleCreate = () => {
        setIsModalOpen(true);
    };

    const handleCloseModal = () => {
        setIsModalOpen(false);
    };

    return (
        <div className="sidebar">
            <h2>Выбор объекта</h2>
            <button onClick={() => setSelectedType("house")}>Дом</button>
            <button onClick={() => setSelectedType("land")}>Земля</button>
            <button onClick={() => setSelectedType("apartment")}>Квартира</button>
            <input
                type="text"
                className="search-input"
                placeholder="Поиск по объектам..."
            />
            <button onClick={handleCreate}>Создать</button>

            {isModalOpen && (
                <RealEstateModal
                    isEditMode={false}
                    onClose={handleCloseModal}
                    onSave={() => {}}
                />
            )}
        </div>
    );
};

export default RealEstateSidebar;
