import axios  from "axios";

export async function fetchUserData(skills) {
  var skillIds = skills.map(skill => skill.Text)
    try {
      const response = await axios({
        method: 'POST',
        url: 'http://localhost:8002/query/order',
        headers: {
          'Content-Type': 'application/json'
        },
        data: {
          skills: skillIds
        }
      });
      return response.data;
    } catch (error) {
      console.error('Error fetching users:', error);
      return null;
    }
  }