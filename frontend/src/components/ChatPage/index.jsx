import React, { useEffect, useState } from 'react'
import { Card } from "antd";
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
    chats.push(<ChatBubble message={text} sender={"user"}/>)
    setChats([...chats])
    await extractUsersForQuery(text)
  }

  return (
    <div className='gradientBackground'>
    <ChatWindow chats = {chats}/>
    <InputBox handleSend = {handleSend}/>

    
  </div>
  )
}

export default ChatPage