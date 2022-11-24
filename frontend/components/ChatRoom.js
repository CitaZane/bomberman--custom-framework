/* @jsx jsx */

import jsx from "../framework/vDom/jsx";
import { store } from "../app";
import { myMonsterColorClass, ws } from "../websocket";

function sendMessage(e) {
  e.preventDefault();
  console.log("state", store.state);
  const formInput = e.target.elements.message;

  const msg = {
    type: "TEXT_MESSAGE",
    creator: store.state.currentPlayerName,
    body: formInput.value,
  };

  ws.send(JSON.stringify(msg));

  formInput.value = "";
}

export function ChatRoom() {
  const messages = store.state.messages;  
  return {    
    template: (
      <div id="chatroom">
        <header>
          <h2 id="chat-header">Chatroom</h2>
        </header>

        <ul id="chat">
          {messages.map((message) => {
            let index = null;
            store.state.lobbyPlayersNames.forEach((lobby_player_name, i) => {              
              if (message.creator == lobby_player_name){
                index = i
              }              
            });
            return (
              <li>
                {/* <p class={`chat-username`}>{message.creator}</p> */}
                <p class={`player-name monster-${index}__color`}>{message.creator}</p>
                {/* <p>{message.body}</p> */}
                <p class={`player-name monster-${index}__color`}>{message.body}</p>
              </li>
            );
          })}
        </ul>

        <form id="send-message" onSubmit={sendMessage}>
          <input type="text" name="message" placeholder="Send message"></input>
          <button>&gt;</button>
        </form>
      </div>
    ),
  };
}
