import React from 'react';
import './Contacts.css'
import userImg from '../../../static/img/user.png'

export function Contacts() {
  return (
    <div className="contacts-column">
      <div className="contact-card selected cursor-pointer">
        <div className="contact-img">
          <img className="user-img" src={userImg} alt="user-img.png" />
        </div>
        <div className="contact-info">
          <div className="contact-name">
            <h3 className="user-name"><b>Contact Name</b></h3>
          </div>
          <div className="contact-last-seen">
            Last seen 12/12/2022
          </div>
          <div className="contact-last-msg">
            Last msg...
          </div>
        </div>
      </div>
      <div className="contact-card cursor-pointer">
        <div className="contact-img">
          <img className="user-img" src={userImg} alt="user-img.png" />
        </div>
        <div className="contact-info">
          <div className="contact-name">
            <h3 className="user-name"><b>Contact Name</b></h3>
          </div>
          <div className="contact-last-seen">
            Last seen 12/12/2022
          </div>
          <div className="contact-last-msg">
            Last msg...
          </div>
        </div>
      </div>
    </div>
  );
}