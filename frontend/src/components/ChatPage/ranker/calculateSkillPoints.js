export function calculateSkillPoints(user, requirements, tags) {
    let points = 0;
    if(!requirements.IsSkillPresent)
        return 0;

    const skillNameList = user.skills.map(skill => skill.skillId)
    
    if(skillNameList.includes(requirements.Skills[0].Text))
    {
        points+=5000
        tags.push(requirements.Skills[0].Name)
    }

    if(requirements.Skills.length > 1 && skillNameList.includes(requirements.Skills[1].Text))
    {
        if(requirements.Skills[1].Score>0.45)
        points+=5000
        else
        points+=300

        tags.push(requirements.Skills[1].Name)
    }

    if(requirements.Skills.length > 2 && skillNameList.includes(requirements.Skills[2].Text))
    {
        if(requirements.Skills[1].Score>0.45)
            points+=1000
            else
            points+=100

        tags.push(requirements.Skills[1].Name)
    }
    

    

    if(user)
    return points;
}

