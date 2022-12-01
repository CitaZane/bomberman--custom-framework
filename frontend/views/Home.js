/* @jsx jsx */
import jsx from "../framework/vDom/jsx";
import { store } from "../app";

import { defineWebSocket } from "../websocket";
import { Overlay } from "../components/Overlay";

function createWebSocketConn(e) {
  e.preventDefault();

  const inputElem = e.target.elements["name"];
  if (inputElem.value === "") {
    console.log("Empty name");
    return;
  }
  defineWebSocket(inputElem.value);

  store.dispatch("savePlayerName", inputElem.value);
}

export function HomeView() {
  return {
    template: (
      <div id="home-layout">
        <h1>Bomberman</h1>
        <form onSubmit={createWebSocketConn} id="username-form">
          <label for="name">Enter your name</label>
          <input type="text" id="name" required maxlength="9"></input>
          <button class="btn">Join lobby</button>
        </form>
      </div>
    ),
  };
}
