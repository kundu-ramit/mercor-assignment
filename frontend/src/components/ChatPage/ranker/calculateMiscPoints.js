import { getMiscTags } from "./getMiscTags";

export function calculateMiscellaneousPoints(user, requirements, tags) {
    let points =0;
    if (requirements.IsMiscellanousPresent) {
        const miscValues = requirements.Miscellanous.map(val => val.Text);
        points += getMiscTags(user.companiesWorkedAt, user.schools, tags, miscValues);
    }
    return points;
}