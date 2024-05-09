import { getMiscTags } from "./getMiscTags";

const MISC_POINTS = {
    0: 300,
    1: 100,
    2: 50
};

export function calculateMiscellaneousPoints(user, requirements, tags) {
    let points =0;
    if (requirements.IsMiscellanousPresent) {
        const miscValues = requirements.Miscellanous;
        points += getMiscTags(user.companiesWorkedAt, user.schools, tags, miscValues);
    }
    return points;
}