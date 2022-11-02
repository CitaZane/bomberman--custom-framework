/* @jsx jsx */

import jsx from "../../framework/vDom/jsx";
import { defineWebSocket } from "../websocket";
import { store } from "../app";

function createWebSocketConn(e) {
  e.preventDefault();

  const inputElem = e.target.elements["name"];

  defineWebSocket(inputElem.value);

  store.dispatch("savePlayerName", inputElem.value)
  // console.log(store.state.currentPlayerName)
}


export function HomeView() {
  return {
    template: (
      <form onSubmit={createWebSocketConn}>
        <label for="name">Enter your username: </label>
        <input type="text" id="name"></input>
        <button>Start game</button>
      </form>
    ),
  };
}
