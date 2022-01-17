import React, { useState, useCallback } from 'react';
import './Dashboard.css'
import { Contacts } from './contacts/Contacts'
import { Chat } from './chat/Chat'
import { Menu } from './menu/Menu'

export function Dashboard(props) {
  const [chatState, setChatState] = useState({
    selectedChatIndex: 0,
    connectedUsers: []
    // {username: string, messages: {content: string, mine: bool}}
  })

  const updateChatMessagesCallback = useCallback(message => {
    let connectedUsersAux = chatState.connectedUsers
    connectedUsersAux[chatState.selectedChatIndex].message.append({ mine: true, message })
    setChatState({
      selectedChatIndex: chatState.selectedChatIndex,
      connectedUsers: connectedUsersAux
    })
  }, [setChatState, chatState])

  const updateSelectedChatIndexCallback = useCallback(selectedChatIndex => {
    setChatState({ ...chatState, selectedChatIndex })
  }, [setChatState, chatState])

  return (
    <div className="dashboard">
      <div className="app-container bg">
        <div className="side-bar">
          {props.menuBarStatus
            ? <Menu
              usernameSetter={props.usernameCallback}
              username={props.username}
            />
            : <Contacts
              username={props.username}
              connectedUsers={chatState.connectedUsers}
              selectedChatIndex={chatState.selectedChatIndex}
              selectedChat={chatState.connectedUsers[chatState.selectedChatIndex]}
              updateSelectedChatIndexCallback={updateSelectedChatIndexCallback}
            />}
        </div>
        <div className="chat">
          <Chat
            username={props.username}
            selectedChat={chatState.connectedUsers[chatState.selectedChat]}
            connectedUsersNumber={chatState.connectedUsers.length}
            updateChatMessagesCallback={updateChatMessagesCallback}
          />
        </div>
      </div>
    </div>
  );
}