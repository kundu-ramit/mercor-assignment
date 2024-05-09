export async function calculateGithubPoints(ocrGithubUsername, tags) {
    try {
        // Fetch the user's GitHub activity
        const response = await fetch(`https://api.github.com/users/${ocrGithubUsername}/events/public`);
        const data = await response.json();

        // Filter the events to include only PushEvent (commits)
        const pushEvents = data.filter(event => event.type === "PushEvent");

        // Get the timestamps of the last year
        const lastYearTimestamp = new Date(Date.now() - 365 * 24 * 60 * 60 * 1000).getTime();

        // Count the number of commits made in the last year
        const commitsLastYear = pushEvents.reduce((total, event) => {
            const eventTimestamp = new Date(event.created_at).getTime();
            return eventTimestamp > lastYearTimestamp ? total + event.payload.size : total;
        }, 0);

        // If the number of commits made in the last year is more than 20, return 100 points, otherwise return 0 points
        if (commitsLastYear > 20) {
            tags.push('ACTIVE CONTRIBUTOR')
            return 100;
        } else {
            return 0;
        }
    } catch (error) {
        // Handle errors
        console.error("Error fetching GitHub activity:", error);
        return 0; // Return 0 points in case of an error
    }
}