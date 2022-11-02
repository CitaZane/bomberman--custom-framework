/* @jsx jsx */
import jsx from "../../framework/vDom/jsx"
import { BombSprite } from "../components/BombSprite";
import { MonsterSprite } from "../components/MonsterSprite";
import { store } from "../app";
import { players } from "../websocket";
export function GameView() {
    let bombDrop = store.state.bomb.drop;
    return {
        template: (
            <div id="home">
                <h1>Game View</h1>
                {players.map(player => <MonsterSprite y={player.y} />)}
                {/* {bombDrop && <BombSprite />} */}
            </div >
        )
    }
}
