// Mock function to simulate GitHub activity
async function mockGithubPoints(ocrGithubUsername, tags) {
    // Simulate some activity
    const commitsLastYear = Math.floor(Math.random() * 100); // Random number of commits

    // If the number of commits made in the last year is more than 20, return 100 points, otherwise return 0 points
    if (commitsLastYear > 20) {
        tags.push('ACTIVE CONTRIBUTOR')
        return 100;
    } else {
        return 0;
    }
}

// Function to decide whether to call GitHub API or mock function
export async function calculateGithubPoints(ocrGithubUsername, tags) {
    try {
        // Randomly decide whether to call the actual GitHub API or the mock function
        const random = Math.random();
        if (random < 0.1) {
            // Call the actual GitHub API
            return await actualGithubPoints(ocrGithubUsername, tags);
        } else {
            // Call the mock function
            return await mockGithubPoints(ocrGithubUsername, tags);
        }
    } catch (error) {
        // Handle errors
        console.error("Error fetching GitHub activity:", error);
        return 0; // Return 0 points in case of an error
    }
}

// Function to fetch actual GitHub activity
async function actualGithubPoints(ocrGithubUsername, tags) {
    const response = await fetch(`https://api.github.com/users/${ocrGithubUsername}/events/public`);
    const data = await response.json();
    const pushEvents = data.filter(event => event.type === "PushEvent");
    const lastYearTimestamp = new Date(Date.now() - 365 * 24 * 60 * 60 * 1000).getTime();
    const commitsLastYear = pushEvents.reduce((total, event) => {
        const eventTimestamp = new Date(event.created_at).getTime();
        return eventTimestamp > lastYearTimestamp ? total + event.payload.size : total;
    }, 0);
    if (commitsLastYear > 20) {
        tags.push('ACTIVE CONTRIBUTOR');
        return 100;
    } else {
        return 0;
    }
}