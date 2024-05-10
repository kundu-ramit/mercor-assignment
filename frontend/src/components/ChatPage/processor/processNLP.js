import { fetchUserData } from '../api/fetchUsers';
import {decodeNLP} from '../api/decodeNLP'
import {getRankedList} from '../ranker/calculateRank'
export async function processNLPQuery(query) {
  try {
    // First, decode the NLP query
    const response = await decodeNLP(query);

    // Extracting data from the response
    const skills = response.Skills.Responses;
    const budget = response.Budget.Responses;
    const experience = response.Experience.Responses;
    const miscellanous = response.Miscellanous.Responses;

    // Processing Skills
    const sortedSkills = skills.slice(0, 3);
    const isSkillPresent = response.Skills.IsPresent;

    // Processing Budget
    const sortedBudget = budget[0];
    const isBudgetPresent = response.Budget.IsPresent;

    // Processing Experience
    const sortedExperience = experience.length > 0 ? experience[0] : null;
    const isExperiencePresent = response.Experience.IsPresent;

    // Processing Miscellanous
    const isMiscellanousPresent = response.Miscellanous.IsPresent;
    
    return {
      Skills: sortedSkills,
      Budget: sortedBudget,
      Experience: sortedExperience,
      Miscellanous: miscellanous,
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

export async function extractUsersForQuery(queryData){
  var data =  await fetchUserData(queryData.Skills)
  var rankedUsers = await getRankedList(data, queryData)
  return rankedUsers;
}
