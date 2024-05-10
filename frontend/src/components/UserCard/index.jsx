import React from 'react';
import { Card } from 'antd';
import './index.css';

export function generateUserCard(userData) {
    if (!userData) {
        return null; // Return null if userData is not provided
    }

    const { name, email, ocrGithubUsername, fullTimeAvailability, partTimeAvailability, fullTimeSalary, partTimeSalary, tags } = userData;

    // Function to render tags content
    const renderTags = () => {
        if (!tags || tags.length === 0) {
            return null; // Return null if tags are not available
        }

        return tags.map(tag => (
            <Card key={tag} className="tagCard">
                <div className="tagCardContent">{tag}</div>
            </Card>
        ));
    };

    return (
        <Card className="tagCard" title={name || "Unknown"} >
            <Card>
                <p style={{ fontWeight: 'bold', fontSize: '26px' }}>Email: {email || "Not provided"}</p>
                <p style={{ fontWeight: 'bold', fontSize: '26px' }}>Github: {ocrGithubUsername || "Not provided"}</p>
                <p style={{ fontWeight: 'bold', fontSize: '26px' }}>Full-time availability: {fullTimeAvailability || "Not provided"}</p>
                <p style={{ fontWeight: 'bold', fontSize: '26px' }}>Part-time availability: {partTimeAvailability || "Not provided"}</p>
                <p style={{ fontWeight: 'bold', fontSize: '26px' }}>Full-time compensation: ${fullTimeSalary || "Not provided"}</p>
                <p style={{ fontWeight: 'bold', fontSize: '26px' }}>Part-time compensation: ${partTimeSalary || "Not provided"}</p>
            </Card>
            <div className="tagsContainer">{renderTags()}</div>
        </Card>
    );
}

export default generateUserCard;
