import React from 'react';
import './Dashboard.css'
import { Contacts } from './contacts/Contacts'
import { Chat } from './chat/Chat'
import { Menu } from './menu/Menu'

export function Dashboard(props) {
  return (
    <div className="dashboard">
      <div className="app-container bg">
        <div className="side-bar">
          {props.menuBarStatus
            ? <Menu
              usernameSetter={props.usernameCallback}
              username={props.chatState.username}
              connectedUsersp={props.chatState.connectedUsers}
            />
            : <Contacts
              username={props.chatState.username}
              selectedChat={props.chatState.selectedChat}
            />}
        </div>
        <div className="chat">
          <Chat username={props.chatState.username} />
        </div>
      </div>
    </div>
  );
}