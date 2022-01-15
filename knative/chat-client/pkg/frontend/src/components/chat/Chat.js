import React, { useState } from 'react';
import sendBtn from '../../static/img/send-btn.png'
import './Chat.css'

export function Chat(props) {
  const [chatContent, setInputMessage] = useState({
    messageInput: '',
    messages: props.messages
  })

  const chatInputHandler = (e) => {
    setInputMessage({ messageInput: e.target.value })
  }

  const handleChatInputKeyDown = (e) => {
    if (e.key === 'Enter') {
      sendButton()
    }
  }

  const sendButton = () => {
    if (chatContent.messageInput !== '') {
      alert(chatContent.messageInput)
      setInputMessage({ messageInput: '' })
    }
  }

  return (
    <div className="chat-column">
      <div className="messages">
        <div className="messages-bubble">
          <div className="bubbles-bg">
            <div className="contact-message">
              This is a mesage...
            </div>
            <div className="user-message">
              This is my mesage...
            </div>
          </div>
        </div>
      </div>
      <div className="chat-section">
        <div className="chat-bubble">
          <div className="bubbles-bg">
            <img src={sendBtn} className="send-btn" alt="send-btn.png" onClick={sendButton} />
            <input
              type="text"
              name="message"
              className="chat-input"
              placeholder="Type a message"
              value={chatContent.messageInput}
              onChange={chatInputHandler}
              onKeyDown={handleChatInputKeyDown}
            />
          </div>
        </div>
      </div>
    </div>
  );
}