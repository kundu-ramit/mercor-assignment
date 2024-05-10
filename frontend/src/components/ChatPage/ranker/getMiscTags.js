// Function to generate miscellaneous tags

import eliteSchools from './data/elite_schools.json'
import mncList from './data/mnc_list.json'

export function getMiscTags(companiesWorkedAt, schools, tags, miscValues) {
    const miscValueList = miscValues.map(val => val.Text)
    const miscMap = {}
    for(var i in miscValueList)
        miscMap[miscValueList[i].Text] = miscValueList[i].Score
    var points=0;
    if(miscValueList.includes("TopUniversity"))
    {
        for(var i=0;i<schools.length;i++)
        {
            if(eliteSchools.includes(schools[i]) ){
                tags.push("TopUniversity")
                if(miscMap["TopUniversity"]>0.5)
                points+=2000
            else 
            points+=400
                break;
            }
        }
    }
    if(miscValueList.includes("MNC"))
        {
        for(i=0;i<companiesWorkedAt.length;i++)
            {
                if(mncList.includes(companiesWorkedAt[i])){
                    if(miscMap["MNC"]>0.5)
                        points+=2000
                    else 
                    points+=400
                    tags.push("MNC")
                    break;
                }
            }
        }

    if(miscValueList.includes("STARTUP"))
        {
        for(i=0;i<companiesWorkedAt.length;i++)
            {
                if(!mncList.includes(companiesWorkedAt[i])){
                    if(miscMap["STARTUP"]>0.5)
                        points+=2000
                    else 
                    points+=400
                    tags.push("STARTUP")
                    break;
                }
            }
        }

        return points
    
}        