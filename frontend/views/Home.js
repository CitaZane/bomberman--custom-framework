/* @jsx jsx */
import jsx from "../framework/vDom/jsx";
import { store } from "../app";

import { defineWebSocket } from "../websocket";

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
        <div class="left-sidebar">
          <div class="players">
            <div>
              <p class="player-name">Player 1</p>
              <div class="player-monster" id="monster-1"></div>
              <div class="lives">
                <img class="life" src="../assets/heart.png"></img>
                <img class="life"></img>
                <img class="life"></img>
              </div>
              <div class="player-power-ups">
                <div>
                  <div class="powerup increase-speed"></div>
                  <span id="increase-speed-count">1</span>
                </div>
                <div>
                  <div class="powerup increase-bombs"></div>
                  <span id="increase-bombs-count">2</span>
                </div>
                <div>
                  <div class="powerup increase-flame"></div>
                  <span id="increase-flame-count">2</span>
                </div>
              </div>
            </div>
            {/* <button id="quit" class="btn">Quit</button> */}
          </div>

        </div>

        <h1>Bomberman</h1>
        <form onSubmit={createWebSocketConn} id="username-form">
          <label for="name">Enter your name</label>
          <input type="text" id="name" required></input>
          <button class="btn">Join lobby</button>
        </form>


      </div>
    ),
  };
}
