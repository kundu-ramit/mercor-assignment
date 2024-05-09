import React from 'react';
import { Avatar } from 'antd';
import './index.css';

const ChatBubble = ({ message}) => {
  return (
    <div className={`chat-bubble`}>
      <span className="message">{message}</span>
    </div>
  );
};

export default ChatBubble;