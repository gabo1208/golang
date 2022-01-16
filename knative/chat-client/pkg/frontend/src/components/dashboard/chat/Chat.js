import React, { useState } from 'react';
import sendBtn from '../../../static/img/send-btn.png'
import './Chat.css'

export function Chat(props) {
  const [chatContent, setInputMessage] = useState({
    messageInput: '',
    messages: props.selectedChat?.messages || []
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

  const chatDefaultMsg = () => {
    return props.username ? "Start chatting with !"
      : "Please choose an username to start chatting!"
  }

  const renderMessages = (messagesList) => {
    return messagesList.map(message => {
      return <div className={(message.mine ? "user" : "contact") + "-message"}>
        {message.content}
      </div>
    })
  }

  return (
    <div className="chat-column">
      <div className="messages">
        <div className="messages-bubble">
          <div className="bubbles-bg">
            {chatContent.messages?.length
              ? renderMessages(chatContent.messages)
              : <div className="empty-chat">
                <h3>{chatDefaultMsg()}</h3>
              </div>}
          </div>
        </div>
      </div>
      <div className={`chat-section ${!props.username && "blocked"}`}>
        <div className="chat-bubble">
          <div className="bubbles-bg">
            <img
              src={sendBtn}
              className="send-btn"
              alt="send-btn.png"
              onClick={sendButton}
            />
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