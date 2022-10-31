/* @jsx jsx */

import jsx from "../../framework/vDom/jsx";
import { defineWebSocket } from "../websocket";
import { ws } from "../websocket";
function createWebSocketConn(e) {
  e.preventDefault();

  const inputElem = e.target.elements["name"];

  defineWebSocket(inputElem.value);
  console.log("WS", ws)
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
