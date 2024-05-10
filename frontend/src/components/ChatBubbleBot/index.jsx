import React from 'react';
import './index.css';

const ChatBubbleBot = ({message}) => {
  return (
    <div className={`chat-bubble`}>
       <div className={`chat-bubble`}>
      <span className="messageBot">{message}</span>
    </div>
    </div>
  );
};

export default ChatBubbleBot;