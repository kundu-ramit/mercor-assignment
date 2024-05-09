import React, { useEffect, useState } from 'react'
import './index.css'
import Intro from '../Intro';
import ChatWindow from '../ChatWindow';
import InputBox from '../InputBox';
import ChatBubble from '../ChatBubble';
import { extractUsersForQuery } from "./processor/processNLP.js";

function ChatPage() {
  const [chats,setChats] = useState([])

  useEffect(() => {
    setChats([<Intro/>])
  }, []);

  async function handleSend(text) {
    chats.push(<ChatBubble message={text}/>)
    setChats([...chats])
    const rankedUsers = await extractUsersForQuery(text)
    console.log(rankedUsers)
    for(var i=0;i<rankedUsers.length;i++)
      {
        chats.push(<ChatBubble message={JSON.stringify(rankedUsers[i])}/>)
        setChats([...chats])
      }
  }

  return (
    <div className='gradientBackground'>
    <ChatWindow chats = {chats}/>
    <InputBox handleSend = {handleSend}/>

    
  </div>
  )
}

export default ChatPage