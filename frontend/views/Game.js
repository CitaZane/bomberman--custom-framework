/* @jsx jsx */
import jsx from "../../framework/vDom/jsx";
import { store } from "../app";

// async function getPlayerCount() {
//   const response = await fetch("http://localhost:8080/queuePlayerCount");
//   const data = await response.text();
//   store.commit("updateUserQueueCount", data);
//   console.log("data", data);
// }

export function GameView() {
  let userQueueCount = store.state.userQueueCount;

  return {
    template: (
      <div>
        <h1>Game room</h1>
        <h3>Users joined: {String(userQueueCount)}</h3>
      </div>
    ),
  };
}
