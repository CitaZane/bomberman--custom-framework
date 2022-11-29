/* @jsx jsx */

import jsx from "../framework/vDom/jsx";
import { store } from "../app";
import { ws } from "../websocket";

function focus() {
  store.commit("updateFocusOnChat", true);
}
function blur() {
  store.commit("updateFocusOnChat", false);
}
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
  formInput.blur();
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
          {messages.map((message) => {
            let index = lobbyPlayersNames.findIndex(
              (name) => name === message.creator
            );

            return (
              <li>
                <p class={`chat-username player-name monster-${index}__color`}>
                  {message.creator}
                </p>
                <p>{message.body}</p>
              </li>
            );
          })}
        </ul>

        <form id="send-message" onSubmit={sendMessage}>
          <input
            type="text"
            name="message"
            placeholder="Send message"
            onFocus={focus}
            onBlur={blur}
          ></input>
          <button>&gt;</button>
        </form>
      </div>
    ),
  };
}
