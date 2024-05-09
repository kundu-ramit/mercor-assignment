import React from 'react';
import './index.css';

const ChatBubbleBot = ({result}) => {
  return (
    <div className={`chat-bubble`}>
      {result}
    </div>
  );
};

export default ChatBubbleBot;