import createRouter from "../framework/router";
import createStore from "../framework/store";
import routes from "./router";
import storeObj from "./store/index";
import { ws } from "./websocket";

const store = createStore(storeObj);
const router = createRouter(routes);
let gameStarted = false;

// add keyup and KeyDown listeners
export function setupGame() {
  document.addEventListener("keyup", (e) => {
    store.dispatch("registerKeyUp", e.code);
  });
  document.addEventListener("keydown", (e) => {
    if (!gameStarted) {
      gameStarted = true;
      animate();
    }
    store.dispatch("registerKeyDown", e.code);
  });
}

// main game loop
let gameFrame = 0;
function animate() {
  let inputs = store.state.inputs;
  let currentPlayerName = store.state.currentPlayerName;

  // let index = store.state.currentPlayerIndex;

  if (inputs["ArrowLeft"]) {
    ws.send(
      JSON.stringify({
        type: "PLAYER_MOVE",
        creator: currentPlayerName,
        body: "LEFT",
      })
    );
  } else if (inputs["ArrowDown"]) {
    ws.send(
      JSON.stringify({
        type: "PLAYER_MOVE",
        creator: currentPlayerName,
        body: "DOWN",
      })
    );
  } else if (inputs["ArrowUp"]) {
    ws.send(
      JSON.stringify({
        type: "PLAYER_MOVE",
        creator: currentPlayerName,
        body: "UP",
      })
    );
  } else if (inputs["ArrowRight"]) {
    ws.send(
      JSON.stringify({
        type: "PLAYER_MOVE",
        creator: currentPlayerName,
        body: "RIGHT",
      })
    );
  } else if (inputs["Space"]) {
    console.log("Bomb drop");
    // ws.send(JSON.stringify({type: 'PLAYER_DROP_BOMB',creator: String(index), body: '?coordinates?'}))
  }

  gameFrame++;
  requestAnimationFrame(animate);
}

export { store, router, gameFrame };
