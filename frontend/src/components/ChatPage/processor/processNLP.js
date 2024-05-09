import { fetchUserData } from '../api/fetchUsers';
import {decodeNLP} from '../api/decodeNLP'
async function processNLPQuery(query) {
  try {
    // First, decode the NLP query
    const response = await decodeNLP(query);

    // Extracting data from the response
    const skills = response.Skills.Responses;
    const budget = response.Budget.Responses;
    const experience = response.Experience.Responses;
    const miscellanous = response.Miscellanous.Responses;

    // Processing Skills
    const sortedSkills = skills.slice(0, 3).map(skill => skill.Text);
    const isSkillPresent = response.Skills.IsPresent;

    // Processing Budget
    const sortedBudget = budget[0].Text;
    const isBudgetPresent = response.Budget.IsPresent;

    // Processing Experience
    const sortedExperience = experience.length > 0 ? experience[0].Text : null;
    const isExperiencePresent = response.Experience.IsPresent;

    // Processing Miscellanous
    const sortedMiscellanous = miscellanous.map(item => item.Text);
    const isMiscellanousPresent = response.Miscellanous.IsPresent;

    return {
      Skills: sortedSkills,
      Budget: sortedBudget,
      Experience: sortedExperience,
      Miscellanous: sortedMiscellanous,
      IsSkillPresent: isSkillPresent,
      IsBudgetPresent: isBudgetPresent,
      IsExperiencePresent: isExperiencePresent,
      IsMiscellanousPresent: isMiscellanousPresent
    };
  } catch (error) {
    console.error('Error processing response:', error);
    return null;
  }
}

export async function extractUsersForQuery(query){
   var queryData =  processNLPQuery(query);

  var data =  await fetchUserData(queryData.Skills)
    console.log(data)
}
