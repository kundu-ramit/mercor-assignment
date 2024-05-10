import React from 'react';
import { Card, Button } from 'antd';

function BotSuggestions({data, prompt , setInputBoxValue}) {
  const { IsSkillPresent, IsBudgetPresent, IsExperiencePresent } = data
  return (
    <Card style={{borderRadius:30, marginBottom:"20px"}}>
      <p style={{ fontSize: '1.5em' }}>
        {!IsSkillPresent && 'Skill is mandatory for correct results. '}
        {!IsBudgetPresent && 'Budget is not present.'}
        {!IsExperiencePresent && 'Experience is not present'}
        {(!IsSkillPresent || !IsBudgetPresent || !IsExperiencePresent) && 'Please add them.'}
        {(IsSkillPresent || IsBudgetPresent || IsExperiencePresent) && 'You have added all necessary parameters.Edit this prompt or enter a new one'}
      </p>
      <Button type="primary" onClick={()=>setInputBoxValue(prompt)} style={{ marginTop: '1em', zoom: 1.2, borderRadius: '8px', boxShadow: '0 2px 8px rgba(0, 0, 0, 0.1)', border: 'none' }}>EDIT PROMPT</Button>
    </Card>
  );
}

export default BotSuggestions;