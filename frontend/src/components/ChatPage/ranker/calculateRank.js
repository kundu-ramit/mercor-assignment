import { calculateBudgetPoints } from "./calculateBudgetPoints";
import { calculateMiscellaneousPoints } from "./calculateMiscPoints";
import { calculateSkillPoints } from "./calculateSkillPoints";
import { calculateGithubPoints } from "./calculateGithubPoints";

async function calculateRank(user, requirements) {
    let points = 0;
    let tags = [];
    points += calculateSkillPoints(user,requirements, tags);
    points += calculateBudgetPoints(user, requirements, tags);
    points += calculateMiscellaneousPoints(user, requirements, tags);
    points += await calculateGithubPoints(user.ocrGithubUsername, tags);

    return { "userId": user.userId,"rank": points,"tags" : tags };
}

export async function getRankedList(users,requirements) {
    let rankArray = []

    for(var i=0;i<users.length;i++)
    rankArray.push(await calculateRank(users[i],requirements))
    rankArray.sort((a, b) => {
        return b.rank - a.rank;
    });

    const userHash = {}
    for(i=0;i<users.length;i++)
    {
       userHash[users[i].userId] = users[i];
    }

    if(rankArray.length>3)
    rankArray = rankArray.slice(0, 3);

    var result = []

    for(i=0;i<rankArray.length;i++)
    {
        result.push({...userHash[rankArray[i].userId],"tags": rankArray[i].tags})
    }
    return result
}
