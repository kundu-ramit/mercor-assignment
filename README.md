READ THIS NOTION DOC INSTEAD : https://pitch-lion-eec.notion.site/Mercor-Assignment-Log-and-Notes-fd253486ed8344de97dac5f470d138b0


# Mercor Assignment Log and Notes

**Hi,**

**In this doc I will explain how I have done the mercor assignment**

[https://mercor.notion.site/Mercor-Software-Engineer-Trial-a26eeb14f36340239c8dc680dc9288ba](https://www.notion.so/Mercor-Software-Engineer-Trial-a26eeb14f36340239c8dc680dc9288ba?pvs=21)

There will be 6 portions :

---

### Github Link and how to run

[https://github.com/kundu-ramit/mercor-assignment](https://github.com/kundu-ramit/mercor-assignment)

How to run 

Ask me for .env file

Paste it in backend terminal

Hit go run main.go server

Frontend do npm install and npm start

### **Software and language I have used**

Backend in Golang 

Frontend in React

SingleStore for vector database

OpenAi for Embedding Model

Reason for use of these technologies :

I wanted a strongly typed but simple backend and I wanted it to be in a language that supports multithreading . Golang is the best choice since goroutines allow us to execute simple ops like querying multiple vector tables in parallel.

For frontend I wanted to move fast and javascript was the best choice. I wanted a non typed language due to time constraints.

For vector databases I have gone with singlestore instead of pinecone and others. Singlestore is a vector database that lets us store coordinates in the form of a relational table and query them. For my usecase it was the best solution . As we move through this doc it will be clearer why.

OpenAi was the obvious choice . I did not have the time to learn and execute a custom embedding. So went with it. (In production will do the opposite to save costs and increase speed)

### **Vector Database**

I have used single store as a vector database.

In single store I have created 4 tables 

![Screenshot 2024-05-11 050302.png](Mercor%20Assignment%20Log%20and%20Notes%20fd253486ed8344de97dac5f470d138b0/Screenshot_2024-05-11_050302.png)

Now I am seeding through these scripts

![Screenshot 2024-05-11 050426.png](Mercor%20Assignment%20Log%20and%20Notes%20fd253486ed8344de97dac5f470d138b0/Screenshot_2024-05-11_050426.png)

 Why seperate tables ?

Simple while testing the data whenever I encounter an issue I can figure out exactly which

table to fix. Also data is partitioned well and I can train my model  more easily.

Cons :

Querying multiple table . But latency is in ms nothing compared to OPENAI and i can do it in parallel with goroutine.

**Since our usecase is limited and we can always put data together like miscellanous table it seems**

**the best solution.**

### **API**

Finally I have exposed 2 apis 

```powershell
curl --location 'localhost:8002/query/nlp' \
--header 'Content-Type: application/json' \
--data '{
    "query":"I want a python and java developer under 3000 dollars. Want full time. Experience 5 to 7 years"
}'
```

```powershell
curl --location 'localhost:8002/query/order' \
--header 'Content-Type: application/json' \
--data '{
    "skills" : ["0bfbda41-7746-11ee-accf-42010a400014","1da65325-7746-11ee-accf-42010a400014","5f9e6469-7517-11ee-accf-42010a400014","5f9e7cf5-7517-11ee-accf-42010a400014"]
}'
```

I will explain in detail in the demo

### Frontend

The frontend is coded with React. Frontend does more heavy lifting in this system. The ranking logic is entirely implemented there.

### Why ranking logic in frontend

This is not something I will be doing in production. Backend should return a ranked list so frontend could be light and load faster. However my backend is in go which is strongly typed 

and due to the volume of code I decided to keep it in the frontend to save on time. Its easier to write large chunks of temporary code in javascript. 

### A 40 minute explanation and testing video

[https://player.vimeo.com/video/945060791?badge=0&amp;autopause=0&amp;player_id=0&amp;app_id=58479](https://player.vimeo.com/video/945060791?badge=0&amp;autopause=0&amp;player_id=0&amp;app_id=58479)

### A basic system design diagram

![FRONTEND.svg](Mercor%20Assignment%20Log%20and%20Notes%20fd253486ed8344de97dac5f470d138b0/FRONTEND.svg)

### **Possible different approach**

A possible better and simpler approach might have been to seed all user data from MySql to vector db. Then use the query to figure out the closest users.
Reasons why I did not think of it.
I have not used vector databases before. So i did not know how it would perform or how optimum it would be once i did a vector database heavy implementation.
I preferred to trust vector Db less and keep more logic on the code since I did not have time for a spike. 
I ran simple fetch queries using Ids in Mysql . In a standard system these queries are supposed to execute in ms.
So although I noted that the system was giving slow sql alerts I ignored it . I failed to understand I was needed to optimize it. 
I also think the above approach might pose unique challenges like figuring out how to maintain consistency during reads and writes to the mysql database. Although nothing too difficult to figure out but did not want to think in that direction. I wanted to build a solution exactly in my zone so I could answer questions on how to scale it to something actually used in production.

Since data is small 100k user data could easily fit in 30mb I believe **redis** would also be an amazing option if we have to move from mysql and want high speed .

### Hardcoded tags

I have used hardcoded words for miscellanous tags like “MNC” and “STARTUP” and “TopEducation” . 

We can enter something like : I want a harvard grad who knows python -

And it will match with TOPEDUCATION and PYTHON. According to me the intent was more important to capture so if someone wants a harvard grad he might be ok with someone from MIT also . So I believe the hardcoded tags capture that intent . And due to limited scope of this chatbot(to be used to find freelancers) this seemed a good approach to me.

### How accurate is my current solution and how did I test and improve it.

Once the solution was built I spent a day testing it . I would enter random text and try to see the output . If the output was not satisfactory I would tweak the parameters in backend and frontend to make it work. That way I have tweaked the model until I was able to run it for 30 minutes without breaking it.

Under a more restrictive scope I believe it currently gives accurate results in more than 90% circumstances.