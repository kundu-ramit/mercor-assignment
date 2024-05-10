import React from 'react';
import { Card } from 'antd';
import './index.css';

export function generateUserCard(userData) {
    const { name, email, ocrGithubUsername, fullTimeAvailability, partTimeAvailability, fullTimeSalary, partTimeSalary, tags } = userData;

    const tagsContent = tags.map(tag => (
        <Card key={tag} className="tagCard">
            <div className="tagCardContent">{tag}</div>
        </Card>
    ));

    return (
        <Card className="tagCard"  title={name}>
            <Card>
            <p style={{ fontWeight: 'bold', fontSize: '26px' }}>Email: {email}</p>
            <p style={{ fontWeight: 'bold', fontSize: '26px' }}>Github: {ocrGithubUsername}</p>
            <p style={{ fontWeight: 'bold', fontSize: '26px' }}>Full-time availability: {fullTimeAvailability}</p>
            <p style={{ fontWeight: 'bold', fontSize: '26px' }}>Part-time availability: {partTimeAvailability}</p>
            <p style={{ fontWeight: 'bold', fontSize: '26px' }}>Full-time compensation: ${fullTimeSalary}</p>
            <p style={{ fontWeight: 'bold', fontSize: '26px' }}>Part-time compensation: ${partTimeSalary}</p>
            </Card>
            <div className="tagsContainer">{tagsContent}</div>
        </Card>
    );
}

export default generateUserCard;