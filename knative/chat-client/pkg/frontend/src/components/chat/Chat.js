import React from 'react';
import './Chat.css'

export function Chat() {
  return (
    <div className="chat-column">
      <div className="messages">
        <div className="messages-bubble">
          <div className="bubbles-bg"></div>
        </div>
      </div>
      <div className="chat-section">
        <div className="chat-bubble">
          <div className="bubbles-bg"></div>
        </div>
      </div>
    </div>
  );
}