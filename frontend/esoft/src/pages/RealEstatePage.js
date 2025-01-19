import React, { useState } from 'react';
import RealEstateHouseTable from '../components/RealEstateHouseTable';
import RealEstateApartmentTable from '../components/RealEstateApartmentTable';
import RealEstateLandTable from '../components/RealEstateLandTable';
import RealEstateModal from '../components/RealEstateModal';
import {Link} from "react-router-dom";

const RealEstatePage = () => {
    const [properties, setProperties] = useState({
        house: [
            { city: 'Москва', street: 'Тверская', houseNumber: '10', floorCount: 2, roomCount: 5, area: 200, latitude: '55.7558', longitude: '37.6173' }
        ],
        apartment: [
            { city: 'Санкт-Петербург', street: 'Невский', houseNumber: '25', apartmentNumber: '12', floorCount: 5, roomCount: 3, area: 80, latitude: '59.9343', longitude: '30.3351' }
        ],
        land: [
            { city: 'Сочи', street: 'Донская', houseNumber: '15', area: 500, latitude: '43.6028', longitude: '39.7342' }
        ],
    });

    const [selectedType, setSelectedType] = useState('house');
    const [isModalOpen, setIsModalOpen] = useState(false);
    const [isEditMode, setIsEditMode] = useState(false);
    const [currentProperty, setCurrentProperty] = useState(null);

    const handleCreate = () => {
        setIsModalOpen(true);
        setIsEditMode(false);
    };

    const handleEdit = (property) => {
        setCurrentProperty(property);
        setIsModalOpen(true);
        setIsEditMode(true);
    };

    const handleDelete = (property) => {
        const updatedProperties = properties[selectedType].filter(p => p !== property);
        setProperties({ ...properties, [selectedType]: updatedProperties });
    };

    const handleSave = (property) => {
        if (isEditMode) {
            const updatedProperties = properties[selectedType].map(p => (p === currentProperty ? property : p));
            setProperties({ ...properties, [selectedType]: updatedProperties });
        } else {
            setProperties({
                ...properties,
                [selectedType]: [...properties[selectedType], property]
            });
        }
        setIsModalOpen(false);
    };

    const renderTable = () => {
        switch (selectedType) {
            case 'house':
                return <RealEstateHouseTable properties={properties.house} onEdit={handleEdit} onDelete={handleDelete} />;
            case 'apartment':
                return <RealEstateApartmentTable properties={properties.apartment} onEdit={handleEdit} onDelete={handleDelete} />;
            case 'land':
                return <RealEstateLandTable properties={properties.land} onEdit={handleEdit} onDelete={handleDelete} />;
            default:
                return null;
        }
    };

    return (
        <div>
            <header className="app-header">
                <div className="header-content">
                    <h1 style={{textAlign: 'center', flex: 1}}>Объекты недвижимости</h1>
                    <Link to="/" className="logo" title="На главную">
                        🏠
                    </Link>
                </div>
            </header>
            <div>
                <button onClick={() => setSelectedType('house')}>Дома</button>
                <button onClick={() => setSelectedType('apartment')}>Квартиры</button>
                <button onClick={() => setSelectedType('land')}>Земля</button>
            </div>
            <div>
                <button onClick={handleCreate}>Создать новый объект</button>
            </div>
            {renderTable()}
            {isModalOpen && (
                <RealEstateModal
                    property={currentProperty}
                    isEditMode={isEditMode}
                    onClose={() => setIsModalOpen(false)}
                    onSave={handleSave}
                />
            )}
        </div>
    );
};

export default RealEstatePage;
