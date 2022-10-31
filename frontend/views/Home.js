/* @jsx jsx */
import jsx from "../../framework/vDom/jsx"
import {store} from "../app"

import { MonsterSprite } from "../components/MonsterSprite"
import { BombSprite } from "../components/BombSprite";
import { FireSprite } from "../components/FireSprite";

export function HomeView() {
    function switchType(){
        let monsterType = store.state.monster.type;
        monsterType = monsterType == 0? 1: 0;
        store.dispatch("updateMonsterType", monsterType)
    }
    return {
        template: (
            <div id="home">
                <h1>Hello monster</h1>
                <button onClick={switchType}>Switch</button>
                <MonsterSprite/>
                <BombSprite/>
                <FireSprite/>
            </div>
        )
    }
}