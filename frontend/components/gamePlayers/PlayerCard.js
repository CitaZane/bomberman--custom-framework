/* @jsx jsx */
import jsx from "../../framework/vDom/jsx";

function createLivesHTML(player) {
  let lives = [];
  for (let i = 0; i < player.lives; i++) {
    lives.push(<img src="../assets/heart.png"></img>);
  }

  return lives;
}

export function PlayerCard({ player, id }) {
  let placeholderPlayer = false;

  if (!player) {
    placeholderPlayer = true;
    player = {
      name: "Default",
      lives: 3,
      active_powerups: {
        speed: 0,
        flames: 0,
        bombs: 0,
      },
    };
  }

  let lives = createLivesHTML(player);
  let playerDead = player.lives < 1;

  return {
    template: (
      <div class={placeholderPlayer || playerDead ? "invincible" : ""}>
        <p class="player-name">{player.name}</p>
        <div class="player-status">
          <div class="player-monster" id={`monster-${placeholderPlayer ? 1 : id}`}></div>
          {!playerDead && <div class="lives">{lives}</div>}
        </div>

        <div class="player-power_ups">
          <div>
            <img src="../assets/increase_speed.png"></img>
            <span class="power-up__count" id="increase-speed-count">
              {String(player["active_powerups"].speed)}
            </span>
          </div>

          <div>
            <img src="../assets/increase_flames.png"></img>
            <span class="power-up__count" id="increase-speed-count">
              {String(player["active_powerups"].flames)}
            </span>
          </div>

          <div>
            <img src="../assets/increase_bombs.png"></img>
            <span class="power-up__count" id="increase-speed-count">
              {String(player["active_powerups"].bombs)}
            </span>
          </div>
        </div>
      </div>
    ),
  };
}
