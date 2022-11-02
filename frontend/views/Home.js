/* @jsx jsx */
import jsx from "../../framework/vDom/jsx"
import { store } from "../app"

import { defineWebSocket } from "../websocket";
import { MonsterSprite } from "../components/MonsterSprite"
import { BombSprite } from "../components/BombSprite";
// import { FireSprite } from "../components/FireSprite";

// export function HomeView() {

// let bombDrop = store.state.bomb.drop;
// return {
//   template: (
//     <div id="home">
//       <h1>Hello monster</h1>
//       <MonsterSprite />
//       {/* {bombDrop && <BombSprite />} */}
//       {/* <FireSprite/> */}
//     </div>
//   )
// }
// }



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
        <button>Enter queue</button>
      </form>
    ),
  };
}
