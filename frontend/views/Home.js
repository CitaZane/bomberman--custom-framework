/* @jsx jsx */
import jsx from "../../framework/vDom/jsx"
import { store } from "../app"

import { defineWebSocket } from "../websocket";


function createWebSocketConn(e) {
  e.preventDefault();

  const inputElem = e.target.elements["name"];

  defineWebSocket(inputElem.value);

  store.dispatch("savePlayerName", inputElem.value)
}


export function HomeView() {
  return {
    template: (
      <form onSubmit={createWebSocketConn}>
        <label for="name">Enter your username: </label>
        <input type="text" id="name"></input>
        <button>Enter queue</button>
      </form>
    ),
  };
}
