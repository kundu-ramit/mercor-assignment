// Function to generate miscellaneous tags

import eliteSchools from './data/elite_schools.json'
import mncList from './data/mnc_list.json'

export function getMiscTags(companiesWorkedAt, schools, tags, miscValues) {
    var points=0;
    if(miscValues.includes("TopUniversity"))
    {
    for(var i=0;i<schools.length;i++)
    {
        if(eliteSchools.includes(schools[i]) ){
            tags.push("TopUniversity")
            points+=500
            break;
        }
    }
}
if(miscValues.includes("MNC"))
    {
    for(var i=0;i<companiesWorkedAt.length;i++)
        {
            if(mncList.includes(companiesWorkedAt[i])){
                points+=500
                tags.push("MNC")
                break;
            }
        }
    }

    if(miscValues.includes("STARTUP"))
        {
        for(var i=0;i<companiesWorkedAt.length;i++)
            {
                if(!mncList.includes(companiesWorkedAt[i])){
                    points+=500
                    tags.push("STARTUP")
                    break;
                }
            }
        }

        return points
    
}        