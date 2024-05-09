import axios  from "axios";

export async function fetchUserData(skills) {
    try {
      const response = await axios({
        method: 'GET',
        url: 'http://localhost:8002/query/order',
        headers: {
          'Content-Type': 'application/json'
        },
        data: {
          skills: skills
        }
      });
      return response.data;
    } catch (error) {
      console.error('Error fetching users:', error);
      return null;
    }
  }