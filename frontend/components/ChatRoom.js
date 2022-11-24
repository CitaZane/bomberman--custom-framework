/* @jsx jsx */

import jsx from "../framework/vDom/jsx";
import { store } from "../app";
import { myMonsterColorClass, ws } from "../websocket";

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
  return {
    template: (
      <div id="chatroom">
        <header>
          <h2 id="chat-header">Chatroom</h2>
        </header>

        <ul id="chat">
          {messages.map((message) => {
            return (
              <li>
                <p class={`chat-username`}>{message.creator}</p>
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
