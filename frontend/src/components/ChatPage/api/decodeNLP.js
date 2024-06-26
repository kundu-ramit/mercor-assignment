import axios  from "axios";

export async function decodeNLP(query) {
  try {
    const response = await axios({
      method: 'POST',
      url: 'http://localhost:8002/query/nlp',
      headers: {
        'Content-Type': 'application/json'
      },
      data: {
        'query': query
      }
    });
    return response.data;
  } catch (error) {
    console.error('Error sending curl request:', error);
    return null;
  }
}