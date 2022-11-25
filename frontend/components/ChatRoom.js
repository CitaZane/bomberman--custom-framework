/* @jsx jsx */

import jsx from "../framework/vDom/jsx";
import { store } from "../app";
import { ws } from "../websocket";

function sendMessage(e) {
  e.preventDefault();
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
  const lobbyPlayersNames = store.state.lobbyPlayersNames;
  return {
    template: (
      <div id="chatroom">
        <header>
          <h3 id="chat-header">Chatroom</h3>
        </header>

        <ul id="chat">
          {/* <li>
            <p class={`chat-username player-name monster-1__color`}>Thomas</p>
            <p>Boys</p>
          </li>
          <li>
            <p class={`chat-username player-name monster-1__color`}>Thomas</p>
            <p>Boys</p>
          </li>
          <li>
            <p class={`chat-username player-name monster-1__color`}>Thomas</p>
            <p>Boys</p>
          </li>
          <li>
            <p class={`chat-username player-name monster-1__color`}>Thomas</p>
            <p>Boys</p>
          </li>
          <li>
            <p class={`chat-username player-name monster-1__color`}>Thomas</p>
            <p>Boys</p>
          </li>
          <li>
            <p class={`chat-username player-name monster-1__color`}>Thomas</p>
            <p>Boys</p>
          </li>
          <li>
            <p class={`chat-username player-name monster-1__color`}>Thomas</p>
            <p>Boys</p>
          </li>
          <li>
            <p class={`chat-username player-name monster-1__color`}>Thomas</p>
            <p>Boys</p>
          </li>
          <li>
            <p class={`chat-username player-name monster-1__color`}>Thomas</p>
            <p>Boys</p>
          </li>
          <li>
            <p class={`chat-username player-name monster-1__color`}>Thomas</p>
            <p>Boys</p>
          </li> */}

          {messages.map((message) => {
            let index = lobbyPlayersNames.findIndex((name) => name === message.creator);

            return (
              <li>
                <p class={`chat-username player-name monster-${index}__color`}>{message.creator}</p>
                <p>{message.body}</p>
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
