import React from 'react';
import './Contacts.css'
import userImg from '../../user.png'

export function Contacts() {
  return (
    <div className="contacts-column">
      <div className="contact-card">
        <div className="contact-img">
          <img className="usr-pic" src={userImg} alt="user-img.png" />
        </div>
        <div className="contact-info">
          <div className="contact-name">
            <h3 className="user-name"><b>Contact Name</b></h3>
          </div>
          <div className="contact-last-msg">
            Last msg...
          </div>
        </div>
      </div>
    </div>
  );
}