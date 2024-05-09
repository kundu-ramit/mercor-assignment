import { Card } from 'antd'
import React from 'react'

function ChatWindow({chats}) {
  return (
    <Card style={{ height : '1020px' , overflow:'scroll', scrollbarWidth: 'none', border: 'none', boxShadow: 'none', background: 'none' , width:"50%", margin:"auto"}}>
      {chats}
    </Card>
  )
}

export default ChatWindow