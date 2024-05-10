import React from 'react';
import { Card, Button } from 'antd';

const buttonStyle = { 
  border: '1px solid #ccc', 
  borderRadius: '5px', 
   
  cursor: 'pointer', 
  boxShadow: '0px 2px 4px rgba(0, 0, 0, 0.1)', 
  background: '#fff',
  display: 'flex',
  lineHeight: '0',
  padding : '20px',
  fontSize: '20px',
  fontWeight: 'bold',
  marginRight : '20px'
}

function BotSuggestions({data, prompt , setInputBoxValue, startIndex, pushToChat,rankedUsers, clearChat}) {
  const { IsSkillPresent, IsBudgetPresent, IsExperiencePresent } = data
  return (
    <Card style={{fontSize : "30px" ,borderRadius:30, marginBottom:"20px"}}>
      <p style={{ fontFamily: "'Roboto', sans-serif", fontWeight: 600 }}>
        {!IsSkillPresent && 'Skill is mandatory for correct results. It is missing. '}
        {!IsBudgetPresent && 'Budget is not present. '}
        {!IsExperiencePresent && 'Experience is not present. '}
        {(!IsSkillPresent || !IsBudgetPresent || !IsExperiencePresent) && 'Please add them.'}
        {(IsSkillPresent || IsBudgetPresent || IsExperiencePresent) && 'You have added all necessary parameters.Edit this prompt or enter a new one'}
      </p>
      <div style={{display:"flex"}}>
      <Button className='displayButton' onClick={() => setInputBoxValue(prompt)} style={buttonStyle}>EDIT PROMPT</Button>
      {IsSkillPresent?<Button className='displayButton' onClick={() => pushToChat(data,prompt, rankedUsers, startIndex+3)} style={buttonStyle}>SHOW MORE</Button>:<></>}
      <Button className='displayButton' onClick={() => clearChat()} style={buttonStyle}>CLEAR ALL</Button>
    </div>
    </Card>
  );
}

export default BotSuggestions;