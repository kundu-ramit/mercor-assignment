import React from 'react';
import { Card } from 'antd';
import './index.css';
import styled from 'styled-components';

const StyledCard = styled(Card)`
  .ant-card-head-title {
    font-size: 30px;
    font-family: cursive;
  }
`;

export function generateUserCard(userData) {
    if (!userData) {
        return null; // Return null if userData is not provided
    }

    const { name,  tags } = userData;

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
        <StyledCard className="tagCard" title={name || "Unknown"} >
             <div style={{ fontSize: "30px", fontFamily: "cursive", fontWeight: "bold" }}>{generateDescription(userData)}</div>

            <div className="tagsContainer">{renderTags()}</div>
        </StyledCard>
    );
}

function generateDescription(inputObject) {
    let description = "";

    // Determine the introduction based on the name
    if (inputObject.name) {
        description += `${inputObject.name} `;
    }

    // Add total work experience
    if (inputObject.totalWorkExperience) {
        if (inputObject.totalWorkExperience === 1) {
            description += `has ${inputObject.totalWorkExperience} year of experience `;
        } else {
            description += `has ${inputObject.totalWorkExperience} years of experience `;
        }
    }

    // Add companies worked at
    if (inputObject.companiesWorkedAt && inputObject.companiesWorkedAt.length > 0) {
        description += `at companies like `;
        inputObject.companiesWorkedAt.forEach((company, index) => {
            if (index === inputObject.companiesWorkedAt.length - 1) {
                description += `and ${company} `;
            } else {
                description += `${company}, `;
            }
        });
    }

    // Add email
    if (inputObject.email) {
        description += `Her email is ${inputObject.email}. `;
    }

    // Add schools
    if (inputObject.schools && inputObject.schools.length > 0) {
        description += `She has studied at `;
        inputObject.schools.forEach((school, index) => {
            if (index === inputObject.schools.length - 1) {
                description += `and ${school}. `;
            } else {
                description += `${school}, `;
            }
        });
    }

    // Add phone
    if (inputObject.phone) {
        description += `Her phone number is ${inputObject.phone} and `;
    }

    // Add work availability
    if (inputObject.workAvailability) {
        description += `her work availability is ${inputObject.workAvailability}. `;
    }

    // Add full-time status
    if (inputObject.fullTimeStatus) {
        if (inputObject.fullTimeStatus.toLowerCase() === "yes") {
            description += `She is available for full-time `;
        } else if (inputObject.fullTimeStatus.toLowerCase() === "no") {
            description += `She is not available for full-time `;
        }
    }

    // Add skills
    if (inputObject.skills && inputObject.skills.length > 0) {
        description += `and is skilled in `;
        inputObject.skills.forEach((skill, index) => {
            if (index === inputObject.skills.length - 1) {
                description += `${skill.skillName}: ${skill.skillValue}.`;
            } else {
                description += `${skill.skillName}: ${skill.skillValue}, `;
            }
        });
    }

    return description;
}


export default generateUserCard;
