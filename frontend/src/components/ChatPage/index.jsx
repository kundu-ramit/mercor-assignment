import React, { useEffect, useState } from 'react'
import './index.css'
import Intro from '../Intro';
import ChatWindow from '../ChatWindow';
import InputBox from '../InputBox';
import ChatBubble from '../ChatBubble';
import { extractUsersForQuery } from "./processor/processNLP.js";
import { generateUserCard } from '../UserCard/index.jsx';
import { processNLPQuery } from './processor/processNLP.js';
import BotSuggestions from './botSuggesions/index.jsx';

function ChatPage() {
  const [chats,setChats] = useState([])
  const [inputBoxValue,setInputBoxValue] = useState("")

  useEffect(() => {
    setChats([<Intro/>])
  }, []);

  async function handleSend(text) {
    setInputBoxValue("");
    chats.push(<ChatBubble message={text}/>)
    setChats([...chats])
    var queryData =  await processNLPQuery(text);
    if(!queryData.IsSkillPresent){
      chats.push(<BotSuggestions data={queryData} prompt ={text} setInputBoxValue={setInputBoxValue}/>)
      setChats([...chats])
      return;
    }

    const rankedUsers = await extractUsersForQuery(queryData)
    for(var i=0;i<rankedUsers.length;i++)
      {
        chats.push(generateUserCard(rankedUsers[i]))
        setChats([...chats])
      }
      chats.push(<BotSuggestions data={queryData} prompt ={text} setInputBoxValue={setInputBoxValue}/>)
      setChats([...chats])
  }

  return (
    <div className='gradientBackground'>
    <ChatWindow chats = {chats}/>
    <InputBox handleSend = {handleSend} value={inputBoxValue} setValue={setInputBoxValue}/>

    
  </div>
  )
}

export default ChatPage