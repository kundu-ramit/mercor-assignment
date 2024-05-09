const FULL_TIME_POINTS = 500;
const WITHIN_BUDGET_POINTS = 500;

export function calculateBudgetPoints(user, requirements, tags) {
    let points=0;
    if (requirements.IsBudgetPresent) {
        const budgetRegex = /([0-9]+)\s*\+?\)/;
        const budgetMatch = requirements.Budget.match(budgetRegex);
        if (budgetMatch) {
            const budget = parseInt(budgetMatch[1]);
            const fullTimeBudget = user.fullTimeSalaryCurrency === "USD" ? parseFloat(user.fullTimeSalary) : 0;
            const partTimeBudget = user.partTimeSalaryCurrency === "USD" ? parseFloat(user.partTimeSalary) : 0;

            if (requirements.Budget.includes("FT")) {
                if (user.fullTime) {
                    points += FULL_TIME_POINTS; // FT availability
                    if (fullTimeBudget <= budget) {
                        points += WITHIN_BUDGET_POINTS; // Within budget for FT
                        tags.push(`Under $${budget} (Full Time)`);
                    }
                }
            }

            if (requirements.Budget.includes("PT")) {
                if (user.partTime) {
                    points += FULL_TIME_POINTS; // PT availability
                    if (partTimeBudget <= budget) {
                        points += WITHIN_BUDGET_POINTS; // Within budget for PT
                        tags.push(`Under $${budget} (Part Time)`);
                    }
                }
            }
        }
    }
    return points;
}