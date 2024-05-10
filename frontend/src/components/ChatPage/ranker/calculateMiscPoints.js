import { getMiscTags } from "./getMiscTags";

export function calculateMiscellaneousPoints(user, requirements, tags) {
    let points =0;
    if (requirements.IsMiscellanousPresent) {
        points += getMiscTags(user.companiesWorkedAt, user.schools, tags, requirements.Miscellanous);
    }
    return points;
}